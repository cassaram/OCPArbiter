package gvocp

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cassaram/ocparbiter/common"
	pci "github.com/cassaram/ocparbiter/protocols/pci"
	"github.com/cassaram/ocparbiter/settings"
)

type GVOCP struct {
	systemSettings      common.SystemSettings
	connection          pci.PCIProtocol
	ocp_grp_id          uint8
	rxCount             uint16
	txCount             uint16
	initializedSettings bool
	ocpInitialized      bool
	cam                 common.Camera
	updatedCamValues    chan updateValue
	rxDataMessages      chan pci.DataMessage
	settings            []settings.Setting
}

func (ocp *GVOCP) Initialize(loadedSettings []settings.Setting) {
	// Load default settings
	ocp.loadDefaultSettings()

	// Load changed settings
	for _, val := range loadedSettings {
		ocp.SetSetting(val)
	}
	ocp.initializedSettings = true

	ocp.updatedCamValues = make(chan updateValue, 20)
	ocp.rxDataMessages = make(chan pci.DataMessage, 20)

	port, _ := ocp.GetSetting("serial_port")

	ocp.connection = pci.PCIProtocol{
		SerialPort:         port.Value.Text,
		DataMessageReceive: ocp.rxDataMessages,
		AppInitFunction:    ocp.initializeOCPValues,
	}
	ocp.connection.StartServices()

	ocp.rxCount = 0
	ocp.txCount = 0

	ocp.ocpInitialized = false

	// Start go routines for serial comms
	go ocp.rxLoop()
	go ocp.txLoop()
}

func (ocp *GVOCP) GetSystemSettings() common.SystemSettings {
	return ocp.systemSettings
}

func (ocp *GVOCP) SetSystemSettings(set common.SystemSettings) {
	ocp.systemSettings = set
}

func (ocp *GVOCP) GetSettings() []settings.Setting {
	if !ocp.initializedSettings {
		ocp.loadDefaultSettings()
	}
	return ocp.settings
}

func (ocp *GVOCP) GetSetting(id string) (settings.Setting, error) {
	for _, val := range ocp.settings {
		if val.Id == id {
			return val, nil
		}
	}

	return settings.Setting{}, errors.New("setting not found")
}

func (ocp *GVOCP) SetSetting(newSetting settings.Setting) {
	for idx, val := range ocp.settings {
		if val.Id == newSetting.Id {
			ocp.settings[idx] = newSetting
		}
	}
}

func (ocp *GVOCP) loadDefaultSettings() {
	fpath, _ := filepath.Abs("../controllers/gvocp/DefaultSettings.json")
	file, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteFile, _ := ioutil.ReadAll(file)

	var settings []settings.Setting
	json.Unmarshal(byteFile, &settings)

	// Load defaults into values
	for idx, _ := range settings {
		settings[idx].Value = settings[idx].Default
	}

	ocp.settings = settings
}

func (ocp *GVOCP) UpdateValue(pkt common.ControllerCommand) {
	// Switch command to handle odd cases
	switch pkt.Function {
	case common.BlackMaster:
		var params []byte
		/*if ocp.ocpInitialized {
			// Use extended 12-bit communication
			evalue := make([]byte, 2)
			binary.BigEndian.PutUint16(evalue, uint16(pkt.Value))

			params = append(params, byte(MBLACK_12BIT_LEVEL))
			params = append(params, evalue...)
		} else {*/
		// We must use standard 8-bit because 12-bit is broken on this
		// one specific command for some reason
		// Use standard 8-bit communication
		params = append(params, byte(ABS_VALUE_CMD))
		params = append(params, byte(MASTER_BLACK_LEVEL))
		params = append(params, byte(pkt.Value/16))
		//}

		ocp.scheduleDataMessage(
			byte(ocp.ocp_grp_id),
			params,
		)
	case common.Iris:
		params := []byte{byte(ABS_VALUE_CMD)}
		if ocp.ocpInitialized {
			// Use extended 12-bit communication
			evalue := make([]byte, 2)
			binary.BigEndian.PutUint16(evalue, uint16(pkt.Value))

			params = append(params, byte(IRIS_12BIT_LEVEL))
			params = append(params, evalue...)
		} else {
			// Use standard 8-bit communication
			params = append(params, byte(IRIS_LEVEL))
			params = append(params, byte(pkt.Value))
		}

		ocp.scheduleDataMessage(
			byte(ocp.ocp_grp_id),
			params,
		)
	case common.GainMaster:
		params := []byte{byte(ABS_VALUE_CMD)}
		// Use extended 12-bit communication
		evalue := make([]byte, 2)
		binary.BigEndian.PutUint16(evalue, uint16(pkt.Value))

		params = append(params, byte(VAR_MGAIN_LEVEL))
		params = append(params, evalue...)

		ocp.scheduleDataMessage(
			byte(ocp.ocp_grp_id),
			params,
		)
	case common.FStop:
		// Has odd mode selection behavior
		// Get value (0.01 * x)
		value := float32(pkt.Value) / float32(100)
		// Utilize helper / recursive search function
		enum := fstopToEnum(value)

		ocp.scheduleDataMessage(
			byte(ocp.ocp_grp_id),
			[]byte{
				byte(MODE_CMD),
				byte(FSTOP_SELECT),
				byte(enum),
			},
		)

	default:
		// Most likely a common 0-255 single value or mode command
		// Utilize utility function to get data
		mcmd, cmd := commonFunctionToGrassFunction(pkt.Function)
		ocp.scheduleDataMessage(
			byte(ocp.ocp_grp_id),
			[]byte{
				byte(mcmd),
				byte(cmd),
				byte(pkt.Value),
			},
		)

	}
}

func (ocp *GVOCP) SetConnectedCamera(cam common.Camera) {
	if ocp.cam != nil {
		ocp.cam.InformControllerRemove(ocp)
		ocp.cam = nil
	}

	ocp.cam = cam
	ocp.cam.InformControllerAdd(ocp)
}

func (ocp *GVOCP) GetConnectedCamera() common.Camera {
	return ocp.cam
}

// Main loop / go routine to handle incoming data from the serial port
func (ocp *GVOCP) rxLoop() {
	for {
		for val := range ocp.rxDataMessages {
			ocp.rxCount++
			ocp.handleDataMessage(val.Group, val.Params)
		}
	}
}

// Main loop / go routine to handle outgoing data to the serial port
func (ocp *GVOCP) txLoop() {
	for {
		for val := range ocp.updatedCamValues {
			ocp.txCount++
			ocp.connection.SendDataMessage(
				val.group,
				val.params,
			)
		}
	}
}

func (ocp *GVOCP) scheduleDataMessage(group byte, params []byte) {
	ocp.updatedCamValues <- updateValue{
		group:  group,
		params: params,
	}
}

// Function to initialize OCP with all values from the camera
func (ocp *GVOCP) initializeOCPValues(group byte) {
	// Set group and id
	ocp.ocp_grp_id = uint8(group)

	// Get and queue values
	vals := ocp.cam.RequestAllValues()

	for _, val := range vals {
		ocp.UpdateValue(val)
	}

	ocp.ocpInitialized = true
}

func (ocp *GVOCP) handleDataMessage(group byte, params []byte) {
	// Increment rx packets
	ocp.rxCount++

	command := params[0]

	//fmt.Println("Received data message:", command, "Params:", params[1:])

	switch GVCommand(command) {
	case ABS_VALUE_CMD:
		switch GVValueCommand(params[1]) {
		case GAIN_RED_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainRed,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case GAIN_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainGreen,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case GAIN_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainBlue,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case BLACK_RED_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackRed,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case BLACK_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackGreen,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case BLACK_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackBlue,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case FLARE_RED_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareRed,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case FLARE_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareGreen,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case FLARE_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareBlue,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case NOTCH_LEVEL:
		case SOFT_CONT_LEVEL:
		case SKIN_CONT_LEVEL:
		case SKIN1_WIDTH_RED:
		case SKIN1_WIDTH_BLUE:
		case SKIN1_COLOR_RED:
		case SKIN1_COLOR_BLUE:
		case SKIN2_WIDTH_RED:
		case SKIN2_WIDTH_BLUE:
		case SKIN2_COLOR_RED:
		case SKIN2_COLOR_BLUE:
		case MATRIX_RG:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixRedGreen,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MATRIX_RB:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixRedBlue,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MATRIX_GR:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenRed,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MATRIX_GB:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenBlue,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MATRIX_BR:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueRed,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MATRIX_BG:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueGreen,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MBLACK_12BIT_LEVEL:
			// Get params
			value := int(binary.BigEndian.Uint16(params[2:3]))
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackMaster,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case GAMMA_RED_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaRed,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case GAMMA_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaGreen,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case GAMMA_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaBlue,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case MASTER_GAMMA_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaMaster,
				Value:      value,
				Adjustment: common.Absolute,
			})

		// Protocol message updates
		case PCI_PANEL_RX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(ABS_VALUE_CMD), byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.scheduleDataMessage(group, txParams)
		case PCI_PANEL_TX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(ABS_VALUE_CMD), byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.scheduleDataMessage(group, txParams)
		}

	// Value adjustment commands
	case VALUE_CMD:
		switch GVValueCommand(params[1]) {
		case VAR_MGAIN_LEVEL:
			// Get params
			value := int(int16(binary.BigEndian.Uint16([]byte{params[2], params[3]})))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainMaster,
				Value:      value,
				Adjustment: common.Relative,
			})
		case GAIN_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainRed,
				Value:      value,
				Adjustment: common.Relative,
			})
		case GAIN_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainGreen,
				Value:      value,
				Adjustment: common.Relative,
			})
		case GAIN_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainBlue,
				Value:      value,
				Adjustment: common.Relative,
			})
		case BLACK_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackRed,
				Value:      value,
				Adjustment: common.Relative,
			})
		case BLACK_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackGreen,
				Value:      value,
				Adjustment: common.Relative,
			})
		case BLACK_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackBlue,
				Value:      value,
				Adjustment: common.Relative,
			})
		case FLARE_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareRed,
				Value:      value,
				Adjustment: common.Relative,
			})
		case FLARE_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareGreen,
				Value:      value,
				Adjustment: common.Relative,
			})
		case FLARE_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareBlue,
				Value:      value,
				Adjustment: common.Relative,
			})
		case NOTCH_LEVEL:
		case SOFT_CONT_LEVEL:
		case SKIN_CONT_LEVEL:
		case SKIN1_WIDTH_RED:
		case SKIN1_WIDTH_BLUE:
		case SKIN1_COLOR_RED:
		case SKIN1_COLOR_BLUE:
		case SKIN2_WIDTH_RED:
		case SKIN2_WIDTH_BLUE:
		case SKIN2_COLOR_RED:
		case SKIN2_COLOR_BLUE:
		case MATRIX_RG:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixRedGreen,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MATRIX_RB:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixRedBlue,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MATRIX_GR:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenRed,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MATRIX_GB:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenBlue,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MATRIX_BR:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueRed,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MATRIX_BG:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueGreen,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MBLACK_12BIT_LEVEL:
			value := int(int16(binary.LittleEndian.Uint16(params[2:4])))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackMaster,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MASTER_BLACK_LEVEL:
			// Get params (and convert to 12-bit resolution / 0-4095 scale)
			value := int(int8(params[2])) * 16
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackMaster,
				Value:      value,
				Adjustment: common.Relative,
			})
		case KNEE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.KneeLevel,
				Value:      value,
				Adjustment: common.Relative,
			})
		case IRIS_LEVEL:
			// Get params (and convert to 12-bit resolution / 0-4095 scale)
			value := int(int8(params[2])) * 16
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.Iris,
				Value:      value,
				Adjustment: common.Relative,
			})
		case IRIS_12BIT_LEVEL:
			// Get params
			value := int(int16(binary.LittleEndian.Uint16(params[2:4])))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.Iris,
				Value:      value,
				Adjustment: common.Relative,
			})
		case GAMMA_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaRed,
				Value:      value,
				Adjustment: common.Relative,
			})
		case GAMMA_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaGreen,
				Value:      value,
				Adjustment: common.Relative,
			})
		case GAMMA_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaBlue,
				Value:      value,
				Adjustment: common.Relative,
			})
		case MASTER_GAMMA_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GammaMaster,
				Value:      value,
				Adjustment: common.Relative,
			})
		}
	case ABS_SWITCH_CMD:
		switch GVSwitchParams(params[1]) {
		case NOTCH:
		case STANDARD:
		case REM_AUDIO_LEVEL:
		case SOFT_CONTOUR:
		case OCP_LOCK:
		case OCP_CONNECTED:
		case SKIN_VIEW:
		case AUTOSKIN:
		case MCP_AVAILABLE:
		case ASPECT_RATIO:
		case SCAN_REVERSE:
		case LAMP_OFF:
		case REM_ASPECT_RATIO:
		case MATRIX_GAMMA:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGamma,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case KNEE_DESAT:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.KneeDesaturationLevel,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case AUTO_IRIS:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.IrisAuto,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case CALL_SIG:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.CallSignal,
				Value:      value,
				Adjustment: common.Absolute,
			})
		case BLACKSTRETCH_TYPE:
		case CAMERA_DISABLE:
		case STANDBY:
		case VF_STANDBY:
		case STUDIO_MODE:
		case SWITCHABLE_VIDEO_TRANSMISSION:
		case SWITCHABLE_PRIVATE_DATA:
		case SWITCHABLE_INTERCOM:
		case COLOUR_BAR:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.ColorBar,
				Value:      value,
				Adjustment: common.Absolute,
			})
		}
	}

}

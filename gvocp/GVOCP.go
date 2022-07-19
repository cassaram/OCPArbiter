package gvocp

import (
	"encoding/binary"
	"fmt"

	"github.com/cassaram/ocparbiter/common"
	pci "github.com/cassaram/ocparbiter/gvocp/PCI"
)

type GVOCP struct {
	connection       pci.PCI
	rxCount          uint16
	txCount          uint16
	ocpInitialized   bool
	cam              common.Camera
	updatedCamValues chan updateValue
}

func (ocp *GVOCP) InitOCP(camera common.Camera, port string) {
	ocp.cam = camera

	ocp.connection.SetPort(port, 1)
	ocp.connection.SetDataMessageHandler(ocp.handleDataMessage)
	ocp.connection.SetInitConnectionHandler(ocp.initializeOCPValues)

	ocp.rxCount = 0
	ocp.txCount = 0

	ocp.ocpInitialized = false

	// Start running loop for OCP
	go ocp.masterLoop()
}

func (ocp *GVOCP) UpdateValue(pkt common.ControllerCommand) {
	// Switch command to handle odd cases
	switch pkt.Function {
	case common.BlackMaster:
		var params []byte
		if ocp.ocpInitialized {
			// Use extended 12-bit communication
			evalue := make([]byte, 2)
			binary.BigEndian.PutUint16(evalue, uint16(pkt.Value))

			params = append(params, byte(MBLACK_12BIT_LEVEL))
			params = append(params, evalue...)
		} else {
			// Use standard 8-bit communication
			params = append(params, byte(MASTER_BLACK_LEVEL))
			params = append(params, byte(pkt.Value))
		}

		ocp.scheduleDataMessage(
			byte(pkt.Metadata["pci_d_id"]),
			byte(pkt.Metadata["pci_group"]),
			byte(ABS_VALUE_CMD),
			params,
		)
	case common.Iris:
		var params []byte
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
			byte(pkt.Metadata["pci_d_id"]),
			byte(pkt.Metadata["pci_group"]),
			byte(ABS_VALUE_CMD),
			params,
		)
	case common.GainMaster:
		var params []byte
		// Use extended 12-bit communication
		evalue := make([]byte, 2)
		binary.BigEndian.PutUint16(evalue, uint16(pkt.Value))

		params = append(params, byte(VAR_MGAIN_LEVEL))
		params = append(params, evalue...)

		ocp.scheduleDataMessage(
			byte(pkt.Metadata["pci_d_id"]),
			byte(pkt.Metadata["pci_group"]),
			byte(ABS_VALUE_CMD),
			params,
		)
	case common.FStop:
		// Has odd mode selection behavior
		// Get value (0.01 * x)
		value := float32(pkt.Value) / float32(100)
		// Utilize helper / recursive search function
		enum := fstopToEnum(value)

		ocp.scheduleDataMessage(
			byte(pkt.Metadata["pci_d_id"]),
			byte(pkt.Metadata["pci_group"]),
			byte(MODE_CMD),
			[]byte{
				byte(FSTOP_SELECT),
				byte(enum),
			},
		)

	default:
		// Most likely a common 0-255 single value or mode command
		// Utilize utility function to get data
		mcmd, cmd := commonFunctionToGrassFunction(pkt.Function)
		ocp.scheduleDataMessage(
			byte(pkt.Metadata["pci_d_id"]),
			byte(pkt.Metadata["pci_group"]),
			byte(mcmd),
			[]byte{
				byte(cmd),
				byte(pkt.Value),
			},
		)

	}
}

func (ocp *GVOCP) masterLoop() {
	// Main loop
	for {
		// Handle outgoing data
		for val := range ocp.updatedCamValues {
			ocp.txCount++
			ocp.connection.SendDataMessage(
				val.d_id,
				val.group,
				val.masterCommand,
				val.params,
			)
		}

		// Handle incoming data
		ocp.connection.HandleData()
	}
}

func (ocp *GVOCP) scheduleDataMessage(d_id byte, group byte, command byte, params []byte) {
	ocp.updatedCamValues <- updateValue{
		d_id:          d_id,
		group:         group,
		masterCommand: command,
		params:        params,
	}
}

// Function to initialize OCP with all values from the camera
func (ocp *GVOCP) initializeOCPValues(s_id byte, group byte) {
	vals := ocp.cam.RequestAllValues()

	for _, val := range vals {
		val.Metadata = makeMetadataMap(s_id, group)
		ocp.UpdateValue(val)
	}

	ocp.ocpInitialized = true
}

func (ocp *GVOCP) handleDataMessage(s_id byte, group byte, params []byte) {
	// Increment rx packets
	ocp.rxCount++

	command := params[0]

	fmt.Println("Received data message:", command, "Params:", params[1:])

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
				Metadata:   makeMetadataMap(s_id, group),
			})
		case GAIN_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainGreen,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case GAIN_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainBlue,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case BLACK_RED_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackRed,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case BLACK_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackGreen,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case BLACK_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackBlue,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case FLARE_RED_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareRed,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case FLARE_GREEN_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareGreen,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case FLARE_BLUE_LEVEL:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareBlue,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
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
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_RB:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixRedBlue,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_GR:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenRed,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_GB:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenBlue,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_BR:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueRed,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_BG:
			// Get params
			value := int(params[2])
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueGreen,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MBLACK_12BIT_LEVEL:
			// Get params
			value := int(binary.BigEndian.Uint16(params[2:3]))
			// Force update value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackMaster,
				Value:      value,
				Adjustment: common.Absolute,
				Metadata:   makeMetadataMap(s_id, group),
			})

		// Protocol message updates
		case PCI_PANEL_RX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.scheduleDataMessage(s_id, group, byte(ABS_VALUE_CMD), txParams)
		case PCI_PANEL_TX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.scheduleDataMessage(s_id, group, byte(ABS_VALUE_CMD), txParams)
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
				Metadata:   makeMetadataMap(s_id, group),
			})
		case GAIN_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainRed,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case GAIN_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainGreen,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case GAIN_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.GainBlue,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case BLACK_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackRed,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case BLACK_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackGreen,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case BLACK_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackBlue,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case FLARE_RED_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareRed,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case FLARE_GREEN_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareGreen,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case FLARE_BLUE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.FlareBlue,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
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
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_RB:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixRedBlue,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_GR:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenRed,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_GB:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixGreenBlue,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_BR:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueRed,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MATRIX_BG:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.MatrixBlueGreen,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MBLACK_12BIT_LEVEL:
			value := int(int16(binary.LittleEndian.Uint16(params[2:4])))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackMaster,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case MASTER_BLACK_LEVEL:
			// Get params (and convert to 12-bit resolution / 0-4095 scale)
			value := int(int8(params[2])) * 16
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.BlackMaster,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case KNEE_LEVEL:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.KneeLevel,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case IRIS_LEVEL:
			// Get params (and convert to 12-bit resolution / 0-4095 scale)
			value := int(int8(params[2])) * 16
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.Iris,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case IRIS_12BIT_LEVEL:
			// Get params
			value := int(int16(binary.LittleEndian.Uint16(params[2:4])))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.Iris,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
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
		case KNEE_DESAT:
		case AUTO_IRIS:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.IrisAuto,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		case CALL_SIG:
			// Get params
			value := int(int8(params[2]))
			// Update Value
			ocp.cam.UpdateValue(common.CameraCommand{
				Function:   common.CallSignal,
				Value:      value,
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
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
				Adjustment: common.Relative,
				Metadata:   makeMetadataMap(s_id, group),
			})
		}
	}

}

func makeMetadataMap(d_id byte, group byte) map[string]int {
	m := make(map[string]int)

	m["pci_d_id"] = int(uint8(d_id))
	m["pci_group"] = int(uint8(group))

	return m
}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func intToBool(i int) bool {
	if i != 0 {
		return true
	} else {
		return false
	}
}

func calcAdjustedValue(original int, adjustment int, minClamp int, maxClamp int) (newVal int) {
	newVal = original + adjustment
	if newVal < minClamp {
		newVal = minClamp
	} else if newVal > maxClamp {
		newVal = maxClamp
	}
	return
}

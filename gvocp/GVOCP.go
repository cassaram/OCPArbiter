package gvocp

import (
	"encoding/binary"
	"fmt"

	"github.com/cassaram/ocparbiter/common"
	pci "github.com/cassaram/ocparbiter/gvocp/PCI"
)

type GVOCP struct {
	connection     pci.PCI
	rxCount        uint16
	txCount        uint16
	ocpInitialized bool
	cam            common.Cam
	camFeatures    common.CamFeatureSet
}

func (ocp *GVOCP) InitOCP(camera common.Cam, port string) {
	ocp.cam = camera
	ocp.camFeatures = ocp.cam.GetFeatureSet()

	ocp.connection.SetPort(port, 1)
	ocp.connection.SetDataMessageHandler(ocp.handleDataMessage)
	ocp.connection.SetInitConnectionHandler(ocp.initializeOCP)

	ocp.rxCount = 0
	ocp.txCount = 0

	ocp.ocpInitialized = false

	for {
		ocp.connection.HandleData()
	}
}

func (ocp *GVOCP) handleDataMessage(s_id byte, group byte, params []byte) {
	// Increment rx packets
	ocp.rxCount++

	// Check if we have initialized the OCP
	if !ocp.ocpInitialized {
		ocp.initializeOCP(s_id, group)
	}

	command := params[0]

	fmt.Println("Received data message:", command, "Params:", params[1:])

	switch GVCommand(command) {
	case ABS_VALUE_CMD:
		switch GVCommandParam(params[1]) {
		case GAIN_RED_LEVEL:
			if ocp.camFeatures.GainR {
				// Get params
				gainR := int(params[2])
				// Force update value
				ocp.cam.SetGainR(gainR)
				// Update ocp
				ocp.updateGainR(s_id, group, ocp.cam.GetGainR())
			}
		case GAIN_GREEN_LEVEL:
			if ocp.camFeatures.GainG {
				// Get params
				gainG := int(params[2])
				// Force update value
				ocp.cam.SetGainG(gainG)
				// Update ocp
				ocp.updateGainG(s_id, group, ocp.cam.GetGainG())
			}
		case GAIN_BLUE_LEVEL:
			if ocp.camFeatures.GainB {
				// Get params
				gainB := int(params[2])
				// Force update value
				ocp.cam.SetGainB(gainB)
				// Update ocp
				ocp.updateGainB(s_id, group, ocp.cam.GetGainB())
			}
		case BLACK_RED_LEVEL:
			if ocp.camFeatures.BlackR {
				// Get params
				blackR := int(params[2])
				// Force update value
				ocp.cam.SetBlackR(blackR)
				// Update ocp
				ocp.updateBlackR(s_id, group, ocp.cam.GetBlackR())
			}
		case BLACK_GREEN_LEVEL:
			if ocp.camFeatures.BlackG {
				// Get params
				blackG := int(params[2])
				// Force update value
				ocp.cam.SetBlackG(blackG)
				// Update ocp
				ocp.updateBlackG(s_id, group, ocp.cam.GetBlackG())
			}
		case BLACK_BLUE_LEVEL:
			if ocp.camFeatures.BlackB {
				// Get params
				blackB := int(params[2])
				// Force update value
				ocp.cam.SetBlackB(blackB)
				// Update ocp
				ocp.updateBlackB(s_id, group, ocp.cam.GetBlackB())
			}
		case FLARE_RED_LEVEL:
			if ocp.camFeatures.FlareR {
				// Get params
				flareR := int(params[2])
				// Force update value
				ocp.cam.SetFlareR(flareR)
				// Update ocp
				ocp.updateFlareR(s_id, group, ocp.cam.GetFlareR())
			}
		case FLARE_GREEN_LEVEL:
			if ocp.camFeatures.FlareG {
				// Get params
				flareG := int(params[2])
				// Force update value
				ocp.cam.SetFlareG(flareG)
				// Update ocp
				ocp.updateFlareG(s_id, group, ocp.cam.GetFlareG())
			}
		case FLARE_BLUE_LEVEL:
			if ocp.camFeatures.FlareB {
				// Get params
				flareB := int(params[2])
				// Force update value
				ocp.cam.SetFlareB(flareB)
				// Update ocp
				ocp.updateFlareB(s_id, group, ocp.cam.GetFlareB())
			}
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
			if ocp.camFeatures.MatrixRG {
				// Get params
				matrix := int(params[2])
				// Force update value
				ocp.cam.SetMatrixRG(matrix)
				// Update ocp
				ocp.updateMatrixRG(s_id, group, ocp.cam.GetMatrixRG())
			}
		case MATRIX_RB:
			if ocp.camFeatures.MatrixRB {
				// Get params
				matrix := int(params[2])
				// Force update value
				ocp.cam.SetMatrixRB(matrix)
				// Update ocp
				ocp.updateMatrixRB(s_id, group, ocp.cam.GetMatrixRB())
			}
		case MATRIX_GR:
			if ocp.camFeatures.MatrixGR {
				// Get params
				matrix := int(params[2])
				// Force update value
				ocp.cam.SetMatrixGR(matrix)
				// Update ocp
				ocp.updateMatrixGR(s_id, group, ocp.cam.GetMatrixGR())
			}
		case MATRIX_GB:
			if ocp.camFeatures.MatrixGB {
				// Get params
				matrix := int(params[2])
				// Force update value
				ocp.cam.SetMatrixGB(matrix)
				// Update ocp
				ocp.updateMatrixGB(s_id, group, ocp.cam.GetMatrixGB())
			}
		case MATRIX_BR:
			if ocp.camFeatures.MatrixBR {
				// Get params
				matrix := int(params[2])
				// Force update value
				ocp.cam.SetMatrixBR(matrix)
				// Update ocp
				ocp.updateMatrixBR(s_id, group, ocp.cam.GetMatrixBR())
			}
		case MATRIX_BG:
			if ocp.camFeatures.MatrixBG {
				// Get params
				matrix := int(params[2])
				// Force update value
				ocp.cam.SetMatrixBG(matrix)
				// Update ocp
				ocp.updateMatrixBG(s_id, group, ocp.cam.GetMatrixBG())
			}
		case MBLACK_12BIT_LEVEL:
			if ocp.camFeatures.BlackMaster {
				// Get params
				black := int(binary.BigEndian.Uint16(params[2:3]))
				// Force update value
				ocp.cam.SetBlackMaster(black)
				// Update ocp
				ocp.updateMBlackL(s_id, group, ocp.cam.GetBlackMaster(), true)
			}

		// Protocol message updates
		case PCI_PANEL_RX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.connection.SendDataMessage(s_id, group, byte(ABS_VALUE_CMD), txParams)
		case PCI_PANEL_TX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.connection.SendDataMessage(s_id, group, byte(ABS_VALUE_CMD), txParams)
		}

	// Value adjustment commands
	case VALUE_CMD:
		switch GVCommandParam(params[1]) {
		case VAR_MGAIN_LEVEL:
			// Get params
			gain := int(int16(binary.BigEndian.Uint16([]byte{params[2], params[3]})))
			// Calculate adjusted value
			gain = calcAdjustedValue(ocp.cam.GetGainMaster(), gain, 0, 4095)
			// Force update value
			ocp.cam.SetGainMaster(gain)
			// Update ocp
			ocp.updateMGainL(s_id, group, ocp.cam.GetGainMaster())
		case GAIN_RED_LEVEL:
			// Get params
			gainR := int(int8(params[2]))
			// Calculate adjusted value
			gainR = calcAdjustedValue(ocp.cam.GetGainR(), gainR, 0, 255)
			// Force update value
			ocp.cam.SetGainR(gainR)
			// Update ocp
			ocp.updateGainR(s_id, group, ocp.cam.GetGainR())
		case GAIN_GREEN_LEVEL:
			// Get params
			gainG := int(int8(params[2]))
			// Calculate adjusted value
			gainG = calcAdjustedValue(ocp.cam.GetGainG(), gainG, 0, 255)
			// Force update value
			ocp.cam.SetGainG(gainG)
			// Update ocp
			ocp.updateGainG(s_id, group, ocp.cam.GetGainG())
		case GAIN_BLUE_LEVEL:
			// Get params
			gainB := int(int8(params[2]))
			// Calculate adjusted value
			gainB = calcAdjustedValue(ocp.cam.GetGainB(), gainB, 0, 255)
			// Force update value
			ocp.cam.SetGainB(gainB)
			// Update ocp
			ocp.updateGainB(s_id, group, ocp.cam.GetGainB())
		case BLACK_RED_LEVEL:
			// Get params
			blackR := int(int8(params[2]))
			// Calculate adjusted value
			blackR = calcAdjustedValue(ocp.cam.GetBlackR(), blackR, 0, 255)
			// Force update value
			ocp.cam.SetBlackR(blackR)
			// Update ocp
			ocp.updateBlackR(s_id, group, ocp.cam.GetBlackR())
		case BLACK_GREEN_LEVEL:
			// Get params
			blackG := int(int8(params[2]))
			// Calculate adjusted value
			blackG = calcAdjustedValue(ocp.cam.GetBlackG(), blackG, 0, 255)
			// Force update value
			ocp.cam.SetBlackG(blackG)
			// Update ocp
			ocp.updateBlackG(s_id, group, ocp.cam.GetBlackG())
		case BLACK_BLUE_LEVEL:
			// Get params
			blackB := int(int8(params[2]))
			// Calculate adjusted value
			blackB = calcAdjustedValue(ocp.cam.GetBlackB(), blackB, 0, 255)
			// Force update value
			ocp.cam.SetBlackB(blackB)
			// Update ocp
			ocp.updateBlackB(s_id, group, ocp.cam.GetBlackB())
		case FLARE_RED_LEVEL:
			// Get params
			flareR := int(int8(params[2]))
			// Calculate adjusted value
			flareR = calcAdjustedValue(ocp.cam.GetFlareR(), flareR, 0, 255)
			// Force update value
			ocp.cam.SetFlareR(flareR)
			// Update ocp
			ocp.updateFlareR(s_id, group, ocp.cam.GetFlareR())
		case FLARE_GREEN_LEVEL:
			// Get params
			flareG := int(int8(params[2]))
			// Calculate adjusted value
			flareG = calcAdjustedValue(ocp.cam.GetFlareG(), flareG, 0, 255)
			// Force update value
			ocp.cam.SetFlareG(flareG)
			// Update ocp
			ocp.updateFlareG(s_id, group, ocp.cam.GetFlareG())
		case FLARE_BLUE_LEVEL:
			// Get params
			flareB := int(int8(params[2]))
			// Calculate adjusted value
			flareB = calcAdjustedValue(ocp.cam.GetFlareB(), flareB, 0, 255)
			// Force update value
			ocp.cam.SetFlareB(flareB)
			// Update ocp
			ocp.updateFlareB(s_id, group, ocp.cam.GetFlareB())
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
			if ocp.camFeatures.MatrixRG {
				// Get params
				value := int(int8(params[2]))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetMatrixRG(), value, 0, 255)
				// Force update value
				ocp.cam.SetMatrixRG(value)
				// Update ocp
				ocp.updateMatrixRG(s_id, group, ocp.cam.GetMatrixRG())
			}
		case MATRIX_RB:
			if ocp.camFeatures.MatrixRB {
				// Get params
				value := int(int8(params[2]))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetMatrixRB(), value, 0, 255)
				// Force update value
				ocp.cam.SetMatrixRB(value)
				// Update ocp
				ocp.updateMatrixRB(s_id, group, ocp.cam.GetMatrixRB())
			}
		case MATRIX_GR:
			if ocp.camFeatures.MatrixGR {
				// Get params
				value := int(int8(params[2]))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetMatrixGR(), value, 0, 255)
				// Force update value
				ocp.cam.SetMatrixGR(value)
				// Update ocp
				ocp.updateMatrixGR(s_id, group, ocp.cam.GetMatrixGR())
			}
		case MATRIX_GB:
			if ocp.camFeatures.MatrixGB {
				// Get params
				value := int(int8(params[2]))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetMatrixGB(), value, 0, 255)
				// Force update value
				ocp.cam.SetMatrixGB(value)
				// Update ocp
				ocp.updateMatrixGB(s_id, group, ocp.cam.GetMatrixGB())
			}
		case MATRIX_BR:
			if ocp.camFeatures.MatrixBR {
				// Get params
				value := int(int8(params[2]))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetMatrixBR(), value, 0, 255)
				// Force update value
				ocp.cam.SetMatrixBR(value)
				// Update ocp
				ocp.updateMatrixBR(s_id, group, ocp.cam.GetMatrixBR())
			}
		case MATRIX_BG:
			if ocp.camFeatures.MatrixBG {
				// Get params
				value := int(int8(params[2]))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetMatrixBG(), value, 0, 255)
				// Force update value
				ocp.cam.SetMatrixBG(value)
				// Update ocp
				ocp.updateMatrixBG(s_id, group, ocp.cam.GetMatrixBG())
			}
		case MBLACK_12BIT_LEVEL:
			if ocp.camFeatures.BlackMaster {
				// Get params
				value := int(int16(binary.LittleEndian.Uint16(params[2:4])))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetBlackMaster(), value, 0, 4095)
				// Force update value
				ocp.cam.SetBlackMaster(value)
				// Update ocp
				ocp.updateMBlackL(s_id, group, ocp.cam.GetBlackMaster(), false)
			}
		case MASTER_BLACK_LEVEL:
			if ocp.camFeatures.BlackMaster {
				// Get params (and convert to 12-bit resolution / 0-4095 scale)
				value := int(int8(params[2])) * 16
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetBlackMaster(), value, 0, 4095)
				// Force update value
				ocp.cam.SetBlackMaster(value)
				// Update ocp
				ocp.updateMBlackL(s_id, group, ocp.cam.GetBlackMaster(), false)
			}
		case KNEE_LEVEL:
		case IRIS_LEVEL:
		case IRIS_12BIT_LEVEL:
			if ocp.camFeatures.Iris {
				// Get params
				value := int(int16(binary.LittleEndian.Uint16(params[2:4])))
				// Calculate adjusted value
				value = calcAdjustedValue(ocp.cam.GetIris(), value, 0, 4095)
				// Force update value
				ocp.cam.SetIris(value)
				// Update ocp
				ocp.updateIrisL(s_id, group, ocp.cam.GetIris(), true)
				ocp.updateFStop(s_id, group, ocp.cam.GetFStop())
			}
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
			if ocp.camFeatures.IrisAuto {
				// Get params
				value := int(int8(params[2]))
				// Force update value
				ocp.cam.SetIrisAuto(value)
				// Update ocp
				ocp.updateIrisAuto(s_id, group, GVModeIrisAuto(ocp.cam.GetIrisAuto()))
			}
		case CALL_SIG:
			if ocp.camFeatures.CallSignal {
				// Get params
				value := int(int8(params[2]))
				// Force update value
				ocp.cam.SetCallSig(value)
				// Update ocp
				ocp.updateCallSignal(s_id, group, GVModeCallSignal(ocp.cam.GetCallSig()))
			}
		case BLACKSTRETCH_TYPE:
		case CAMERA_DISABLE:
		case STANDBY:
		case VF_STANDBY:
		case STUDIO_MODE:
		case SWITCHABLE_VIDEO_TRANSMISSION:
		case SWITCHABLE_PRIVATE_DATA:
		case SWITCHABLE_INTERCOM:
		case COLOUR_BAR:
			if ocp.camFeatures.ColorBar {
				// Get params
				value := int(int8(params[2]))
				// Force update value
				ocp.cam.SetColorBar(value)
				// Update ocp
				ocp.updateColorBar(s_id, group, GVModeColorBar(ocp.cam.GetColorBar()))
			}
		}
	}

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

func (ocp *GVOCP) initializeOCP(d_id byte, group byte) {
	if ocp.camFeatures.CamNumber {
		ocp.updateCamNumber(d_id, group, ocp.cam.GetCamNumber())
	}

	// Other opt-in features
	if ocp.camFeatures.CallSignal {
		ocp.updateCallSignal(d_id, group, GVModeCallSignal(ocp.cam.GetCallSig()))
	}
	if ocp.camFeatures.ColorBar {
		ocp.updateColorBar(d_id, group, GVModeColorBar(ocp.cam.GetColorBar()))
	}
	if ocp.camFeatures.GainMaster {
		ocp.updateMGainL(d_id, group, ocp.cam.GetGainMaster())
	}
	if ocp.camFeatures.GainR {
		ocp.updateGainR(d_id, group, ocp.cam.GetGainR())
	}
	if ocp.camFeatures.GainG {
		ocp.updateGainG(d_id, group, ocp.cam.GetGainG())
	}
	if ocp.camFeatures.GainB {
		ocp.updateGainB(d_id, group, ocp.cam.GetGainB())
	}
	if ocp.camFeatures.BlackR {
		ocp.updateBlackR(d_id, group, ocp.cam.GetBlackR())
	}
	if ocp.camFeatures.BlackG {
		ocp.updateBlackG(d_id, group, ocp.cam.GetBlackG())
	}
	if ocp.camFeatures.BlackB {
		ocp.updateBlackB(d_id, group, ocp.cam.GetBlackB())
	}
	if ocp.camFeatures.BlackMaster {
		ocp.updateMBlackL(d_id, group, ocp.cam.GetBlackMaster(), false)
	}
	if ocp.camFeatures.FlareR {
		ocp.updateFlareR(d_id, group, ocp.cam.GetFlareR())
	}
	if ocp.camFeatures.FlareG {
		ocp.updateFlareG(d_id, group, ocp.cam.GetFlareG())
	}
	if ocp.camFeatures.FlareB {
		ocp.updateFlareB(d_id, group, ocp.cam.GetFlareB())
	}
	if ocp.camFeatures.Iris {
		ocp.updateIrisL(d_id, group, ocp.cam.GetIris(), false)
	}
	if ocp.camFeatures.FStop {
		ocp.updateFStop(d_id, group, ocp.cam.GetFStop())
	}
	if ocp.camFeatures.IrisAuto {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(SWITCH_CMD),
			[]byte{
				byte(AUTO_IRIS),
				byte(ocp.cam.GetIrisAuto()),
			},
		)
	}
	if ocp.camFeatures.GammaLevel {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(MASTER_GAMMA_LEVEL),
				byte(ocp.cam.GetGammaLevel()),
			},
		)
	}
	if ocp.camFeatures.GammaR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAMMA_RED_LEVEL),
				byte(ocp.cam.GetGammaR()),
			},
		)
	}
	if ocp.camFeatures.GammaG {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAMMA_GREEN_LEVEL),
				byte(ocp.cam.GetGammaG()),
			},
		)
	}
	if ocp.camFeatures.GammaB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAMMA_BLUE_LEVEL),
				byte(ocp.cam.GetGammaB()),
			},
		)
	}
	if ocp.camFeatures.WBR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(WH_BAL_RED_LEVEL),
				byte(ocp.cam.GetWBR()),
			},
		)
	}
	if ocp.camFeatures.WBB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(WH_BAL_BLUE_LEVEL),
				byte(ocp.cam.GetWBB()),
			},
		)
	}

	ocp.ocpInitialized = true
}

func (ocp *GVOCP) updateGainR(d_id byte, group byte, gain int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(GAIN_RED_LEVEL),
			byte(gain),
		},
	)
}

func (ocp *GVOCP) updateGainG(d_id byte, group byte, gain int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(GAIN_GREEN_LEVEL),
			byte(gain),
		},
	)
}

func (ocp *GVOCP) updateGainB(d_id byte, group byte, gain int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(GAIN_BLUE_LEVEL),
			byte(gain),
		},
	)
}

func (ocp *GVOCP) updateBlackR(d_id byte, group byte, black int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(BLACK_RED_LEVEL),
			byte(black),
		},
	)
}

func (ocp *GVOCP) updateBlackG(d_id byte, group byte, black int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(BLACK_GREEN_LEVEL),
			byte(black),
		},
	)
}

func (ocp *GVOCP) updateBlackB(d_id byte, group byte, black int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(BLACK_BLUE_LEVEL),
			byte(black),
		},
	)
}

func (ocp *GVOCP) updateFlareR(d_id byte, group byte, flare int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(FLARE_RED_LEVEL),
			byte(flare),
		},
	)
}

func (ocp *GVOCP) updateFlareG(d_id byte, group byte, flare int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(FLARE_GREEN_LEVEL),
			byte(flare),
		},
	)
}

func (ocp *GVOCP) updateFlareB(d_id byte, group byte, flare int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(FLARE_BLUE_LEVEL),
			byte(flare),
		},
	)
}

func (ocp *GVOCP) updateNotchL(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(NOTCH_LEVEL),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSoftContL(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SOFT_CONT_LEVEL),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkinContL(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN_CONT_LEVEL),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin1WidthR(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN1_WIDTH_RED),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin1WidthB(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN1_WIDTH_BLUE),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin1ColorR(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN1_COLOR_RED),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin1ColorB(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN1_COLOR_BLUE),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin2WidthR(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN2_WIDTH_RED),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin2WidthB(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN2_WIDTH_BLUE),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin2ColorR(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN2_COLOR_RED),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateSkin2ColorB(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SKIN2_COLOR_BLUE),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMatrixRG(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(MATRIX_RG),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMatrixRB(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(MATRIX_RB),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMatrixGR(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(MATRIX_GR),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMatrixGB(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(MATRIX_GB),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMatrixBR(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(MATRIX_BR),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMatrixBG(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(MATRIX_BG),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateMBlackL(d_id byte, group byte, value int, extended bool) {
	if extended {
		// 12-bit
		// Separate int into 12-bit byte slice
		txParams := make([]byte, 2)
		binary.BigEndian.PutUint16(txParams, uint16(value))
		//txParams[1] &= 0x0F

		ocp.txCount++
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(MBLACK_12BIT_LEVEL),
				txParams[0],
				txParams[1],
			},
		)
	} else {
		// 8-bit
		// Condense from 12-bit (0-4095) -> (0-255)
		value = value / 16
		ocp.txCount++
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(MASTER_BLACK_LEVEL),
				byte(value),
			},
		)
	}

}

func (ocp *GVOCP) updateIrisL(d_id byte, group byte, value int, extended bool) {
	if extended {
		// 12-bit Level

		// Separate int into 12-bit byte slice
		txParams := make([]byte, 2)
		binary.BigEndian.PutUint16(txParams, uint16(value))
		txParams[0] &= 0x0F

		ocp.txCount++
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(IRIS_12BIT_LEVEL),
				txParams[0],
				txParams[1],
			},
		)
	} else {
		// 8-bit level
		// Condense from 12-bit (0-4095) -> (0-255)
		value = value / 16
		ocp.txCount++
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(IRIS_LEVEL),
				byte(value),
			},
		)
	}
}

func (ocp *GVOCP) updateMGainL(d_id byte, group byte, value int) {
	// Separate int into 12-bit byte slice
	txParams := make([]byte, 2)
	binary.BigEndian.PutUint16(txParams, uint16(value))
	txParams[0] &= 0x0F

	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(VAR_MGAIN_LEVEL),
			txParams[0],
			txParams[1],
		},
	)
}

func (ocp *GVOCP) updateCamNumber(d_id byte, group byte, value int) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(ABS_VALUE_CMD),
		[]byte{
			byte(SERIAL_CAMERA_NUMBER),
			byte(value),
		},
	)
}

// ----- Switch Commands -----

func (ocp *GVOCP) updateIrisAuto(d_id byte, group byte, value GVModeIrisAuto) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(SWITCH_CMD),
		[]byte{
			byte(AUTO_IRIS),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateCallSignal(d_id byte, group byte, value GVModeCallSignal) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(SWITCH_CMD),
		[]byte{
			byte(CALL_SIG),
			byte(value),
		},
	)
}

func (ocp *GVOCP) updateColorBar(d_id byte, group byte, value GVModeColorBar) {
	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(SWITCH_CMD),
		[]byte{
			byte(COLOUR_BAR),
			byte(value),
		},
	)
}

// ----- Mode Commands -----

func (ocp *GVOCP) updateFStop(d_id byte, group byte, value float32) {
	enum := fstopToEnum(value)

	ocp.txCount++
	ocp.connection.SendDataMessage(
		d_id,
		group,
		byte(MODE_CMD),
		[]byte{
			byte(FSTOP_SELECT),
			byte(enum),
		},
	)
}

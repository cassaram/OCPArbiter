package gvserialbasic

import (
	"encoding/binary"

	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/protocols/pciprotocol"
)

// Function to handle converting a PCI Data Message to a camera command
func GVSBToCommon(msg pciprotocol.DataMessage) common.CameraCommand {
	switch msg.Params[0] {
	case byte(ABS_VALUE_CMD):
		return gvsbValueToCommon(msg)
	case byte(VALUE_CMD):
		return gvsbValueToCommon(msg)
	case byte(SWITCH_CMD):
		return gvsbSwitchToCommon(msg)
	case byte(MODE_CMD):
		return gvsbModeToCommon(msg)
	case byte(ABS_SWITCH_CMD):
		return gvsbSwitchToCommon(msg)
	default:
		return common.CameraCommand{
			Function: common.NOFUNC_ERR,
		}
	}
}

// Function to tell which command types are relative or absolute
func gvsbToAdjustment(msg pciprotocol.DataMessage) common.CameraCommandAdjustment {
	switch GVCommand(msg.Params[0]) {
	case VALUE_CMD:
		return common.Relative
	case SWITCH_CMD:
		return common.Relative
	default:
		return common.Absolute
	}
}

// Sub-function which handles (ABS_)VALUE_CMD commands
func gvsbValueToCommon(msg pciprotocol.DataMessage) common.CameraCommand {
	switch byte(msg.Params[1]) {
	case byte(VAR_MGAIN_LEVEL):
		val := int(int16(binary.BigEndian.Uint16(msg.Params[2:4])))
		return common.CameraCommand{
			Function:   common.GainMaster,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(GAIN_RED_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GainRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(GAIN_GREEN_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GainGreen,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(GAIN_BLUE_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GainBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(MBLACK_12BIT_LEVEL):
		val := int(int16(binary.LittleEndian.Uint16(msg.Params[2:4])))
		return common.CameraCommand{
			Function:   common.BlackMaster,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(BLACK_RED_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.BlackRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(BLACK_GREEN_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.BlackGreen,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(BLACK_BLUE_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.BlackBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(FLARE_RED_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.FlareRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(FLARE_GREEN_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.FlareGreen,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(FLARE_BLUE_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.FlareBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(IRIS_12BIT_LEVEL):
		val := int(int16(binary.BigEndian.Uint16(msg.Params[2:4])))
		return common.CameraCommand{
			Function:   common.Iris,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeLevel,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_DESAT_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeDesaturationLevel,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_SLOPE_R):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeSlopeRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_SLOPE_B):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeSlopeBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_ATTACK_M):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeAttack,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_ATTACK_R):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeAttackRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_ATTACK_B):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeAttackBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(KNEE_POINT_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneePoint,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(MASTER_GAMMA_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GammaMaster,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(GAMMA_RED_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GammaRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(GAMMA_GREEN_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GammaGreen,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(GAMMA_BLUE_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GammaBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(WH_BAL_RED_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.WhiteBalanceRed,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case byte(WH_BAL_BLUE_LEVEL):
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.WhiteBalanceBlue,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	default:
		return common.CameraCommand{
			Function: common.NOFUNC_ERR,
		}
	}
}

// Sub-function which handles (ABS_)SWITCH_CMD commands
func gvsbSwitchToCommon(msg pciprotocol.DataMessage) common.CameraCommand {
	switch GVCommand(msg.Params[1]) {
	case AUTO_IRIS:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.IrisAuto,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case EXTENDED_IRIS:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.IrisExtended,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case CALL_SIG:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.CallSignal,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case COLOUR_BAR:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.ColorBar,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case AUTOWHITE:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.WhiteBalanceAutoAction,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case FLARE_SELECT:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.FlareMode,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	default:
		return common.CameraCommand{
			Function: common.NOFUNC_ERR,
		}
	}
}

// Sub-function which handles MODE_CMD commands
func gvsbModeToCommon(msg pciprotocol.DataMessage) common.CameraCommand {
	switch GVCommand(msg.Params[1]) {
	case COLOUR_TEMP_SELECT:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.WhiteBalanceMode,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case KNEE_SELECT:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.KneeMode,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	case GAMMA_SELECT:
		val := int(int8(msg.Params[2]))
		return common.CameraCommand{
			Function:   common.GammaMode,
			Value:      val,
			Adjustment: gvsbToAdjustment(msg),
		}
	default:
		return common.CameraCommand{
			Function: common.NOFUNC_ERR,
		}
	}
}

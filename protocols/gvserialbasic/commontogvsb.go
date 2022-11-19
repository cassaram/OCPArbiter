package gvserialbasic

import (
	"encoding/binary"

	"github.com/cassaram/ocparbiter/common"
)

func CommonToGVSB(cmd common.ControllerCommand) []byte {
	switch cmd.Function {
	case common.CameraNumber:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(SERIAL_CAMERA_NUMBER), byte(val)}
	case common.CallSignal:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_SWITCH_CMD), byte(CALL_SIG), byte(val)}
	case common.ColorBar:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_SWITCH_CMD), byte(COLOUR_BAR), byte(val)}
	case common.GainMaster:
		val := make([]byte, 2)
		binary.BigEndian.PutUint16(val, uint16(cmd.Value))
		return []byte{byte(ABS_VALUE_CMD), byte(VAR_MGAIN_LEVEL), val[0], val[1]}
	case common.GainRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(GAIN_RED_LEVEL), byte(val)}
	case common.GainGreen:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(GAIN_GREEN_LEVEL), byte(val)}
	case common.GainBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(GAIN_BLUE_LEVEL), byte(val)}
	case common.BlackMaster:
		val := common.ConvertRange(0, 4095, 0, 255, cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(MASTER_BLACK_LEVEL), byte(val)}
	case common.BlackRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(BLACK_RED_LEVEL), byte(val)}
	case common.BlackGreen:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(BLACK_GREEN_LEVEL), byte(val)}
	case common.BlackBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(BLACK_BLUE_LEVEL), byte(val)}
	case common.FlareMode:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_SWITCH_CMD), byte(FLARE), byte(val)}
	case common.FlareRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(FLARE_RED_LEVEL), byte(val)}
	case common.FlareGreen:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(FLARE_GREEN_LEVEL), byte(val)}
	case common.FlareBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(FLARE_BLUE_LEVEL), byte(val)}
	case common.Iris:
		val := make([]byte, 2)
		binary.BigEndian.PutUint16(val, uint16(cmd.Value))
		return []byte{byte(ABS_VALUE_CMD), byte(VAR_MGAIN_LEVEL), val[0], val[1]}
	case common.FStop:
		val := fstopToEnum(float32(cmd.Value) / float32(10))
		return []byte{byte(MODE_CMD), byte(FSTOP_SELECT), byte(val)}
	case common.IrisAuto:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_SWITCH_CMD), byte(AUTO_IRIS), byte(val)}
	case common.IrisExtended:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_SWITCH_CMD), byte(EXTENDED_IRIS), byte(val)}
	case common.KneeMode:
		val := uint8(cmd.Value)
		return []byte{byte(MODE_CMD), byte(KNEE_SELECT), byte(val)}
	case common.KneeLevel:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_LEVEL), byte(val)}
	case common.KneeDesaturationLevel:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_DESAT_LEVEL), byte(val)}
	case common.KneeSlopeRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_SLOPE_R), byte(val)}
	case common.KneeSlopeBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_SLOPE_B), byte(val)}
	case common.KneeAttack:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_ATTACK_M), byte(val)}
	case common.KneeAttackRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_ATTACK_R), byte(val)}
	case common.KneeAttackBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_ATTACK_B), byte(val)}
	case common.KneePoint:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(KNEE_POINT_LEVEL), byte(val)}
	case common.GammaMode:
		val := uint8(cmd.Value)
		return []byte{byte(MODE_CMD), byte(GAMMA_SELECT), byte(val)}
	case common.GammaMaster:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(MASTER_GAMMA_LEVEL), byte(val)}
	case common.GammaRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(GAMMA_RED_LEVEL), byte(val)}
	case common.GammaGreen:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(GAMMA_GREEN_LEVEL), byte(val)}
	case common.GammaBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(GAMMA_BLUE_LEVEL), byte(val)}
	case common.WhiteBalanceMode:
		val := uint8(cmd.Value)
		return []byte{byte(MODE_CMD), byte(COLOUR_TEMP_SELECT), byte(val)}
	case common.WhiteBalanceRed:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(WH_BAL_RED_LEVEL), byte(val)}
	case common.WhiteBalanceBlue:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_VALUE_CMD), byte(WH_BAL_BLUE_LEVEL), byte(val)}
	case common.WhiteBalanceAutoAction:
		val := uint8(cmd.Value)
		return []byte{byte(ABS_SWITCH_CMD), byte(AUTOWHITE), byte(val)}
	default:
		return []byte{}
	}
}

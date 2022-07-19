package gvocp

import "github.com/cassaram/ocparbiter/common"

// Library of helper functions for the grass command definitions

// Struct to hold data for calculating enum from float for fstops
type fstopEnhancedEnum struct {
	enum  GVModeFStop
	focal float32
}

// Function to convert an f-stop value as a float to a define
func fstopToEnum(fstop float32) (enum GVModeFStop) {
	enums := []fstopEnhancedEnum{
		{FCT_FSTOP_010, 1.0},
		{FCT_FSTOP_012, 1.2},
		{FCT_FSTOP_013, 1.3},
		{FCT_FSTOP_014, 1.4},
		{FCT_FSTOP_015, 1.5},
		{FCT_FSTOP_017, 1.7},
		{FCT_FSTOP_018, 1.8},
		{FCT_FSTOP_020, 2.0},
		{FCT_FSTOP_022, 2.2},
		{FCT_FSTOP_024, 2.4},
		{FCT_FSTOP_026, 2.6},
		{FCT_FSTOP_028, 2.8},
		{FCT_FSTOP_031, 3.1},
		{FCT_FSTOP_034, 3.4},
		{FCT_FSTOP_037, 3.7},
		{FCT_FSTOP_040, 4.0},
		{FCT_FSTOP_044, 4.4},
		{FCT_FSTOP_048, 4.8},
		{FCT_FSTOP_052, 5.2},
		{FCT_FSTOP_056, 5.6},
		{FCT_FSTOP_062, 6.2},
		{FCT_FSTOP_067, 6.7},
		{FCT_FSTOP_073, 7.3},
		{FCT_FSTOP_080, 8.0},
		{FCT_FSTOP_087, 8.7},
		{FCT_FSTOP_095, 9.5},
		{FCT_FSTOP_100, 10.0},
		{FCT_FSTOP_110, 11.0},
		{FCT_FSTOP_120, 12.0},
		{FCT_FSTOP_130, 13.0},
		{FCT_FSTOP_150, 15.0},
		{FCT_FSTOP_160, 16.0},
		{FCT_FSTOP_170, 17.0},
		{FCT_FSTOP_190, 19.0},
		{FCT_FSTOP_210, 21.0},
		{FCT_FSTOP_220, 22.0},
		{FCT_FSTOP_250, 25.0},
		{FCT_FSTOP_270, 27.0},
		{FCT_FSTOP_290, 29.0},
		{FCT_FSTOP_320, 32.0},
		{FCT_FSTOP_350, 35.0},
		{FCT_FSTOP_380, 38.0},
		{FCT_FSTOP_420, 42.0},
		{FCT_FSTOP_450, 45.0},
		{FCT_FSTOP_490, 49.0},
		{FCT_FSTOP_530, 53.0},
		{FCT_FSTOP_590, 59.0},
		{FCT_FSTOP_640, 64.0},
		{FCT_FSTOP_16M, 16000000.0},
		{FCT_FSTOP_22M, 22000000.0},
		{FCT_FSTOP_27M, 27000000.0},
		{FCT_FSTOP_32M, 32000000.0},
	}

	return recursiveFStopToEnum(enums, fstop)
}

func recursiveFStopToEnum(enums []fstopEnhancedEnum, value float32) GVModeFStop {
	// Recursive break
	if len(enums) == 1 {
		return enums[0].enum
	}

	// Binary search
	left := 0
	right := len(enums)
	mid := (left + right) / 2
	if value < enums[mid].focal {
		return recursiveFStopToEnum(enums[left:mid], value)
	} else {
		return recursiveFStopToEnum(enums[mid:right], value)
	}
}

func commonFunctionToGrassFunction(function common.CameraFunction) (GVCommand, int) {
	// Returns all commands which are straightforward [{GVCMD}, {GVFUNC}, {VAL (0-255)}] format
	// Used to simplify other functions
	/* Current edgecases:
	 * - MBLACK_12BIT_LEVEL
	 * - MASTER_BLACK_LEVEL
	 * - IRIS_LEVEL
	 * - IRIS_12BIT_LEVEL
	 * - FSTOP_SELECT
	 */

	switch function {
	case common.CameraNumber:
		return ABS_VALUE_CMD, int(BS_CAMERA_NUMBER)
	case common.CallSignal:
		return ABS_VALUE_CMD, int(CALL_SIG)
	case common.ColorBar:
		return ABS_VALUE_CMD, int(COLOUR_BAR)
	case common.GainRed:
		return ABS_VALUE_CMD, int(GAIN_RED_LEVEL)
	case common.GainGreen:
		return ABS_VALUE_CMD, int(GAIN_GREEN_LEVEL)
	case common.GainBlue:
		return ABS_VALUE_CMD, int(GAIN_BLUE_LEVEL)
	case common.BlackRed:
		return ABS_VALUE_CMD, int(BLACK_RED_LEVEL)
	case common.BlackGreen:
		return ABS_VALUE_CMD, int(BLACK_GREEN_LEVEL)
	case common.BlackBlue:
		return ABS_VALUE_CMD, int(BLACK_BLUE_LEVEL)
	case common.FlareRed:
		return ABS_VALUE_CMD, int(FLARE_RED_LEVEL)
	case common.FlareGreen:
		return ABS_VALUE_CMD, int(FLARE_GREEN_LEVEL)
	case common.FlareBlue:
		return ABS_VALUE_CMD, int(FLARE_BLUE_LEVEL)
	case common.MatrixRedGreen:
		return ABS_VALUE_CMD, int(MATRIX_RG)
	case common.MatrixRedBlue:
		return ABS_VALUE_CMD, int(MATRIX_RB)
	case common.MatrixGreenRed:
		return ABS_VALUE_CMD, int(MATRIX_GR)
	case common.MatrixGreenBlue:
		return ABS_VALUE_CMD, int(MATRIX_GB)
	case common.MatrixBlueRed:
		return ABS_VALUE_CMD, int(MATRIX_BR)
	case common.MatrixBlueGreen:
		return ABS_VALUE_CMD, int(MATRIX_BG)
	case common.MatrixGamma:
		return ABS_VALUE_CMD, int(MATRIX_GAMMA)
	case common.IrisAuto:
		return ABS_SWITCH_CMD, int(AUTO_IRIS)
	case common.IrisExtended:
		return ABS_SWITCH_CMD, int(EXTENDED_IRIS)
	case common.KneeLevel:
		return ABS_VALUE_CMD, int(KNEE_LEVEL)
	case common.KneeDesaturationLevel:
		return ABS_VALUE_CMD, int(KNEE_DESAT_LEVEL)
	case common.KneeSlope:
		return ABS_VALUE_CMD, int(KNEE_SLOPE_M)
	case common.KneeSlopeRed:
		return ABS_VALUE_CMD, int(KNEE_SLOPE_R)
	case common.KneeSlopeBlue:
		return ABS_VALUE_CMD, int(KNEE_SLOPE_B)
	case common.KneeAttack:
		return ABS_VALUE_CMD, int(KNEE_ATTACK_M)
	case common.KneeAttackRed:
		return ABS_VALUE_CMD, int(KNEE_ATTACK_R)
	case common.KneeAttackBlue:
		return ABS_VALUE_CMD, int(KNEE_ATTACK_B)
	case common.KneePoint:
		return ABS_VALUE_CMD, int(KNEE_POINT_LEVEL)
	case common.GammaMaster:
		return ABS_VALUE_CMD, int(MASTER_GAMMA_LEVEL)
	case common.GammaRed:
		return ABS_VALUE_CMD, int(GAMMA_RED_LEVEL)
	case common.GammaGreen:
		return ABS_VALUE_CMD, int(GAMMA_GREEN_LEVEL)
	case common.GammaBlue:
		return ABS_VALUE_CMD, int(GAMMA_BLUE_LEVEL)
	case common.WhiteBalanceRed:
		return ABS_VALUE_CMD, int(WH_BAL_RED_LEVEL)
	case common.WhiteBalanceBlue:
		return ABS_VALUE_CMD, int(WH_BAL_BLUE_LEVEL)
	default:
		return ABS_VALUE_CMD, int(-1)
	}
}

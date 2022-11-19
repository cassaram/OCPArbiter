package gvserialbasic

// A collection of functions which are useful for converting values to enums

// Struct to hold data for calculating enum from float for fstops
type fstopEnhancedEnum struct {
	enum  GVCommand
	focal float32
}

// Function to convert an f-stop value as a float to a define
func fstopToEnum(fstop float32) (enum GVCommand) {
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

func recursiveFStopToEnum(enums []fstopEnhancedEnum, value float32) GVCommand {
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

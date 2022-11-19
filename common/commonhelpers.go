package common

func ConvertRange(OldLow int, OldHigh int, NewLow int, NewHigh int, Value int) int {
	if OldHigh-OldLow == 0 {
		// Old range is zero
		return NewLow
	} else {
		return (((Value - OldLow) * (NewHigh - NewLow)) / (OldHigh - OldLow)) + NewLow
	}
}

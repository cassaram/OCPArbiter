package settings

func ValidateSetting(s Setting) bool {
	// Check basic settings
	if s.Id == "" || s.Label == "" {
		return false
	}

	switch s.Type {
	case Enum:
		// Ensure there are enum options
		if len(s.EnumList) < 1 {
			return false
		}

		// Ensure value is a listed enum
		matchedNumber := false
		for _, val := range s.EnumList {
			if val.Value == s.Value {
				matchedNumber = true
				break
			}
		}

		if !matchedNumber {
			return false
		}
	case TextInput:
		// Ensure
	case TextDisplay:
	case Toggle:
	case Number:
	default:
		return false
	}

	return true
}

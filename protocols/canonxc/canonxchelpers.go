package canonxc

import (
	"bufio"
	"strings"
)

func extractKeyValue(root string) (string, string) {
	// Find potential := / ==
	for i := 0; i < len(root)-1; i++ {
		substr := root[i : i+2]
		if substr == ":=" || substr == "==" {
			key := root[:i]
			value := root[i+2:]

			return key, value
		}
	}
	return "", ""
}

func readResponse(response string) map[string]string {
	result := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(response))
	for scanner.Scan() {
		line := scanner.Text()

		// Handle body header junk
		if strings.HasPrefix(line, "timestamp=") {
			continue
		}
		if strings.HasPrefix(line, "realtime=") {
			continue
		}

		// Extract key value pair
		key, val := extractKeyValue(line)
		result[key] = val
	}

	return result
}

func convertRange(ALow int, AHigh int, BLow int, BHigh int, val int) int {
	A := AHigh - ALow
	B := BHigh - BLow

	if A == 0 {
		return BLow
	}

	return (((val - ALow) * B) / A) + BLow
}

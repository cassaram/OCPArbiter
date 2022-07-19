package common

type ControllerCommand struct {
	Function CameraFunction
	Value    int
	Metadata map[string]int
}

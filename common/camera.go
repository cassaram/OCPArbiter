package common

// Implements an interface for a camera

type Camera interface {
	// Initialization
	Initialize()

	// System
	InformControllerAdd(Controller)
	InformControllerRemove(Controller)
	UpdateValue(CameraCommand)
	RequestAllValues() []ControllerCommand
}

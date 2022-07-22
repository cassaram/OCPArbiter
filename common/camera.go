package common

// Implements an interface for a camera

type Camera interface {
	// Initialization
	Initialize()

	// Systems
	GetSystemSettings() SystemSettings
	SetSystemSettings(SystemSettings)

	// Controller-specific
	InformControllerAdd(Controller)
	InformControllerRemove(Controller)
	UpdateValue(CameraCommand)
	RequestAllValues() []ControllerCommand
}

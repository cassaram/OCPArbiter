package common

type ControllerCameraState uint8

const (
	CTRL_CAM_UNASSIGNED ControllerCameraState = iota
	CTRL_CAM_CONNECTING
	CTRL_CAM_CONNECTED
)

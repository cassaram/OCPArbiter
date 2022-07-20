package common

type CameraCommand struct {
	Function   CameraFunction
	Value      int
	Adjustment CameraCommandAdjustment
}

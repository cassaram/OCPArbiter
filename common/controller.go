package common

type Controller interface {
	Initialize()

	GetConnectedCamera() Camera
	SetConnectedCamera(Camera)
}

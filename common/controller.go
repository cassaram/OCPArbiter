package common

import "github.com/cassaram/ocparbiter/settings"

type Controller interface {
	Initialize([]settings.Setting)

	GetConnectedCamera() Camera
	SetConnectedCamera(Camera)

	UpdateValue(ControllerCommand)
}

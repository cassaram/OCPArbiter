package common

import "github.com/cassaram/ocparbiter/settings"

type Controller interface {
	Start()
	Stop()
	Restart()

	GetConfig() ControllerConfig
	UpdateDeviceSettings([]settings.Setting)

	GetConnectedCamera() Camera
	SetConnectedCamera(Camera)

	UpdateValue(ControllerCommand)
}

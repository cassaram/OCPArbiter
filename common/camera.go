package common

import (
	"github.com/cassaram/ocparbiter/settings"
	"github.com/google/uuid"
)

// Implements an interface for a camera

type Camera interface {
	// Handle service
	Start()
	Stop()
	Restart()

	// Configuration and Settings
	GetConfig() CameraConfig
	UpdateDeviceSettings([]settings.Setting)
	GetID() uuid.UUID

	// Controller-specific
	ControllerAdd(Controller)
	ControllerRemove(Controller)

	UpdateValue(CameraCommand)
	SendAllValues()
}

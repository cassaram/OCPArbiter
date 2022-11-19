package common

import (
	"github.com/cassaram/ocparbiter/settings"
	"github.com/google/uuid"
)

type CameraConfig struct {
	Type            string             `json:"type"`
	ID              uuid.UUID          `json:"camera_id"`
	System_Settings SystemSettings     `json:"system_settings"`
	Device_Settings []settings.Setting `json:"device_settings"`
}

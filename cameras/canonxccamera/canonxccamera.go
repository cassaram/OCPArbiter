package canonxccamera

import (
	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/protocols/canonxc"
	"github.com/cassaram/ocparbiter/settings"
	"github.com/google/uuid"
)

type CanonXCCamera struct {
	SystemSettings   common.SystemSettings
	deviceSettings   settings.SafeSettings
	id               uuid.UUID
	controllers      []common.Controller
	protocol         canonxc.CanonXCProtocol
	cache            safeMapCameraFunctionInt
	controllersQueue chan common.ControllerCommand
	cameraQueue      chan common.CameraCommand
	serviceStop      chan bool
}

// Function to start the camera interface
func (c *CanonXCCamera) Start() {
	// Ensure service is stopped
	c.Stop()

	// Check if we have an ID
	if c.id == uuid.NullUUID {
		// Ignore error
		c.id, _ = uuid.NewUUID()
	}

	// Initialize variables
	c.cache = *newSafeMapCameraFunctionInt()
	c.controllersQueue = make(chan common.ControllerCommand, 100)
	c.cameraQueue = make(chan common.CameraCommand, 100)
	c.serviceStop = make(chan bool)

	// Start services
	go c.incomingCommandService(c.serviceStop)
	go c.controllerTransmitService(c.serviceStop)
	go c.cacheUpdateService(c.serviceStop)
}

// Function to stop the camera interface
func (c *CanonXCCamera) Stop() {
	c.serviceStop <- true
}

func (c *CanonXCCamera) Restart() {
	c.Stop()
	c.Start()
}

// Configuration
// Get current camera configuration (all values needed to restart with)
func (c *CanonXCCamera) GetConfig() common.CameraConfig {
	return common.CameraConfig{
		Type:            "CanonXCCamera",
		ID:              c.id,
		System_Settings: c.SystemSettings,
		Device_Settings: c.deviceSettings.Export(),
	}
}

// Update multiple device settings
func (c *CanonXCCamera) UpdateDeviceSettings(sets []settings.Setting) {
	for _, set := range sets {
		c.deviceSettings.Change(set.Id, set.Value)
	}
}

// Get camera ID
func (c *CanonXCCamera) GetID() uuid.UUID {
	return c.id
}

// Service to send incoming commands from controllers to the protocol
func (c *CanonXCCamera) incomingCommandService(stop chan bool) {
	for {
		// Select used to stop service
		select {
		case <-stop:
			return
		default:
			// Check if there is an command, if so handle it
			select {
			case cmd, ok := <-c.cameraQueue:
				if ok {
					c.protocol.UpdateValue(cmd)
				}
			}
		}
	}
}

// Service to send outgoing controller commands to the controllers
func (c *CanonXCCamera) controllerTransmitService(stop chan bool) {
	for {
		// Select used to stop service
		select {
		case <-stop:
			return
		default:
			// Check if there is an command, if so handle it
			select {
			case cmd, ok := <-c.controllersQueue:
				if ok {
					for _, ctrl := range c.controllers {
						ctrl.UpdateValue(cmd)
					}
				}
			}
		}
	}
}

// Service to handle checking and updating controllers with changed values
func (c *CanonXCCamera) cacheUpdateService(stop chan bool) {
	for {
		// Select used to stop service
		select {
		case <-stop:
			return
		default:
			// Get current values
			current := c.protocol.GetAllValues()

			// Update cache and inform controllers
			for key, val := range current {
				cVal, cOk := c.cache.Get(key)
				if cOk {
					if cVal != val {
						// Update cache
						c.cache.Set(key, val)
						// Inform controllers of update
						c.controllersQueue <- common.ControllerCommand{
							Function: key,
							Value:    val,
						}
					}
				}
			}
		}
	}
}

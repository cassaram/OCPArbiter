package gvserialbasiccontroller

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/protocols/gvserialbasic"
	"github.com/cassaram/ocparbiter/protocols/pciprotocol"
	"github.com/cassaram/ocparbiter/settings"
)

// Used for a GV Serial basic controller such as an OCP400
type GVSerialBasicController struct {
	SystemSettings        common.SystemSettings
	deviceSettings        settings.SafeSettings
	connectedCamera       common.Camera
	cameraReturnChannel   chan common.ControllerCommand
	incomingSerialCommand chan pciprotocol.DataMessage
	outgoingSerialCommand chan pciprotocol.DataMessage
	protocol              pciprotocol.PCIProtocol
	pcigroup              uint8
	serviceTerminator     chan bool
}

func (d *GVSerialBasicController) Start() {
	d.loadSettings()

	// Start PCI Protocol
	d.protocol = pciprotocol.PCIProtocol{
		SerialPort:         d.deviceSettings.Get("serial_port"),
		DataMessageReceive: d.incomingSerialCommand,
		AppInitFunction:    d.initController,
	}
	d.protocol.Start()

	// Start sub services
	go d.cameraReturnService(d.serviceTerminator)
	go d.transmitService(d.serviceTerminator)
	go d.receiveService(d.serviceTerminator)
}

func (d *GVSerialBasicController) Stop() {
	// Stop internal services
	d.serviceTerminator <- true

	// Stop PCI services
	d.protocol.Stop()
}

func (d *GVSerialBasicController) Restart() {
	d.Stop()
	d.Start()
}

func (d *GVSerialBasicController) cameraReturnService(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			for cmd := range d.cameraReturnChannel {
				cmdByte := gvserialbasic.CommonToGVSB(cmd)
				msg := pciprotocol.DataMessage{Group: d.pcigroup, Params: cmdByte}
				d.scheduleDataMessage(msg)
			}
		}
	}
}

func (d *GVSerialBasicController) transmitService(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			for cmd := range d.outgoingSerialCommand {
				d.protocol.SendDataMessage(cmd.Group, cmd.Params)
			}
		}
	}
}

func (d *GVSerialBasicController) receiveService(stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			for cmd := range d.incomingSerialCommand {
				commonCmd := gvserialbasic.GVSBToCommon(cmd)
				d.connectedCamera.UpdateValue(commonCmd)
			}
		}
	}
}

// Configs
// Exports the config of the controller
func (d *GVSerialBasicController) GetConfig() common.ControllerConfig {
	return common.ControllerConfig{
		Type:            "GVSerialBasicController",
		Camera_ID:       d.connectedCamera.GetID(),
		System_Settings: d.SystemSettings,
		Device_Settings: d.deviceSettings.Export(),
	}
}

// Updates device settings on the controller
func (d *GVSerialBasicController) UpdateDeviceSettings(sets []settings.Setting) {
	for _, set := range sets {
		d.deviceSettings.Change(set.Id, set.Value)
	}
}

// Load all default settings
func (d *GVSerialBasicController) loadSettings() {
	// Load default settings
	fpath, _ := filepath.Abs("../controllers/gvserialbasiccontroller/DefaultSettings.json")
	file, err := os.Open(fpath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteFile, _ := io.ReadAll(file)

	var defaultSettings []settings.Setting
	json.Unmarshal(byteFile, &defaultSettings)

	for _, set := range defaultSettings {
		d.deviceSettings.Add(set)
		d.deviceSettings.Change(set.Id, set.Default) // Load default as value
	}
}

// Controller - Camera communications
// Returns the currently connected camera
func (d *GVSerialBasicController) GetConnectedCamera() common.Camera {
	return d.connectedCamera
}

// Sets the currently connected camera
func (d *GVSerialBasicController) SetConnectedCamera(cam common.Camera) {
	// Set connected camera
	d.connectedCamera = cam

	// Request current values from camera
	d.connectedCamera.SendAllValues()
}

// Controller and camera data communications
// Schedule a data message to be sent to the controller
func (d *GVSerialBasicController) scheduleDataMessage(msg pciprotocol.DataMessage) {
	d.outgoingSerialCommand <- msg
}

// Callback function for when the controller needs to be initialized with all values
func (d *GVSerialBasicController) initController(group uint8) {
	d.pcigroup = group

	d.connectedCamera.SendAllValues()
}

// Function for other devices to send a command to forward to the controller
func (d *GVSerialBasicController) UpdateValue(command common.ControllerCommand) {
	d.cameraReturnChannel <- command
}

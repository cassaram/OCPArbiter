package testcam

import "github.com/cassaram/ocparbiter/common"

type TestCam struct {
	controllers      []common.Controller
	cameraInterface  testCamCameraInterface
	cache            map[common.CameraFunction]int
	controllersQueue chan common.ControllerCommand
}

func (c *TestCam) Initialize() {
	c.cameraInterface = testCamCameraInterface{}
	c.cameraInterface.Initialize()

	// Start main service
	go c.mainLoop()
}

func (c *TestCam) InformControllerAdd(ctrl common.Controller) {
	// Ensure controller not already in list
	for _, _ctrller := range c.controllers {
		if _ctrller == ctrl {
			// Exit function and do not add another copy
			return
		}
	}

	c.controllers = append(c.controllers, ctrl)
}

func (c *TestCam) InformControllerRemove(ctrl common.Controller) {
	// Find controller in slice
	for _index, _ctrller := range c.controllers {
		// remove controller and break
		if _ctrller == ctrl {
			c.controllers = append(c.controllers[:_index], c.controllers[_index+1:]...)
			break
		}
	}
}

func (c *TestCam) UpdateValue(pkt common.CameraCommand) {
	// Get all functions
	m := c.getSupportedFunctions()

	// Call setter function
	m[pkt.Function].setter(pkt.Value)
}

func (c *TestCam) RequestAllValues() []common.ControllerCommand {
	var result []common.ControllerCommand

	for key, val := range c.cache {
		result = append(result, common.ControllerCommand{
			Function: key,
			Value:    val,
		})
	}

	return result
}

func (c *TestCam) mainLoop() {
	for {
		// Update cache
		c.updateCache()

		// Update connected controllers
		for pkt := range c.controllersQueue {
			for _, ctrl := range c.controllers {
				ctrl.UpdateValue(pkt)
			}
		}
	}
}

type camFunctionStruct struct {
	getter func() int
	setter func(int)
}

func (c *TestCam) getSupportedFunctions() map[common.CameraFunction]camFunctionStruct {
	result := map[common.CameraFunction]camFunctionStruct{
		common.CameraNumber: {
			getter: c.cameraInterface.GetCamNumber,
			setter: c.cameraInterface.SetCamNumber,
		},
		common.CallSignal: {
			getter: c.cameraInterface.GetCallSig,
			setter: c.cameraInterface.SetCallSig,
		},
		common.ColorBar: {
			getter: c.cameraInterface.GetColorBar,
			setter: c.cameraInterface.SetColorBar,
		},
		common.GainMaster: {
			getter: c.cameraInterface.GetGainMaster,
			setter: c.cameraInterface.SetGainMaster,
		},
		common.GainRed: {
			getter: c.cameraInterface.GetGainR,
			setter: c.cameraInterface.SetGainR,
		},
		common.GainGreen: {
			getter: c.cameraInterface.GetGainG,
			setter: c.cameraInterface.SetGainG,
		},
		common.BlackMaster: {
			getter: c.cameraInterface.GetBlackMaster,
			setter: c.cameraInterface.SetBlackMaster,
		},
		common.BlackRed: {
			getter: c.cameraInterface.GetBlackR,
			setter: c.cameraInterface.SetBlackR,
		},
		common.BlackGreen: {
			getter: c.cameraInterface.GetBlackG,
			setter: c.cameraInterface.SetBlackG,
		},
		common.BlackBlue: {
			getter: c.cameraInterface.GetBlackB,
			setter: c.cameraInterface.SetBlackB,
		},
	}

	return result
}

func (c *TestCam) updateCache() {
	for key, s := range c.getSupportedFunctions() {
		val := s.getter()
		if c.cache[key] != val {
			// Update value
			c.cache[key] = val
			// Queue update
			c.controllersQueue <- common.ControllerCommand{
				Function: key,
				Value:    val,
			}
		}
	}
}

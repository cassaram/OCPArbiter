package testcam

import "github.com/cassaram/ocparbiter/common"

type TestCam struct {
	controllers     []common.Controller
	cameraInterface testCamCameraInterface
	cache           map[common.CameraFunction]int
}

func (c *TestCam) Initialize() {
	c.cameraInterface = testCamCameraInterface{}
	c.cameraInterface.Initialize()
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
	// Switch on function
	switch pkt.Function {
	case common.GainMaster:
		c.cameraInterface.SetGainMaster(pkt.Value)
	case common.GainRed:
		c.cameraInterface.SetGainR(pkt.Value)
	case common.GainGreen:
		c.cameraInterface.SetGainG(pkt.Value)
	case common.GainBlue:
		c.cameraInterface.SetGainB(pkt.Value)
	}
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
	}
}

type camFunctionStruct struct {
	function common.CameraFunction
	getter   func() int
	setter   func(int)
}

func (c *TestCam) getSupportedFunctions() []camFunctionStruct {
	result := []camFunctionStruct{
		{
			function: common.CameraNumber,
			getter:   c.cameraInterface.GetCamNumber,
			setter:   c.cameraInterface.SetCamNumber,
		}, {
			function: common.CallSignal,
			getter:   c.cameraInterface.GetCallSig,
			setter:   c.cameraInterface.SetCallSig,
		}, {
			function: common.ColorBar,
			getter:   c.cameraInterface.GetColorBar,
			setter:   c.cameraInterface.SetColorBar,
		}, {
			function: common.GainMaster,
			getter:   c.cameraInterface.GetGainMaster,
			setter:   c.cameraInterface.SetGainMaster,
		}, {
			function: common.GainRed,
			getter:   c.cameraInterface.GetGainR,
			setter:   c.cameraInterface.SetGainR,
		}, {
			function: common.GainGreen,
			getter:   c.cameraInterface.GetGainG,
			setter:   c.cameraInterface.SetGainG,
		}, {
			function: common.GainBlue,
			getter:   c.cameraInterface.GetGainB,
			setter:   c.cameraInterface.SetGainB,
		},
	}

	return result
}

func (c *TestCam) updateCache() {
	for _, s := range c.getSupportedFunctions() {
		val := s.getter()
		if c.cache[s.function] != val {
			// Update value
			c.cache[s.function] = val
			// Queue update
		}
	}
}

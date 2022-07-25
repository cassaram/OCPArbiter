package testcam

import (
	"github.com/cassaram/ocparbiter/common"
)

type TestCam struct {
	systemSettings   common.SystemSettings
	controllers      []common.Controller
	cameraInterface  testCamCameraInterface
	cache            safeCameraFunctionInt
	controllersQueue chan common.ControllerCommand
	cameraQueue      chan common.CameraCommand
	cacheInitialized bool
}

func (c *TestCam) Initialize() {
	// Initialize variables
	c.cache = *newSafeCameraFunctionInt()
	c.controllersQueue = make(chan common.ControllerCommand, 20)
	c.cameraQueue = make(chan common.CameraCommand, 20)
	c.cacheInitialized = false

	c.cameraInterface = testCamCameraInterface{}
	c.cameraInterface.Initialize()

	// Start services
	go c.cacheLoop()
	go c.rxLoop()
	go c.txLoop()
}

func (c *TestCam) GetSystemSettings() common.SystemSettings {
	return c.systemSettings
}

func (c *TestCam) SetSystemSettings(s common.SystemSettings) {
	c.systemSettings = s
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
	c.cameraQueue <- pkt
}

func (c *TestCam) RequestAllValues() []common.ControllerCommand {
	var result []common.ControllerCommand

	for key, val := range c.getSupportedFunctions() {
		result = append(result, common.ControllerCommand{
			Function: key,
			Value:    val.getter(),
		})
	}

	return result
}

// Start go routine for monitoring changes and relaying to controllers
func (c *TestCam) txLoop() {
	for {
		// Update controllers
		// Ensure controller is available
		if len(c.controllers) > 0 {
			for pkt := range c.controllersQueue {
				for _, ctrl := range c.controllers {
					ctrl.UpdateValue(pkt)
				}
			}
		}
	}
}

// Go routine for handling updating the cache of camera values and propogating
func (c *TestCam) cacheLoop() {
	for {
		if !c.cacheInitialized {
			c.initCache()
		}

		c.updateCache()
	}
}

// Start go routine for handling receiving control packets from a controller
func (c *TestCam) rxLoop() {
	for {
		functions := c.getSupportedFunctions()
		for pkt := range c.cameraQueue {
			fun := functions[pkt.Function]
			val := calcAdjustedValue(fun.getter(), pkt)
			fun.setter(val)
		}
	}
}

func calcAdjustedValue(oldVal int, pkt common.CameraCommand) int {
	val := pkt.Value
	if pkt.Adjustment == common.Relative {
		val += oldVal
	}

	// Enforce constraints
	lowl, upl := common.GetCameraFunctionLimits(pkt.Function)

	if val < lowl {
		val = lowl
	} else if val > upl {
		val = upl
	}

	return val
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
		common.GainBlue: {
			getter: c.cameraInterface.GetGainB,
			setter: c.cameraInterface.SetGainB,
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
		common.FlareRed: {
			getter: c.cameraInterface.GetFlareR,
			setter: c.cameraInterface.SetFlareR,
		},
		common.FlareGreen: {
			getter: c.cameraInterface.GetFlareG,
			setter: c.cameraInterface.SetFlareG,
		},
		common.FlareBlue: {
			getter: c.cameraInterface.GetFlareB,
			setter: c.cameraInterface.SetFlareB,
		},
		common.MatrixRedGreen: {
			getter: c.cameraInterface.GetMatrixRG,
			setter: c.cameraInterface.SetMatrixRG,
		},
		common.MatrixRedBlue: {
			getter: c.cameraInterface.GetMatrixRB,
			setter: c.cameraInterface.SetMatrixRB,
		},
		common.MatrixGreenRed: {
			getter: c.cameraInterface.GetMatrixGR,
			setter: c.cameraInterface.SetMatrixGR,
		},
		common.MatrixGreenBlue: {
			getter: c.cameraInterface.GetMatrixGB,
			setter: c.cameraInterface.SetMatrixGB,
		},
		common.MatrixBlueRed: {
			getter: c.cameraInterface.GetMatrixBR,
			setter: c.cameraInterface.SetMatrixBR,
		},
		common.MatrixBlueGreen: {
			getter: c.cameraInterface.GetMatrixBG,
			setter: c.cameraInterface.SetMatrixBG,
		},
		common.MatrixGamma: {
			getter: c.cameraInterface.GetMatrixGamma,
			setter: c.cameraInterface.SetMatrixGamma,
		},
		common.Iris: {
			getter: c.cameraInterface.GetIris,
			setter: c.cameraInterface.SetIris,
		},
		common.FStop: {
			getter: c.cameraInterface.GetFStop,
			setter: c.cameraInterface.SetFStop,
		},
		common.IrisAuto: {
			getter: c.cameraInterface.GetIrisAuto,
			setter: c.cameraInterface.SetIrisAuto,
		},
		common.IrisExtended: {
			getter: c.cameraInterface.GetIrisExtended,
			setter: c.cameraInterface.SetIrisExtended,
		},
		common.KneeLevel: {
			getter: c.cameraInterface.GetKneeLvl,
			setter: c.cameraInterface.SetKneeLvl,
		},
		common.KneeDesaturationLevel: {
			getter: c.cameraInterface.GetKneeDesatLvl,
			setter: c.cameraInterface.SetKneeDesatLvl,
		},
		common.KneeSlope: {
			getter: c.cameraInterface.GetKneeSlope,
			setter: c.cameraInterface.SetKneeSlope,
		},
		common.KneeSlopeRed: {
			getter: c.cameraInterface.GetKneeSlopeR,
			setter: c.cameraInterface.SetKneeSlopeR,
		},
		common.KneeSlopeBlue: {
			getter: c.cameraInterface.GetKneeSlopeB,
			setter: c.cameraInterface.SetKneeSlopeB,
		},
		common.KneeAttack: {
			getter: c.cameraInterface.GetKneeAttack,
			setter: c.cameraInterface.SetKneeAttack,
		},
		common.KneeAttackRed: {
			getter: c.cameraInterface.GetKneeAttackR,
			setter: c.cameraInterface.SetKneeAttackR,
		},
		common.KneeAttackBlue: {
			getter: c.cameraInterface.GetKneeAttackB,
			setter: c.cameraInterface.SetKneeAttackB,
		},
		common.KneePoint: {
			getter: c.cameraInterface.GetKneePoint,
			setter: c.cameraInterface.SetKneePoint,
		},
		common.GammaMaster: {
			getter: c.cameraInterface.GetGammaLevel,
			setter: c.cameraInterface.SetGammaLevel,
		},
		common.GammaRed: {
			getter: c.cameraInterface.GetGammaR,
			setter: c.cameraInterface.SetGammaR,
		},
		common.GammaGreen: {
			getter: c.cameraInterface.GetGammaG,
			setter: c.cameraInterface.SetGammaG,
		},
		common.GammaBlue: {
			getter: c.cameraInterface.GetGammaB,
			setter: c.cameraInterface.SetGammaB,
		},
		common.WhiteBalanceRed: {
			getter: c.cameraInterface.GetWBR,
			setter: c.cameraInterface.SetWBR,
		},
		common.WhiteBalanceBlue: {
			getter: c.cameraInterface.GetWBB,
			setter: c.cameraInterface.SetWBB,
		},
	}

	return result
}

func (c *TestCam) initCache() {
	for key, s := range c.getSupportedFunctions() {
		val := s.getter()
		// Update value
		c.cache.Set(key, val)
	}

	c.cacheInitialized = true
}

func (c *TestCam) updateCache() {
	for key, s := range c.getSupportedFunctions() {
		val := s.getter()
		if cVal, _ := c.cache.Get(key); cVal != val {
			// Update value
			c.cache.Set(key, val)
			// Queue update
			c.controllersQueue <- common.ControllerCommand{
				Function: key,
				Value:    val,
			}
		}
	}
}

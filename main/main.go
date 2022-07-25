package main

import (
	"github.com/cassaram/ocparbiter/cameras/testcam"
	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/controllers/gvocp"
	"github.com/cassaram/ocparbiter/settings"
)

func main() {
	var ocp gvocp.GVOCP
	var testCam testcam.TestCam
	cam := common.Camera(&testCam)

	cam.Initialize()
	cam.UpdateValue(common.CameraCommand{
		Function:   common.CameraNumber,
		Value:      1,
		Adjustment: common.Absolute,
	})

	ocp.SetConnectedCamera(cam)
	ocp.Initialize([]settings.Setting{})
	//fmt.Println(ocp.GetSettings())

	// Dead loop
	for {

	}
}

func getControllers() {

}

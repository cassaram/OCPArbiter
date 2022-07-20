package main

import (
	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/gvocp"
	"github.com/cassaram/ocparbiter/testcam"
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

	ocp.InitOCP(cam, "COM4")

	// Dead loop
	for {

	}
}

func getControllers() {

}

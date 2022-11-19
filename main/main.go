package main

import (
	"github.com/cassaram/ocparbiter/cameras/testcam"
	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/controllers/gvserialbasiccontroller"
)

func main() {
	var ocp gvserialbasiccontroller.GVSerialBasicController
	controller := common.Controller(&ocp)
	var testCam testcam.TestCam
	cam := common.Camera(&testCam)

	cam.Initialize()
	cam.UpdateValue(common.CameraCommand{
		Function:   common.CameraNumber,
		Value:      1,
		Adjustment: common.Absolute,
	})

	cam.InformControllerAdd(controller)
	//fmt.Println(ocp.GetSettings())

	// Dead loop
	for {

	}
}

func getControllers() {

}

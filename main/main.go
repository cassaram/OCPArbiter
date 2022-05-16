package main

import (
	"github.com/cassaram/ocparbiter/common"
	"github.com/cassaram/ocparbiter/gvocp"
	"github.com/cassaram/ocparbiter/testcam"
)

func main() {
	var ocp gvocp.GVOCP
	var testCam testcam.TestCam
	cam := common.Cam(&testCam)

	cam.Initialize()
	ocp.InitOCP(cam, "COM4")
}

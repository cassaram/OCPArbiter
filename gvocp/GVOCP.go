package gvocp

import (
	"encoding/binary"
	"fmt"

	"github.com/cassaram/ocparbiter/common"
	pci "github.com/cassaram/ocparbiter/gvocp/PCI"
)

type GVOCP struct {
	connection     pci.PCI
	rxCount        uint16
	txCount        uint16
	ocpInitialized bool
	cam            common.Cam
	camFeatures    common.CamFeatureSet
}

func (ocp *GVOCP) InitOCP(camera common.Cam, port string) {
	ocp.cam = camera
	ocp.camFeatures = ocp.cam.GetFeatureSet()

	ocp.connection.SetPort(port, 1)
	ocp.connection.SetDataMessageHandler(ocp.handleDataMessage)
	ocp.connection.SetInitConnectionHandler(ocp.initializeOCP)

	ocp.rxCount = 0
	ocp.txCount = 0

	ocp.ocpInitialized = false

	fmt.Println("Loop called")
	for {
		ocp.connection.HandleData()
	}
}

func (ocp *GVOCP) handleDataMessage(s_id byte, group byte, params []byte) {
	// Increment rx packets
	ocp.rxCount++

	// Check if we have initialized the OCP
	if !ocp.ocpInitialized {
		ocp.initializeOCP(s_id, group)
	}

	command := params[0]

	fmt.Println("Received data message:", command, "Params:", params[1:])

	switch GVCommand(command) {
	case ABS_VALUE_CMD:
		switch GVCommandParam(params[1]) {
		case PCI_PANEL_RX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.connection.SendDataMessage(s_id, group, byte(ABS_VALUE_CMD), txParams)
		case PCI_PANEL_TX_MSG_NR:
			// Increment tx packets
			ocp.txCount++
			// Send data
			cnt := make([]byte, 2)
			binary.BigEndian.PutUint16(cnt, ocp.rxCount)
			txParams := append([]byte{}, byte(PCI_CAM_RX_MSG_NR))
			txParams = append(txParams, cnt[:]...)
			ocp.connection.SendDataMessage(s_id, group, byte(ABS_VALUE_CMD), txParams)
		}
	}

}

func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func intToBool(i int) bool {
	if i != 0 {
		return true
	} else {
		return false
	}
}

func (ocp *GVOCP) initializeOCP(d_id byte, group byte) {
	if ocp.camFeatures.CallSignal {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(SWITCH_CMD),
			[]byte{
				byte(CALL_SIG),
				byte(ocp.cam.GetCallSig()),
			},
		)
	}
	if ocp.camFeatures.ColorBar {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(SWITCH_CMD),
			[]byte{
				byte(COLOUR_BAR),
				byte(ocp.cam.GetColorBar()),
			},
		)
	}
	if ocp.camFeatures.GainMaster {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(VAR_MGAIN_LEVEL),
				byte(ocp.cam.GetGainMaster()),
			},
		)
	}
	if ocp.camFeatures.GainR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAIN_RED_LEVEL),
				byte(ocp.cam.GetGainR()),
			},
		)
	}
	if ocp.camFeatures.GainG {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAIN_GREEN_LEVEL),
				byte(ocp.cam.GetGainG()),
			},
		)
	}
	if ocp.camFeatures.GainB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAIN_BLUE_LEVEL),
				byte(ocp.cam.GetGainB()),
			},
		)
	}
	if ocp.camFeatures.BlackR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(BLACK_RED_LEVEL),
				byte(ocp.cam.GetBlackR()),
			},
		)
	}
	if ocp.camFeatures.BlackG {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(BLACK_GREEN_LEVEL),
				byte(ocp.cam.GetBlackG()),
			},
		)
	}
	if ocp.camFeatures.BlackB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(BLACK_BLUE_LEVEL),
				byte(ocp.cam.GetBlackB()),
			},
		)
	}
	if ocp.camFeatures.BlackMaster {
		ebitLvl := make([]byte, 2)
		binary.BigEndian.PutUint16(ebitLvl, uint16(ocp.cam.GetBlackMaster()))
		ebitLvl[0] &= 0x0F
		txParams := append([]byte{}, byte(MBLACK_12BIT_LEVEL))
		txParams = append(txParams, ebitLvl[:]...)

		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			txParams,
		)
	}
	if ocp.camFeatures.FlareR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(FLARE_RED_LEVEL),
				byte(ocp.cam.GetFlareR()),
			},
		)
	}
	if ocp.camFeatures.FlareG {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(FLARE_GREEN_LEVEL),
				byte(ocp.cam.GetFlareG()),
			},
		)
	}
	if ocp.camFeatures.FlareB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(FLARE_BLUE_LEVEL),
				byte(ocp.cam.GetFlareB()),
			},
		)
	}
	if ocp.camFeatures.Iris {
		ebitLvl := make([]byte, 2)
		binary.BigEndian.PutUint16(ebitLvl, uint16(ocp.cam.GetIris()))
		ebitLvl[0] &= 0x0F
		txParams := append([]byte{}, byte(IRIS_12BIT_LEVEL))
		txParams = append(txParams, ebitLvl[:]...)

		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			txParams,
		)
	}
	if ocp.camFeatures.IrisAuto {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(SWITCH_CMD),
			[]byte{
				byte(AUTO_IRIS),
				byte(ocp.cam.GetIrisAuto()),
			},
		)
	}
	if ocp.camFeatures.GammaLevel {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(MASTER_GAMMA_LEVEL),
				byte(ocp.cam.GetGammaLevel()),
			},
		)
	}
	if ocp.camFeatures.GammaR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAMMA_RED_LEVEL),
				byte(ocp.cam.GetGammaR()),
			},
		)
	}
	if ocp.camFeatures.GammaG {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAMMA_GREEN_LEVEL),
				byte(ocp.cam.GetGammaG()),
			},
		)
	}
	if ocp.camFeatures.GammaB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(GAMMA_BLUE_LEVEL),
				byte(ocp.cam.GetGammaB()),
			},
		)
	}
	if ocp.camFeatures.WBR {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(WH_BAL_RED_LEVEL),
				byte(ocp.cam.GetWBR()),
			},
		)
	}
	if ocp.camFeatures.WBB {
		ocp.connection.SendDataMessage(
			d_id,
			group,
			byte(ABS_VALUE_CMD),
			[]byte{
				byte(WH_BAL_BLUE_LEVEL),
				byte(ocp.cam.GetWBB()),
			},
		)
	}

	ocp.ocpInitialized = true
}

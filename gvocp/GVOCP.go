package gvocp

import (
	"encoding/binary"
	"fmt"

	PCI "github.com/cassaram/ocparbiter/gvocp/PCI"
)

type GVOCP struct {
	connection     PCI.PCI
	rxCount        uint16
	txCount        uint16
	ocpInitialized bool
}

func (ocp *GVOCP) InitOCP() {
	ocp.connection.SetPort("COM2", 1)
	ocp.connection.SetDataMessageHandler(ocp.handleDataMessage)

	ocp.rxCount = 0
	ocp.txCount = 0

	ocp.ocpInitialized = false

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

			// DEBUG TEST
			ocp.connection.SendDataMessage(s_id, group, byte(ABS_VALUE_CMD), []byte{byte(GAIN_RED_LEVEL), 0})
		}
	}

}

func (ocp *GVOCP) initializeOCP(d_id byte, group byte) {
	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(SERIAL_CAMERA_NUMBER), 6})
	ocp.connection.SendDataMessage(d_id, group, byte(ABS_SWITCH_CMD), []byte{byte(COLOUR_BAR), 1})

	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(GAIN_RED_LEVEL), 0})
	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(GAIN_GREEN_LEVEL), 0})
	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(GAIN_BLUE_LEVEL), 0})

	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(BLACK_RED_LEVEL), 0})
	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(BLACK_GREEN_LEVEL), 0})
	ocp.connection.SendDataMessage(d_id, group, byte(ABS_VALUE_CMD), []byte{byte(BLACK_BLUE_LEVEL), 0})

	ocp.ocpInitialized = true
}

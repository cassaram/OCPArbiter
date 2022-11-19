package pciprotocol

// Defines various enums and types used by the pci service

// Enum for PCI command types
type PCI_CMD uint8

const (
	DATA_MESSAGE        PCI_CMD = 0
	TEST_RING           PCI_CMD = 1
	ASSIGN_TO_GROUP     PCI_CMD = 2
	CONNECTED_TO_PART   PCI_CMD = 3
	DEASSIGN_FROM_GROUP PCI_CMD = 4
	MULTICAST_MODE      PCI_CMD = 5
	MULTIPLEXED         PCI_CMD = 7
)

// Enum for PCI connection states
type PCI_CONNECTION_STATE uint8

const (
	PCI_START          PCI_CONNECTION_STATE = 1
	PCI_WAIT_TEST_RING PCI_CONNECTION_STATE = 2
	PCI_WAIT_ASSIGN    PCI_CONNECTION_STATE = 4
	PCI_IDENTIFIED     PCI_CONNECTION_STATE = 5
	PCI_CONNECTED      PCI_CONNECTION_STATE = 6
)

// Struct for PCI data messages to be passed into calling module
type DataMessage struct {
	Group  uint8
	Params []byte
}

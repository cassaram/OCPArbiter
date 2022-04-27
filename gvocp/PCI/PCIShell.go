package pci

type GVCommand int

const (
	ABS_VALUE_CMD      GVCommand = 1
	VALUE_CMD          GVCommand = 2
	SWITCH_CMD         GVCommand = 3
	MODE_CMD           GVCommand = 4
	CUSTOM_VALUE_CMD   GVCommand = 6
	SIGNAL_CMD         GVCommand = 7
	CUSTOM_SWITCH_CMD  GVCommand = 8
	CUSTOM_MODE_CMD    GVCommand = 9
	PROMCODE_CMD       GVCommand = 10
	ABS_SWITCH_CMD     GVCommand = 11
	COMM_CONTROL       GVCommand = 14
	INIT_CONTROL       GVCommand = 15
	SCENE_FILE_CMD     GVCommand = 16
	STATUS_REQUEST_CMD GVCommand = 17
	C_VAL_REQUEST_CMD  GVCommand = 18
	RECALL_FUNC_CMD    GVCommand = 19
	X100_INTERNAL_CMD  GVCommand = 20
	X100_INTERNAL_CMD1 GVCommand = 21
	X100_INTERNAL_CMD2 GVCommand = 22
	X100_INTERNAL_CMD3 GVCommand = 23
	MODEVALREQ_CMD     GVCommand = 24
	MODEVALTXT_CMD     GVCommand = 25
	WLTP_CMD           GVCommand = 26
)

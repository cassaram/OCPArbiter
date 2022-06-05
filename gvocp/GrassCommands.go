package gvocp

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

type GVCommandParam int

const (
	GAIN_RED_LEVEL         GVCommandParam = 1
	GAIN_GREEN_LEVEL       GVCommandParam = 2
	GAIN_BLUE_LEVEL        GVCommandParam = 3
	BLACK_RED_LEVEL        GVCommandParam = 4
	BLACK_GREEN_LEVEL      GVCommandParam = 5
	BLACK_BLUE_LEVEL       GVCommandParam = 6
	FLARE_RED_LEVEL        GVCommandParam = 7
	FLARE_GREEN_LEVEL      GVCommandParam = 8
	FLARE_BLUE_LEVEL       GVCommandParam = 9
	NOTCH_LEVEL            GVCommandParam = 10
	SOFT_CONT_LEVEL        GVCommandParam = 11
	SKIN_CONT_LEVEL        GVCommandParam = 12
	SKIN1_WIDTH_RED        GVCommandParam = 13
	SKIN1_WIDTH_BLUE       GVCommandParam = 14
	SKIN1_COLOR_RED        GVCommandParam = 15
	SKIN1_COLOR_BLUE       GVCommandParam = 16
	SKIN2_WIDTH_RED        GVCommandParam = 17
	SKIN2_WIDTH_BLUE       GVCommandParam = 18
	SKIN2_COLOR_RED        GVCommandParam = 19
	SKIN2_COLOR_BLUE       GVCommandParam = 20
	MATRIX_RG              GVCommandParam = 21
	MATRIX_RB              GVCommandParam = 22
	MATRIX_GR              GVCommandParam = 23
	MATRIX_GB              GVCommandParam = 24
	MATRIX_BR              GVCommandParam = 25
	MATRIX_BG              GVCommandParam = 26
	MBLACK_12BIT_LEVEL     GVCommandParam = 27
	MASTER_BLACK_LEVEL     GVCommandParam = 28
	KNEE_LEVEL             GVCommandParam = 29
	IRIS_LEVEL             GVCommandParam = 30
	IRIS_12BIT_LEVEL       GVCommandParam = 31
	KNEE_DESAT_LEVEL       GVCommandParam = 32
	LAMP_DIM_LEVEL         GVCommandParam = 33
	CONTOUR_LEVEL          GVCommandParam = 34
	IRIS_PA_LEVEL          GVCommandParam = 35
	MBLACK_LOW_LEVEL       GVCommandParam = 36
	ZOOM_CONTROL_LEVEL     GVCommandParam = 37
	FOCUS_CONTROL_LEVEL    GVCommandParam = 40
	SHAD_DC_RED_LEVEL      GVCommandParam = 41
	SHAD_DC_GREEN_LEVEL    GVCommandParam = 42
	SHAD_DC_BLUE_LEVEL     GVCommandParam = 43
	MATRIX_RG_VAR_2        GVCommandParam = 44
	MATRIX_RB_VAR_2        GVCommandParam = 45
	MATRIX_GR_VAR_2        GVCommandParam = 46
	MATRIX_GB_VAR_2        GVCommandParam = 47
	MATRIX_BR_VAR_2        GVCommandParam = 48
	MATRIX_BG_VAR_2        GVCommandParam = 49
	FL_GAIN_R_LEVEL        GVCommandParam = 50
	FL_GAIN_B_LEVEL        GVCommandParam = 51
	AFS_LEVEL              GVCommandParam = 52
	CONTOUR_BLACK_LEVEL    GVCommandParam = 53
	CONTOUR_LOWMID_LEVEL   GVCommandParam = 54
	CONTOUR_MID_LEVEL      GVCommandParam = 55
	CONTOUR_WHITE_LEVEL    GVCommandParam = 56
	CONTOUR_BLACK_POS      GVCommandParam = 57
	CONTOUR_LOWMID_POS     GVCommandParam = 58
	CONTOUR_MID_POS        GVCommandParam = 59
	CONTOUR_WHITE_POS      GVCommandParam = 60
	DIAG_MIX_CONTOUR_LEVEL GVCommandParam = 61
	SHUTTER_ANGLE_LEVEL    GVCommandParam = 62
	C2IP_CLIENTS           GVCommandParam = 63
	LDK_CONNECT_CLIENT     GVCommandParam = 64
	GAMMA_RED_LEVEL        GVCommandParam = 71
	MASTER_GAMMA_LEVEL     GVCommandParam = 72
	GAMMA_BLUE_LEVEL       GVCommandParam = 73
	GAMMA_GREEN_LEVEL      GVCommandParam = 74
	LDK6000_SDTV_C_LEVEL   GVCommandParam = 75
	LDK6000_SDTV_SC_LEVEL  GVCommandParam = 76
	LDK6000_SDTV_IE_LEVEL  GVCommandParam = 77
	LDK6000_SDTV_VC_LEVEL  GVCommandParam = 78
	LDK6000_SDTV_NS_LEVEL  GVCommandParam = 79
	LDK6000_SDTV_LD_LEVEL  GVCommandParam = 80
	GRADIENT_CENTRE        GVCommandParam = 81
	GRADIENT_DEPTH_R       GVCommandParam = 82
	GRADIENT_DEPTH_G       GVCommandParam = 83
	GRADIENT_DEPTH_B       GVCommandParam = 84
	SOFTFCS_RADIUS         GVCommandParam = 85
	SOFTFCS_DC_LEVEL       GVCommandParam = 86
	SOFTFCS_TRANSIT_LEVEL  GVCommandParam = 87
	SOFTFCS_DC_FADE        GVCommandParam = 88
	SOFTFCS_X_POS          GVCommandParam = 89
	SOFTFCS_Y_POS          GVCommandParam = 90
	SOFTFCS_ASP_RATIO      GVCommandParam = 91
	MONOTONE_RED           GVCommandParam = 92
	MONOTONE_BLUE          GVCommandParam = 93
	MONOTONE_DEPTH         GVCommandParam = 94
	VAR_MGAIN_LEVEL        GVCommandParam = 95
	VAR_CTEMP_LEVEL        GVCommandParam = 96
	SERIAL_CAMERA_NUMBER   GVCommandParam = 97
	BS_CAMERA_NUMBER       GVCommandParam = 98
	BLACKSTRETCH_LEVEL     GVCommandParam = 100
	H_PHASE_LEVEL          GVCommandParam = 101
	SUBC_FINE_LEVEL        GVCommandParam = 102
	SATURATION_LEVEL       GVCommandParam = 103
	HSAW_RED_LEVEL         GVCommandParam = 107
	HSAW_GREEN_LEVEL       GVCommandParam = 108
	HSAW_BLUE_LEVEL        GVCommandParam = 109
	HPAR_RED_LEVEL         GVCommandParam = 110
	HPAR_GREEN_LEVEL       GVCommandParam = 111
	HPAR_BLUE_LEVEL        GVCommandParam = 112
	VSAW_RED_LEVEL         GVCommandParam = 113
	VSAW_GREEN_LEVEL       GVCommandParam = 114
	VSAW_BLUE_LEVEL        GVCommandParam = 115
	VPAR_RED_LEVEL         GVCommandParam = 116
	VPAR_GREEN_LEVEL       GVCommandParam = 117
	VPAR_BLUE_LEVEL        GVCommandParam = 118
	GP_ANALOG0_LEVEL       GVCommandParam = 119
	GP_ANALOG1_LEVEL       GVCommandParam = 120
	ZOOM_FOLLOW_LEVEL      GVCommandParam = 121
	FOCUS_FOLLOW_LEVEL     GVCommandParam = 122
	VOLUME_PROG_LEVEL      GVCommandParam = 123
	VOLUME_ENG_LEVEL       GVCommandParam = 124
	AW1_GAIN_R_LEVEL       GVCommandParam = 125
	MA_GAIN_R1_LEVEL       GVCommandParam = 125
	AW1_GAIN_B_LEVEL       GVCommandParam = 126
	MA_GAIN_B1_LEVEL       GVCommandParam = 126
	AW2_GAIN_R_LEVEL       GVCommandParam = 127
	MA_GAIN_R2_LEVEL       GVCommandParam = 127
	AW2_GAIN_B_LEVEL       GVCommandParam = 128
	MA_GAIN_B2_LEVEL       GVCommandParam = 128
	CONT_HVBAL_LEVEL       GVCommandParam = 130
	CONT_BWBAL_LEVEL       GVCommandParam = 131
	CONT_INEDGE_LEVEL      GVCommandParam = 132
	CONT_NOISE_SLICER      GVCommandParam = 133
	CONT_LEVDEP_LEVEL      GVCommandParam = 134
	CONT_INBAND_LEVEL      GVCommandParam = 135
	CONT_EDGEBAND_LEVEL    GVCommandParam = 136
	CONT_VERTCONT_LEVEL    GVCommandParam = 137
	KNEE_SLOPE_M           GVCommandParam = 29
	KNEE_SLOPE_R           GVCommandParam = 140
	KNEE_SLOPE_B           GVCommandParam = 141
	KNEE_ATTACK_M          GVCommandParam = 142
	KNEE_ATTACK_R          GVCommandParam = 143
	KNEE_ATTACK_B          GVCommandParam = 144
	WHITE_LIMIT_M          GVCommandParam = 145
	WHITE_LIMIT_R          GVCommandParam = 146
	WHITE_LIMIT_G          GVCommandParam = 147
	WHITE_LIMIT_B          GVCommandParam = 148
	EXPOSURE_LEVEL         GVCommandParam = 149
	EXP_INDICATION         GVCommandParam = 150
	EXP_INDICATION_HS      GVCommandParam = 151
	EXPOSURE_LEVEL_2       GVCommandParam = 152
	COLOUR_FILTER          GVCommandParam = 153
	PCI_PANEL_RX_MSG_NR    GVCommandParam = 154
	PCI_PANEL_TX_MSG_NR    GVCommandParam = 155
	PCI_CAM_RX_MSG_NR      GVCommandParam = 156
	PCI_CAM_TX_MSG_NR      GVCommandParam = 157
	PCI_COBS_TX_COUNT      GVCommandParam = 158
	PCI_COBS_RX_COUNT      GVCommandParam = 159
	PCI_COBS_ERR_COUNT     GVCommandParam = 160
	PCI_DTCP_ERR_COUNT     GVCommandParam = 161
	BATTERY_LEVEL          GVCommandParam = 162
	COLCORR_SET_4_COLOR    GVCommandParam = 185
	COLCORR_SET_4_WIDTH    GVCommandParam = 186
	COLCORR_SET_4_HUE      GVCommandParam = 187
	COLCORR_SET_4_SAT      GVCommandParam = 188
	COLCORR_SET_4_LUM      GVCommandParam = 189
	COLCORR_SET_5_COLOR    GVCommandParam = 190
	COLCORR_SET_5_WIDTH    GVCommandParam = 191
	COLCORR_SET_5_HUE      GVCommandParam = 192
	COLCORR_SET_5_SAT      GVCommandParam = 193
	COLCORR_SET_5_LUM      GVCommandParam = 194
	COLCORR_SET_6_COLOR    GVCommandParam = 195
	COLCORR_SET_6_WIDTH    GVCommandParam = 196
	COLCORR_SET_6_HUE      GVCommandParam = 197
	COLCORR_SET_6_SAT      GVCommandParam = 198
	COLCORR_SET_6_LUM      GVCommandParam = 199
	KNEE_POINT_LEVEL       GVCommandParam = 200
	WH_BAL_RED_LEVEL       GVCommandParam = 201
	WH_BAL_BLUE_LEVEL      GVCommandParam = 202
	HSAW_RED_LEVEL2        GVCommandParam = 207
	HSAW_GREEN_LEVEL2      GVCommandParam = 208
	HSAW_BLUE_LEVEL2       GVCommandParam = 209
	HPAR_RED_LEVEL2        GVCommandParam = 210
	HPAR_GREEN_LEVEL2      GVCommandParam = 211
	HPAR_BLUE_LEVEL2       GVCommandParam = 212
	VSAW_RED_LEVEL2        GVCommandParam = 213
	VSAW_GREEN_LEVEL2      GVCommandParam = 214
	VSAW_BLUE_LEVEL2       GVCommandParam = 215
	VPAR_RED_LEVEL2        GVCommandParam = 216
	VPAR_GREEN_LEVEL2      GVCommandParam = 217
	VPAR_BLUE_LEVEL2       GVCommandParam = 218
	RE_RED_SAW             GVCommandParam = 219
	RE_GREEN_SAW           GVCommandParam = 220
	RE_BLUE_SAW            GVCommandParam = 221
	VF_CONTOUR_LEVEL       GVCommandParam = 222
	ZEBRA_LEV              GVCommandParam = 223
	ZEBRA_RANGE            GVCommandParam = 224
	V_APERTURE_LEVEL       GVCommandParam = 225
	H_APERTURE_LEVEL       GVCommandParam = 226
	CAM_TEMP_LEVEL         GVCommandParam = 227
	CAM_CRITICAL_TEMP      GVCommandParam = 228
	COLCORR_COLOR          GVCommandParam = 229
	COLCORR_WIDTH          GVCommandParam = 230
	COLCORR_HUE            GVCommandParam = 231
	COLCORR_SAT            GVCommandParam = 232
	COLCORR_LUM            GVCommandParam = 233
	COLCORR_SETS_ON        GVCommandParam = 234
	COLCORR_SET_1_COLOR    GVCommandParam = 235
	COLCORR_SET_1_WIDTH    GVCommandParam = 236
	COLCORR_SET_1_HUE      GVCommandParam = 237
	COLCORR_SET_1_SAT      GVCommandParam = 238
	COLCORR_SET_1_LUM      GVCommandParam = 239
	COLCORR_SET_2_COLOR    GVCommandParam = 240
	COLCORR_SET_2_WIDTH    GVCommandParam = 241
	COLCORR_SET_2_HUE      GVCommandParam = 242
	COLCORR_SET_2_SAT      GVCommandParam = 243
	COLCORR_SET_2_LUM      GVCommandParam = 244
	COLCORR_SET_3_COLOR    GVCommandParam = 245
	COLCORR_SET_3_WIDTH    GVCommandParam = 246
	COLCORR_SET_3_HUE      GVCommandParam = 247
	COLCORR_SET_3_SAT      GVCommandParam = 248
	COLCORR_SET_3_LUM      GVCommandParam = 249
	SKIN1_CONT_LEVEL       GVCommandParam = 250
	SKIN2_CONT_LEVEL       GVCommandParam = 251
)

type GVSwitchParams uint8

const (
	NOTCH                         GVSwitchParams = 1
	STANDARD                      GVSwitchParams = 2
	REM_AUDIO_LEVEL               GVSwitchParams = 3
	SOFT_CONTOUR                  GVSwitchParams = 4
	OCP_LOCK                      GVSwitchParams = 5
	OCP_CONNECTED                 GVSwitchParams = 6
	SKIN_VIEW                     GVSwitchParams = 7
	AUTOSKIN                      GVSwitchParams = 8
	MCP_AVAILABLE                 GVSwitchParams = 9
	ASPECT_RATIO                  GVSwitchParams = 10
	SCAN_REVERSE                  GVSwitchParams = 11
	LAMP_OFF                      GVSwitchParams = 12
	REM_ASPECT_RATIO              GVSwitchParams = 13
	MATRIX_GAMMA                  GVSwitchParams = 14
	KNEE_DESAT                    GVSwitchParams = 15
	AUTO_IRIS                     GVSwitchParams = 16
	CALL_SIG                      GVSwitchParams = 17
	BLACKSTRETCH_TYPE             GVSwitchParams = 18
	CAMERA_DISABLE                GVSwitchParams = 19
	STANDBY                       GVSwitchParams = 20
	VF_STANDBY                    GVSwitchParams = 21
	STUDIO_MODE                   GVSwitchParams = 22
	SWITCHABLE_VIDEO_TRANSMISSION GVSwitchParams = 23
	SWITCHABLE_PRIVATE_DATA       GVSwitchParams = 24
	SWITCHABLE_AUX_VIDEO          GVSwitchParams = 25
	SWITCHABLE_INTERCOM           GVSwitchParams = 26
	COLOUR_BAR                    GVSwitchParams = 27
)

type GVModeParams uint8

const (
	SKIN_CONTOUR       GVModeParams = 1
	GAIN_SELECT        GVModeParams = 2
	KNEE_SELECT        GVModeParams = 3
	GAMMA_SELECT       GVModeParams = 4
	SYSTEMID_SELECT    GVModeParams = 5
	FILTER_SELECT      GVModeParams = 6
	AUDIO_FILTER       GVModeParams = 7
	LENS_TYPE          GVModeParams = 8
	MONITORING_SELECT  GVModeParams = 9
	USER_LEVEL         GVModeParams = 10
	CRTSCAN_SELECT     GVModeParams = 11
	GAMMA_CURVE_SELECT GVModeParams = 12
	KNEE_SOURCE_SELECT GVModeParams = 13
	FLARE_SELECT       GVModeParams = 14
	FSTOP_SELECT       GVModeParams = 15
)

type GVModeFStop uint8

const (
	FCT_FSTOP_010   GVModeFStop = 1
	FCT_FSTOP_012   GVModeFStop = 2
	FCT_FSTOP_013   GVModeFStop = 3
	FCT_FSTOP_014   GVModeFStop = 5
	FCT_FSTOP_015   GVModeFStop = 6
	FCT_FSTOP_017   GVModeFStop = 8
	FCT_FSTOP_018   GVModeFStop = 9
	FCT_FSTOP_020   GVModeFStop = 11
	FCT_FSTOP_022   GVModeFStop = 12
	FCT_FSTOP_024   GVModeFStop = 43
	FCT_FSTOP_026   GVModeFStop = 13
	FCT_FSTOP_028   GVModeFStop = 14
	FCT_FSTOP_031   GVModeFStop = 15
	FCT_FSTOP_034   GVModeFStop = 44
	FCT_FSTOP_037   GVModeFStop = 16
	FCT_FSTOP_040   GVModeFStop = 17
	FCT_FSTOP_044   GVModeFStop = 18
	FCT_FSTOP_048   GVModeFStop = 45
	FCT_FSTOP_052   GVModeFStop = 19
	FCT_FSTOP_056   GVModeFStop = 20
	FCT_FSTOP_062   GVModeFStop = 21
	FCT_FSTOP_067   GVModeFStop = 46
	FCT_FSTOP_073   GVModeFStop = 22
	FCT_FSTOP_080   GVModeFStop = 23
	FCT_FSTOP_087   GVModeFStop = 24
	FCT_FSTOP_095   GVModeFStop = 47
	FCT_FSTOP_100   GVModeFStop = 25
	FCT_FSTOP_110   GVModeFStop = 26
	FCT_FSTOP_120   GVModeFStop = 27
	FCT_FSTOP_130   GVModeFStop = 48
	FCT_FSTOP_150   GVModeFStop = 28
	FCT_FSTOP_160   GVModeFStop = 29
	FCT_FSTOP_170   GVModeFStop = 49
	FCT_FSTOP_190   GVModeFStop = 50
	FCT_FSTOP_210   GVModeFStop = 30
	FCT_FSTOP_220   GVModeFStop = 51
	FCT_FSTOP_250   GVModeFStop = 31
	FCT_FSTOP_270   GVModeFStop = 52
	FCT_FSTOP_290   GVModeFStop = 53
	FCT_FSTOP_320   GVModeFStop = 32
	FCT_FSTOP_350   GVModeFStop = 33
	FCT_FSTOP_CLOSE GVModeFStop = 34
	FCT_FSTOP_380   GVModeFStop = 36
	FCT_FSTOP_420   GVModeFStop = 37
	FCT_FSTOP_450   GVModeFStop = 38
	FCT_FSTOP_490   GVModeFStop = 39
	FCT_FSTOP_530   GVModeFStop = 40
	FCT_FSTOP_590   GVModeFStop = 41
	FCT_FSTOP_640   GVModeFStop = 42
	FCT_FSTOP_16M   GVModeFStop = 72
	FCT_FSTOP_22M   GVModeFStop = 73
	FCT_FSTOP_27M   GVModeFStop = 74
	FCT_FSTOP_32M   GVModeFStop = 75
)

type GVModeIrisAuto uint8

const (
	AUTOIRIS_OFF GVModeIrisAuto = 0
	AUTOIRIS_ON  GVModeIrisAuto = 1
)

type GVModeCallSignal uint8

const (
	CALLSIG_OFF GVModeCallSignal = 0
	CALLSIG_ON  GVModeCallSignal = 1
)

type GVModeColorBar uint8

const (
	COLORBAR_OFF GVModeColorBar = 0
	COLORBAR_ON  GVModeColorBar = 1
)

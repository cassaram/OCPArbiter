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

type GVValueCommand int

const (
	GAIN_RED_LEVEL         GVValueCommand = 1
	GAIN_GREEN_LEVEL       GVValueCommand = 2
	GAIN_BLUE_LEVEL        GVValueCommand = 3
	BLACK_RED_LEVEL        GVValueCommand = 4
	BLACK_GREEN_LEVEL      GVValueCommand = 5
	BLACK_BLUE_LEVEL       GVValueCommand = 6
	FLARE_RED_LEVEL        GVValueCommand = 7
	FLARE_GREEN_LEVEL      GVValueCommand = 8
	FLARE_BLUE_LEVEL       GVValueCommand = 9
	NOTCH_LEVEL            GVValueCommand = 10
	SOFT_CONT_LEVEL        GVValueCommand = 11
	SKIN_CONT_LEVEL        GVValueCommand = 12
	SKIN1_WIDTH_RED        GVValueCommand = 13
	SKIN1_WIDTH_BLUE       GVValueCommand = 14
	SKIN1_COLOR_RED        GVValueCommand = 15
	SKIN1_COLOR_BLUE       GVValueCommand = 16
	SKIN2_WIDTH_RED        GVValueCommand = 17
	SKIN2_WIDTH_BLUE       GVValueCommand = 18
	SKIN2_COLOR_RED        GVValueCommand = 19
	SKIN2_COLOR_BLUE       GVValueCommand = 20
	MATRIX_RG              GVValueCommand = 21
	MATRIX_RB              GVValueCommand = 22
	MATRIX_GR              GVValueCommand = 23
	MATRIX_GB              GVValueCommand = 24
	MATRIX_BR              GVValueCommand = 25
	MATRIX_BG              GVValueCommand = 26
	MBLACK_12BIT_LEVEL     GVValueCommand = 27
	MASTER_BLACK_LEVEL     GVValueCommand = 28
	KNEE_LEVEL             GVValueCommand = 29
	IRIS_LEVEL             GVValueCommand = 30
	IRIS_12BIT_LEVEL       GVValueCommand = 31
	KNEE_DESAT_LEVEL       GVValueCommand = 32
	LAMP_DIM_LEVEL         GVValueCommand = 33
	CONTOUR_LEVEL          GVValueCommand = 34
	IRIS_PA_LEVEL          GVValueCommand = 35
	MBLACK_LOW_LEVEL       GVValueCommand = 36
	ZOOM_CONTROL_LEVEL     GVValueCommand = 37
	FOCUS_CONTROL_LEVEL    GVValueCommand = 40
	SHAD_DC_RED_LEVEL      GVValueCommand = 41
	SHAD_DC_GREEN_LEVEL    GVValueCommand = 42
	SHAD_DC_BLUE_LEVEL     GVValueCommand = 43
	MATRIX_RG_VAR_2        GVValueCommand = 44
	MATRIX_RB_VAR_2        GVValueCommand = 45
	MATRIX_GR_VAR_2        GVValueCommand = 46
	MATRIX_GB_VAR_2        GVValueCommand = 47
	MATRIX_BR_VAR_2        GVValueCommand = 48
	MATRIX_BG_VAR_2        GVValueCommand = 49
	FL_GAIN_R_LEVEL        GVValueCommand = 50
	FL_GAIN_B_LEVEL        GVValueCommand = 51
	AFS_LEVEL              GVValueCommand = 52
	CONTOUR_BLACK_LEVEL    GVValueCommand = 53
	CONTOUR_LOWMID_LEVEL   GVValueCommand = 54
	CONTOUR_MID_LEVEL      GVValueCommand = 55
	CONTOUR_WHITE_LEVEL    GVValueCommand = 56
	CONTOUR_BLACK_POS      GVValueCommand = 57
	CONTOUR_LOWMID_POS     GVValueCommand = 58
	CONTOUR_MID_POS        GVValueCommand = 59
	CONTOUR_WHITE_POS      GVValueCommand = 60
	DIAG_MIX_CONTOUR_LEVEL GVValueCommand = 61
	SHUTTER_ANGLE_LEVEL    GVValueCommand = 62
	C2IP_CLIENTS           GVValueCommand = 63
	LDK_CONNECT_CLIENT     GVValueCommand = 64
	GAMMA_RED_LEVEL        GVValueCommand = 71
	MASTER_GAMMA_LEVEL     GVValueCommand = 72
	GAMMA_BLUE_LEVEL       GVValueCommand = 73
	GAMMA_GREEN_LEVEL      GVValueCommand = 74
	LDK6000_SDTV_C_LEVEL   GVValueCommand = 75
	LDK6000_SDTV_SC_LEVEL  GVValueCommand = 76
	LDK6000_SDTV_IE_LEVEL  GVValueCommand = 77
	LDK6000_SDTV_VC_LEVEL  GVValueCommand = 78
	LDK6000_SDTV_NS_LEVEL  GVValueCommand = 79
	LDK6000_SDTV_LD_LEVEL  GVValueCommand = 80
	GRADIENT_CENTRE        GVValueCommand = 81
	GRADIENT_DEPTH_R       GVValueCommand = 82
	GRADIENT_DEPTH_G       GVValueCommand = 83
	GRADIENT_DEPTH_B       GVValueCommand = 84
	SOFTFCS_RADIUS         GVValueCommand = 85
	SOFTFCS_DC_LEVEL       GVValueCommand = 86
	SOFTFCS_TRANSIT_LEVEL  GVValueCommand = 87
	SOFTFCS_DC_FADE        GVValueCommand = 88
	SOFTFCS_X_POS          GVValueCommand = 89
	SOFTFCS_Y_POS          GVValueCommand = 90
	SOFTFCS_ASP_RATIO      GVValueCommand = 91
	MONOTONE_RED           GVValueCommand = 92
	MONOTONE_BLUE          GVValueCommand = 93
	MONOTONE_DEPTH         GVValueCommand = 94
	VAR_MGAIN_LEVEL        GVValueCommand = 95
	VAR_CTEMP_LEVEL        GVValueCommand = 96
	SERIAL_CAMERA_NUMBER   GVValueCommand = 97
	BS_CAMERA_NUMBER       GVValueCommand = 98
	BLACKSTRETCH_LEVEL     GVValueCommand = 100
	H_PHASE_LEVEL          GVValueCommand = 101
	SUBC_FINE_LEVEL        GVValueCommand = 102
	SATURATION_LEVEL       GVValueCommand = 103
	HSAW_RED_LEVEL         GVValueCommand = 107
	HSAW_GREEN_LEVEL       GVValueCommand = 108
	HSAW_BLUE_LEVEL        GVValueCommand = 109
	HPAR_RED_LEVEL         GVValueCommand = 110
	HPAR_GREEN_LEVEL       GVValueCommand = 111
	HPAR_BLUE_LEVEL        GVValueCommand = 112
	VSAW_RED_LEVEL         GVValueCommand = 113
	VSAW_GREEN_LEVEL       GVValueCommand = 114
	VSAW_BLUE_LEVEL        GVValueCommand = 115
	VPAR_RED_LEVEL         GVValueCommand = 116
	VPAR_GREEN_LEVEL       GVValueCommand = 117
	VPAR_BLUE_LEVEL        GVValueCommand = 118
	GP_ANALOG0_LEVEL       GVValueCommand = 119
	GP_ANALOG1_LEVEL       GVValueCommand = 120
	ZOOM_FOLLOW_LEVEL      GVValueCommand = 121
	FOCUS_FOLLOW_LEVEL     GVValueCommand = 122
	VOLUME_PROG_LEVEL      GVValueCommand = 123
	VOLUME_ENG_LEVEL       GVValueCommand = 124
	AW1_GAIN_R_LEVEL       GVValueCommand = 125
	MA_GAIN_R1_LEVEL       GVValueCommand = 125
	AW1_GAIN_B_LEVEL       GVValueCommand = 126
	MA_GAIN_B1_LEVEL       GVValueCommand = 126
	AW2_GAIN_R_LEVEL       GVValueCommand = 127
	MA_GAIN_R2_LEVEL       GVValueCommand = 127
	AW2_GAIN_B_LEVEL       GVValueCommand = 128
	MA_GAIN_B2_LEVEL       GVValueCommand = 128
	CONT_HVBAL_LEVEL       GVValueCommand = 130
	CONT_BWBAL_LEVEL       GVValueCommand = 131
	CONT_INEDGE_LEVEL      GVValueCommand = 132
	CONT_NOISE_SLICER      GVValueCommand = 133
	CONT_LEVDEP_LEVEL      GVValueCommand = 134
	CONT_INBAND_LEVEL      GVValueCommand = 135
	CONT_EDGEBAND_LEVEL    GVValueCommand = 136
	CONT_VERTCONT_LEVEL    GVValueCommand = 137
	KNEE_SLOPE_M           GVValueCommand = 29
	KNEE_SLOPE_R           GVValueCommand = 140
	KNEE_SLOPE_B           GVValueCommand = 141
	KNEE_ATTACK_M          GVValueCommand = 142
	KNEE_ATTACK_R          GVValueCommand = 143
	KNEE_ATTACK_B          GVValueCommand = 144
	WHITE_LIMIT_M          GVValueCommand = 145
	WHITE_LIMIT_R          GVValueCommand = 146
	WHITE_LIMIT_G          GVValueCommand = 147
	WHITE_LIMIT_B          GVValueCommand = 148
	EXPOSURE_LEVEL         GVValueCommand = 149
	EXP_INDICATION         GVValueCommand = 150
	EXP_INDICATION_HS      GVValueCommand = 151
	EXPOSURE_LEVEL_2       GVValueCommand = 152
	COLOUR_FILTER          GVValueCommand = 153
	PCI_PANEL_RX_MSG_NR    GVValueCommand = 154
	PCI_PANEL_TX_MSG_NR    GVValueCommand = 155
	PCI_CAM_RX_MSG_NR      GVValueCommand = 156
	PCI_CAM_TX_MSG_NR      GVValueCommand = 157
	PCI_COBS_TX_COUNT      GVValueCommand = 158
	PCI_COBS_RX_COUNT      GVValueCommand = 159
	PCI_COBS_ERR_COUNT     GVValueCommand = 160
	PCI_DTCP_ERR_COUNT     GVValueCommand = 161
	BATTERY_LEVEL          GVValueCommand = 162
	COLCORR_SET_4_COLOR    GVValueCommand = 185
	COLCORR_SET_4_WIDTH    GVValueCommand = 186
	COLCORR_SET_4_HUE      GVValueCommand = 187
	COLCORR_SET_4_SAT      GVValueCommand = 188
	COLCORR_SET_4_LUM      GVValueCommand = 189
	COLCORR_SET_5_COLOR    GVValueCommand = 190
	COLCORR_SET_5_WIDTH    GVValueCommand = 191
	COLCORR_SET_5_HUE      GVValueCommand = 192
	COLCORR_SET_5_SAT      GVValueCommand = 193
	COLCORR_SET_5_LUM      GVValueCommand = 194
	COLCORR_SET_6_COLOR    GVValueCommand = 195
	COLCORR_SET_6_WIDTH    GVValueCommand = 196
	COLCORR_SET_6_HUE      GVValueCommand = 197
	COLCORR_SET_6_SAT      GVValueCommand = 198
	COLCORR_SET_6_LUM      GVValueCommand = 199
	KNEE_POINT_LEVEL       GVValueCommand = 200
	WH_BAL_RED_LEVEL       GVValueCommand = 201
	WH_BAL_BLUE_LEVEL      GVValueCommand = 202
	HSAW_RED_LEVEL2        GVValueCommand = 207
	HSAW_GREEN_LEVEL2      GVValueCommand = 208
	HSAW_BLUE_LEVEL2       GVValueCommand = 209
	HPAR_RED_LEVEL2        GVValueCommand = 210
	HPAR_GREEN_LEVEL2      GVValueCommand = 211
	HPAR_BLUE_LEVEL2       GVValueCommand = 212
	VSAW_RED_LEVEL2        GVValueCommand = 213
	VSAW_GREEN_LEVEL2      GVValueCommand = 214
	VSAW_BLUE_LEVEL2       GVValueCommand = 215
	VPAR_RED_LEVEL2        GVValueCommand = 216
	VPAR_GREEN_LEVEL2      GVValueCommand = 217
	VPAR_BLUE_LEVEL2       GVValueCommand = 218
	RE_RED_SAW             GVValueCommand = 219
	RE_GREEN_SAW           GVValueCommand = 220
	RE_BLUE_SAW            GVValueCommand = 221
	VF_CONTOUR_LEVEL       GVValueCommand = 222
	ZEBRA_LEV              GVValueCommand = 223
	ZEBRA_RANGE            GVValueCommand = 224
	V_APERTURE_LEVEL       GVValueCommand = 225
	H_APERTURE_LEVEL       GVValueCommand = 226
	CAM_TEMP_LEVEL         GVValueCommand = 227
	CAM_CRITICAL_TEMP      GVValueCommand = 228
	COLCORR_COLOR          GVValueCommand = 229
	COLCORR_WIDTH          GVValueCommand = 230
	COLCORR_HUE            GVValueCommand = 231
	COLCORR_SAT            GVValueCommand = 232
	COLCORR_LUM            GVValueCommand = 233
	COLCORR_SETS_ON        GVValueCommand = 234
	COLCORR_SET_1_COLOR    GVValueCommand = 235
	COLCORR_SET_1_WIDTH    GVValueCommand = 236
	COLCORR_SET_1_HUE      GVValueCommand = 237
	COLCORR_SET_1_SAT      GVValueCommand = 238
	COLCORR_SET_1_LUM      GVValueCommand = 239
	COLCORR_SET_2_COLOR    GVValueCommand = 240
	COLCORR_SET_2_WIDTH    GVValueCommand = 241
	COLCORR_SET_2_HUE      GVValueCommand = 242
	COLCORR_SET_2_SAT      GVValueCommand = 243
	COLCORR_SET_2_LUM      GVValueCommand = 244
	COLCORR_SET_3_COLOR    GVValueCommand = 245
	COLCORR_SET_3_WIDTH    GVValueCommand = 246
	COLCORR_SET_3_HUE      GVValueCommand = 247
	COLCORR_SET_3_SAT      GVValueCommand = 248
	COLCORR_SET_3_LUM      GVValueCommand = 249
	SKIN1_CONT_LEVEL       GVValueCommand = 250
	SKIN2_CONT_LEVEL       GVValueCommand = 251
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
	SKIN_CONTOUR          GVModeParams = 1
	GAIN_SELECT           GVModeParams = 2
	KNEE_SELECT           GVModeParams = 3
	GAMMA_SELECT          GVModeParams = 4
	SYSTEMID_SELECT       GVModeParams = 5
	FILTER_SELECT         GVModeParams = 6
	AUDIO_FILTER          GVModeParams = 7
	LENS_TYPE             GVModeParams = 8
	MONITORING_SELECT     GVModeParams = 9
	USER_LEVEL            GVModeParams = 10
	CRTSCAN_SELECT        GVModeParams = 11
	GAMMA_CURVE_SELECT    GVModeParams = 12
	KNEE_SOURCE_SELECT    GVModeParams = 13
	FLARE_SELECT          GVModeParams = 14
	FSTOP_SELECT          GVModeParams = 15
	SKIN_VIEW_SELECT      GVModeParams = 16
	MATRIX_DIG_SELECT     GVModeParams = 17
	MATRIX_SELECT         GVModeParams = 18
	EXPOSURE9000          GVModeParams = 19
	BSMENU_CONTROL        GVModeParams = 20
	DISKRECORDING_IF      GVModeParams = 21
	FLICKER_REDUCTION     GVModeParams = 22
	VIDEO_MODE            GVModeParams = 23
	FRAME_FORMAT          GVModeParams = 24
	TEMPORAL_FREQUENCY    GVModeParams = 25
	COLOUR_TEMP_SELECT    GVModeParams = 26
	EXPOSURE_SELECT       GVModeParams = 27
	MONITOR_VF_SELECT     GVModeParams = 28
	CAM_ENG_PROD_SELECT   GVModeParams = 29
	FLOOR_ENG_PROD_SELECT GVModeParams = 30
	CAMERA_TYPE_SELECT    GVModeParams = 31
	CONTOUR_SELECT        GVModeParams = 32
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

type GVModeColorTemp uint8

const (
	FCT_CTEMP_7500K GVModeColorTemp = 1
	FCT_CTEMP_5600K GVModeColorTemp = 2
	FCT_CTEMP_3200K GVModeColorTemp = 3
	FCT_CTEMP_AW1   GVModeColorTemp = 4
	FCT_CTEMP_AW2   GVModeColorTemp = 5
	FCT_CTEMP_TL    GVModeColorTemp = 6
	FCT_CTEMP_AWC   GVModeColorTemp = 7
	FCT_CTEMP_4700K GVModeColorTemp = 8
	FCT_CTEMP_FL50  GVModeColorTemp = 9
	FCT_CTEMP_FL60  GVModeColorTemp = 10
	FCT_CTEMP_THRU  GVModeColorTemp = 11
	FCT_CTEMP_VAR   GVModeColorTemp = 12
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

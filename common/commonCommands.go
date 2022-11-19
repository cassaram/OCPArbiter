package common

type CameraFunction uint8

const (
	NOFUNC_ERR CameraFunction = iota
	CameraNumber
	CallSignal
	ColorBar
	GainMaster
	GainRed
	GainGreen
	GainBlue
	BlackMaster
	BlackRed
	BlackGreen
	BlackBlue
	FlareMode
	FlareRed
	FlareGreen
	FlareBlue
	Iris
	FStop
	IrisAuto
	IrisExtended
	KneeMode
	KneeLevel
	KneeDesaturationLevel
	KneeSlopeRed
	KneeSlopeBlue
	KneeAttack
	KneeAttackRed
	KneeAttackBlue
	KneePoint
	GammaMode
	GammaMaster
	GammaRed
	GammaGreen
	GammaBlue
	WhiteBalanceMode
	WhiteBalanceRed
	WhiteBalanceBlue
	WhiteBalanceAutoAction
)

type CameraCommandAdjustment uint8

const (
	Absolute CameraCommandAdjustment = iota
	Relative
)

func GetCameraFunctionLimits(fun CameraFunction) (int, int) {
	switch fun {
	case CallSignal:
		return 0, 1
	case ColorBar:
		return 0, 1
	case GainMaster:
		return 0, 4095
	case BlackMaster:
		return 0, 4095
	case Iris:
		return 0, 4095
	case IrisAuto:
		return 0, 1
	case IrisExtended:
		return 0, 1
	default:
		return 0, 255
	}
}

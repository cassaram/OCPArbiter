package common

type CameraFunction uint8

const (
	CameraNumber CameraFunction = iota
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
	FlareRed
	FlareGreen
	FlareBlue
	MatrixRedGreen
	MatrixRedBlue
	MatrixGreenRed
	MatrixGreenBlue
	MatrixBlueRed
	MatrixBlueGreen
	MatrixGamma
	Iris
	FStop
	IrisAuto
	IrisExtended
	KneeLevel
	KneeDesaturationLevel
	KneeSlope
	KneeSlopeRed
	KneeSlopeBlue
	KneeAttack
	KneeAttackRed
	KneeAttackBlue
	KneePoint
	GammaMaster
	GammaRed
	GammaGreen
	GammaBlue
	WhiteBalanceRed
	WhiteBalanceBlue
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

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

type CameraCommand struct {
	Command    CameraFunction
	Value      int
	Adjustment CameraCommandAdjustment
}

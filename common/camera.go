package common

// Implements an interface for a camera

type CamFeatureSet struct {
	CamNumber    bool
	CallSignal   bool
	ColorBar     bool
	GainMaster   bool
	GainR        bool
	GainG        bool
	GainB        bool
	BlackR       bool
	BlackG       bool
	BlackB       bool
	FlareR       bool
	FlareG       bool
	FlareB       bool
	MatrixRG     bool
	MatrixRB     bool
	MatrixGR     bool
	MatrixGB     bool
	MatrixBR     bool
	MatrixBG     bool
	MatrixGamma  bool
	BlackMaster  bool
	Iris         bool
	FStop        bool
	IrisAuto     bool
	KneeLvl      bool
	KneeDesatLvl bool
	KneeSlope    bool
	KneeSlopeR   bool
	KneeSlopeB   bool
	KneeAttack   bool
	KneeAttackR  bool
	KneeAttackB  bool
	KneePoint    bool
	GammaLevel   bool
	GammaR       bool
	GammaG       bool
	GammaB       bool
	WBMode       bool
	WBR          bool
	WBB          bool
}

type Camera interface {
	// Initialization
	Initialize()
	GetFeatureSet() CamFeatureSet

	// System
	GetCamNumber() int
	SetCamNumber(int)
	GetCallSig() int
	SetCallSig(int)
	GetColorBar() int
	SetColorBar(int)

	// Gain
	GetGainMaster() int
	SetGainMaster(int)
	GetGainR() int
	SetGainR(int)
	GetGainG() int
	SetGainG(int)
	GetGainB() int
	SetGainB(int)

	// Black
	GetBlackR() int
	SetBlackR(int)
	GetBlackG() int
	SetBlackG(int)
	GetBlackB() int
	SetBlackB(int)

	// Flare
	GetFlareR() int
	SetFlareR(int)
	GetFlareG() int
	SetFlareG(int)
	GetFlareB() int
	SetFlareB(int)

	// Matrix
	GetMatrixRG() int
	SetMatrixRG(int)
	GetMatrixRB() int
	SetMatrixRB(int)
	GetMatrixGR() int
	SetMatrixGR(int)
	GetMatrixGB() int
	SetMatrixGB(int)
	GetMatrixBR() int
	SetMatrixBR(int)
	GetMatrixBG() int
	SetMatrixBG(int)
	GetMatrixGamma() int
	SetMatrixGamma(int)

	// Master black
	GetBlackMaster() int
	SetBlackMaster(int)

	// Iris
	GetIris() int
	SetIris(int)
	GetFStop() float32
	SetFStop(float32)
	GetIrisAuto() int
	SetIrisAuto(int)

	// Knee
	GetKneeLvl() int
	SetKneeLvl(int)
	GetKneeDesatLvl() int
	SetKneeDesatLvl(int)
	GetKneeSlope() int
	SetKneeSlope(int)
	GetKneeSlopeR() int
	SetKneeSlopeR(int)
	GetKneeSlopeB() int
	SetKneeSlopeB(int)
	GetKneeAttack() int
	SetKneeAttack(int)
	GetKneeAttackR() int
	SetKneeAttackR(int)
	GetKneeAttackB() int
	SetKneeAttackB(int)
	GetKneePoint() int
	SetKneePoint(int)

	// Gamma
	GetGammaLevel() int
	SetGammaLevel(int)
	GetGammaR() int
	SetGammaR(int)
	GetGammaG() int
	SetGammaG(int)
	GetGammaB() int
	SetGammaB(int)

	// White Balance
	GetWBMode() int
	SetWBMode(int)
	GetWBR() int
	SetWBR(int)
	GetWBB() int
	SetWBB(int)
}

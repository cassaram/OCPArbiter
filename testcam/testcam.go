package testcam

import "github.com/cassaram/ocparbiter/common"

// Implements a simple "virtual" camera destination for testing controls and arbitration

type TestCam struct {
	featureSet   common.CamFeatureSet
	camNumber    int
	callSignal   bool
	colorBar     bool
	gainMaster   int
	gainR        int
	gainG        int
	gainB        int
	blackR       int
	blackG       int
	blackB       int
	blackMaster  int
	flareR       int
	flareG       int
	flareB       int
	matrixRG     int
	matrixRB     int
	matrixGR     int
	matrixGB     int
	matrixBR     int
	matrixBG     int
	matrixGamma  int
	iris         int
	irisAuto     bool
	kneeLvl      int
	kneeDesatLvl int
	kneeSlope    int
	kneeSlopeR   int
	kneeSlopeB   int
	kneeAttack   int
	kneeAttackR  int
	kneeAttackB  int
	kneePoint    int
	gammaLevel   int
	gammaR       int
	gammaG       int
	gammaB       int
	wbR          int
	wbB          int
}

func (c *TestCam) Initialize() {
	// Setup features
	c.featureSet.CamNumber = true
	c.featureSet.CallSignal = true
	c.featureSet.ColorBar = true
	c.featureSet.GainMaster = true
	c.featureSet.GainR = true
	c.featureSet.GainG = true
	c.featureSet.GainB = true
	c.featureSet.BlackR = true
	c.featureSet.BlackG = true
	c.featureSet.BlackB = true
	c.featureSet.FlareR = true
	c.featureSet.FlareG = true
	c.featureSet.FlareB = true
	c.featureSet.MatrixRG = true
	c.featureSet.MatrixRB = true
	c.featureSet.MatrixGR = true
	c.featureSet.MatrixGB = true
	c.featureSet.MatrixBR = true
	c.featureSet.MatrixBG = true
	c.featureSet.MatrixGamma = true
	c.featureSet.BlackMaster = true
	c.featureSet.Iris = true
	c.featureSet.IrisAuto = true
	c.featureSet.KneeLvl = true
	c.featureSet.KneeDesatLvl = true
	c.featureSet.KneeSlope = true
	c.featureSet.KneeSlopeR = true
	c.featureSet.KneeSlopeB = true
	c.featureSet.KneeAttack = true
	c.featureSet.KneeAttackR = true
	c.featureSet.KneeAttackB = true
	c.featureSet.KneePoint = true
	c.featureSet.GammaLevel = true
	c.featureSet.GammaR = true
	c.featureSet.GammaG = true
	c.featureSet.GammaB = true
	c.featureSet.WBR = true
	c.featureSet.WBB = true

	// Variables init
	c.callSignal = false
	c.colorBar = false
	c.gainR = 0
	c.gainG = 0
	c.gainB = 0
	c.gainMaster = 0
	c.blackR = 0
	c.blackG = 0
	c.blackB = 0
	c.blackMaster = 0
	c.flareR = 0
	c.flareG = 0
	c.flareB = 0
	c.iris = 0
	c.irisAuto = false
}

func (c *TestCam) GetFeatureSet() common.CamFeatureSet {
	return c.featureSet
}

// System
func (c *TestCam) GetCamNumber() int {
	return c.camNumber
}

func (c *TestCam) SetCamNumber(num int) {
	c.camNumber = num
}

func (c *TestCam) GetCallSig() int {
	if c.callSignal {
		return 1
	} else {
		return 0
	}
}

func (c *TestCam) SetCallSig(call int) {
	if call == 0 {
		c.callSignal = false
	} else {
		c.callSignal = true
	}
}

func (c *TestCam) GetColorBar() int {
	if c.colorBar {
		return 1
	} else {
		return 0
	}
}

func (c *TestCam) SetColorBar(bar int) {
	if bar == 0 {
		c.colorBar = false
	} else {
		c.colorBar = true
	}
}

// Gain
func (c *TestCam) GetGainMaster() int {
	return c.gainMaster
}

func (c *TestCam) SetGainMaster(gain int) {
	c.gainMaster = gain
}

func (c *TestCam) GetGainR() int {
	return c.gainR
}

func (c *TestCam) SetGainR(gr int) {
	c.gainR = gr
}

func (c *TestCam) GetGainG() int {
	return c.gainG
}

func (c *TestCam) SetGainG(gg int) {
	c.gainG = gg
}

func (c *TestCam) GetGainB() int {
	return c.gainB
}

func (c *TestCam) SetGainB(gb int) {
	c.gainB = gb
}

// Black
func (c *TestCam) GetBlackR() int {
	return c.blackR
}

func (c *TestCam) SetBlackR(br int) {
	c.blackR = br
}

func (c *TestCam) GetBlackG() int {
	return c.blackG
}

func (c *TestCam) SetBlackG(bg int) {
	c.blackG = bg
}

func (c *TestCam) GetBlackB() int {
	return c.blackB
}

func (c *TestCam) SetBlackB(bb int) {
	c.blackB = bb
}

// Flare
func (c *TestCam) GetFlareR() int {
	return c.flareR
}

func (c *TestCam) SetFlareR(fr int) {
	c.flareR = fr
}

func (c *TestCam) GetFlareG() int {
	return c.flareG
}

func (c *TestCam) SetFlareG(fg int) {
	c.flareG = fg
}

func (c *TestCam) GetFlareB() int {
	return c.flareB
}

func (c *TestCam) SetFlareB(fb int) {
	c.flareB = fb
}

// Matrix
func (c *TestCam) GetMatrixRG() int {
	return c.matrixRG
}

func (c *TestCam) SetMatrixRG(rg int) {
	c.matrixRG = rg
}

func (c *TestCam) GetMatrixRB() int {
	return c.matrixRB
}

func (c *TestCam) SetMatrixRB(rb int) {
	c.matrixRB = rb
}

func (c *TestCam) GetMatrixGR() int {
	return c.matrixGR
}

func (c *TestCam) SetMatrixGR(gr int) {
	c.matrixGR = gr
}

func (c *TestCam) GetMatrixGB() int {
	return c.matrixGB
}

func (c *TestCam) SetMatrixGB(gb int) {
	c.matrixGB = gb
}

func (c *TestCam) GetMatrixBR() int {
	return c.matrixBR
}

func (c *TestCam) SetMatrixBR(br int) {
	c.matrixBR = br
}

func (c *TestCam) GetMatrixBG() int {
	return c.matrixGB
}

func (c *TestCam) SetMatrixBG(bg int) {
	c.matrixBG = bg
}

func (c *TestCam) GetMatrixGamma() int {
	return c.matrixGamma
}

func (c *TestCam) SetMatrixGamma(gamma int) {
	c.matrixGamma = gamma
}

// Master black
func (c *TestCam) GetBlackMaster() int {
	return c.blackMaster
}

func (c *TestCam) SetBlackMaster(mb int) {
	c.blackMaster = mb
}

// Iris
func (c *TestCam) GetIris() int {
	return c.iris
}

func (c *TestCam) SetIris(iris int) {
	c.iris = iris
}

func (c *TestCam) GetIrisAuto() int {
	if c.irisAuto {
		return 1
	} else {
		return 0
	}
}

func (c *TestCam) SetIrisAuto(auto int) {
	if auto == 0 {
		c.irisAuto = false
	} else {
		c.irisAuto = true
	}
}

// Knee
func (c *TestCam) GetKneeLvl() int {
	return c.kneeLvl
}

func (c *TestCam) SetKneeLvl(lvl int) {
	c.kneeLvl = lvl
}

func (c *TestCam) GetKneeDesatLvl() int {
	return c.kneeDesatLvl
}

func (c *TestCam) SetKneeDesatLvl(lvl int) {
	c.kneeDesatLvl = lvl
}

func (c *TestCam) GetKneeSlope() int {
	return c.kneeSlope
}

func (c *TestCam) SetKneeSlope(slope int) {
	c.kneeSlope = slope
}

func (c *TestCam) GetKneeSlopeR() int {
	return c.kneeSlopeR
}

func (c *TestCam) SetKneeSlopeR(slopeR int) {
	c.kneeSlopeR = slopeR
}

func (c *TestCam) GetKneeSlopeB() int {
	return c.kneeSlopeB
}

func (c *TestCam) SetKneeSlopeB(slopeB int) {
	c.kneeSlopeB = slopeB
}

func (c *TestCam) GetKneeAttack() int {
	return c.kneeAttack
}

func (c *TestCam) SetKneeAttack(atk int) {
	c.kneeAttack = atk
}

func (c *TestCam) GetKneeAttackR() int {
	return c.kneeAttackR
}

func (c *TestCam) SetKneeAttackR(atkR int) {
	c.kneeAttackR = atkR
}

func (c *TestCam) GetKneeAttackB() int {
	return c.kneeAttackB
}

func (c *TestCam) SetKneeAttackB(atkB int) {
	c.kneeAttackB = atkB
}

func (c *TestCam) GetKneePoint() int {
	return c.kneePoint
}

func (c *TestCam) SetKneePoint(pt int) {
	c.kneePoint = pt
}

// Gamma
func (c *TestCam) GetGammaLevel() int {
	return c.gammaLevel
}

func (c *TestCam) SetGammaLevel(gamma int) {
	c.gammaLevel = gamma
}

func (c *TestCam) GetGammaR() int {
	return c.gammaR
}

func (c *TestCam) SetGammaR(gammaR int) {
	c.gammaR = gammaR
}

func (c *TestCam) GetGammaG() int {
	return c.gammaG
}

func (c *TestCam) SetGammaG(gammaG int) {
	c.gammaG = gammaG
}

func (c *TestCam) GetGammaB() int {
	return c.gammaB
}

func (c *TestCam) SetGammaB(gammaB int) {
	c.gammaB = gammaB
}

// White Balance
func (c *TestCam) GetWBR() int {
	return c.wbR
}

func (c *TestCam) SetWBR(wbr int) {
	c.wbR = wbr
}

func (c *TestCam) GetWBB() int {
	return c.wbB
}

func (c *TestCam) SetWBB(wbb int) {
	c.wbB = wbb
}

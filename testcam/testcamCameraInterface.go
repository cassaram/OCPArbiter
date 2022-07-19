package testcam

// Implements a simple "virtual" camera destination for testing controls and arbitration

type testCamCameraInterface struct {
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
	fstop        float32
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
	wbMode       int
	wbR          int
	wbB          int
}

func (c *testCamCameraInterface) Initialize() {
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
	c.fstop = 1.0
	c.irisAuto = false
	c.wbMode = 12 // Variable WB
	c.wbR = 0
	c.wbB = 0

}

// System
func (c *testCamCameraInterface) GetCamNumber() int {
	return c.camNumber
}

func (c *testCamCameraInterface) SetCamNumber(num int) {
	c.camNumber = num
}

func (c *testCamCameraInterface) GetCallSig() int {
	if c.callSignal {
		return 1
	} else {
		return 0
	}
}

func (c *testCamCameraInterface) SetCallSig(call int) {
	if call == 0 {
		c.callSignal = false
	} else {
		c.callSignal = true
	}
}

func (c *testCamCameraInterface) GetColorBar() int {
	if c.colorBar {
		return 1
	} else {
		return 0
	}
}

func (c *testCamCameraInterface) SetColorBar(bar int) {
	if bar == 0 {
		c.colorBar = false
	} else {
		c.colorBar = true
	}
}

// Gain
func (c *testCamCameraInterface) GetGainMaster() int {
	return c.gainMaster
}

func (c *testCamCameraInterface) SetGainMaster(gain int) {
	c.gainMaster = gain
}

func (c *testCamCameraInterface) GetGainR() int {
	return c.gainR
}

func (c *testCamCameraInterface) SetGainR(gr int) {
	c.gainR = gr
}

func (c *testCamCameraInterface) GetGainG() int {
	return c.gainG
}

func (c *testCamCameraInterface) SetGainG(gg int) {
	c.gainG = gg
}

func (c *testCamCameraInterface) GetGainB() int {
	return c.gainB
}

func (c *testCamCameraInterface) SetGainB(gb int) {
	c.gainB = gb
}

// Black
func (c *testCamCameraInterface) GetBlackR() int {
	return c.blackR
}

func (c *testCamCameraInterface) SetBlackR(br int) {
	c.blackR = br
}

func (c *testCamCameraInterface) GetBlackG() int {
	return c.blackG
}

func (c *testCamCameraInterface) SetBlackG(bg int) {
	c.blackG = bg
}

func (c *testCamCameraInterface) GetBlackB() int {
	return c.blackB
}

func (c *testCamCameraInterface) SetBlackB(bb int) {
	c.blackB = bb
}

// Flare
func (c *testCamCameraInterface) GetFlareR() int {
	return c.flareR
}

func (c *testCamCameraInterface) SetFlareR(fr int) {
	c.flareR = fr
}

func (c *testCamCameraInterface) GetFlareG() int {
	return c.flareG
}

func (c *testCamCameraInterface) SetFlareG(fg int) {
	c.flareG = fg
}

func (c *testCamCameraInterface) GetFlareB() int {
	return c.flareB
}

func (c *testCamCameraInterface) SetFlareB(fb int) {
	c.flareB = fb
}

// Matrix
func (c *testCamCameraInterface) GetMatrixRG() int {
	return c.matrixRG
}

func (c *testCamCameraInterface) SetMatrixRG(rg int) {
	c.matrixRG = rg
}

func (c *testCamCameraInterface) GetMatrixRB() int {
	return c.matrixRB
}

func (c *testCamCameraInterface) SetMatrixRB(rb int) {
	c.matrixRB = rb
}

func (c *testCamCameraInterface) GetMatrixGR() int {
	return c.matrixGR
}

func (c *testCamCameraInterface) SetMatrixGR(gr int) {
	c.matrixGR = gr
}

func (c *testCamCameraInterface) GetMatrixGB() int {
	return c.matrixGB
}

func (c *testCamCameraInterface) SetMatrixGB(gb int) {
	c.matrixGB = gb
}

func (c *testCamCameraInterface) GetMatrixBR() int {
	return c.matrixBR
}

func (c *testCamCameraInterface) SetMatrixBR(br int) {
	c.matrixBR = br
}

func (c *testCamCameraInterface) GetMatrixBG() int {
	return c.matrixGB
}

func (c *testCamCameraInterface) SetMatrixBG(bg int) {
	c.matrixBG = bg
}

func (c *testCamCameraInterface) GetMatrixGamma() int {
	return c.matrixGamma
}

func (c *testCamCameraInterface) SetMatrixGamma(gamma int) {
	c.matrixGamma = gamma
}

// Master black
func (c *testCamCameraInterface) GetBlackMaster() int {
	return c.blackMaster
}

func (c *testCamCameraInterface) SetBlackMaster(mb int) {
	c.blackMaster = mb
}

// Iris
func (c *testCamCameraInterface) GetIris() int {
	return c.iris
}

func (c *testCamCameraInterface) SetIris(iris int) {
	c.iris = iris
}

func (c *testCamCameraInterface) GetFStop() float32 {
	// Calculate f-stop number from iris 0-4055 value
	// Interprets as a f2.8 - f22 lens because limits
	pct := c.iris * 100 / 4095
	rng := (220 - 28)
	return float32(pct*rng)/1000 + 2.8

}

func (c *testCamCameraInterface) SetFStop(fstop float32) {
	c.fstop = fstop
}

func (c *testCamCameraInterface) GetIrisAuto() int {
	if c.irisAuto {
		return 1
	} else {
		return 0
	}
}

func (c *testCamCameraInterface) SetIrisAuto(auto int) {
	if auto == 0 {
		c.irisAuto = false
	} else {
		c.irisAuto = true
	}
}

// Knee
func (c *testCamCameraInterface) GetKneeLvl() int {
	return c.kneeLvl
}

func (c *testCamCameraInterface) SetKneeLvl(lvl int) {
	c.kneeLvl = lvl
}

func (c *testCamCameraInterface) GetKneeDesatLvl() int {
	return c.kneeDesatLvl
}

func (c *testCamCameraInterface) SetKneeDesatLvl(lvl int) {
	c.kneeDesatLvl = lvl
}

func (c *testCamCameraInterface) GetKneeSlope() int {
	return c.kneeSlope
}

func (c *testCamCameraInterface) SetKneeSlope(slope int) {
	c.kneeSlope = slope
}

func (c *testCamCameraInterface) GetKneeSlopeR() int {
	return c.kneeSlopeR
}

func (c *testCamCameraInterface) SetKneeSlopeR(slopeR int) {
	c.kneeSlopeR = slopeR
}

func (c *testCamCameraInterface) GetKneeSlopeB() int {
	return c.kneeSlopeB
}

func (c *testCamCameraInterface) SetKneeSlopeB(slopeB int) {
	c.kneeSlopeB = slopeB
}

func (c *testCamCameraInterface) GetKneeAttack() int {
	return c.kneeAttack
}

func (c *testCamCameraInterface) SetKneeAttack(atk int) {
	c.kneeAttack = atk
}

func (c *testCamCameraInterface) GetKneeAttackR() int {
	return c.kneeAttackR
}

func (c *testCamCameraInterface) SetKneeAttackR(atkR int) {
	c.kneeAttackR = atkR
}

func (c *testCamCameraInterface) GetKneeAttackB() int {
	return c.kneeAttackB
}

func (c *testCamCameraInterface) SetKneeAttackB(atkB int) {
	c.kneeAttackB = atkB
}

func (c *testCamCameraInterface) GetKneePoint() int {
	return c.kneePoint
}

func (c *testCamCameraInterface) SetKneePoint(pt int) {
	c.kneePoint = pt
}

// Gamma
func (c *testCamCameraInterface) GetGammaLevel() int {
	return c.gammaLevel
}

func (c *testCamCameraInterface) SetGammaLevel(gamma int) {
	c.gammaLevel = gamma
}

func (c *testCamCameraInterface) GetGammaR() int {
	return c.gammaR
}

func (c *testCamCameraInterface) SetGammaR(gammaR int) {
	c.gammaR = gammaR
}

func (c *testCamCameraInterface) GetGammaG() int {
	return c.gammaG
}

func (c *testCamCameraInterface) SetGammaG(gammaG int) {
	c.gammaG = gammaG
}

func (c *testCamCameraInterface) GetGammaB() int {
	return c.gammaB
}

func (c *testCamCameraInterface) SetGammaB(gammaB int) {
	c.gammaB = gammaB
}

// White Balance
func (c *testCamCameraInterface) GetWBMode() int {
	return c.wbMode
}

func (c *testCamCameraInterface) SetWBMode(value int) {
	c.wbMode = value
}

func (c *testCamCameraInterface) GetWBR() int {
	return c.wbR
}

func (c *testCamCameraInterface) SetWBR(wbr int) {
	c.wbR = wbr
}

func (c *testCamCameraInterface) GetWBB() int {
	return c.wbB
}

func (c *testCamCameraInterface) SetWBB(wbb int) {
	c.wbB = wbb
}

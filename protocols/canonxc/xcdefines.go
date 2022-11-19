package canonxc

type XCControlCommand string

const (
	CTRL_SHOOTING_MODE      XCControlCommand = "c.1.shooting"
	CTRL_EXPOSURE_MODE      XCControlCommand = "c.1.exp"
	CTRL_AUTO_GAINLIMIT_MAX XCControlCommand = "c.1.ae.gainlimit.max"
	CTRL_AUTO_FLICKERREDUCT XCControlCommand = "c.1.ae.flickerreduct"
	CTRL_AUTO_RESPONSE      XCControlCommand = "c.1.ae.resp"
	CTRL_SHUTTER_SPEED      XCControlCommand = "c.1.me.shutter"
	CTRL_SHUTTER_MODE       XCControlCommand = "c.1.me.shutter.mode"
	CTRL_SHUTTER_CLEARSCAN  XCControlCommand = "c.1.me.clearscan"
	CTRL_IRIS               XCControlCommand = "c.1.me.iris"
	CTRL_APUTURE            XCControlCommand = "c.1.me.diaphragm"
	CTRL_APUTURE_MODE       XCControlCommand = "c.1.me.diaphragm.mode"
	CTRL_GAIN               XCControlCommand = "c.1.me.gain"
	CTRL_GAIN_MODE          XCControlCommand = "c.1.me.gain.mode"
	CTRL_GAINLIMIT_MAX      XCControlCommand = "c.1.me.gainlimit.max"
	CTRL_BRIGHTNESS         XCControlCommand = "c.1.me.brightness"
	CTRL_PHOTOMETRY         XCControlCommand = "c.1.me.photometry"
	CTRL_RESPONSE           XCControlCommand = "c.1.me.resp"
)

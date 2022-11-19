package canonxc

import (
	"github.com/cassaram/ocparbiter/common"
)

func commonFunctionToCanonXC(cf common.CameraFunction) string {
	switch cf {
	case common.ColorBar:
		return "c.1.colorbar"
	case common.GainMaster:
		return "c.1.me.gain"
	case common.GainRed:
		return "c.1.wb.shift.rgain"
	case common.GainBlue:
		return "c.1.wb.shift.bgain"
	case common.BlackMaster:
		return "c.1.blacklevel"
	case common.BlackRed:
		return "c.1.blacklevel.red"
	case common.BlackBlue:
		return "c.1.blacklevel.blue"
	case common.Iris:
		return "c.1.me.iris"
	case common.FStop:
		return "c.1.me.diaphragm"
	case common.IrisAuto:
		return "c.1.me.diaphragm.mode"
	case common.WhiteBalanceMode:
		return "c.1.wb"
	default:
		return ""
	}
}

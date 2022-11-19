package canonxc

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/cassaram/ocparbiter/common"
)

type CanonXCProtocol struct {
	IPAddress string
}

func (p *CanonXCProtocol) setValue(key string, val string) {
	url := "http://" + p.IPAddress + "/-wvhttp-01-/control.cgi?" + key + "=" + val

	_, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		return
	}
}

func (p *CanonXCProtocol) getValue(key string) string {
	url := "http://" + p.IPAddress + "/-wvhttp-01-/info.cgi?item=" + key

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	body, _ := io.ReadAll(resp.Body)
	respMap := readResponse(string(body))

	val := respMap[key]
	return val
}

func (p *CanonXCProtocol) GetAllValues() map[common.CameraFunction]int {
	basicMap := p.getAllValues()
	result := make(map[common.CameraFunction]int)

	for key, val := range basicMap {
		fun, valInt := p.xcToCommonFunction(key, val)
		result[fun] = valInt
	}

	return result
}

func (p *CanonXCProtocol) getAllValues() map[string]string {
	url := "http://" + p.IPAddress + "/-wvhttp-01-/info.cgi"

	resp, err := http.Get(url)
	if err != nil {
		return make(map[string]string)
	}

	body, _ := io.ReadAll(resp.Body)
	respMap := readResponse(string(body))

	return respMap
}

func (p *CanonXCProtocol) UpdateValue(cmd common.CameraCommand) {
	switch cmd.Function {
	case common.GainMaster:
		p.setGainMaster(cmd)
	case common.BlackMaster:
		p.setBlackMaster(cmd)
	case common.BlackRed:
		p.setBlackRed(cmd)
	case common.BlackBlue:
		p.setBlackBlue(cmd)
	}
}

// Function to convert xc commands to common commands
func (p *CanonXCProtocol) xcToCommonFunction(key string, value string) (common.CameraFunction, int) {
	switch key {
	case "c.1.colorbar":
		if value == "true" {
			return common.ColorBar, 1
		} else {
			return common.ColorBar, 0
		}
	case "c.1.me.gain":
		valInt, _ := strconv.Atoi(value)
		return common.GainMaster, valInt
	case "c.1.wb.shift.rgain":
		valInt, _ := strconv.Atoi(value)
		AHigh, _ := strconv.Atoi(p.getValue("c.1.me.shift.rgain.max"))
		ALow, _ := strconv.Atoi(p.getValue("c.1.me.shift.rgain.min"))
		BLow, BHigh := common.GetCameraFunctionLimits(common.GainRed)
		result := convertRange(ALow, AHigh, BLow, BHigh, valInt)
		return common.GainRed, result
	case "c.1.wb.shift.bgain":
		valInt, _ := strconv.Atoi(value)
		AHigh, _ := strconv.Atoi(p.getValue("c.1.me.shift.bgain.max"))
		ALow, _ := strconv.Atoi(p.getValue("c.1.me.shift.bgain.min"))
		BLow, BHigh := common.GetCameraFunctionLimits(common.GainBlue)
		result := convertRange(ALow, AHigh, BLow, BHigh, valInt)
		return common.GainBlue, result
	case "c.1.blacklevel":
		valInt, _ := strconv.Atoi(value)
		AHigh, _ := strconv.Atoi(p.getValue("c.1.blacklevel.max"))
		ALow, _ := strconv.Atoi(p.getValue("c.1.blacklevel.min"))
		BLow, BHigh := common.GetCameraFunctionLimits(common.BlackMaster)
		result := convertRange(ALow, AHigh, BLow, BHigh, valInt)
		return common.BlackMaster, result
	case "c.1.blacklevel.red":
		valInt, _ := strconv.Atoi(value)
		AHigh, _ := strconv.Atoi(p.getValue("c.1.blacklevel.red.max"))
		ALow, _ := strconv.Atoi(p.getValue("c.1.blacklevel.red.min"))
		BLow, BHigh := common.GetCameraFunctionLimits(common.BlackRed)
		result := convertRange(ALow, AHigh, BLow, BHigh, valInt)
		return common.BlackRed, result
	case "c.1.blacklevel.blue":
		valInt, _ := strconv.Atoi(value)
		AHigh, _ := strconv.Atoi(p.getValue("c.1.blacklevel.blue.max"))
		ALow, _ := strconv.Atoi(p.getValue("c.1.blacklevel.blue.min"))
		BLow, BHigh := common.GetCameraFunctionLimits(common.BlackBlue)
		result := convertRange(ALow, AHigh, BLow, BHigh, valInt)
		return common.BlackBlue, result
	case "c.1.me.iris":
		valInt, _ := strconv.Atoi(value)
		AHigh, _ := strconv.Atoi(p.getValue("c.1.me.iris.max"))
		ALow, _ := strconv.Atoi(p.getValue("c.1.me.iris.min"))
		BLow, BHigh := common.GetCameraFunctionLimits(common.Iris)
		result := convertRange(ALow, AHigh, BLow, BHigh, valInt)
		return common.Iris, result
	case "c.1.me.diaphragm":
		valInt, _ := strconv.Atoi(value)
		return common.FStop, valInt
	case "c.1.me.diaphragm.mode":
		if value == "manual" {
			return common.IrisAuto, 0
		} else if value == "auto" {
			return common.IrisAuto, 1
		}
	case "c.1.me.wb":
		switch value {
		case "auto":
			return common.WhiteBalanceMode, 7
		case "manual":
			return common.WhiteBalanceMode, 12
		case "wb_a":
			return common.WhiteBalanceMode, 4
		case "wb_b":
			return common.WhiteBalanceMode, 5
		case "daylight":
			return common.WhiteBalanceMode, 2
		case "tungsten":
			return common.WhiteBalanceMode, 3
		case "kelvin":
			return common.WhiteBalanceMode, 12
		}
	}

	return common.NOFUNC_ERR, -1
}

func (p *CanonXCProtocol) setGainMaster(cmd common.CameraCommand) {
	val := cmd.Value
	if cmd.Adjustment == common.Relative {
		curVal, _ := strconv.Atoi(p.getValue("c.1.me.gain"))
		val += curVal
	}

	p.setValue("c.1.me.gain", strconv.Itoa(val))
}

func (p *CanonXCProtocol) setBlackMaster(cmd common.CameraCommand) {
	val := cmd.Value

	BLow, _ := strconv.Atoi(p.getValue("c.1.blacklevel.min"))
	BHigh, _ := strconv.Atoi(p.getValue("c.1.blacklevel.max"))
	ALow, AHigh := common.GetCameraFunctionLimits(cmd.Function)
	newVal := convertRange(ALow, AHigh, BLow, BHigh, val)

	if cmd.Adjustment == common.Relative {
		curVal, _ := strconv.Atoi(p.getValue("c.1.blacklevel"))
		newVal += curVal
	}

	p.setValue("c.1.blacklevel", strconv.Itoa(newVal))
}

func (p *CanonXCProtocol) setBlackRed(cmd common.CameraCommand) {
	val := cmd.Value

	BLow, _ := strconv.Atoi(p.getValue("c.1.blacklevel.red.min"))
	BHigh, _ := strconv.Atoi(p.getValue("c.1.blacklevel.red.max"))
	ALow, AHigh := common.GetCameraFunctionLimits(cmd.Function)
	newVal := convertRange(ALow, AHigh, BLow, BHigh, val)

	if cmd.Adjustment == common.Relative {
		curVal, _ := strconv.Atoi(p.getValue("c.1.blacklevel.red"))
		newVal += curVal
	}

	p.setValue("c.1.blacklevel.red", strconv.Itoa(newVal))
}

func (p *CanonXCProtocol) setBlackBlue(cmd common.CameraCommand) {
	val := cmd.Value

	BLow, _ := strconv.Atoi(p.getValue("c.1.blacklevel.blue.min"))
	BHigh, _ := strconv.Atoi(p.getValue("c.1.blacklevel.blue.max"))
	ALow, AHigh := common.GetCameraFunctionLimits(cmd.Function)
	newVal := convertRange(ALow, AHigh, BLow, BHigh, val)

	if cmd.Adjustment == common.Relative {
		curVal, _ := strconv.Atoi(p.getValue("c.1.blacklevel.blue"))
		newVal += curVal
	}

	p.setValue("c.1.blacklevel.blue", strconv.Itoa(newVal))
}

func (p *CanonXCProtocol) setWhiteBalanceMode(mode int) {
	switch mode {
	case 2:
		// 5600k
		p.setValue("c.1.wb", "daylight")
	case 3:
		// 3200k
		p.setValue("c.1.wb", "tungsten")
	case 4:
		// AW1
		p.setValue("c.1.wb", "wb_a")
	case 5:
		// AW2
		p.setValue("c.1.wb", "wb_b")
	case 7:
		// AWC
		p.setValue("c.1.wb", "auto")
	case 12:
		// VAR
		p.setValue("c.1.wb", "manual")
	}
}

func (p *CanonXCProtocol) setIris(iris int) {
	BHigh, _ := strconv.Atoi(p.getValue("c.1.me.iris.max"))
	BLow, _ := strconv.Atoi(p.getValue("c.1.me.iris.min"))
	ALow, AHigh := common.GetCameraFunctionLimits(common.Iris)
	result := convertRange(ALow, AHigh, BLow, BHigh, iris)

	p.setValue("c.1.me.iris", strconv.Itoa(result))
}

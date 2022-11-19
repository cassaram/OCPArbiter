package settings

import "strconv"

type Setting struct {
	Type     SettingType     `json:"type"`
	Id       string          `json:"id"`
	Label    string          `json:"label"`
	Value    string          `json:"value"`
	Default  string          `json:"default"`
	Regex    string          `json:"regex"`
	Visible  func([]Setting) `json:"visible"`
	EnumList []EnumValue     `json:"enum_list"`
}

type EnumValue struct {
	Value     string `json:"value"`
	Text      string `json:"text"`
	ShortText string `json:"short_text"`
}

func (s *Setting) ValueInt() int {
	val, err := strconv.Atoi(s.Value)
	if err != nil {
		return -1
	}
	return val
}

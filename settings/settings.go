package settings

type Setting struct {
	Type    SettingType     `json:"type"`
	Id      string          `json:"id"`
	Label   string          `json:"label"`
	Value   SettingValue    `json:"value"`
	Default SettingValue    `json:"default"`
	Regex   string          `json:"regex"`
	Visible func([]Setting) `json:"visible"`
}

type SettingValue struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

type EnumValue struct {
	Value     int    `json:"value"`
	Text      string `json:"text"`
	ShortText string `json:"short_text"`
}

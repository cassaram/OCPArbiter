package settings

type SettingType string

const (
	Enum        SettingType = "enum"
	TextInput   SettingType = "text_input"
	TextDisplay SettingType = "text_display"
	Toggle      SettingType = "toggle"
	Number      SettingType = "number"
)

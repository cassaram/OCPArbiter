package settings

type SettingType string

const (
	Enum           SettingType = "enum"
	TextInput      SettingType = "string"
	DisplayText    SettingType = "data"
	Toggle         SettingType = "integer"
	Number         SettingType = "octets"
	OctetsReadOnly SettingType = "octets_read_only"
)

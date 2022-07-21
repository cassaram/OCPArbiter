package settings

type Setting struct {
	ParamType         ParamTypes         `json:"param_type"`
	DescriptorType    ParamTypes         `json:"descriptor_type"`
	ParamID           string             `json:"param_id"`
	ParamName         string             `json:"param_name"`
	PersistenceType   PersistenceTypes   `json:"persistence_type"`
	RegisterType      string             `json:"register_type"`
	FactoryResetType  string             `json:"warm"`
	Relations         Relations          `json:"relations"`
	ClassNames        []string           `json:"class_names"`
	StringAttributes  []StringAttribute  `json:"string_attributes"`
	IntegerAttributes []IntegerAttribute `json:"integer_attributes"`
	EnumValues        []EnumValue        `json:"enum_values"`
	MinValue          int                `json:"min_value"`
	MaxValue          int                `json:"max_value"`
	DefaultValue      int                `json:"default_value"`
	AdjustBy          int                `json:"adjust_by"`
	ScaleBy           int                `json:"scale_by"`
	OffsetBy          int                `json:"offset_by"`
	DisplayPrecision  int                `json:"display_precision"`
	Units             string             `json:"units"`
}

type Relations struct {
}

type StringAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type IntegerAttribute struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type EnumValue struct {
	Value     int    `json:"value"`
	Text      string `json:"text"`
	ShortText string `json:"short_text"`
}

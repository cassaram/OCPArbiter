package settings

type ParamTypes string

const (
	Enum           ParamTypes = "enum"
	String         ParamTypes = "string"
	Data           ParamTypes = "data"
	Integer        ParamTypes = "integer"
	Octets         ParamTypes = "octets"
	OctetsReadOnly ParamTypes = "octets_read_only"
)

type PersistenceTypes string

const (
	Persistent PersistenceTypes = "persistent"
	Ephemeral  PersistenceTypes = "ephemeral"
)

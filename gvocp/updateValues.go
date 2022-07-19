package gvocp

type updateValue struct {
	d_id          uint8
	group         uint8
	masterCommand uint8
	params        []byte
}

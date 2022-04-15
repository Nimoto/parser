package reader

type Data struct {
	id       uint32
	name     string
	required bool
}

func (d *Data) SetData(name string, required bool) *Data {
	return &Data{
		name:     name,
		required: required,
	}
}

func (d *Data) GetId() uint32 {
	return d.id
}

func (d *Data) GetName() string {
	return d.name
}

func (d *Data) IsRequired() bool {
	return d.required
}

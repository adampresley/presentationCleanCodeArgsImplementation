package arguments

type StringArgumentMarshaler struct {
	argumentID     string
	valueToMarshal string
	value          string
}

func (marshaler *StringArgumentMarshaler) Value() interface{} {
	return marshaler.value
}

func (marshaler *StringArgumentMarshaler) Marshal() error {
	marshaler.value = marshaler.valueToMarshal
	return nil
}

func NewStringArgumentMarshaler(argumentID string, value string) *StringArgumentMarshaler {
	return &StringArgumentMarshaler{
		argumentID:     argumentID,
		valueToMarshal: value,
	}
}

package arguments

import "strconv"

type IntegerArgumentMarshaler struct {
	argumentID     string
	valueToMarshal string
	value          int
}

func (marshaler *IntegerArgumentMarshaler) Value() interface{} {
	return marshaler.value
}

func (marshaler *IntegerArgumentMarshaler) Marshal() error {
	var err error
	var result int

	if result, err = strconv.Atoi(marshaler.valueToMarshal); err != nil {
		return InvalidIntegerException(marshaler.argumentID, marshaler.valueToMarshal)
	}

	marshaler.value = result
	return nil
}

func NewIntegerArgumentMarshaler(argumentID string, value string) *IntegerArgumentMarshaler {
	return &IntegerArgumentMarshaler{
		argumentID:     argumentID,
		valueToMarshal: value,
	}
}

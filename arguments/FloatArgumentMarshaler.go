package arguments

import "strconv"

type FloatArgumentMarshaler struct {
	argumentID     string
	valueToMarshal string
	value          float64
}

func (marshaler *FloatArgumentMarshaler) Value() interface{} {
	return marshaler.value
}

func (marshaler *FloatArgumentMarshaler) Marshal() error {
	var err error
	var result float64

	if result, err = strconv.ParseFloat(marshaler.valueToMarshal, 64); err != nil {
		return InvalidFloatException(marshaler.argumentID, marshaler.valueToMarshal)
	}

	marshaler.value = result
	return nil
}

func NewFloatArgumentMarshaler(argumentID string, value string) *FloatArgumentMarshaler {
	return &FloatArgumentMarshaler{
		argumentID:     argumentID,
		valueToMarshal: value,
	}
}

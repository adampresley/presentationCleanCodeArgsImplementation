package arguments

import "strconv"

type BooleanArgumentMarshaler struct {
	argumentID     string
	valueToMarshal string
	value          bool
}

func (marshaler *BooleanArgumentMarshaler) Value() interface{} {
	return marshaler.value
}

func (marshaler *BooleanArgumentMarshaler) Marshal() error {
	var err error
	var result bool

	if result, err = strconv.ParseBool(marshaler.valueToMarshal); err != nil {
		return InvalidBooleanException(marshaler.argumentID, marshaler.valueToMarshal)
	}

	marshaler.value = result
	return nil
}

func NewBooleanArgumentMarshaler(argumentID string, value string) *BooleanArgumentMarshaler {
	return &BooleanArgumentMarshaler{
		argumentID:     argumentID,
		valueToMarshal: value,
	}
}

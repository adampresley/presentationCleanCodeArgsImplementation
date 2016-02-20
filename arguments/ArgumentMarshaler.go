package arguments

/*
ArgumentMarshaler is an interface to be implement by specific type marshalers.
Each type marshaler will parse the argument and place
the result into the implementing structure's value. It may then be read
using Value()
*/
type ArgumentMarshaler interface {
	Marshal() error
	Value() interface{}
}

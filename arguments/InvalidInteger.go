package arguments

import "fmt"

/*
InvalidInteger describes an error where an integer parameter gets a
value that is not the right type
*/
type InvalidInteger struct {
	ArgumentID string
	Parameter  string
}

func InvalidIntegerException(argumentID, parameter string) InvalidInteger {
	return InvalidInteger{ArgumentID: argumentID, Parameter: parameter}
}

func (exception InvalidInteger) Error() string {
	return fmt.Sprintf("Argument -%s expects an integer but was '%s'.", exception.ArgumentID, exception.Parameter)
}

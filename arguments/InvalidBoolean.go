package arguments

import "fmt"

/*
InvalidBoolean describes an error where a boolean parameter gets a
value that is not the right type
*/
type InvalidBoolean struct {
	ArgumentID string
	Parameter  string
}

func InvalidBooleanException(argumentID, parameter string) InvalidBoolean {
	return InvalidBoolean{ArgumentID: argumentID, Parameter: parameter}
}

func (exception InvalidBoolean) Error() string {
	return fmt.Sprintf("Argument -%s expects a boolean but was '%s'.", exception.ArgumentID, exception.Parameter)
}

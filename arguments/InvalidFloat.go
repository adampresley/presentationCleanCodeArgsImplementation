package arguments

import "fmt"

/*
InvalidFloat describes an error where a float parameter gets a
value that is not the right type
*/
type InvalidFloat struct {
	ArgumentID string
	Parameter  string
}

func InvalidFloatException(argumentID, parameter string) InvalidFloat {
	return InvalidFloat{ArgumentID: argumentID, Parameter: parameter}
}

func (exception InvalidFloat) Error() string {
	return fmt.Sprintf("Argument -%s expects a float but was '%s'.", exception.ArgumentID, exception.Parameter)
}

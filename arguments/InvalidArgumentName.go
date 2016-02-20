package arguments

import "fmt"

/*
InvalidArgumentName describes an error where a schema argument name contains
invalid characters
*/
type InvalidArgumentName struct {
	ArgumentID string
}

func InvalidArgumentNameException(argumentID string) InvalidArgumentName {
	return InvalidArgumentName{ArgumentID: argumentID}
}

func (exception InvalidArgumentName) Error() string {
	return fmt.Sprintf("'%s' is not a valid argument name.", exception.ArgumentID)
}

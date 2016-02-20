package arguments

import "fmt"

/*
InvalidArgumentFormat describes an error where a parameter is in the wrong format
*/
type InvalidArgumentFormat struct {
	Parameter string
}

func InvalidArgumentFormatException(parameter string) InvalidArgumentFormat {
	return InvalidArgumentFormat{Parameter: parameter}
}

func (exception InvalidArgumentFormat) Error() string {
	return fmt.Sprintf("'%s' is not a valid argument format.", exception.Parameter)
}

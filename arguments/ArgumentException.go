package arguments

import "fmt"

type ArgumentExceptionCode int

const (
	OK ArgumentExceptionCode = iota
	ARGUMENT_COUNT_MISMATCH
	INVALID_ARGUMENT_FORMAT
	UNEXPECTED_ARGUMENT
	INVALID_ARGUMENT_NAME
	MISSING_STRING
	MISSING_INTEGER
	MISSING_FLOAT
	INVALID_INTEGER
	INVALID_FLOAT
	INVALID_BOOLEAN
)

type ArgumentException struct {
	ArgumentID string
	Code       ArgumentExceptionCode
	Parameter  string
}

func NewArgumentException(errorCode ArgumentExceptionCode, argumentID string, parameter string) ArgumentException {
	return ArgumentException{
		ArgumentID: argumentID,
		Code:       errorCode,
		Parameter:  parameter,
	}
}

func (exception ArgumentException) Error() string {
	switch exception.Code {
	case OK:
		return "TILT: Should not get here."

	case ARGUMENT_COUNT_MISMATCH:
		return "The number of arguments does not match the number of elements in the schema"

	case UNEXPECTED_ARGUMENT:
		return fmt.Sprintf("Argument -%s unexpected.", exception.ArgumentID)

	case MISSING_STRING:
		return fmt.Sprintf("Could not find string parameter for '-%s'.", exception.ArgumentID)

	case INVALID_INTEGER:
		return fmt.Sprintf("Argument -%s expects an integer but was '%s'.", exception.ArgumentID, exception.Parameter)

	case MISSING_INTEGER:
		return fmt.Sprintf("Could not find integer parameter for '-%s'.", exception.ArgumentID)

	case INVALID_FLOAT:
		return fmt.Sprintf("Argument -%s expects a float but was '%s'.", exception.ArgumentID, exception.Parameter)

	case MISSING_FLOAT:
		return fmt.Sprintf("Could not find float parameter for '-%s'.", exception.ArgumentID)

	case INVALID_ARGUMENT_NAME:
		return fmt.Sprintf("'%s' is not a valid argument name.", exception.ArgumentID)

	case INVALID_ARGUMENT_FORMAT:
		return fmt.Sprintf("'%s' is not a valid argument format.", exception.Parameter)

	case INVALID_BOOLEAN:
		return fmt.Sprintf("Argument -%s expects a boolean but was '%s'.", exception.ArgumentID, exception.Parameter)

	default:
		return ""
	}
}

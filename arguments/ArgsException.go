package arguments

import (
	"fmt"

	"github.com/adampresley/presentationCleanCodeArgsImplementation/arguments/errorCodes"
)

type ArgsException struct {
	ErrorArgumentID string
	ErrorCode       errorCodes.ErrorCode
	ErrorParameter  string
}

func NewArgsException(errorCode errorCodes.ErrorCode) ArgsException {
	return ArgsException{
		ErrorCode: errorCode,
	}
}

func NewArgsExceptionWithParameter(errorCode errorCodes.ErrorCode, parameter string) ArgsException {
	return ArgsException{
		ErrorCode:      errorCode,
		ErrorParameter: parameter,
	}
}

func NewArgsExceptionWithArgumentID(errorCode errorCodes.ErrorCode, argumentID string, parameter string) ArgsException {
	return ArgsException{
		ErrorArgumentID: argumentID,
		ErrorCode:       errorCode,
		ErrorParameter:  parameter,
	}
}

func (argsException ArgsException) Error() string {
	switch argsException.ErrorCode {
	case errorCodes.OK:
		return "TILT: Should not get here."

	case errorCodes.UNEXPECTED_ARGUMENT:
		return fmt.Sprintf("Argument -%s unexpected.", argsException.ErrorArgumentID)

	case errorCodes.MISSING_STRING:
		return fmt.Sprintf("Could not find string parameter for '-%s'.", argsException.ErrorArgumentID)

	case errorCodes.INVALID_INTEGER:
		return fmt.Sprintf("Argument -%s expects an integer but was '%s'.", argsException.ErrorArgumentID, argsException.ErrorParameter)

	case errorCodes.MISSING_INTEGER:
		return fmt.Sprintf("Could not find integer parameter for '-%s'.", argsException.ErrorArgumentID)

	case errorCodes.INVALID_DOUBLE:
		return fmt.Sprintf("Argument -%s expects a double but was '%s'.", argsException.ErrorArgumentID, argsException.ErrorParameter)

	case errorCodes.MISSING_DOUBLE:
		return fmt.Sprintf("Could not find double parameter for '-%s'.", argsException.ErrorArgumentID)

	case errorCodes.INVALID_ARGUMENT_NAME:
		return fmt.Sprintf("'%s' is not a valid argument name.", argsException.ErrorArgumentID)

	case errorCodes.INVALID_ARGUMENT_FORMAT:
		return fmt.Sprintf("'%s' is not a valid argument format.", argsException.ErrorParameter)

	default:
		return ""
	}
}

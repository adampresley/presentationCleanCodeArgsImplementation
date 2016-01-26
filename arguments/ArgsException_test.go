package arguments

import (
	"testing"

	"github.com/adampresley/presentationCleanCodeArgsImplementation/arguments/errorCodes"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArgsException(t *testing.T) {
	Convey("An ArgsException implements the error interface", t, func() {
		argsException := ArgsException{}
		implementsError := false

		_, implementsError = interface{}(argsException).(error)
		So(implementsError, ShouldEqual, true)
	})

	Convey("ArgsException constructors...", t, func() {
		Convey("NewArgsException creates a new ArgsException with the specified error code", func() {
			expected := ArgsException{
				ErrorArgumentID: "",
				ErrorCode:       errorCodes.OK,
				ErrorParameter:  "",
			}

			actual := NewArgsException(errorCodes.OK)
			So(actual, ShouldResemble, expected)
		})

		Convey("NewArgsExceptionWithParameter creates a new ArgsException with an error code and parameter name", func() {
			expected := ArgsException{
				ErrorArgumentID: "",
				ErrorCode:       errorCodes.OK,
				ErrorParameter:  "*",
			}

			actual := NewArgsExceptionWithParameter(errorCodes.OK, "*")
			So(actual, ShouldResemble, expected)
		})

		Convey("NewArgsExceptionWithArgumentID creates a new ArgsException with an error code and argument identifier and parameter name", func() {
			expected := ArgsException{
				ErrorArgumentID: "p",
				ErrorCode:       errorCodes.OK,
				ErrorParameter:  "*",
			}

			actual := NewArgsExceptionWithArgumentID(errorCodes.OK, "p", "*")
			So(actual, ShouldResemble, expected)
		})
	})

	Convey("ArgException error messages...", t, func() {
		Convey("An UNEXPECTED_ARGUMENT reports the ID of the unexpected argument", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.UNEXPECTED_ARGUMENT, "p", "*")
			expected := "Argument -p unexpected."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An MISSING_STRING reports the ID of the string parameter that could not be found", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.MISSING_STRING, "p", "*")
			expected := "Could not find string parameter for '-p'."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An INVALID_INTEGER reports that an ID is not of type integer", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.INVALID_INTEGER, "p", "*")
			expected := "Argument -p expects an integer but was '*'."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An MISSING_INTEGER reports that an ID of type integer could not be found", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.MISSING_INTEGER, "p", "#")
			expected := "Could not find integer parameter for '-p'."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An INVALID_DOUBLE reports that an ID is not of type double", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.INVALID_DOUBLE, "p", "*")
			expected := "Argument -p expects a double but was '*'."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An MISSING_DOUBLE reports that an ID of type double could not be found", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.MISSING_DOUBLE, "p", "##")
			expected := "Could not find double parameter for '-p'."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An INVALID_ARGUMENT_NAME will report the name of an invalid argument name", func() {
			argException := NewArgsExceptionWithArgumentID(errorCodes.INVALID_ARGUMENT_NAME, "1", "*")
			expected := "'1' is not a valid argument name."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("An INVALID_ARGUMENT_FORMAT will report an invalid format specifier", func() {
			argException := NewArgsExceptionWithParameter(errorCodes.INVALID_ARGUMENT_FORMAT, ")")
			expected := "')' is not a valid argument format."
			actual := argException.Error()

			So(actual, ShouldEqual, expected)
		})
	})
}

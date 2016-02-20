package arguments

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestArgumentExceptions(t *testing.T) {
	Convey("NewArgumentException satisfies the Error interface", t, func() {
		var implementsInterface bool
		argumentException := NewArgumentException(OK, "", "")

		_, implementsInterface = interface{}(argumentException).(error)
		So(implementsInterface, ShouldEqual, true)
	})

	Convey("Regarding error messages by code...", t, func() {
		Convey("OK should return \"TILT: Should not get here.\"", func() {
			argumentException := NewArgumentException(OK, "", "")

			expected := "TILT: Should not get here."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("ARGUMENT_COUNT_MISMATCH should return \"The number of arguments does not match the number of elements in the schema\"", func() {
			argumentException := NewArgumentException(ARGUMENT_COUNT_MISMATCH, "", "")

			expected := "The number of arguments does not match the number of elements in the schema"
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("UNEXPECTED_ARGUMENT should return \"Argument -a unexpected.\"", func() {
			argumentException := NewArgumentException(UNEXPECTED_ARGUMENT, "a", "")

			expected := "Argument -a unexpected."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("MISSING_STRING should return \"Could not find string parameter for '-a'.\"", func() {
			argumentException := NewArgumentException(MISSING_STRING, "a", "")

			expected := "Could not find string parameter for '-a'."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("INVALID_INTEGER should return \"Argument -a expects an integer but was 'a'.\"", func() {
			argumentException := NewArgumentException(INVALID_INTEGER, "a", "a")

			expected := "Argument -a expects an integer but was 'a'."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("MISSING_INTEGER should return \"Could not find integer parameter for '-a'.\"", func() {
			argumentException := NewArgumentException(MISSING_INTEGER, "a", "")

			expected := "Could not find integer parameter for '-a'."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("INVALID_FLOAT should return \"Argument -a expects a float but was 'a'.\"", func() {
			argumentException := NewArgumentException(INVALID_FLOAT, "a", "a")

			expected := "Argument -a expects a float but was 'a'."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("MISSING_FLOAT should return \"Could not find float parameter for '-a'.\"", func() {
			argumentException := NewArgumentException(MISSING_FLOAT, "a", "")

			expected := "Could not find float parameter for '-a'."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("INVALID_ARGUMENT_NAME should return \"'2' is not a valid argument name.\"", func() {
			argumentException := NewArgumentException(INVALID_ARGUMENT_NAME, "2", "")

			expected := "'2' is not a valid argument name."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("INVALID_ARGUMENT_FORMAT should return \"'^' is not a valid argument format.\"", func() {
			argumentException := NewArgumentException(INVALID_ARGUMENT_FORMAT, "", "^")

			expected := "'^' is not a valid argument format."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})

		Convey("INVALID_BOOLEAN should return \"Argument -a expects a boolean but was 'b'.\"", func() {
			argumentException := NewArgumentException(INVALID_BOOLEAN, "a", "b")

			expected := "Argument -a expects a boolean but was 'b'."
			actual := argumentException.Error()

			So(actual, ShouldEqual, expected)
		})
	})
}

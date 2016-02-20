package arguments

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBooleanArgumentMarshaler(t *testing.T) {
	Convey("BooleanArgumentMarshaler implements the ArgumentMarshaler interface", t, func() {
		var implementsInterface bool
		marshaler := NewBooleanArgumentMarshaler("", "")

		_, implementsInterface = interface{}(marshaler).(ArgumentMarshaler)
		So(implementsInterface, ShouldEqual, true)
	})

	Convey("The Value() function returns a boolean value", t, func() {
		marshaler := NewBooleanArgumentMarshaler("argumentID", "value")

		expected := false
		actual := marshaler.Value().(bool)

		So(actual, ShouldEqual, expected)
	})

	Convey("The Marshal() function...", t, func() {
		Convey("Returns an ArgumentException when the value cannot be marshaled to a boolean", func() {
			marshaler := NewBooleanArgumentMarshaler("argumentID", "a")
			var isArgumentException bool

			err := marshaler.Marshal()
			_, isArgumentException = interface{}(err).(ArgumentException)

			So(isArgumentException, ShouldEqual, true)
		})

		Convey("Returns nil when the value is marshaled successfully", func() {
			marshaler := NewBooleanArgumentMarshaler("argumentID", "1")

			actual := marshaler.Marshal()
			So(actual, ShouldBeNil)
		})

		Convey("Sets value to true when parsing a truthy value", func() {
			marshaler := NewBooleanArgumentMarshaler("argumentID", "1")
			marshaler.Marshal()

			expected := true
			actual := marshaler.Value().(bool)

			So(actual, ShouldEqual, expected)
		})
	})
}

package arguments

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerArgumentMarshaler(t *testing.T) {
	Convey("IntegerArgumentMarshaler implements the ArgumentMarshaler interface", t, func() {
		var implementsInterface bool
		marshaler := NewIntegerArgumentMarshaler("", "")

		_, implementsInterface = interface{}(marshaler).(ArgumentMarshaler)
		So(implementsInterface, ShouldEqual, true)
	})

	Convey("The Marshal() function...", t, func() {
		Convey("Returns an ArgumentException when the value cannot be marshaled to an integer", func() {
			marshaler := NewIntegerArgumentMarshaler("argumentID", "a")
			var isArgumentException bool

			err := marshaler.Marshal()
			_, isArgumentException = interface{}(err).(ArgumentException)

			So(isArgumentException, ShouldEqual, true)
		})

		Convey("Returns nil when the value is marshaled successfully", func() {
			marshaler := NewIntegerArgumentMarshaler("argumentID", "2")

			actual := marshaler.Marshal()
			So(actual, ShouldBeNil)
		})

		Convey("Sets value to 2 when parsed", func() {
			marshaler := NewIntegerArgumentMarshaler("argumentID", "2")
			marshaler.Marshal()

			expected := 2
			actual := marshaler.Value().(int)

			So(actual, ShouldEqual, expected)
		})
	})
}

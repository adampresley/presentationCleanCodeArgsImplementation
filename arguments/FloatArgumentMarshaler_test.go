package arguments

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFloatArgumentMarshaler(t *testing.T) {
	Convey("FloatArgumentMarshaler implements the ArgumentMarshaler interface", t, func() {
		var implementsInterface bool
		marshaler := NewFloatArgumentMarshaler("", "")

		_, implementsInterface = interface{}(marshaler).(ArgumentMarshaler)
		So(implementsInterface, ShouldEqual, true)
	})

	Convey("The Marshal() function...", t, func() {
		Convey("Returns an ArgumentException when the value cannot be marshaled to a float", func() {
			marshaler := NewFloatArgumentMarshaler("argumentID", "a")
			var isArgumentException bool

			err := marshaler.Marshal()
			_, isArgumentException = interface{}(err).(ArgumentException)

			So(isArgumentException, ShouldEqual, true)
		})

		Convey("Returns nil when the value is marshaled successfully", func() {
			marshaler := NewFloatArgumentMarshaler("argumentID", "1.1")

			actual := marshaler.Marshal()
			So(actual, ShouldBeNil)
		})

		Convey("Sets value to 1.1 when parsed", func() {
			marshaler := NewFloatArgumentMarshaler("argumentID", "1.1")
			marshaler.Marshal()

			expected := 1.1
			actual := marshaler.Value().(float64)

			So(actual, ShouldEqual, expected)
		})
	})
}

package arguments

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringArgumentMarshaler(t *testing.T) {
	Convey("StringArgumentMarshaler implements the ArgumentMarshaler interface", t, func() {
		var implementsInterface bool
		marshaler := NewStringArgumentMarshaler("", "")

		_, implementsInterface = interface{}(marshaler).(ArgumentMarshaler)
		So(implementsInterface, ShouldEqual, true)
	})

	Convey("The Marshal() function...", t, func() {
		Convey("Returns nil when the value is marshaled successfully", func() {
			marshaler := NewStringArgumentMarshaler("argumentID", "abc")

			actual := marshaler.Marshal()
			So(actual, ShouldBeNil)
		})

		Convey("Sets value to 'abc' when parsed", func() {
			marshaler := NewStringArgumentMarshaler("argumentID", "abc")
			marshaler.Marshal()

			expected := "abc"
			actual := marshaler.Value().(string)

			So(actual, ShouldEqual, expected)
		})
	})
}

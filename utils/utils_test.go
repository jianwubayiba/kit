package utils

import (
	"runtime"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToCamelCase(t *testing.T) {
	Convey("Test if ToCamelCase works", t, func() {
		So(ToCamelCase("test -to camel Case"), ShouldEqual, "TestToCamelCase")
	})
}
func TestToLowerFirstCamelCase(t *testing.T) {
	Convey("Test if ToLowerFirstCamelCase works", t, func() {
		So(ToLowerFirstCamelCase("test -to lowerCamel Case"), ShouldEqual, "testToLowerCamelCase")
	})
}

func TestToLowerSnakeCase(t *testing.T) {
	Convey("Test if ToLowerSnakeCase works", t, func() {
		So(ToLowerSnakeCase("test -to lowerCamel Case"), ShouldEqual, "test-to-lower-camel-case")
	})
}

func TestToUpperFirst(t *testing.T) {
	Convey("Test if ToUpperFirst works", t, func() {
		So(ToUpperFirst("test"), ShouldEqual, "Test")
	})
}

func TestIsExist(t *testing.T) {
	invalidP := "/@^*(*(&%^&"
	existedAbsoluteP := "c:"
	if runtime.GOOS != "windows" {
		existedAbsoluteP = "/etc"
	}
	Convey("Test if IsExist works", t, func() {
		So(IsExist(invalidP), ShouldEqual, false)
		So(IsExist(existedAbsoluteP), ShouldEqual, true)
	})
}

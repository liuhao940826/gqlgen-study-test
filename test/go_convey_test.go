package test

import (
	"github.com/smartystreets/goconvey/convey"
	"self_gqlgen/models"
	"self_gqlgen/service"
	"testing"
)

type testStruct struct {
	Name string
}

func TestTestName(t *testing.T) {
	convey.Convey("mock TestTestName", t, func() {
		// mock

		model := models.RegisterUser{Name: "liuhao", Age: 15}

		convey.So(service.PrintPlayerName(model), convey.ShouldBeNil)
	})
}

//func TestSpec(t *testing.T) {
//
//	// Only pass t into top-level Convey calls
//	convey.Convey("Given some integer with a starting value", t, func() {
//		x := 1
//
//		convey.Convey("When the integer is incremented", func() {
//			x++
//
//			convey.Convey("The value should be greater by one", func() {
//				convey.So(x, convey.ShouldEqual, 2)
//			})
//		})
//	})
//}

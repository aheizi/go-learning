package bdd

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	convey.Convey("Given 2 even numbers", t, func() {
		a := 3
		b := 4
		convey.Convey("When add the two numbers", func() {
			c := a + b
			convey.Convey("The result is still even", func() {
				convey.So(c%2, convey.ShouldEqual, 0)
			})
		})
	})
}

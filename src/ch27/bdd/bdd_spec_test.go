package bdd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	// only pass t into top-level convey calls
	Convey("given 2 even numbers", t, func() {
		a := 2
		b := 4

		Convey("when add the two numbers", func() {
			c := a + b

			Convey("then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}

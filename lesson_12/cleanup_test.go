package main

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	convey.Convey("Something works", t, func() {
		convey.So(1, convey.ShouldEqual, 1)
		convey.So(2*2, convey.ShouldEqual, 4)
		convey.Convey("More test", func() {
			convey.So("test", convey.ShouldEqual, "test")
			convey.So(10, convey.ShouldEqual, 10)
		})
	})
	x := 0
	convey.Convey("A", t, func() {
		x++
		convey.Convey("A-B", func() {
			x++
			convey.Convey("A-B-C1", func() {
				convey.So(x, convey.ShouldAlmostEqual, 1.9, .11)
			})
			//convey.Convey("A-B-C2", func() {
			//	convey.So([]int{1, 2, 3}, convey.,)
		})
	})
}

package main

import (
	"github.com/prashantv/gostub"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

var x int

func TestExec(t *testing.T) {
	convey.Convey("Test Exec Func.", t, func() {
		stubs := gostub.Stub(&x, 100)
		defer stubs.Reset()
		y := 50
		convey.So(x+y, convey.ShouldEqual, 150)
	})
}

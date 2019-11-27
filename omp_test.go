package omp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const vasTestAddr = "127.0.0.1:9390"

func TestOMPNew(t *testing.T) {
	Convey("new omp", t, func() {
		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		omp, err = New(vasTestAddr, "admin", "wrongpass")
		So(err, ShouldNotBeNil)
		So(omp, ShouldBeNil)
	})
}

package omp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuth(t *testing.T) {
	addr := "127.0.0.1:9390"

	Convey("create connector", t, func() {
		c, err := newConnector(addr)
		So(err, ShouldBeNil)

		Convey("authenticate with wrong password ", func() {
			err = c.Auth("admin", "superpassword41")
			So(err, ShouldNotBeNil)
		})

		Convey("authenticate with correct password", func() {
			err = c.Auth("admin", "securepassword41")
			So(err, ShouldBeNil)
		})
	})
}

func TestVersion(t *testing.T) {
	addr := "127.0.0.1:9390"

	Convey("create connector", t, func() {
		c, err := newConnector(addr)
		So(err, ShouldBeNil)

		Convey("get version ", func() {
			v, err := c.GetVersion()
			So(err, ShouldBeNil)
			So(v, ShouldEqual, "7.0")
		})
	})
}

package omp

import (
	"encoding/xml"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRequestAuth(t *testing.T) {
	Convey("construct auth request", t, func() {
		d := requestAuth{}
		d.Credentials.Username = "username"
		d.Credentials.Password = "password"

		Convey("marshal auth request", func() {
			b, err := xml.Marshal(d)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, "<authenticate><credentials><username>username</username><password>password</password></credentials></authenticate>")
		})
	})
}

func TestRequestVersion(t *testing.T) {
	Convey("construct version request", t, func() {
		d := requestVersion{}

		Convey("marshal version request", func() {
			b, err := xml.Marshal(d)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, "<get_version></get_version>")
		})
	})
}

package omp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCredential(t *testing.T) {
	Convey("new omp", t, func() {

		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		// create
		Convey("create credential", func() {

			credID, err := omp.CreateCredential(Credential{
				Name:     "credential test",
				Login:    "admin",
				Password: "superpassword41",
				Comment:  "test login ",
			})
			So(err, ShouldBeNil)
			fmt.Println("credential ID: ", credID)

			// get
			creds, err := omp.GetCredentials()
			So(err, ShouldBeNil)
			for i := range creds {
				fmt.Println(i, creds[i])
			}

		})

	})
}

package omp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("new omp", t, func() {

		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		Convey("get configs", func() {
			configs := []Config{}
			configs, err := omp.GetConfigs()
			So(err, ShouldBeNil)
			c := configs[1]

			for i := range configs {
				fmt.Println(configs[i])
			}

			Convey("create config", func() {
				id, err := omp.CreateConfig(c.ID, "testConfig1")
				So(err, ShouldBeNil)
				fmt.Println("config id:", id)
			})
		})
	})

}

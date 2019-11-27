package omp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTarget(t *testing.T) {
	Convey("new omp", t, func() {
		var id string
		var err error

		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		Convey("create target", func() {
			target := Target{
				Name:  "test target",
				Hosts: []string{"localhost"},
			}
			id, err = omp.CreateTarget(&target)
			So(err, ShouldBeNil)
			fmt.Println("target id:", id)

			Convey("modify target", func() {
				target := Target{
					ID:           id,
					Hosts:        []string{"localhost"},
					ExcludeHosts: []string{"127.0.0.1"},
				}
				err := omp.ModifyTarget(&target)
				So(err, ShouldBeNil)

				Convey("delete target", func() {
					err := omp.DeleteTarget(id)
					So(err, ShouldBeNil)
				})
			})
		})

	})
}

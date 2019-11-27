package omp

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScanner(t *testing.T) {
	Convey("new omp", t, func() {

		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		// Convey("create target", func() {
		// 	target := Target{
		// 		Name:  "test target",
		// 		Hosts: []string{"localhost"},
		// 	}
		// 	targetID, err := omp.CreateTarget(&target)
		// 	So(err, ShouldBeNil)
		// 	fmt.Println("target id:", id)
		// 	scannerID, err := omp.CreateScanner(&Scanner{
		// 		Name: "Default Scanner",
		// 		Host: "localhost",
		// 		Port: "9391",
		// 		Type: "2",
		// 	})
		// 	fmt.Println("scanner id: ", scannerID)
		// })
	})
}

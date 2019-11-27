package omp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResult(t *testing.T) {
	Convey("new omp", t, func() {
		var err error

		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		Convey("get result", func() {
			results, err := omp.GetResults(nil, "")
			So(err, ShouldBeNil)
			for i := range results {
				fmt.Println(results[i])
			}
		})
	})
}

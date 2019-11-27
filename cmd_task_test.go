package omp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTask(t *testing.T) {
	Convey("new omp", t, func() {
		omp, err := New(vasTestAddr, "admin", "securepassword41")
		So(err, ShouldBeNil)
		So(omp, ShouldNotBeNil)

		Convey("create task", func() {

			configs, err := omp.GetConfigs()
			So(err, ShouldBeNil)

			var targetID string
			targets, err := omp.GetTargets()
			So(err, ShouldBeNil)
			if len(targets) == 0 {
				targetID, err = omp.CreateTarget(&Target{
					Name:  "test target",
					Hosts: []string{"localhost"},
				})
				So(err, ShouldBeNil)
			} else {
				targetID = targets[0].ID
			}

			creds, err := omp.GetCredentials()
			So(err, ShouldBeNil)
			var credID string
			if len(creds) != 0 {
				credID = creds[0].ID
			} else {
				credID, err = omp.CreateCredential(Credential{
					Name:     "credential test",
					Login:    "admin",
					Password: "superpassword41",
					Comment:  "test login ",
				})
				So(err, ShouldBeNil)
			}

			scanners, err := omp.GetScanners()
			So(err, ShouldBeNil)
			var scannerID string
			if len(scanners) == 0 {
				scannerID, err = omp.CreateScanner(&Scanner{
					Name: "Default Scanner",
					Host: "localhost",
					Port: "9391",
					Type: "2",
				}, credID)
				So(err, ShouldBeNil)
			} else {
				scannerID = scanners[0].ID
			}

			configID := configs[0].ID

			task := Task{
				Name:    "test task",
				Comment: "test scan task",
				Config: &Config{
					ID: configID,
				},
				Target: &Target{
					ID: targetID,
				},
				Scanner: &Scanner{
					ID: scannerID,
				},
			}

			taskID, err := omp.CreateTask(&task)
			So(err, ShouldBeNil)
			fmt.Println("task id: ", taskID)

			Convey("modify task", func() {
				task := Task{
					ID:      taskID,
					Comment: "test new comment for modifying task",
				}

				err := omp.ModifyTask(&task)
				So(err, ShouldBeNil)
			})

			Convey("start task", func() {
				reportID, err := omp.StartTask(taskID)
				So(err, ShouldBeNil)
				fmt.Println("report id: ", reportID)

			})

			Convey("stop task", func() {
				err := omp.StopTask(&task)
				So(err, ShouldNotBeNil)
			})

			Convey("get task", func() {
				// tasks
				_, err := omp.GetTasks(taskID)
				//TODO print the content of tasks
				So(err, ShouldBeNil)
			})

			Convey("delete task", func() {
				err := omp.DeleteTask(taskID)
				So(err, ShouldBeNil)
			})
		})
	})
}

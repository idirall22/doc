package app_test

import (
	"fmt"
	"testing"

	app "github.com/idirall22/doc/application"
)

func TestMain(t *testing.T) {
	t.Run("Single Interview Application", testSingleInterviewApplication)
	// t.Run("Multi Interview Application", testMultiInterviewApplication)
	// t.Run("Schedule Interview Application", testScheduleInterviewApplication)
}

func testSingleInterviewApplication(t *testing.T) {

	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 1, false)
	fmt.Println(c.ListJobs()[0])
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
}

// func testMultiInterviewApplication(t *testing.T) {
// 	c := app.Candidate{}
// 	r := app.Recruiter{}
// 	appStore := app.NewApplicationStore()
// 	a := app.NewApplication(c, r, 2, false)
// 	appStore.Add(a)
// 	r.UpdateApplication(a, app.Accept)
// 	c.UpdateApplication(a, app.Accept)
// 	r.UpdateApplication(a, app.Accept)
// 	c.UpdateApplication(a, app.Accept)
// 	r.UpdateApplication(a, app.Accept)
// 	c.UpdateApplication(a, app.Accept)
// }

// func testScheduleInterviewApplication(t *testing.T) {
// 	c := app.Candidate{}
// 	r := app.Recruiter{}
// 	appStore := app.NewApplicationStore()
// 	a := app.NewApplication(c, r, 1, true)
// 	appStore.Add(a)
// 	r.UpdateApplication(a, app.Accept)
// 	c.Schedule(a, "2020-12-23", "2020-12-24", "2020-12-25")
// 	r.FixDate(a, 0)
// 	c.UpdateApplication(a, app.Accept)
// 	r.UpdateApplication(a, app.Accept)
// 	c.UpdateApplication(a, app.Accept)
// }

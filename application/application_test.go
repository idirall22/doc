package app_test

import (
	"testing"
	"time"

	app "github.com/idirall22/doc/application"
)

func TestMain(t *testing.T) {
	t.Run("Full Single Interview Application", testApplication)
}

func testApplication(t *testing.T) {
	c1 := app.NewJobConstraint(app.AfterApply, app.StartSchedule)
	c2 := app.NewJobConstraint(app.AfterInterview, app.StartInterview)
	job := app.NewJob(1, app.Recruiter{}, []*app.JobConstraint{c1, c2})
	a := app.NewApplication(&app.Candidate{}, job)
	a.SetState(app.Recruiter{}, app.Accept)
	a.ProposeDates(time.Now())
	a.RescheduleDate()
	a.ProposeDates(time.Now())
	a.FixDate(0)
	a.SetState(app.Candidate{}, app.Accept)
	a.SetState(app.Recruiter{}, app.Accept)
	a.SetState(app.Candidate{}, app.Accept)
	a.SetState(app.Recruiter{}, app.Accept)
	a.SetState(app.Candidate{}, app.Accept)
	a.Print()
}

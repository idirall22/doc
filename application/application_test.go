package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	app "github.com/idirall22/doc/application"
)

func TestMain(t *testing.T) {
	t.Run("Full Single Interview Application", testFullSingleInterviewApplication)
	t.Run("Reject Single Interview Application", testRejectSingleInterviewApplication)
	t.Run("Full Multi Interview Application", testFullMultiInterviewApplication)
	t.Run("Skip Multi Interview Application", testSkipMultiInterviewApplication)
	t.Run("Schedule Interview Application", testScheduleInterviewApplication)
}

func testFullSingleInterviewApplication(t *testing.T) {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 1, false)
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
	assert.Equal(t, len(a.ListSteps()), 6)
}

func testRejectSingleInterviewApplication(t *testing.T) {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 1, false)
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Reject)
	assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
	assert.Equal(t, len(a.ListSteps()), 3)
}

func testFullMultiInterviewApplication(t *testing.T) {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 2, false)
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
	assert.Equal(t, len(a.ListSteps()), 8)
}

func testSkipMultiInterviewApplication(t *testing.T) {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 2, false)
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Skip)
	c.UpdateApplication(a.ID, app.Accept)
	assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
	assert.Equal(t, len(a.ListSteps()), 7)
}

func testScheduleInterviewApplication(t *testing.T) {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 2, true)
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Accept)
	c.Schedule(a, []string{"2020-12-23", "2020-12-24", "2020-12-25"})
	r.UpdateApplication(a.ID, app.Reschedule)
	c.Schedule(a, []string{"2020-12-23", "2020-12-24", "2020-12-25"})
	r.FixDate(a, 0)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
	assert.Equal(t, len(a.ListSteps()), 12)
	// a.GetHistory()
}

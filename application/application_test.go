package app_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	app "github.com/idirall22/doc/application"
)

func TestMain(t *testing.T) {
	t.Run("Single Interview Application", testFullSingleInterviewApplication)
	t.Run("Multi Interview Application", testMultiInterviewApplication)
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

func testMultiInterviewApplication(t *testing.T) {
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
	// assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
}

func testScheduleInterviewApplication(t *testing.T) {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)
	r.CreateJob("Wanted a software developer", 2, true)
	a := c.Apply(1)
	r.UpdateApplication(a.ID, app.Accept)
	c.Schedule(a, []string{"2020-12-23", "2020-12-24", "2020-12-25"})
	r.FixDate(a, 0)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	r.UpdateApplication(a.ID, app.Accept)
	c.UpdateApplication(a.ID, app.Accept)
	assert.Equal(t, a.GetCurrentStep().GetStatus(), app.Closed)
	// a.GetHistory()
}

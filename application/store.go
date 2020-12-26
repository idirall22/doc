package app

import (
	"github.com/fatih/color"
)

// Store application store.
type Store struct {
	Jobs []*Job
}

// Add add application to datastore and return index
func (s *Store) Add(job *Job) int {
	id := len(s.Jobs) + 1
	job.ID = id
	s.Jobs = append(s.Jobs, job)
	return id
}

// List list jobs
func (s Store) List(u interface{}) []*Job {
	switch u.(type) {
	case Candidate:
		out := []*Job{}
		for _, job := range s.Jobs {
			if job.Open {
				out = append(out, job)
			}
		}
		return out
	}
	return s.Jobs
}

// Apply create a new application to a job.
func (s *Store) Apply(c *Candidate, jobID int) *Application {
	var app *Application
	for _, job := range s.Jobs {
		if job.ID == jobID {
			app = NewApplication(c, job)
			job.applications = append(job.applications, app)
		}
	}
	return app
}

// NewStore create a new Store.
func NewStore() *Store {
	return &Store{Jobs: []*Job{}}
}

// Job model.
type Job struct {
	ID           int
	r            *Recruiter
	applications []*Application
	Description  string
	Interview    uint8
	Schedule     bool
	Open         bool
}

// NewJob create job.
func NewJob(r *Recruiter, desc string, interview uint8, schedule bool) *Job {
	return &Job{
		r:            r,
		Description:  desc,
		Interview:    interview,
		Schedule:     schedule,
		applications: []*Application{},
		Open:         true,
	}
}

// ListApplication list all application for a specific job.
func (j Job) ListApplication() []*Application {
	return j.applications
}

// Print print job description.
func (j Job) Print() {
	status := "Open"
	if !j.Open {
		status = "Closed"
	}
	d := color.New(color.FgWhite, color.BgHiBlack)
	d.Printf(`--------------------------------------------------
Job ID: %d
Status: %s
Job description: %s
Interview step: %d
Schedule: %v
--------------------------------------------------
`, j.ID, status, j.Description, j.Interview, j.Schedule)
}

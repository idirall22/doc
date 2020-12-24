package app

import "fmt"

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
	fmt.Println("\nJob ID:", j.ID)
	fmt.Println("Job description:", j.Description)
	fmt.Println("---> How many Interview?", j.Interview)
	fmt.Println("---> Need Schedule?", j.Schedule)
	fmt.Println()
}

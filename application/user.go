package app

// User model.
type User struct{}

// Recruiter model.
type Recruiter struct {
	s *Store
}

// NewRecruiter create new Recruiter
func NewRecruiter(s *Store) *Recruiter {
	return &Recruiter{s: s}
}

// FixDate fix date.
func (r *Recruiter) FixDate(app *Application, dateIndex int) {
	app.fixDate = dateIndex
	app.setState(Recruiter{}, Fixdate)
}

// Reschedule request reschedule.
func (r *Recruiter) Reschedule(app *Application) {
	app.fixDate = 0
	app.reschedule = true
	app.setState(Recruiter{}, Reschedule)
}

// UpdateApplication update application process.
func (r *Recruiter) UpdateApplication(appID int, act Action) {
	for _, job := range r.s.Jobs {
		for _, app := range job.ListApplication() {
			if app.ID == appID {
				app.setState(Recruiter{}, act)
			}
		}
	}
}

// CreateJob create a new job.
func (r *Recruiter) CreateJob(desc string, interview uint8, schedule bool) bool {
	id := r.s.Add(NewJob(r, desc, interview, schedule))
	if id == 0 {
		return false
	}
	return true
}

// GetJobs all jobs created by the recruiter.
func (r *Recruiter) GetJobs() []*Job {
	jobs := []*Job{}
	for _, j := range r.s.Jobs {
		if j.r == r {
			jobs = append(jobs, j)
		}
	}
	return jobs
}

// Candidate model.
type Candidate struct {
	s *Store
}

// NewCandidate create new candidate
func NewCandidate(s *Store) *Candidate {
	return &Candidate{s: s}
}

// Schedule schedule some dates for interview.
func (c *Candidate) Schedule(app *Application, dates ...string) {
	app.scheduledDates = dates
	app.reschedule = false
	app.setState(Candidate{}, Submitdate)
}

// UpdateApplication update application process.
func (c *Candidate) UpdateApplication(appID int, act Action) {
	for _, job := range c.s.Jobs {
		for _, app := range job.ListApplication() {
			if app.ID == appID {
				app.setState(Candidate{}, act)
			}
		}
	}
}

// ListJobs list all jobs.
func (c Candidate) ListJobs() []*Job {
	return c.s.Jobs
}

// Apply for a job.
func (c *Candidate) Apply(jobID int) *Application {
	app := c.s.Apply(c, jobID)
	return app
}

// GetMyApplication return list of application ids.
func (c Candidate) GetMyApplication() []*Application {
	applications := []*Application{}
	for _, job := range c.s.Jobs {
		for _, app := range job.applications {
			applications = append(applications, app)
		}
	}
	return applications
}

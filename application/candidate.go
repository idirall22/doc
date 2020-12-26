package app

// NewCandidate create new candidate
func NewCandidate(s *Store) *Candidate {
	return &Candidate{s: s}
}

// Candidate model.
type Candidate struct {
	s *Store
}

// Schedule schedule some dates for interview.
func (c *Candidate) Schedule(app *Application, dates []string) {
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

// Apply for a job.
func (c *Candidate) Apply(jobID int) *Application {
	for _, job := range c.s.List(Candidate{}) {
		if job.ID == jobID {
			if job.Open {
				return c.s.Apply(c, jobID)
			} else {
				return nil
			}
		}
	}
	return nil
}

// ListApplication return list of application ids.
func (c Candidate) ListApplication() []*Application {
	applications := []*Application{}
	for _, job := range c.s.List(Candidate{}) {
		for _, app := range job.applications {
			applications = append(applications, app)
		}
	}
	return applications
}

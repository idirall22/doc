package app

// NewRecruiter create new Recruiter
func NewRecruiter(s *Store) *Recruiter {
	return &Recruiter{s: s}
}

// Recruiter model.
type Recruiter struct {
	s *Store
}

// FixDate fix date.
func (r *Recruiter) FixDate(app *Application, dateIndex int) {
	app.fixDate = dateIndex
	app.setState(Recruiter{}, Fixdate)
}

// UpdateApplication update application process.
func (r *Recruiter) UpdateApplication(appID int, act Action) {
	for _, job := range r.s.Jobs {
		for _, app := range job.ListApplication() {
			if app.ID == appID {
				app.setState(Recruiter{}, act)
				if act == Reschedule {
					app.fixDate = -1
					app.reschedule = true
				}
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

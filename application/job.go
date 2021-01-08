package app

// Job struct
type Job struct {
	id           int
	r            Recruiter
	applications []*Application
	conditions   []*JobConstraint
	open         bool
}

// JobConstraint struct
type JobConstraint struct {
	status ApplicationStatus
	next   ApplicationStatus
	done   bool
}

// NewJobConstraint new job constraint.
func NewJobConstraint(status, next ApplicationStatus) *JobConstraint {
	return &JobConstraint{
		status: status,
		next:   next,
		done:   false,
	}
}

// NewJob new job.
func NewJob(id int, r Recruiter, conn []*JobConstraint) *Job {
	return &Job{
		id:           id,
		r:            r,
		applications: []*Application{},
		conditions:   conn,
		open:         true,
	}
}

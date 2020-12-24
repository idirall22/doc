package app

// ApplicationStatus application process status.
type ApplicationStatus int

const (
	startApply ApplicationStatus = iota
	afterApply
	startSchedule
	afterSchedule
	startInterview
	afterInterview
	offer
	closed
)

var statusMap = map[ApplicationStatus]string{
	startApply:     "startApply ",
	afterApply:     "afterApply",
	startSchedule:  "startSchedule",
	afterSchedule:  "afterSchedule",
	startInterview: "startInterview",
	afterInterview: "afterInterview",
	offer:          "offer",
	closed:         "closed",
}

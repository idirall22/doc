package app

// ApplicationStatus application process status.
type ApplicationStatus int

const (
	StartApply ApplicationStatus = iota
	AfterApply
	StartSchedule
	AfterSchedule
	StartInterview
	AfterInterview
	Offer
	Closed
)

var statusMap = map[ApplicationStatus]string{
	StartApply:     "startApply ",
	AfterApply:     "afterApply",
	StartSchedule:  "startSchedule",
	AfterSchedule:  "afterSchedule",
	StartInterview: "startInterview",
	AfterInterview: "afterInterview",
	Offer:          "offer",
	Closed:         "closed",
}

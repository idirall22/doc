package app

// ApplicationStatus application process status.
type ApplicationStatus int

const (
	None ApplicationStatus = iota
	Any
	StartApply
	AfterApply
	AfterApplySchedule
	StartSchedule
	AfterSchedule
	StartInterview
	AfterInterview
	Offer
	Closed
)

var statusMap = map[ApplicationStatus]string{
	None:               "none",
	Any:                "any",
	StartApply:         "startApply ",
	AfterApply:         "afterApply",
	AfterApplySchedule: "afterApplySchedule",
	StartSchedule:      "startSchedule",
	AfterSchedule:      "afterSchedule",
	StartInterview:     "startInterview",
	AfterInterview:     "afterInterview",
	Offer:              "offer",
	Closed:             "closed",
}

var stringStatusMap = map[string]ApplicationStatus{
	"none":           None,
	"any":            Any,
	"startApply ":    StartApply,
	"afterApply":     AfterApply,
	"startSchedule":  StartSchedule,
	"afterSchedule":  AfterSchedule,
	"startInterview": StartInterview,
	"afterInterview": AfterInterview,
	"offer":          Offer,
	"closed":         Closed,
}

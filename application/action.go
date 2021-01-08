package app

// Action action that used to update application process.
type Action int

const (
	nothing Action = iota
	Apply
	Recive
	Accept
	Reject
	Reschedule
	Submitdate
	Fixdate
	Decline
	Cancel
	Skip
)

// ActionStringMap a map that allows to get action name.
var ActionStringMap = map[Action]string{
	nothing:    "nothing",
	Apply:      "apply",
	Recive:     "recive",
	Accept:     "accept",
	Reject:     "reject",
	Reschedule: "reschedule",
	Submitdate: "submitdate",
	Fixdate:    "fixdate",
	Decline:    "decline",
	Cancel:     "cancel",
	Skip:       "skip",
}

// StringActionMap get action int from string.
var StringActionMap = map[string]Action{
	"nothing":    nothing,
	"apply":      Apply,
	"recive":     Recive,
	"accept":     Accept,
	"reject":     Reject,
	"reschedule": Reschedule,
	"submitdate": Submitdate,
	"fixdate":    Fixdate,
	"decline":    Decline,
	"cancel":     Cancel,
	"skip":       Skip,
}

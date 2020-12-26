package app

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func newApplicationStep(
	desc string,
	status ApplicationStatus,
	recruiterAction, candidateAction Action,
	interviewID uint8,
	app *Application,
) *ApplicationStep {
	return &ApplicationStep{
		Description:     desc,
		status:          status,
		recruiterAction: recruiterAction,
		candidateAction: candidateAction,
		interviewID:     interviewID,
		app:             app,
	}
}

// ApplicationStep application step.
type ApplicationStep struct {
	Description     string
	app             *Application
	status          ApplicationStatus
	recruiterAction Action
	candidateAction Action
	interviewID     uint8
}

// GetStatus return application status.
func (a ApplicationStep) GetStatus() ApplicationStatus {
	return a.status
}

// PrintStep print current step data.
func (a ApplicationStep) PrintStep() {
	d := color.New(color.FgWhite, color.BgHiBlack)
	d.Printf(`
Desciption: %s
-->Recruiter Action: %s 
-->Candidate Action: %s 
`, a.formatDescription(), ActionIntStringMap[a.recruiterAction],
		ActionIntStringMap[a.candidateAction])
}

func (a ApplicationStep) formatDescription() string {
	rAct := strings.Join(a.app.ListActions(Recruiter{}, a), ",")
	cAct := strings.Join(a.app.ListActions(Candidate{}, a), ",")
	return fmt.Sprintf(a.Description, rAct, cAct)
}

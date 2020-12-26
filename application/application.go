package app

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// NewApplication create new application process.
func NewApplication(c *Candidate, job *Job) *Application {
	id := len(job.ListApplication()) + 1

	app := &Application{
		ID:         id,
		candidate:  c,
		steps:      []*ApplicationStep{},
		fixDate:    -1,
		job:        job,
		stepsCount: 1,
	}
	app.startApplyStep()
	app.afterApplyStep()
	return app
}

// Application application process.
type Application struct {
	ID               int
	candidate        *Candidate
	steps            []*ApplicationStep
	job              *Job
	currentInterview uint8
	reschedule       bool
	scheduledDates   []string
	fixDate          int
	stepsCount       uint
}

// GetScheduledDates get scheduled dates.
func (a Application) GetScheduledDates() []string {
	return a.scheduledDates
}

// GetCurrentStep get current step.
func (a Application) GetCurrentStep() *ApplicationStep {
	return a.steps[len(a.steps)-1]
}

// ListSteps list all application steps.
func (a Application) ListSteps() []*ApplicationStep {
	return a.steps
}

func (a *Application) setState(u interface{}, act Action) {
	// check if the application is finished
	if a.GetCurrentStep().status == Closed {
		fmt.Println("*** The Application process is closed ***")
		return
	}

	step := a.GetCurrentStep()
	if !a.validateAction(u, act, step.status, step.interviewID) {
		fmt.Println(ActionIntStringMap[act], "Action Not valid in the step:", statusMap[step.status])
		return
	}

	switch u.(type) {
	case Recruiter:
		step.recruiterAction = act
	case Candidate:
		step.candidateAction = act
	default:
		return
	}
	a.updateStep(a.GetCurrentStep())
}

// update a step.
func (a *Application) updateStep(step *ApplicationStep) {
	// check if the application process was closed by a user.
	if !a.checkApplicationProcess(step.recruiterAction, step.candidateAction) {
		a.closeAppStep()
		step = a.GetCurrentStep()
	} else if step.recruiterAction == Skip {
		a.offerStep()
		return
	}

	switch step.status {
	case StartApply:

	case AfterApply:
		if a.job.Schedule {
			a.startScheduleStep()
		} else {
			a.startInterviewStep()
		}
		break

	case StartSchedule:
		if len(a.scheduledDates) == 0 {
			fmt.Println("The candidate has to propose dates for interview.")
		} else {
			a.afterScheduleStep()
		}
		break

	case AfterSchedule:
		if a.fixDate == -1 || a.reschedule {
			fmt.Println("recruiter requiest a reschedule")
			a.startScheduleStep()
		} else {
			a.startInterviewStep()
		}
		break

	case StartInterview:
		a.afterInterviewStep()
		break

	case AfterInterview:
		a.currentInterview++
		if a.job.Interview == a.currentInterview {
			a.offerStep()
		} else {
			a.startInterviewStep()
		}
		break

	case Offer:
		a.closeAppStep()
		a.job.Open = false
		break
	}
}

// check if the application process was stoped.
func (a Application) checkApplicationProcess(cAct, rAct Action) bool {
	if cAct == Reject || cAct == Decline || cAct == Cancel ||
		rAct == Reject || rAct == Decline || rAct == Cancel {
		return false
	}
	return true
}

// ListActions list all actions.
func (a *Application) ListActions(u interface{}, step ApplicationStep) []string {
	acts := a.getActions(u, step.status, step.interviewID)
	out := []string{}
	for _, act := range acts {
		out = append(out, ActionIntStringMap[act])
	}
	return out
}

func (a *Application) startApplyStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The apply step, the candidate apply for a job.%s%s",
		StartApply, Recive, Apply, a.currentInterview, a,
	))
}
func (a *Application) afterApplyStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The after apply step, the recuiter has to [%s] the candidate can [%s].",
		AfterApply, nothing, Accept, a.currentInterview, a,
	))
}

func (a *Application) startScheduleStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The start schedule step, the recruiter can [%s] the interview. The candidate has to [%s] the interview.",
		StartSchedule, Accept, nothing, a.currentInterview, a,
	))
}

func (a *Application) afterScheduleStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The after schedule step, the recruiter can [%s] date for interview. The candidate can [%s] the process",
		AfterSchedule, nothing, Accept, a.currentInterview, a,
	))
}
func (a *Application) startInterviewStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The start interview step, the recuiter can [%s] the interview. The candidate can [%s] the interview",
		StartInterview, Accept, nothing, a.currentInterview, a,
	))
}

func (a *Application) afterInterviewStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The after interview step, the recruiter can [%s] the interview. The candidate can [%s] the interview.",
		AfterInterview, nothing, Accept, a.currentInterview, a,
	))
}

func (a *Application) offerStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The offer step, the recruiter can [%s]. The candidate can [%s] the offer.",
		Offer, Accept, nothing, a.currentInterview, a,
	))
}

func (a *Application) closeAppStep() {
	a.steps = append(a.steps, newApplicationStep(
		"The close application step, the application is colsed and archived.%s%s",
		Closed, nothing, nothing, a.currentInterview, a,
	))
}

func (a *Application) getActions(u interface{}, status ApplicationStatus, interviewID uint8) []Action {
	var actions []Action

	switch u.(type) {
	case Recruiter:
		switch status {
		case StartApply:
			actions = []Action{}
			break

		case AfterApply:
			actions = []Action{Accept, Reject}
			break

		case StartSchedule:
			actions = []Action{Cancel}
			break

		case AfterSchedule:
			actions = []Action{Cancel, Fixdate, Reschedule}
			break

		case StartInterview:
			actions = []Action{Cancel}
			break

		case AfterInterview:
			actions = []Action{Accept, Reject}
			break

		case Offer:
			actions = []Action{Cancel}
			break
		case Closed:
			actions = []Action{}
			break
		}
		if interviewID >= 1 && status != Offer {
			actions = append(actions, Skip)
		}
	case Candidate:
		switch status {
		case StartApply:
			actions = []Action{Apply}
			break

		case AfterApply:
			actions = []Action{Cancel}
			break

		case StartSchedule:
			actions = []Action{Submitdate, Cancel}
			break

		case AfterSchedule:
			actions = []Action{Cancel}
			break

		case StartInterview:
			actions = []Action{Accept, Reject}
			break

		case AfterInterview:
			actions = []Action{Decline}
			break

		case Offer:
			actions = []Action{Accept, Decline}
			break
		case Closed:
			actions = []Action{}
			break
		}
	}
	return actions
}

// validate user action related to the current step.
func (a *Application) validateAction(u interface{}, act Action, status ApplicationStatus, interviewID uint8) bool {
	switch u.(type) {
	case Recruiter:
		for _, a := range a.getActions(u, status, interviewID) {
			if a == act {
				return true
			}
		}
		break
	case Candidate:
		for _, a := range a.getActions(u, status, interviewID) {
			if a == act {
				return true
			}
		}
		break
	default:

		return false
	}
	return false
}

// PrintStep print step details.
func (a Application) PrintStep(step *ApplicationStep) {
	if step == nil {
		step = a.GetCurrentStep()
	}
	rAct := strings.Join(a.ListActions(Recruiter{}, *step), ",")
	cAct := strings.Join(a.ListActions(Candidate{}, *step), ",")
	fmt.Printf(step.Description, rAct, cAct)
}

// GetHistory display application history.
func (a Application) GetHistory() {
	a.job.Print()
	d := color.New(color.FgWhite, color.BgBlue)
	for index, step := range a.steps {
		d.Printf("Step %d:", index+1)
		step.PrintStep()
	}
}

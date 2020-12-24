package app

import (
	"fmt"
	"strings"
)

// NewApplication create new application process.
func NewApplication(c *Candidate, job *Job) *Application {
	id := len(job.ListApplication()) + 1
	app := &Application{
		ID:         id,
		c:          c,
		steps:      []ApplicationStep{},
		fixDate:    -1,
		job:        job,
		stepsCount: 1,
	}
	app.startApplyStep()
	app.afterApplyStep()
	// app.PrintDetails()
	return app
}

// ApplicationStep application step.
type ApplicationStep struct {
	Description     string
	status          ApplicationStatus
	recruiterAction Action
	candidateAction Action
}

// GetStatus return application status.
func (a ApplicationStep) GetStatus() ApplicationStatus {
	return a.status
}

// PrintStep print current step data.
func (a ApplicationStep) PrintStep() {
	fmt.Printf("\n")
	fmt.Println("Application Step:", statusMap[a.status])
	fmt.Println("-->Recruiter:", ActionIntStringMap[a.recruiterAction])
	fmt.Println("-->Candidate:", ActionIntStringMap[a.candidateAction])
}

// Application application process.
type Application struct {
	ID               int
	c                *Candidate
	steps            []ApplicationStep
	job              *Job
	currentInterview uint8
	reschedule       bool
	scheduledDates   []string
	fixDate          int
	stepsCount       uint
}

// GetCurrentStep get current step.
func (a Application) GetCurrentStep() *ApplicationStep {
	return &a.steps[len(a.steps)-1]
}

func (a *Application) setState(u interface{}, act Action) {
	// check if the application is finished
	if a.GetCurrentStep().status == closed {
		fmt.Println("*** The Application process is closed ***")
		return
	}

	step := a.GetCurrentStep()
	if !a.validateAction(u, act, step.status) {
		fmt.Println(ActionIntStringMap[act], "Action Not valid in the step:", statusMap[step.status])
		return
	}

	// update user choice.
	switch u.(type) {
	case Recruiter:
		step.recruiterAction = act
	case Candidate:
		step.candidateAction = act
	default:
		return
	}

	// update step.
	if step.recruiterAction != nothing &&
		step.candidateAction != nothing ||
		step.recruiterAction == Skip {
		a.updateStep(a.GetCurrentStep())
	}

	// print step infos.
	step.PrintStep()
	// a.PrintDetails()
}

// update a step.
func (a *Application) updateStep(step *ApplicationStep) {
	// check if the application process was closed by a user.
	if !a.checkApplicationProcess(step.recruiterAction, step.candidateAction) {
		a.closeAppStep()
	}
	// check if
	if step.recruiterAction == Skip {
		a.steps = append(a.steps, ApplicationStep{
			"The Offer step, the recruiter made an offer to the candidate",
			offer, Accept, nothing,
		})
		return
	}

	switch step.status {
	case startApply:

	case afterApply:
		if a.job.Schedule {
			a.startScheduleStep()
		} else {
			a.startInterviewStep()
		}
		break

	case startSchedule:
		if len(a.scheduledDates) == 0 {
			fmt.Println("The candidate has to propose dates for interview.")
		} else {
			a.afterScheduleStep()
		}
		break

	case afterSchedule:
		if a.fixDate == -1 || a.reschedule {
			fmt.Println("recruiter requiest a reschedule")
			a.startScheduleStep()
		} else {
			a.startInterviewStep()
		}
		break

	case startInterview:
		a.afterInterviewStep()
		break

	case afterInterview:
		a.currentInterview++
		if a.job.Interview == a.currentInterview {
			a.offerStep()
		} else {
			a.startInterviewStep()
		}
		break

	case offer:
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
func (a *Application) ListActions(u interface{}, status ApplicationStatus) []string {
	acts := a.getActions(u, status)
	out := []string{}
	for _, act := range acts {
		out = append(out, ActionIntStringMap[act])
	}
	return out
}

func (a *Application) startApplyStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The apply step, the candidate apply for a job.%s%s",
		startApply, Recive, Apply,
	})
}
func (a *Application) afterApplyStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The after apply step, the recuiter has to [%s] the candidate can [%s].",
		afterApply, nothing, Accept,
	})
}

func (a *Application) startScheduleStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The start schedule step, the recruiter can [%s] the interview. The candidate has to [%s] the interview.",
		startSchedule, Accept, nothing,
	})
}

func (a *Application) afterScheduleStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The after schedule step, the recruiter can [%s] date for interview. The candidate can [%s] the process",
		afterSchedule, nothing, Accept,
	})
}
func (a *Application) startInterviewStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The start interview step, the recuiter can [%s] the interview. The candidate can [%s] the interview",
		startInterview, Accept, nothing,
	})
}

func (a *Application) afterInterviewStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The after interview step, the recruiter can [%s] the interview. The candidate can [%s] the interview.",
		afterInterview, nothing, Accept,
	})
}

func (a *Application) offerStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The offer step, the recruiter can [%s]. The candidate can [%s] the offer.",
		offer, Accept, nothing,
	})
}

func (a *Application) closeAppStep() {
	a.steps = append(a.steps, ApplicationStep{
		"The close application step, the application is colsed and archived.%s%s",
		closed, nothing, nothing,
	})
}

func (a *Application) getActions(u interface{}, status ApplicationStatus) []Action {
	var actions []Action

	switch u.(type) {
	case Recruiter:
		switch status {
		case startApply:
			actions = []Action{}
			break

		case afterApply:
			actions = []Action{Accept, Reject}
			break

		case startSchedule:
			actions = []Action{Cancel}
			break

		case afterSchedule:
			actions = []Action{Cancel, Fixdate, Reschedule}
			break

		case startInterview:
			actions = []Action{Cancel}
			break

		case afterInterview:
			actions = []Action{Accept, Reject}
			break

		case offer:
			actions = []Action{Cancel}
			break
		case closed:
			actions = []Action{}
			break
		}
		if a.job.Interview > 1 && a.currentInterview >= 1 && a.GetCurrentStep().status != closed {
			actions = append(actions, Skip)
		}
	case Candidate:
		switch status {
		case startApply:
			actions = []Action{Apply}
			break

		case afterApply:
			actions = []Action{Cancel}
			break

		case startSchedule:
			actions = []Action{Submitdate, Cancel}
			break

		case afterSchedule:
			actions = []Action{Cancel}
			break

		case startInterview:
			actions = []Action{Accept, Reject}
			break

		case afterInterview:
			actions = []Action{Decline}
			break

		case offer:
			actions = []Action{Accept, Decline}
			break
		case closed:
			actions = []Action{}
			break
		}
	}
	return actions
}

// validate user action related to the current step.
func (a *Application) validateAction(u interface{}, act Action, status ApplicationStatus) bool {
	switch u.(type) {
	case Recruiter:
		for _, a := range a.getActions(u, status) {
			if a == act {
				return true
			}
		}
		break
	case Candidate:
		for _, a := range a.getActions(u, status) {
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
	rAct := strings.Join(a.ListActions(Recruiter{}, step.status), ",")
	cAct := strings.Join(a.ListActions(Candidate{}, step.status), ",")
	fmt.Printf(step.Description, rAct, cAct)
}

// GetHistory display application history.
func (a Application) GetHistory() {
	fmt.Println("Job ID:", a.job.ID)
	fmt.Println("Description:", a.job.Description)
	fmt.Println("History Application:")
	for _, step := range a.steps {
		fmt.Println()
		a.PrintStep(&step)
		step.PrintStep()
		fmt.Println()
	}
}

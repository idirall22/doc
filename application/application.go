package app

import (
	"fmt"
	"time"
)

// Application struct.
type Application struct {
	ID            int
	C             *Candidate
	Steps         []*ApplicationStep
	ProposedDates []*ScheduleDate
	Job           *Job
}

// ScheduleDate struct.
type ScheduleDate struct {
	active bool
	date   time.Time
}

// NewApplication create new application process.
func NewApplication(c *Candidate, job *Job) *Application {
	id := len(job.applications) + 1
	app := &Application{
		ID:            id,
		C:             c,
		Steps:         []*ApplicationStep{},
		ProposedDates: []*ScheduleDate{},
		Job:           job,
	}
	app.newStep(store[StartApply])
	app.newStep(store[AfterApply])
	return app
}

// GetCurrentStep get current step.
func (a Application) GetCurrentStep() *ApplicationStep {
	return a.Steps[len(a.Steps)-1]
}

// validateAction validate users actions
func (a *Application) validateAction(u interface{}, status ApplicationStatus, act Action) bool {
	step, ok := store[status]
	if ok {
		switch u.(type) {
		case Recruiter:
			for _, action := range step.recuirterAllowedActions {
				if action == act {
					return true
				}
			}
			break
		case Candidate:
			for _, action := range step.candidateAllowedActions {
				if action == act {
					return true
				}
			}
			break
		default:
			return false
		}
	}
	return false
}

// ProposeDates candidate propose meeting dates.
func (a *Application) ProposeDates(times ...time.Time) {
	dates := []*ScheduleDate{}
	for _, date := range times {
		dates = append(dates, &ScheduleDate{date: date})
	}
	a.ProposedDates = dates
	a.SetState(Candidate{}, Submitdate)
}

// FixDate recruiter fix the interview date.
func (a *Application) FixDate(index int) {
	a.ProposedDates[index].active = true
	a.SetState(Recruiter{}, Fixdate)
}

// RescheduleDate request reschedule
func (a *Application) RescheduleDate() {
	a.Job.conditions = append(a.Job.conditions, NewJobConstraint(AfterSchedule, StartSchedule))
	a.SetState(Recruiter{}, Reschedule)
}

// SetState set application state.
func (a *Application) SetState(u interface{}, act Action) {
	curStep := a.GetCurrentStep()

	// check if the application is finished
	if curStep.Step.name == Closed {
		fmt.Println("*** The Application process is closed ***")
		return
	}
	if !a.validateAction(u, curStep.Step.name, act) {
		fmt.Println(ActionStringMap[act], "Action Not valid in the step")
		return
	}

	switch u.(type) {
	case Recruiter:
		curStep.RecuirterAction = act
	case Candidate:
		curStep.CandidateAction = act
	default:
		return
	}

	a.addStep(curStep)
}

func (a Application) checkApplicationProcess(cAct, rAct Action) bool {
	if cAct == Reject || cAct == Decline || cAct == Cancel ||
		rAct == Reject || rAct == Decline || rAct == Cancel {
		return false
	}
	return true
}

// addStep update step.
func (a *Application) addStep(step *ApplicationStep) {
	next := a.getNextStepStatus(step)
	a.newStep(store[next])
}

func (a *Application) getNextStepStatus(step *ApplicationStep) ApplicationStatus {
	var next ApplicationStatus
	if !a.checkApplicationProcess(step.RecuirterAction, step.CandidateAction) {
		next = Closed

	} else if step.RecuirterAction == Skip {
		next = Offer

	} else {
		next = a.checkCondition(step)
		if next == None {
			next = step.Step.nextStep
		}
	}
	return next
}

func (a *Application) checkCondition(step *ApplicationStep) ApplicationStatus {
	for _, con := range a.Job.conditions {
		if con.status == step.Step.name && !con.done {
			con.done = true
			return con.next
		}
	}
	return None
}

func (a *Application) newStep(step Step) {
	a.Steps = append(a.Steps, newApplicationStep(
		step.desc,
		step.defaultRecuirterAction,
		step.defaultcandidateAction,
		a,
		store[step.name],
	))
}

// Print print application.
func (a Application) Print() {
	fmt.Println("-----------------------------")
	for _, step := range a.Steps {
		step.Print()
		fmt.Println("-----------------------------")
	}
}

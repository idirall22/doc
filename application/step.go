package app

import (
	"fmt"
	"strings"
)

// Step struct
type Step struct {
	name                    ApplicationStatus
	desc                    string
	nextStep                ApplicationStatus
	defaultRecuirterAction  Action
	defaultcandidateAction  Action
	recuirterAllowedActions []Action
	candidateAllowedActions []Action
}

// ApplicationStep struct.
type ApplicationStep struct {
	Desc            string
	RecuirterAction Action
	CandidateAction Action
	Step            Step
	App             *Application
}

// newApplicationStep new applicationStep.
func newApplicationStep(desc string, rAct, cAct Action,
	app *Application, step Step) *ApplicationStep {
	return &ApplicationStep{
		Desc:            desc,
		RecuirterAction: rAct,
		CandidateAction: cAct,
		Step:            step,
	}
}

// Print print step data.
func (a ApplicationStep) Print() {
	fmt.Println(statusMap[a.Step.name])
	fmt.Println("description:", a.getDescription())
	fmt.Println("recruiter:", ActionStringMap[a.RecuirterAction])
	fmt.Println("candidate:", ActionStringMap[a.CandidateAction])
}

func (a ApplicationStep) getDescription() string {
	s1 := []string{}
	for _, act := range a.Step.recuirterAllowedActions {
		s1 = append(s1, ActionStringMap[act])
	}

	s2 := []string{}
	for _, act := range a.Step.candidateAllowedActions {
		s2 = append(s2, ActionStringMap[act])
	}

	return fmt.Sprintf(a.Desc, strings.Join(s1, ","), strings.Join(s2, ","))
}

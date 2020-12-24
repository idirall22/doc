package app

import "fmt"

func newApplicationStep(
	desc string,
	status ApplicationStatus,
	recruiterAction, candidateAction Action,
	interviewID uint8,
) *ApplicationStep {
	return &ApplicationStep{
		Description:     desc,
		status:          status,
		recruiterAction: recruiterAction,
		candidateAction: candidateAction,
		interviewID:     interviewID,
	}
}

// ApplicationStep application step.
type ApplicationStep struct {
	Description     string
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
	fmt.Printf("\n")
	fmt.Println("Application Step:", statusMap[a.status])
	fmt.Println("-->Recruiter:", ActionIntStringMap[a.recruiterAction])
	fmt.Println("-->Candidate:", ActionIntStringMap[a.candidateAction])
}

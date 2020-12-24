package main

import (
	"fmt"
	"log"
	"strconv"

	app "github.com/idirall22/doc/application"
	"github.com/manifoldco/promptui"
)

func main() {

	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)

	// r.CreateJob("Wanted a software developer", 1, false)
	// fmt.Println(c.ListJobs()[0])
	// a := c.Apply(1)
	// r.UpdateApplication(a.ID, app.Accept)
	// c.UpdateApplication(a.ID, app.Accept)
	// r.UpdateApplication(a.ID, app.Accept)
	// c.UpdateApplication(a.ID, app.Accept)
	exit := false
	for !exit {
		prompt := promptui.Select{
			Label: "Users",
			Items: []string{"Recruiter", "Candidate", "Exit"},
		}
		_, user, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		ret := false
		switch user {
		case "Recruiter":
			for !ret {
				items := []string{
					"Post Job",
					"update Application",
					"Return",
				}
				if len(r.GetJobs()) > 0 {
					items = append(items[:len(items)-1], []string{"My Jobs", items[len(items)-1]}...)
				}

				prompt := promptui.Select{
					Label: "Users",
					Items: items,
				}

				_, actions, err := prompt.Run()
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}
				switch actions {
				case "Post Job":
					postJob(r)
				case "My Jobs":
					for _, job := range r.GetJobs() {
						job.Print()
					}
				case "update Application":
					for _, job := range r.GetJobs() {
						for _, a := range job.ListApplication() {
							a.GetHistory()
							items := a.ListActions(app.Recruiter{}, a.GetCurrentStep().GetStatus())
							prompt := promptui.Select{
								Label: "Users",
								Items: items,
							}
							_, action, err := prompt.Run()
							if err != nil {
								fmt.Printf("Prompt failed %v\n", err)
								return
							}
							r.UpdateApplication(a.ID, app.ActionStringActionMap[action])
						}
					}
				case "Return":
					ret = true
				}
			}
		case "Candidate":
			for !ret {
				items := []string{
					"List open Jobs",
					"Apply",
					"Return",
				}
				if len(c.GetMyApplication()) > 0 {
					items = append(items[:len(items)-1],
						[]string{"My Applications", items[len(items)-1]}...)
				}

				prompt := promptui.Select{
					Label: "Users",
					Items: items,
				}

				_, action, err := prompt.Run()
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}

				switch action {
				case "List open Jobs":
					for _, job := range c.ListJobs() {
						job.Print()
					}
				case "My Applications":
					// items := [][]string{}
					for _, a := range c.GetMyApplication() {
						a.GetHistory()
						// a.PrintDetails()
					}
					// items = append(items, a.ListActions(app.Candidate{}, a.GetCurrentStep().GetStatus()))

					// prompt.Label = "Update Application"
					// prompt.Items = appli.ListActions()

					break
				case "Apply":
					prompt := promptui.Prompt{
						Label:   "Job id(DEFAULT=1)",
						Default: "1",
						Pointer: func(in []rune) []rune { return []rune{'|'} },
					}
					jobID, err := prompt.Run()
					if err != nil {
						fmt.Printf("Prompt failed %v\n", err)
						return
					}
					a := c.Apply(parseInt(jobID))
					if a == nil {
						fmt.Println("Could not Apply to this job")
					}
				case "Return":
					ret = true
				}
			}
		case "Exit":
			exit = true
		}
	}
}

func postJob(r *app.Recruiter) {

	// scan for job description.
	prompt := promptui.Prompt{
		Label:   "Job description",
		Default: "React Software developer",
		Pointer: func(in []rune) []rune { return []rune{'|'} },
	}

	description, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	// scan for interviews steps.
	prompt.Label = "How Many Interviews? (DEFAULT=1)"
	prompt.Default = "1"
	interviewSteps, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// scan for schedule
	prompt.Label = "Need to fix Schedule? (DEFAULT=false)"
	prompt.Default = "false"
	schedule, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// create job
	ok := r.CreateJob(
		description,
		uint8(parseInt(interviewSteps)),
		parseBool(schedule),
	)
	if ok {
		fmt.Println("*** New Job created successfully :) ***")
	} else {
		fmt.Println("Error to create a job :(")
	}
}
func parseBool(in string) bool {
	out, err := strconv.ParseBool(in)
	if err != nil {
		log.Println("Error to parse bool value:", err)
		return false
	}
	return out
}

func parseInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		log.Println("Error to parse int value:", err)
		return 1
	}
	return out
}

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	app "github.com/idirall22/doc/application"
	"github.com/manifoldco/promptui"
)

func main() {
	store := app.NewStore()
	r := app.NewRecruiter(store)
	c := app.NewCandidate(store)

	exit := false
	for !exit {
		prompt := promptui.Select{
			Label:   "Users",
			Items:   []string{"Recruiter", "Candidate", "Exit"},
			Pointer: func(in []rune) []rune { return []rune{'|'} },
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
				action := recruiterChoice(store)
				switch action {

				case "Post Job":
					recruiterCreateJob(r)
					ret = true
					break

				case "My Jobs":
					for _, job := range store.List(nil) {
						job.Print()
					}
					break

				case "View/update Application":
					recruiterViewUpdateApplication(store, r)
					ret = true
					break

				case "Return":
					ret = true
					break
				}
			}
		case "Candidate":
			for !ret {
				action := candidateChoice(c)

				switch action {
				case "List open Jobs":
					candidateListOpenJobs(store)
					break

				case "My Applications":
					candidateApplication(c)
					ret = true
					break

				case "Apply":
					candidateApply(c)
					ret = true
					break

				case "Return":
					ret = true
					break
				}
			}
		case "Exit":
			exit = true
		}
	}
}

func recruiterChoice(store *app.Store) string {
	items := []string{
		"Post Job",
		"View/update Application",
		"Return",
	}
	if len(store.List(nil)) > 0 {
		items = append(items[:len(items)-1], []string{"My Jobs", items[len(items)-1]}...)
	}

	prompt := promptui.Select{
		Label: "Recruiter Actions:",
		Items: items,
	}

	_, actions, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return actions
}

func recruiterCreateJob(r *app.Recruiter) {
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

func recruiterViewUpdateApplication(store *app.Store, r *app.Recruiter) {
	for _, job := range store.List(nil) {
		for _, a := range job.ListApplication() {
			a.GetHistory()
			items := a.ListActions(app.Recruiter{}, *a.GetCurrentStep())
			items = append(items, "return")
			prompt := promptui.Select{
				Label: "Actions",
				Items: items,
			}
			// prompt.Label = "Actions"
			// prompt.Items = items

			_, action, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			if action != "return" {
				if action == app.ActionIntStringMap[app.Fixdate] {
					prompt := promptui.Prompt{
						Label:   "Date(DEFAULT=0)",
						Default: "0",
						Pointer: func(in []rune) []rune { return []rune{'|'} },
					}
					date, err := prompt.Run()
					if err != nil {
						fmt.Printf("Prompt failed %v\n", err)
						return
					}
					r.FixDate(a, parseInt(date))
				} else {
					r.UpdateApplication(a.ID, app.ActionStringActionMap[action])
				}
			}
		}
	}
}

func candidateApplication(c *app.Candidate) {
	for _, a := range c.ListApplication() {
		a.GetHistory()
		items := a.ListActions(app.Candidate{}, *a.GetCurrentStep())
		items = append(items, "return")
		prompt := promptui.Select{
			Label: "Actions",
			Items: items,
		}
		_, action, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if action != "return" {
			if action == app.ActionIntStringMap[app.Submitdate] {
				prompt := promptui.Prompt{
					Label:   "Dates(DEFAULT=01-01-2021,02-01-2021,03-01-2021)",
					Default: "01-01-2021,02-01-2021,03-01-2021",
					Pointer: func(in []rune) []rune { return []rune{'|'} },
				}
				dates, err := prompt.Run()
				if err != nil {
					fmt.Printf("Prompt failed %v\n", err)
					return
				}
				c.Schedule(a, strings.Split(dates, ","))
			} else {
				c.UpdateApplication(a.ID, app.ActionStringActionMap[action])
			}
		}
	}
}

func candidateApply(c *app.Candidate) {
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
}

func candidateListOpenJobs(store *app.Store) {
	for _, job := range store.List(app.Candidate{}) {
		job.Print()
	}
}

func candidateChoice(c *app.Candidate) string {
	items := []string{
		"List open Jobs",
		"Apply",
		"Return",
	}
	if len(c.ListApplication()) > 0 {
		items = append(items[:len(items)-1],
			[]string{"My Applications", items[len(items)-1]}...)
	}

	prompt := promptui.Select{
		Label: "Candidate Actions",
		Items: items,
	}

	_, action, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return action
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

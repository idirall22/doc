package main

func main() {}

// 	// create a store
// 	store := app.NewStore()
// 	// create a recuiter
// 	r := app.NewRecruiter(store)
// 	// create a candidate
// 	c := app.NewCandidate(store)
// 	// clear cmd console
// 	clearCMD()

// 	exit := false
// 	for !exit {
// 		user := chooseUser()
// 		ret := false
// 		clearCMD()

// 		// switch between Recruiter and the Candidate
// 		switch user {

// 		case "Recruiter":
// 			for !ret {
// 				action := recruiterChoice(store)
// 				switch action {

// 				case "Post a new job":
// 					recruiterCreateJob(r)
// 					break

// 				case "Job postings":
// 					// list all recuiter jobs.
// 					for _, job := range store.List(nil) {
// 						job.Print()
// 					}
// 					break

// 				case "Manage candidates":
// 					recruiterViewUpdateApplication(store, r)
// 					break

// 				case "Return":
// 					ret = true
// 					clearCMD()
// 					break
// 				}
// 			}
// 		case "Candidate":
// 			for !ret {
// 				action := candidateChoice(c)

// 				switch action {
// 				case "List open Jobs":
// 					candidateListOpenJobs(store)
// 					break

// 				case "My applications":
// 					candidateApplication(c)
// 					break

// 				case "Apply":
// 					candidateApply(store, c)
// 					break

// 				case "Return":
// 					ret = true
// 					clearCMD()
// 					break
// 				}
// 			}
// 		case "Exit":
// 			exit = true
// 		}
// 	}
// }

// func chooseUser() string {
// 	prompt := promptui.Select{
// 		Label:   "Users",
// 		Items:   []string{"Recruiter", "Candidate", "Exit"},
// 		Pointer: func(in []rune) []rune { return []rune{'|'} },
// 	}

// 	_, user, err := prompt.Run()
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	return user
// }
// func recruiterChoice(store *app.Store) string {
// 	items := []string{
// 		"Post a new job",
// 		"Manage candidates",
// 		"Return",
// 	}
// 	if len(store.List(nil)) > 0 {
// 		items = append(items[:len(items)-1], []string{"Job postings", items[len(items)-1]}...)
// 	}

// 	prompt := promptui.Select{
// 		Label: "Recruiter Actions",
// 		Items: items,
// 	}

// 	_, actions, err := prompt.Run()
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	return actions
// }

// // create a new job.
// func recruiterCreateJob(r *app.Recruiter) {
// 	// scan for job description.
// 	prompt := promptui.Prompt{
// 		Label:       "Job description",
// 		Default:     "React Software developer",
// 		Pointer:     func(in []rune) []rune { return []rune{'|'} },
// 		HideEntered: true,
// 	}

// 	description, err := prompt.Run()
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	prompt2 := promptui.Select{
// 		Label: "Rounds of Interviews",
// 		Items: []string{"1", "2", "3"},
// 	}
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	_, interviewSteps, err := prompt2.Run()

// 	prompt3 := promptui.Select{
// 		Label: "Appointment Scheduling?",
// 		Items: []string{"false", "true"},
// 	}

// 	_, schedule, err := prompt3.Run()
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	ok := r.CreateJob(
// 		description,
// 		uint8(parseInt(interviewSteps)),
// 		parseBool(schedule),
// 	)
// 	if ok {
// 		color.Green("***New Job created successfully :)***")
// 	} else {
// 		color.Red("Error, impossible to create a new job :(")
// 	}
// }

// // view and update current applications.
// func recruiterViewUpdateApplication(store *app.Store, r *app.Recruiter) {
// 	application := false
// 	for _, job := range store.List(nil) {
// 		for _, a := range job.ListApplication() {
// 			application = true
// 			a.GetHistory()
// 			items := a.ListActions(app.Recruiter{}, *a.GetCurrentStep())
// 			items = append(items, "return")
// 			prompt := promptui.Select{
// 				Label: "Actions",
// 				Items: items,
// 			}
// 			_, action, err := prompt.Run()
// 			if err != nil {
// 				log.Fatalf("Prompt failed %v\n", err)
// 			}

// 			if action != "return" {
// 				if action == app.ActionIntStringMap[app.Fixdate] {
// 					items := a.GetScheduledDates()
// 					prompt := promptui.Select{
// 						Label: "Choose an appointment date",
// 						Items: items,
// 					}
// 					_, action, err := prompt.Run()
// 					if err != nil {
// 						log.Fatalf("Prompt failed %v\n", err)
// 					}
// 					for index, date := range items {
// 						if action == date {
// 							r.FixDate(a, index)
// 						}
// 					}
// 				} else {
// 					r.UpdateApplication(a.ID, app.ActionStringActionMap[action])
// 				}
// 			}
// 		}
// 	}
// 	if !application {
// 		color.Red("No one has applied yet.")
// 	}
// }

// // list candidate application
// func candidateApplication(c *app.Candidate) {
// 	for _, a := range c.ListApplication(nil) {
// 		a.GetHistory()
// 		items := a.ListActions(app.Candidate{}, *a.GetCurrentStep())
// 		items = append(items, "return")
// 		prompt := promptui.Select{
// 			Label: "Actions",
// 			Items: items,
// 		}
// 		_, action, err := prompt.Run()
// 		if err != nil {
// 			log.Fatalf("Prompt failed %v\n", err)
// 		}
// 		if action != "return" {
// 			if action == app.ActionIntStringMap[app.Submitdate] {
// 				prompt := promptui.Prompt{
// 					Label:   "Dates(DEFAULT=01-01-2021,02-01-2021,03-01-2021)",
// 					Default: "01-01-2021,02-01-2021,03-01-2021",
// 					Pointer: func(in []rune) []rune { return []rune{'|'} },
// 				}
// 				dates, err := prompt.Run()
// 				if err != nil {
// 					log.Fatalf("Prompt failed %v\n", err)
// 				}
// 				c.Schedule(a, strings.Split(dates, ","))
// 			} else {
// 				c.UpdateApplication(a.ID, app.ActionStringActionMap[action])
// 			}
// 		}
// 	}
// }

// // Apply for a job
// func candidateApply(s *app.Store, c *app.Candidate) {
// 	jobs := s.List(app.Candidate{})
// 	items := []int{}
// 	for _, job := range jobs {
// 		items = append(items, job.ID)
// 	}
// 	prompt := promptui.Select{
// 		Label: "Apply to a job",
// 		Items: items,
// 	}
// 	_, jobID, err := prompt.Run()
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	a := c.Apply(parseInt(jobID))
// 	if a == nil {
// 		color.Red("Can not Apply for this job, the job was closed.")
// 	} else {
// 		color.Green("Application successful :)")
// 	}
// }

// // List open jobs
// func candidateListOpenJobs(store *app.Store) {
// 	for _, job := range store.List(app.Candidate{}) {
// 		job.Print()
// 	}
// }

// func candidateChoice(c *app.Candidate) string {

// 	items := []string{
// 		"List open Jobs",
// 		"Apply",
// 		"Return",
// 	}
// 	if len(c.ListApplication(nil)) > 0 {
// 		items = append(items[:len(items)-1],
// 			[]string{"My applications", items[len(items)-1]}...)
// 	}

// 	prompt := promptui.Select{
// 		Label: "Candidate Actions",
// 		Items: items,
// 	}

// 	_, action, err := prompt.Run()
// 	if err != nil {
// 		log.Fatalf("Prompt failed %v\n", err)
// 	}
// 	return action
// }

// func parseBool(in string) bool {
// 	out, err := strconv.ParseBool(in)
// 	if err != nil {
// 		// log.Println("Error to parse bool value:", err)
// 		return false
// 	}
// 	return out
// }

// func parseInt(in string) int {
// 	out, err := strconv.Atoi(in)
// 	if err != nil {
// 		// log.Println("Error to parse int value:", err)
// 		return 1
// 	}
// 	return out
// }

// func clearCMD() {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

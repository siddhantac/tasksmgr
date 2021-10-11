package main

import "log"

func main() {
	/*
		// get google tasks client
		taskClient, err := NewTasksClient()
		if err != nil {
			log.Fatal(err)
		}

		// get tasks list
		tasks, err := taskClient.ListTasks("Home", 10)
		if err != nil {
			log.Fatal(err)
		}
	*/

	// get notion client
	notionClient, err := NewNotionClient()
	if err != nil {
		log.Fatal(err)
	}

	tasks, err := notionClient.ListTasks(10)
	if err != nil {
		log.Fatal(err)
	}

	// debug
	PrintTasks(tasks)

	// get notion tasks/todos
	// perform diff
	// construct list of objects to be
	// - added to tasks
	// - deleted from tasks
	// - update in tasks
	// - added to Notion
	// - deleted from Notion
	// - updated in Notion
}

package main

import (
	"fmt"
	"time"

	"github.com/sidc9/gotion"
	"google.golang.org/api/tasks/v1"
)

type Task struct {
	ID        string
	Title     string
	Due       time.Time
	Deleted   bool
	Completed bool
}

func NewTaskFromGoogleTask(googleTask *tasks.Task) *Task {
	completed := googleTask.Completed != nil
	return &Task{
		ID:        googleTask.Id,
		Title:     googleTask.Title,
		Deleted:   googleTask.Deleted,
		Completed: completed,
	}
}

// TODO: date
func NewTaskFromNotionPage(page *gotion.Page) *Task {
	var completed bool
	if _, ok := page.Properties["Done"]; ok {
		completed = page.Properties["Done"].Checkbox
	}
	return &Task{
		ID:        page.ID,
		Title:     page.Title(),
		Completed: completed,
		// TODO
		// Due:       page.Properties["Date"],
	}
}

func PrintTasks(taskList []*Task) {
	for _, t := range taskList {
		fmt.Printf("\t* %s %q %t\n", t.Title, t.Due, t.Deleted)
	}
}

type TaskList []*Task

func (t TaskList) Diff(task TaskList) {

}

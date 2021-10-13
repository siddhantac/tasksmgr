package main

import (
	"testing"

	"github.com/matryer/is"
	"github.com/sidc9/gotion"
	"google.golang.org/api/tasks/v1"
)

func TestNewTaskFromGoogleTask(t *testing.T) {
	is := is.New(t)

	gt := &tasks.Task{
		Id:        "abc",
		Title:     "This is a test task",
		Deleted:   true,
		Completed: strptr("2021-09-10"),
	}

	got := NewTaskFromGoogleTask(gt)
	expected := &Task{
		ID:        "abc",
		Title:     "This is a test task",
		Deleted:   true,
		Completed: true,
	}

	is.Equal(got, expected)
}

func TestNewTaskFromNotionPage(t *testing.T) {
	is := is.New(t)

	page := &gotion.Page{
		ID: "abc",
		Properties: gotion.PageProperties{
			"title": &gotion.PageProperty{
				Type: "title",
				Title: []*gotion.RichText{
					{PlainText: "This is a test task"},
				},
			},
			"Done": &gotion.PageProperty{
				Checkbox: true,
			},
		},
	}

	got := NewTaskFromNotionPage(page)
	expected := &Task{
		ID:        "abc",
		Title:     "This is a test task",
		Deleted:   false,
		Completed: true,
	}
	is.Equal(got, expected)
}

func TestTaskDiff(t *testing.T) {

}

func strptr(s string) *string { return &s }

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/sidc9/gotion"
)

const databaseID = "539a391b9f83427f933518f5dc2b6c83"

type NotionClient struct {
	client *gotion.Client
}

func NewNotionClient() (*NotionClient, error) {
	apiKey, err := ioutil.ReadFile(".env")
	if err != nil {
		return nil, err
	}

	gotion.Init(string(apiKey), gotion.DefaultURL)
	c := gotion.GetClient()
	return &NotionClient{client: c}, nil
}

func (c *NotionClient) ListTasks(limit int) ([]*Task, error) {
	db, err := c.client.GetDatabase(databaseID)
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	pages, err := db.Query(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	taskList := make([]*Task, 0, len(pages.Results))
	for _, page := range pages.Results {
		taskList = append(taskList, NewTaskFromNotionPage(page))
		// fmt.Printf(" - [%t] %s %s %s\n", done, p.Title(), p.Properties["Date"], p.ID)
	}
	return taskList, nil
}

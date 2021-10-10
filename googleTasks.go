package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

type TasksClient struct {
	svc *tasks.Service
}

func NewTasksClient() (*TasksClient, error) {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file: %w", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %w", err)
	}
	client := getClient(config)

	srv, err := tasks.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve tasks Client %w", err)
	}

	return &TasksClient{svc: srv}, nil

}

func (c *TasksClient) ListTasks(limit int64) ([]*Task, error) {
	r, err := c.svc.Tasklists.List().MaxResults(10).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve task lists. %w", err)
	}

	if len(r.Items) <= 0 {
		return nil, fmt.Errorf("no list found")
	}

	var id string
	for _, i := range r.Items {
		if i.Title == "Home" {
			id = i.Id
			break
		}
	}

	googleTasks, err := c.svc.Tasks.List(id).MaxResults(limit).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}

	taskList := make([]*Task, 0, len(googleTasks.Items))
	for _, item := range googleTasks.Items {
		taskList = append(taskList, NewTask(item))
	}

	return taskList, nil
}

type Task struct {
	ID        string
	Title     string
	Due       time.Time
	Deleted   bool
	Completed bool
}

func NewTask(googleTask *tasks.Task) *Task {
	completed := googleTask.Completed != nil
	return &Task{
		ID:        googleTask.Id,
		Title:     googleTask.Title,
		Deleted:   googleTask.Deleted,
		Completed: completed,
	}
}

func PrintTasks(taskList []*Task) {
	for _, t := range taskList {
		fmt.Printf("\t* %s %q %t\n", t.Title, t.Due, t.Deleted)
	}
}

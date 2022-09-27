package todoist

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TodoistTask struct {
	Id          string   `json:"id"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Due         struct {
		Date     string `json:"date"`
		Datetime string `json:"datetime"`
	} `json:"due"`
}

type TodoistTasks []struct {
	TodoistTask
}

func (c *TodoistClient) GetTasks() (*TodoistTasks, error) {
	resp, err := c.get("https://api.todoist.com/rest/v2/tasks")
	if err != nil {
		return nil, err
	}

	var tasks *TodoistTasks = &TodoistTasks{}
	err = resIntoObj(resp, tasks)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return tasks, err
}

func (c *TodoistClient) GetTask(id string) (*TodoistTask, error) {
	resp, err := c.get("https://api.todoist.com/rest/v2/tasks/" + id)
	if err != nil {
		return nil, err
	}

	var task *TodoistTask = &TodoistTask{}
	err = resIntoObj(resp, task)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(task)

	return task, err
}

func (c *TodoistClient) CloseTask(id string) (resp *http.Response, err error) {
	// 6205926394
	resp, err = c.post("https://api.todoist.com/rest/v2/tasks/"+id+"/close", nil)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return
}

func (c *TodoistClient) ReopenTask() (resp *http.Response, err error) {
	// 6205926394
	resp, err = c.post("https://api.todoist.com/rest/v2/tasks/6205926394/reopen", nil)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	return
}

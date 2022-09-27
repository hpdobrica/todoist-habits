package todoist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type TodoistComment struct {
	Id      string `json:"id,omitempty"`
	Content string `json:"content"`
	TaskId  string `json:"task_id"`
}

type TodoistComments []struct {
	TodoistComment
}

func (c *TodoistClient) GetCommentsForTask(taskId string) (*TodoistComments, error) {
	resp, err := c.get("https://api.todoist.com/rest/v2/comments?task_id=" + taskId)
	if err != nil {
		return nil, err
	}

	var comments *TodoistComments = &TodoistComments{}
	err = resIntoObj(resp, comments)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return comments, err
}

func (c *TodoistClient) CreateComment(comment *TodoistComment) error {

	data, err := json.Marshal(comment)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(data))

	resp, err := c.post("https://api.todoist.com/rest/v2/comments", strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(resBody))

	return err
}

func (c *TodoistClient) UpdateComment(comment *TodoistComment) error {
	if comment.Id == "" {
		return errors.New("comment must have an id to be updated")
	}

	data, err := json.Marshal(comment)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(data))

	resp, err := c.post("https://api.todoist.com/rest/v2/comments/"+comment.Id, strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(resBody))

	return err
}

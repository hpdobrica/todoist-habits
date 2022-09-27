package main

import (
	"github.com/hpdobrica/todoist-habits/todoist"
)

func main() {

	habitsMetaTaskId := "6211149191"

	x := todoist.TodoistClient{}

	// read: 6162945790
	// gtd review: 5959317052
	// yoga: 6162705309
	// every hour test: 6208499463

	// x.GetTask("6162705309")
	// tasks, _ := x.GetTasks()

	// comments, _ := x.GetCommentsForTask(habitsMetaTaskId)

	// comment := &todoist.TodoistComment{
	// 	Content: "some content",
	// 	TaskId:  habitsMetaTaskId,
	// }
	// x.CreateComment(comment)

	newcomment := &todoist.TodoistComment{
		Content: "some content updated",
		TaskId:  habitsMetaTaskId,
		// Id:      "3144055204",
	}

	x.UpdateComment(newcomment)

	// fmt.Println((*comments)[0].Content)

	// x.ReopenTask()

}

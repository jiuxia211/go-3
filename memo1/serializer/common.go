package serializer

import "jiuxia/memo1/model"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
type User struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
type Task struct {
	Tid       uint   `json:"tid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int    `json:"status"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}
type TaskList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildTasks(items []model.Task) []Task {
	var tasks []Task
	for _, item := range items {
		task := Task{
			Tid:       item.ID,
			Title:     item.Title,
			Content:   item.Content,
			Status:    item.Status,
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
		}
		tasks = append(tasks, task)
	}
	return tasks
}

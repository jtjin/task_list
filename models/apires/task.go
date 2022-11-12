package apires

type ListTask struct {
	Result []Task `json:"result"`
}

type Task struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

package entity

type Task struct {
	ID      int64  `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

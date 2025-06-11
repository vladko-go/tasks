package tasksService

type TaskStatus string

const (
	TODO       TaskStatus = "TODO"
	DONE       TaskStatus = "DONE"
	INPROGRESS TaskStatus = "INPROGRESS"
)

type Task struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type TaskCreateRequest struct {
	Task string `json:"task"`
}

type TaskUpdateRequest struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

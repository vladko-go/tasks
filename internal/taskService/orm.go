package taskService

type TaskStatus string

const (
	TODO       TaskStatus = "TODO"
	DONE       TaskStatus = "DONE"
	INPROGRESS TaskStatus = "INPROGRESS"
)

type Task struct {
	ID     string     `gorm: "primaryKey" json:"id"`
	Title  string     `json:title"`
	Status TaskStatus `json:"status"`
}

type TaskCreateRequest struct {
	Title string `json:"title" validate:"required"`
}

type TaskUpdateRequest struct {
	Title  string     `json:"title" validate:"required"`
	Status TaskStatus `json:"status" validate:"oneof=TODO DONE INPROGRESS"`
}

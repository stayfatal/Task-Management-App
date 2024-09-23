package models

type CreateTaskRequest struct {
	Task Task
}

type CreateTaskResponse struct {
	Err error `json:"error,omitempty"`
}

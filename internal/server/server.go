package server

import "github.com/lgu-elo/task/internal/task"

type (
	ProjectHandler task.IHandler

	API struct {
		ProjectHandler
	}
)

func NewAPI(project task.IHandler) *API {
	return &API{project}
}

package Server

import (
	repo "github.com/AlexeyChudov/todoApp/internal/Repository"
	taskSvc "github.com/AlexeyChudov/todoApp/internal/service/Tasks"

	"net/http"
)

type Server struct {
	server *http.Server
	api    taskSvc.TaskService
	repo   repo.TaskRepository
}

func NewServer(api taskSvc.TaskService, repo repo.TaskRepository) {
}

func (s *Server) Run() {

}

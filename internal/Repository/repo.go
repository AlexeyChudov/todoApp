package Repository

import (
	"context"
	taskSvc "github.com/AlexeyChudov/todoApp/pkg/TaskService"
)

type repository struct {
}

type TaskRepository interface {
	List(ctx context.Context, f taskSvc.TaskFilter)
	Create(ctx context.Context, t *taskSvc.Task)
	GetById(ctx context.Context, id int)
	Update(ctx context.Context, id int)
	Delete(ctx context.Context, id int)
}

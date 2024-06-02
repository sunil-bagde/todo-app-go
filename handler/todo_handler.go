package handler

import (
    "sync"
    "github.com/sunil-bagde/go-todo-app/service"
    "github.com/sunil-bagde/go-todo-app/types"
  //  "github.com/sunil-bagde/go-todo-app/utils"
    "github.com/sirupsen/logrus"
)
/**
 * TodoHandler struct
 */
type TodoHandler struct {
    logger *logrus.Logger
    todoService service.TodoService
}

/**
 * NewTodoHandler creates a new TodoHandler
 * @param logger *logrus.Logger
 * @param todoService service.TodoService
 * @return *TodoHandler
 */
func NewTodoHandler(logger *logrus.Logger, todoService service.TodoService) *TodoHandler {
    return &TodoHandler{logger: logger, todoService: todoService}
}

/**
 * FetchTodos fetches todos from the service
 * @param count int
 * @return []*types.Todo
 * @return error
 */
func (h *TodoHandler) FetchTodos(count int) ([]*types.Todo, error){
    var wg sync.WaitGroup
    ch := make(chan *types.Todo)
	errCh := make(chan error, count)


	// Fetching todos concurrently
    for i := 2; i <= count*2; i+=2 {
        wg.Add(1)
		go h.todoService.FetchTodo(i, &wg, ch, errCh)
    }

	// Closing the channel once all the todos are fetched
	go func() {
        wg.Wait()
        close(ch)
        close(errCh)
    }()

	var todos []*types.Todo
    for todo := range ch {
        todos = append(todos, todo)
    }

    for err := range errCh {
        if err != nil {
            return nil, err
        }
    }

    return todos, nil
}

package service

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/sunil-bagde/go-todo-app/config"
	"github.com/sunil-bagde/go-todo-app/types"
	utils "github.com/sunil-bagde/go-todo-app/utils"
)

/**
 * TodoService interface
 */

type TodoService interface {
	FetchTodo(id int, wg *sync.WaitGroup, ch chan<- *types.Todo, errCh chan<- error)
}

/**
 * TodoServiceImpl struct
 */
type TodoServiceImpl struct {
	config *config.Config
	logger *logrus.Logger
}

/**
 * NewTodoService creates a new TodoService
 * @param config *config.Config
 * @param logger *logrus.Logger
 * @return TodoService
 */
func NewTodoService(config *config.Config, logger *logrus.Logger) TodoService {
	return &TodoServiceImpl{config: config, logger: logger}
}

/**
 * FetchTodo fetches a todo from the API
 * @param id int
 * @param wg *sync.WaitGroup
 * @param ch chan<- *types.Todo
 * @param errCh chan<- error
 */
func (s *TodoServiceImpl) FetchTodo(id int, wg *sync.WaitGroup, ch chan<- *types.Todo, errCh chan<- error) {
	defer wg.Done()
	var todo types.Todo
	url := fmt.Sprintf("%s%d", s.config.ApiUrl, id)
	err := utils.FetchJSON(url, &todo)
	if err != nil {
		errCh <- err
		return
	}
	ch <- &todo
}

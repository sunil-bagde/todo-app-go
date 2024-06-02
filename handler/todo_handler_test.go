package handler

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sunil-bagde/go-todo-app/config"
	"github.com/sunil-bagde/go-todo-app/service"
)

func TestFetchTodos(t *testing.T) {
	logger := logrus.New()
	todoService := service.NewTodoService(config.NewConfig(), logger)
	handler := NewTodoHandler(logger, todoService)

	count := 5
	todos, err := handler.FetchTodos(count)

	if err != nil {
		t.Errorf("FetchTodos returned an error: %v", err)
	}

	if len(todos) != count {
		t.Errorf("FetchTodos returned %d todos, expected %d", len(todos), count)
	}
}

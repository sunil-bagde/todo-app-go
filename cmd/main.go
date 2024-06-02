package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/sunil-bagde/go-todo-app/config"
	"github.com/sunil-bagde/go-todo-app/handler"
	"github.com/sunil-bagde/go-todo-app/service"
	"github.com/sunil-bagde/go-todo-app/utils"
	"go.uber.org/fx"
)

/**
 * NewTodoService function
 * @param config *config.Config
 * @param logger *logrus.Logger
 * @return service.TodoService
 */

func NewTodoService(config *config.Config, logger *logrus.Logger) service.TodoService {
	return service.NewTodoService(config, logger)
}

/**
 * rootCmd
 */
var rootCmd = &cobra.Command{
	Use:   "todo-app",
	Short: "Fetches TODOs",
	Long:  `Fetches TODOs from JSONPlaceholder API and prints them`,
	Run: func(cmd *cobra.Command, args []string) {

		color.Cyan("Fetching TODOs...")
		app := fx.New(
			fx.Provide(
				config.NewConfig,
				logrus.New,
				NewTodoService,
				handler.NewTodoHandler,
			),
			fx.Invoke(run),
			fx.Logger(utils.NoOpLogger{}),
		)
		if err := app.Start(context.Background()); err != nil {
			log.Fatalf("Failed to start application: %v", err)
		}

		go func() {
			<-app.Done()

			if err := app.Stop(context.Background()); err != nil {
				log.Fatalf("Failed to stop application: %v", err)
			}
		}()
		//  here

		if err := app.Stop(context.Background()); err != nil {
			log.Fatalf("Failed to stop application: %v", err)
		}
	},
}

/**
 * run function
 * @param lc fx.Lifecycle
 * @param config *config.Config
 * @param logger *logrus.Logger
 * @param todoHandler *handler.TodoHandler
 * @return error
 */
func run(lc fx.Lifecycle, config *config.Config, logger *logrus.Logger, todoHandler *handler.TodoHandler) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			count := config.TodoCount
			if count <= 0 {
				logger.Error("Cannot fetch  TODOs")
				return nil
			}
			if count > 20 {
				logger.Error("Cannot fetch more than 20 TODOs")
				return nil
			}

			todos, err := todoHandler.FetchTodos(count)
			if err != nil {
				logger.Error(err)
				color.Red(err.Error())
				return nil
			}
			for _, todo := range todos {
				if todo.Completed {
					color.Green("Title: %s, Completed: %v\n", todo.Title, todo.Completed)
				} else {
					color.Yellow("Title: %s, Completed: %v\n", todo.Title, todo.Completed)
				}
			}
			fmt.Println()
			color.Cyan("Fetching TODOs Completed.")
			return nil
		},
	})
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

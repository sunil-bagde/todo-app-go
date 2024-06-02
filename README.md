# Todo app

## Description

Todo App built in Go.

## s/w requirements

make 3.81 [Makefile](https://www.gnu.org/software/make/)

go 1.22.3 [Go](https://go.dev/doc/install)

## Installation

To get a local copy up and running follow these simple steps.

```bash
# Clone the repository
git clone https://github.com/sunil-bagde/todo-app-go.git

or use SSH

git clone git@github.com:sunil-bagde/todo-app-go.git

# Navigate to the project directory
cd todo-app-go

# Install Go dependencies
go mod download

```

cp .env.example .env

Compile Run Test same time

```sh
make all
```

Compiles your Go code and creates an executable named todo-app.

```sh
make build or go build -o todo-app cmd/main.go
```

Runs all the tests in your project

```sh
make test   or go test -v ./...
```

Runs your Go application.

```sh
make run   or go run cmd/main.go run
```

Removes the todo-app executable.`

```sh
make clean or rm -f ./todo-app
```

```

```

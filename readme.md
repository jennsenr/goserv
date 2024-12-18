## goserv

A simple and easy to use HTTP server for Go.

## Features

- Simple and easy to use
- Supports global middlewares
- Supports route middlewares
- Supports group middlewares
- Supports group routes
- Supports context values
- Supports path parameters
- Supports JSON binding
- Supports custom error responses

## Installation

```bash
go get github.com/jennsenr/goserv
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/jennsenr/goserv"
	"log"
	"math/rand"
	"strconv"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUser struct {
	Name string `json:"name"`
}

type UpdateUser struct {
	Name string `json:"name"`
}

var (
	users = map[string]*User{}
)

func main() {
	s := goserv.New("/test", "8080")

	s.Use(GlobalMiddleware(), LoggerMiddleware())

	s.GET("/health", handleHealthCheck(), RouteMiddleware())

	u := s.Group("/users", GroupMiddleware())
	u.POST("/", createUser)
	u.GET("/{userID}", findUser)
	u.PUT("/{userID}", updateUser)

	fmt.Printf("running on %v\n", "8080")

	err := s.Start()
	if err != nil {
		log.Panicf("listen and serve: %v", err)
	}
}

func RouteMiddleware() func(req *goserv.Request) *goserv.Response {
	return func(req *goserv.Request) *goserv.Response {
		fmt.Println("route middleware")
		return nil
	}
}

func GroupMiddleware() func(req *goserv.Request) *goserv.Response {
	return func(req *goserv.Request) *goserv.Response {
		fmt.Println("group middleware")
		return nil
	}
}

func LoggerMiddleware() func(req *goserv.Request) *goserv.Response {
	return func(req *goserv.Request) *goserv.Response {
		fmt.Printf("got request on path %v\n", req.Path())
		return nil
	}
}

func GlobalMiddleware() func(req *goserv.Request) *goserv.Response {
	return func(req *goserv.Request) *goserv.Response {
		fmt.Println("global middleware")
		return nil
	}
}

func handleHealthCheck() func(req *goserv.Request) goserv.Response {
	return func(req *goserv.Request) goserv.Response {
		return goserv.NewEmptyResponse()
	}
}

func updateUser(req *goserv.Request) goserv.Response {
	userID := req.GetPathParam("userID")

	user, ok := users[userID]
	if !ok {
		return goserv.NewNotFoundResponse()
	}

	payload := new(UpdateUser)

	err := req.Bind(payload)
	if err != nil {
		return goserv.NewBadRequestResponse()
	}

	user.Name = payload.Name

	return goserv.NewDataResponse(user)
}

func createUser(req *goserv.Request) goserv.Response {
	payload := new(CreateUser)

	err := req.Bind(payload)
	if err != nil {
		return goserv.NewBadRequestResponse()
	}

	user := User{
		ID:   strconv.Itoa(rand.Intn(1000)),
		Name: payload.Name,
	}

	users[user.ID] = &user

	return goserv.NewDataResponse(user)
}

func findUser(req *goserv.Request) goserv.Response {
	userID := req.GetPathParam("userID")

	user, ok := users[userID]
	if !ok {
		return goserv.NewNotFoundResponse()
	}

	return goserv.NewDataResponse(user)
}
```

## License

MIT

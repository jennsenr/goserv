package main

import (
	"fmt"
	"github.com/jennsenr/goserv"
	"log"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateUser struct {
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

	user := users[userID]

	if user == nil {
		return goserv.NewNotFoundResponse()
	}

	payload := new(CreateUser)

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
		ID:   "123",
		Name: payload.Name,
	}

	users[user.ID] = &user

	return goserv.NewDataResponse(user)
}

func findUser(req *goserv.Request) goserv.Response {
	userID := req.GetPathParam("userID")

	user := users[userID]

	if user == nil {
		return goserv.NewNotFoundResponse()
	}

	return goserv.NewDataResponse(user)
}

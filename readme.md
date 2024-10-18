## goserv

A simple and easy to use HTTP server for Go.

## Features

- Simple and easy to use
- Supports global middlewares
- Supports route middlewares
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
	"github.com/jennsenr/goserv"
	"log"
)


func main() {
	s := goserv.New("/users", "8080")

	s.Use(GlobalMiddleware(), LoggerMiddleware())

	s.GET("/health", handleHealthCheck(), RouteMiddleware())
	s.POST("/", createUser)
	s.GET("/{userID}", findUser)
	s.PUT("/{userID}", updateUser)
	
	err := s.Start()
	if err != nil {
		log.Panicf("listen and serve: %v", err)
	}
}   
```

## License

MIT

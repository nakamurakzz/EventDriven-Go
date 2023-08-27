package main

import (
	"time"

	"github.com/nakamurakzz/event-driven-go/triggers"
)

func main() {
	userCreateRequest := triggers.UserCreateRequest{
		Email: "test@example.com",
	}
	triggers.CreateUser(userCreateRequest)
	time.Sleep(3 * time.Second)
}

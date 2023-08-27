package triggers

import "github.com/nakamurakzz/event-driven-go/events"

type UserCreateRequest struct {
	Email string
}

func CreateUser(userCreateRequest UserCreateRequest) {
	events.UserCreated.Trigger(events.UserCreatedPayload{
		Email: userCreateRequest.Email,
	})
}

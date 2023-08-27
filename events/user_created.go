package events

import "time"

var UserCreated userCreated

type UserCreatedPayload struct {
	Email string
	Time  time.Time
}

type userCreated struct {
	handlers []interface{ Handle(UserCreatedPayload) }
}

func (u *userCreated) Register(handler interface{ Handle(UserCreatedPayload) }) {
	u.handlers = append(u.handlers, handler)
}

func (u *userCreated) Trigger(payload UserCreatedPayload) {
	for _, handler := range u.handlers {
		go handler.Handle(payload)
	}
}

package main

import (
	"log"

	"github.com/nakamurakzz/event-driven-go/events"
)

func init() {
	createNotifier := userCreatedNotifier{
		adminEmail: "test@example.com",
	}

	events.UserCreated.Register(createNotifier)
}

type userCreatedNotifier struct {
	adminEmail string
}

func (u userCreatedNotifier) Notify(payload events.UserCreatedPayload) {
	// Send an email to the admin
	log.Printf("Sending email to %s\n", u.adminEmail)
}

func (u userCreatedNotifier) Handle(payload events.UserCreatedPayload) {
	u.Notify(payload)
}

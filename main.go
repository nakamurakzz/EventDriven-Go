package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/nakamurakzz/event-driven-go/component"
	"github.com/nakamurakzz/event-driven-go/pubsub"
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx := context.Background()
	h := component.NewHub()
	components := component.InitializeComponent()

	sigctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	pubsub.Exec()

	go func() {
		for i := range components {
			c := components[i]

			c.Register(&h)
			h.Register(c)

			go func() {
				err := c.Start()
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}()

	<-sigctx.Done()
	return 0
}

package main

import (
	"log"

	"github.com/nakamurakzz/event-driven-go/hub"
	"github.com/nakamurakzz/event-driven-go/observer"
	"golang.org/x/sync/errgroup"
)

func main() {
	h := hub.NewHub()

	observers := initializeObserver()
	for _, o := range observers {
		o.Register(&h)
		h.Register(o)
	}

	g := errgroup.Group{}
	for _, observer := range observers {
		o := observer // capture loop variable
		g.Go(func() error {
			return o.Start()
		})
	}

	err := g.Wait()
	if err != nil {
		log.Println(err)
	}
}

func initializeObserver() []hub.Observer {
	return []hub.Observer{
		observer.NewEnvBackObserver(),
		observer.NewLightBackObserver(),
		observer.NewEnvFrontObserver(),
		observer.NewLightFrontObserver(),
	}
}

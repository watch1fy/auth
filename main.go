package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/watch1fy/auth/application"
)

func main() {
	r := chi.NewRouter()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app := application.NewRouter(r)
	err := app.Start(ctx)

	if err != nil {
		fmt.Printf("Cannot start the server %s", err)
	}

}

package main

import (
	"context"
	"github.com/thejimmyg/greener"
	"log"
)

func main() {
	app := greener.NewDefaultApp(
		greener.NewDefaultServeConfigProviderFromEnvironment(),
		greener.NewDefaultLogger(log.Printf),
		greener.NewDefaultEmptyPageProvider([]greener.Injector{}),
	)
	app.HandleWithServices("/", func(s greener.Services) {
		s.W().Write([]byte(app.Page("Hello", greener.Text("Hello <>!"))))
	})
	app.Serve(context.Background())
}

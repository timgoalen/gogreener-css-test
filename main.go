package main

import (
	"context"
	"log"

	"github.com/thejimmyg/greener"
)

func main() {
	// Create UI Support with the CSS
	uiSupport := greener.NewDefaultUISupport(HomePageCSS, "", "")

	// Set up logging
	logger := greener.NewDefaultLogger(log.Printf)

	// Set up the application
	app := greener.NewDefaultApp(
		greener.NewDefaultServeConfigProviderFromEnvironment(),
		logger,
		greener.NewDefaultEmptyPageProvider([]greener.Injector{
			greener.NewDefaultStyleInjector(logger, []greener.UISupport{uiSupport}), // Cache CSS for 1 hour
		}),
	)

	// Define a simple route handler
	app.HandleWithServices("/", func(s greener.Services) {
		s.W().Write([]byte(app.Page("Hello", greener.HTMLPrintf(HomePageHTML))))
	})

	// Start the server
	app.Serve(context.Background())
}

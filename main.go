package main

import (
	"context"
	"log"

	"github.com/thejimmyg/greener"
)

func main() {
	// HTML for our app
	html := `
		<div class="container">
			<h1>Welcome!</h1>
			<p>This is a simple example showcasing CSS injection in GoGreener.</p>
		</div>`

	// CSS for our app
	css := `
		body {
			font-family: sans-serif;
			margin: 0;
			padding: 0;
			background-color: #f0f0f0;
		}

		.container {
			max-width: 800px;
			margin: 0 auto;
			padding: 20px;
			background-color: #fff;
			box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		}

		h1 {
			color: #333;
		}

		p {
			line-height: 1.6;
			color: #555;
		}

		a {
			color: #007bff;
			text-decoration: none;
		}

		a:hover {
			text-decoration: underline;
		}
		`

	// Create UI Support with the CSS
	uiSupport := greener.NewDefaultUISupport(css, "", "")

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
		s.W().Write([]byte(app.Page("Hello", greener.HTMLPrintf(html))))
	})

	// Start the server
	app.Serve(context.Background())
}

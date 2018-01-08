package main

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/scott-rayapoulle.fr/components/css"
	"github.com/soulcramer/scott-rayapoulle.fr/middleware"
	"github.com/soulcramer/scott-rayapoulle.fr/pages"
)

var app = aero.New()

func main() {
	// Configure and start
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {

	// CSS
	app.SetStyle(css.Bundle())

	// Security
	configureHTTPS(app)

	// Assets
	configureAssets(app)

	// Pages
	pages.Configure(app)

	// Middleware
	app.Use(
		middleware.Log(),
		middleware.Session(),
	)

	return app
}

package main

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/scott-rayapoulle.fr/components/css"
	"github.com/soulcramer/scott-rayapoulle.fr/layout"
	"github.com/soulcramer/scott-rayapoulle.fr/page/frontpage"
	"github.com/soulcramer/scott-rayapoulle.fr/middleware"
)

func main() {
	app := aero.New()
	configure(app).Run()
}

func configure(app *aero.Application) *aero.Application {

	// CSS
	app.SetStyle(css.Bundle())

	// Layout
	app.Layout = layout.Render

	// Ajax routes
	app.Ajax("/", frontpage.Get)

	// Assets
	configureAssets(app)

	// Middleware
	app.Use(
		middleware.Log(),
		middleware.Session(),
	)

	return app
}

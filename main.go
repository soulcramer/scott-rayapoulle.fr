package main

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/eggma.fr/components/css"
	"github.com/soulcramer/eggma.fr/layout"
	"github.com/soulcramer/eggma.fr/page/frontpage"
	"github.com/soulcramer/eggma.fr/middleware"
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

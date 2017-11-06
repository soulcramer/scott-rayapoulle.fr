package main

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/scott-rayapoulle.fr/components/js"
)

// configureAssets adds all the routes used for media assets.
func configureAssets(app *aero.Application) {
	// Script bundle
	scriptBundle := js.Bundle()

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/brand/64", ".png")
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}

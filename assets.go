package main

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/scott-rayapoulle.fr/components/css"
	"github.com/soulcramer/scott-rayapoulle.fr/components/js"
	"io/ioutil"
)

// configureAssets adds all the routes used for media assets.
func configureAssets(app *aero.Application) {
	// Script bundle
	scriptBundle := js.Bundle()

	// Service worker
	serviceWorkerBytes, err := ioutil.ReadFile("sw/service-worker.js")
	serviceWorker := string(serviceWorkerBytes)

	// CSS bundle
	cssBundle := css.Bundle()

	if err != nil {
		panic("Couldn't load service worker")
	}

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/styles", func(ctx *aero.Context) string {
		return ctx.CSS(cssBundle)
	})

	app.Get("/service-worker", func(ctx *aero.Context) string {
		return ctx.JavaScript(serviceWorker)
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.File("images/brand/64.png")
	})

	// Images
	app.Get("/images/*file", func(ctx *aero.Context) string {
		return ctx.File("images" + ctx.Get("file"))
	})

	// Videos
	app.Get("/videos/*file", func(ctx *aero.Context) string {
		return ctx.File("videos" + ctx.Get("file"))
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}

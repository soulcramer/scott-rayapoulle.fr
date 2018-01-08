package pages

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/soulcramer/scott-rayapoulle.fr/pages/frontpage"
	"github.com/soulcramer/scott-rayapoulle.fr/layout"
	"github.com/soulcramer/scott-rayapoulle.fr/pages/experiments"
	"github.com/soulcramer/scott-rayapoulle.fr/pages/articles"
	"github.com/soulcramer/scott-rayapoulle.fr/pages/animations"
)

// Configure registers the page routes in the application.
func Configure(app *aero.Application) {
	l := layout.New(app)

	// Set render function for the layout
	l.Render = fullpage.Render

	// Main menu
	l.Page("/", frontpage.Get)
	l.Page("/experiments", experiments.Get)
	l.Page("/articles", articles.Get)
	l.Page("/animations", animations.Get)
}

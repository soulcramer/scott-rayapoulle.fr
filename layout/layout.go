package fullpage

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/scott-rayapoulle.fr/components"
)

// Render layout.
func Render(ctx *aero.Context, content string) string {
	return components.Layout(ctx.App, content)
}

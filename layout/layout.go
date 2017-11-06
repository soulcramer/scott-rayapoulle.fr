package layout

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/eggma.fr/components"
)

// Render layout.
func Render(ctx *aero.Context, content string) string {
	return components.Layout(ctx.App, content)
}

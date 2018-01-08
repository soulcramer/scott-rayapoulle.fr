package articles

import (
	"github.com/aerogo/aero"
	"github.com/soulcramer/scott-rayapoulle.fr/components"
)

// Get ...
func Get(ctx *aero.Context) string {

	return ctx.HTML(components.Articles())
}

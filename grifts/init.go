package grifts

import (
	"github.com/Jumanjii/dojo/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}

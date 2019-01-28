package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	// Helpers for navbar.
	c.Set("currentResource", "home")
	return c.Render(200, r.HTML("index.html"))
}

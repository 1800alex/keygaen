package stories

import (
	"github.com/1800alex/keygaen/pkg/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type NavbarStory struct {
	Story
}

func (c *NavbarStory) Render() app.UI {
	return c.WithRoot(
		&components.Navbar{},
	)
}

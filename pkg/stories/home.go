package stories

import (
	"github.com/1800alex/keygaen/pkg/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type HomeStory struct {
	Story
}

func (c *HomeStory) Render() app.UI {
	return c.WithRoot(
		&components.Home{},
	)
}

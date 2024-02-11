package stories

import (
	"github.com/1800alex/keygaen/pkg/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type EmptyStateStory struct {
	Story
}

func (c *EmptyStateStory) Render() app.UI {
	return c.WithRoot(
		&components.EmptyState{
			OnCreateKey: func() {
				app.Window().Call("alert", "Created key")
			},
			OnImportKey: func() {
				app.Window().Call("alert", "Imported key")
			},
		},
	)
}

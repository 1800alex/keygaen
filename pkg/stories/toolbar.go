package stories

import (
	"github.com/1800alex/keygaen/pkg/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ToolbarStory struct {
	Story
}

func (c *ToolbarStory) Render() app.UI {
	return c.WithRoot(
		&components.Toolbar{
			OnAddGitRepo: func() {
				app.Window().Call("alert", "Added git repository")
			},
			OnCreateKey: func() {
				app.Window().Call("alert", "Created key")
			},
			OnImportKey: func() {
				app.Window().Call("alert", "Imported key")
			},

			OnEncryptAndSign: func() {
				app.Window().Call("alert", "Encrypted and signed")
			},
			OnDecryptAndVerify: func() {
				app.Window().Call("alert", "Decrypted and verified")
			},
		},
	)
}

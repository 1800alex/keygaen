package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// AddGitRepoModal is a modal which provides the information needed to create a key
type AddGitRepoModal struct {
	app.Compo

	OnSubmit func(
		url string,
		username string,
		password string,
	) // Handler to call to create the key
	OnCancel func(dirty bool, clear chan struct{}) // Handler to call when closing/cancelling the modal

	url      string
	username string
	password string

	dirty bool
}

func (c *AddGitRepoModal) Render() app.UI {
	return &Modal{
		ID:           "add-git-repo-modal",
		Title:        "Add Git Repository",
		DisableFocus: true,
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("add-git-repo-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					// Submit the form
					c.OnSubmit(
						c.url,
						c.username,
						c.password,
					)

					c.clear()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("url-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("URL"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									&Autofocused{
										Component: app.Input().
											Class("pf-c-form-control").
											Required(true).
											Type("text").
											Placeholder("Jean Doe").
											ID("url-input").
											Aria("describedby", "form-demo-basic-name-helper").
											OnInput(func(ctx app.Context, e app.Event) {
												c.url = ctx.JSSrc().Get("value").String()

												c.dirty = true
											}).
											Value(c.url),
									},
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("username-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Text("Email"),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Text("*"),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("username").
												Placeholder("jean@example.com").
												ID("username-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.username = ctx.JSSrc().Get("value").String()

													c.dirty = true
												}).
												Value(c.username),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group").
								Body(
									app.Div().
										Class("pf-c-form__group-label").
										Body(
											app.Label().
												Class("pf-c-form__label").
												For("password-input").
												Body(
													app.Span().
														Class("pf-c-form__label-text").
														Text("Password"),
													app.Span().
														Class("pf-c-form__label-required").
														Aria("hidden", true).
														Text("*"),
												),
										),
									app.Div().
										Class("pf-c-form__group-control").
										Body(
											app.Input().
												Class("pf-c-form-control").
												Type("password").
												ID("password-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.password = ctx.JSSrc().Get("value").String()

													c.dirty = true
												}).
												Value(c.password),
										),
								),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Text("Create key").
				Form("add-git-repo-form"),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text("Cancel").
				OnClick(func(ctx app.Context, e app.Event) {
					c.handleCancel(c.clear, c.dirty, c.OnCancel)
				}),
		},
		OnClose: func() {
			c.handleCancel(c.clear, c.dirty, c.OnCancel)
		},
	}

}

func (c *AddGitRepoModal) handleCancel(clear func(), dirty bool, cancel func(bool, chan struct{})) {
	done := make(chan struct{})

	go func() {
		<-done

		clear()
	}()
	cancel(dirty, done)
}

func (c *AddGitRepoModal) clear() {
	c.url = ""
	c.username = ""
	c.password = ""
	c.dirty = false
}

package core

import (
	"net/http"

	"newser.app/ui/view/pages/auth"
	"newser.app/ui/view/pages/desk"
	"newser.app/ui/view/pages/home"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}
	render(c, w, home.Index())
}

func (app *App) Desk(w http.ResponseWriter, r *http.Request) {
	render(r.Context(), w, desk.Index())
}

func (app *App) Auth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/auth/login/", http.StatusSeeOther)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	render(r.Context(), w, auth.Login())
}

func (app *App) Signup(w http.ResponseWriter, r *http.Request) {
	render(r.Context(), w, auth.Signup())
}

package core

import (
	"net/http"
	"path/filepath"
)

func (app *App) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(staticFs{(http.Dir("./ui/static/"))})

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /desk/", app.Desk)
	mux.HandleFunc("GET /auth/", app.Auth)
	mux.HandleFunc("GET /auth/login/", app.Login)
	mux.HandleFunc("GET /auth/signup/", app.Signup)

	return mux
}

type staticFs struct {
	fs http.FileSystem
}

func (sfs staticFs) Open(path string) (http.File, error) {
	f, err := sfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := sfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}

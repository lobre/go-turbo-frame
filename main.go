package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type application struct {
	indexTmpl *template.Template
	frameTmpl *template.Template
}

//go:embed index.html frame.html
var assets embed.FS

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var app application
	var err error

	app.indexTmpl, err = template.ParseFS(assets, "index.html")
	if err != nil {
		return err
	}

	app.frameTmpl, err = template.ParseFS(assets, "frame.html")
	if err != nil {
		return err
	}

	http.HandleFunc("/", app.index)
	http.HandleFunc("/form", app.form)
	http.HandleFunc("/frame", app.frame)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	err := app.indexTmpl.Execute(w, nil)
	if err != nil {
		serverErr(w)
	}
}

func (app *application) form(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/frame", http.StatusSeeOther)
}

func (app *application) frame(w http.ResponseWriter, r *http.Request) {
	// net/http won’t automatically detect the content-type
	// if the frame is the first element of the payload.
	w.Header().Set("Content-Type", "text/html")

	err := app.frameTmpl.Execute(w, nil)
	if err != nil {
		serverErr(w)
	}
}

func serverErr(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

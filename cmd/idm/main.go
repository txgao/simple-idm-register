package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/tendant/chi-demo/app"
)

func main() {
	myApp := app.Default()
	Routes(myApp.R)
	myApp.Run()

}

func Routes(r *chi.Mux) {
	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, http.StatusText(http.StatusOK))
	})

}

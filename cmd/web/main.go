package main

import (
	"hello-world-server/pkg/config"
	"hello-world-server/pkg/handlers"
	"hello-world-server/pkg/render"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTempleteCache()
	if err != nil {
		log.Fatal("Can not create template cache", err)
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

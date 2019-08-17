package main

import (
	"log"
	"net/http"

	"github.com/betonr/deplus-api/modules/status"
	r "github.com/betonr/go-utils/rest"
	"github.com/bmizerany/pat"
)

// App - Struct default at Application
type App struct {
	Router *pat.PatternServeMux
}

func (app *App) initRoutes() {
	app.Router = pat.New()

	// status group
	app.Router.Get("/status", http.HandlerFunc(status.GetStatus))
}

// Run - make start of Application
func (app *App) Run(port string) {
	handlers := r.EnableCors(app.Router, r.Cors{
		Methods: []string{"GET", "POST", "PUT", "DELETE"},
		Origins: []string{"*"},
	})
	handlers = r.EnableLogs(handlers)

	log.Println("** App started on port:", port, "**")
	log.Fatal(http.ListenAndServe(":"+port, handlers))
}

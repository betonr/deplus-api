package main

import (
	"log"
	"net/http"

	"github.com/betonr/deplus-api/modules/status"
	"github.com/betonr/deplus-api/modules/utils"
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
	handlers := utils.EnableCors(app.Router, utils.Cors{
		Methods: []string{"GET", "POST", "PUT", "DELETE"},
		Origins: []string{"*"},
	})
	handlers = utils.EnableLogs(handlers)

	log.Println("** App started on port:", port, "**")
	log.Fatal(http.ListenAndServe(":"+port, handlers))
}

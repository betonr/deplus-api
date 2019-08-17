package main

import (
	"log"
	"net/http"
	"os"

	"github.com/betonr/deplus-api/modules/status"
	ub "github.com/betonr/go-utils/base"
	ur "github.com/betonr/go-utils/rest"
	"github.com/bmizerany/pat"
)

// Struct default at Application
type app struct {
	Router *pat.PatternServeMux
}

// set routes of Application
func (a *app) initRoutes() {
	a.Router = pat.New()

	// status group
	a.Router.Get("/status", http.HandlerFunc(status.GetStatus))
}

// make start of Application
func main() {
	port := ub.GetBetween([]string{os.Getenv("PORT"), "5000"})

	app := app{}
	app.initRoutes()

	handlers := ur.EnableCors(app.Router, ur.Cors{
		Methods: []string{"GET", "POST", "PUT", "DELETE"},
		Origins: []string{"*"},
	})
	handlers = ur.EnableLogs(handlers)

	log.Println("** App started on port:", port, "**")
	log.Fatal(http.ListenAndServe(":"+port, handlers))
}

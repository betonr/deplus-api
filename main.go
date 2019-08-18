package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/betonr/deplus-api/modules/status"
	ub "github.com/betonr/go-utils/base"
	ur "github.com/betonr/go-utils/rest"
	"github.com/bmizerany/pat"

	_ "github.com/lib/pq"
)

// Struct default at Application
type App struct {
	Router *pat.PatternServeMux
	Bd     *DB
}

// set routes of Application
func (a *App) initRoutes() {
	a.Router = pat.New()

	// status group
	a.Router.Get("/status", http.HandlerFunc(status.GetStatus))
}

// initialize DB of Application
func (a *App) initDB() {
	host := ub.GetBetween([]string{os.Getenv("BD_HOST"), "localhost"})
	port := ub.GetBetween([]string{os.Getenv("BD_POST"), "5432"})
	user := ub.GetBetween([]string{os.Getenv("BD_USER"), "postgres"})
	password := ub.GetBetween([]string{os.Getenv("BD_PASSWORD"), "postgres"})
	dbname := ub.GetBetween([]string{os.Getenv("BD_DB"), "deplus"})

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	app.Db = db
	defer db.Close()
}

// make start of Application
func main() {
	port := ub.GetBetween([]string{os.Getenv("PORT"), "5000"})

	app := App{}
	app.initRoutes()
	app.initDB()

	handlers := ur.EnableCors(app.Router, ur.Cors{
		Methods: []string{"GET", "POST", "PUT", "DELETE"},
		Origins: []string{"*"},
	})
	handlers = ur.EnableLogs(handlers)

	log.Println("** App started on port:", port, "**")
	log.Fatal(http.ListenAndServe(":"+port, handlers))
}

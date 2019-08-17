package main

import (
	"os"

	b "github.com/betonr/go-utils/base"
)

func main() {
	port := b.GetBetween([]string{os.Getenv("PORT"), "5000"})

	app := App{}
	app.initRoutes()

	app.Run(port)
}

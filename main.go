package main

import (
	"os"

	"github.com/betonr/deplus-api/modules/utils"
)

func main() {
	port := utils.GetBetween([]string{os.Getenv("PORT"), "5000"})

	app := App{}
	app.initRoutes()

	app.Run(port)
}

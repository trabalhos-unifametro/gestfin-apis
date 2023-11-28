package main

import (
	"fmt"
	"gestfin-apis/models"
	"gestfin-apis/routes"
	"gestfin-apis/utils"
	"time"
)

func init() {
	utils.LoadENVs()
}

func main() {
	app := routes.Routes()
	time.Local = time.UTC
	models.RunMigrations()

	if err := app.Listen(fmt.Sprint(":", utils.GodotEnv("PORT_APIS"))); err != nil {
		return
	}
}

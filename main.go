package main

import (
	"fmt"
	// "log"
	// "togo_app/config"
	"togo_app/app/controllers"
	"togo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
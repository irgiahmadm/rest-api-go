package main

import (
	"github.com/irgiahmadm/simple-rest-api-go/app"
	"github.com/irgiahmadm/simple-rest-api-go/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":2001")
}

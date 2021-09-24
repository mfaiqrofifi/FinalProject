package main

import (
	"social_media/configs"
	"social_media/routes"
)

func main() {
	configs.InitDB()
	e := routes.NewRoutes()
	e.Start(":8000")
}

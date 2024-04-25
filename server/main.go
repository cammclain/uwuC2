package main

import (
	"server/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run() // default listens on :8080
}

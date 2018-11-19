package main

import "Prem/pkg/router"

var (
	PORT = ":8080"
)

func main() {
	router.RunRouter(PORT)
}

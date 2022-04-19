package main

import (
	"todogohexa/infrastructure"
)

func main() {
	infrastructure.GinRoutes("memory").Run()
}
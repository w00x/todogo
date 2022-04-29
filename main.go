package main

import (
	"todogo/infrastructure"
)

func main() {
	infrastructure.GinRoutes("gorm").Run()
}

package main

import (
	"todogo/infrastructure"
	"todogo/infrastructure/repository"
)

func main() {
	infrastructure.GinRoutes(repository.GORM).Run()
}

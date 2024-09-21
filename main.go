package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	rotar := gin.Default()
	rotar.SetTrustedProxies(nil)

	DB := initdb("tasks.db")
	defer DB.Close()

	RegisterRotas(rotar)

	rotar.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func RegisterRotas(rotar *gin.Engine) {

	rotar.GET("/Tarefas", GetAll)

	rotar.POST("/AddTarefas", AdicionarTarefas)

	rotar.GET("/Tarefas/:id", GetbyId)

	rotar.DELETE("/Tarefas/:id", DeletaTarefa)

	rotar.PUT("/Tarefas/:id", AtualizarTarefa)

}

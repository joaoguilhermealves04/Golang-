package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Tasks struct {
	id    int    `json:"id"`
	Title string `json:"title"`
}

func GetAll(c *gin.Context) {

	rows, err := db.Query("SELECT id ,title FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	defer rows.Close()

	var tasks []Tasks
	for rows.Next() {
		var task Tasks

		if err := rows.Scan(&task.id, &task.Title); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		tasks = append(tasks, task)
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func AdicionarTarefas(c *gin.Context) {

	var newTask Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	newTask.id = int(id)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func GetbyId(c *gin.Context) {
	id := c.Param("id")

	var task Tasks

	row := db.QueryRow("SELECT id,title FROM tasks WHERE id=?", id)

	if err := row.Scan(&task.id, &task.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, task)
}

func DeletaTarefa(c *gin.Context) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM tasks WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"mensagem": " trarefa deletada com sucesso."})
}

func AtualizarTarefa(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var newTasks Tasks

	if err := c.ShouldBind(&newTasks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("update tasks set title=? where id=?", newTasks.Title, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTasks.id = id
	c.IndentedJSON(http.StatusCreated, newTasks)
}

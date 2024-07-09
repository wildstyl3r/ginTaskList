package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) getTasks(c *gin.Context) {
	var tasks []Task
	a.db().Find(&tasks)
	c.IndentedJSON(http.StatusOK, tasks)
}

func (a *App) postTasks(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		log.Print(err)
		return
	}
	result := a.db().Create(&newTask)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to create task"})
		return
	}
	c.IndentedJSON(http.StatusOK, newTask)

}

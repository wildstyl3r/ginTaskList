package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *App) startTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var task Task
	result := a.db().First(&task, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	if task.Started != nil {
		c.IndentedJSON(http.StatusConflict, task)
		return
	}
	t := time.Now()
	task.Started = &t
	a.db().Save(&task)
	c.IndentedJSON(http.StatusOK, task)
}

func (a *App) cancelTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var task Task
	result := a.db().First(&task, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	if task.Cancelled != nil {
		c.IndentedJSON(http.StatusConflict, task)
		return
	}
	t := time.Now()
	task.Cancelled = &t
	a.db().Save(&task)
	c.IndentedJSON(http.StatusOK, task)
}

func (a *App) completeTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var task Task
	result := a.db().First(&task, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	if task.Completed != nil {
		c.IndentedJSON(http.StatusConflict, task)
		return
	}
	t := time.Now()
	task.Completed = &t
	a.db().Save(&task)
	c.IndentedJSON(http.StatusOK, task)
}

func (a *App) getTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	var task Task
	result := a.db().First(&task, id)
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

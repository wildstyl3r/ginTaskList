package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int       `json: "id"`
	Name        string    `json:"name"`
	Description string    `json:"descr"`
	Added       time.Time `json: "added"`
	Started     time.Time `json: "started"`
	Cancelled   time.Time `json: "cancelled"`
	Done        time.Time `json: "done"`
}

var tasks = map[int]Task{
	1: {1, "Todo1", "SOmedetails", time.Now(), time.Time{}, time.Time{}, time.Time{}},
	2: {2, "Todo3", "Important", time.Now(), time.Time{}, time.Time{}, time.Time{}},
	3: {3, "Todo2", "", time.Now(), time.Time{}, time.Time{}, time.Time{}},
}

func postTasks(c *gin.Context) {
	var newTask Task
	if err := c.BindJSON(&newTask); err != nil {
		log.Print(err)
		return
	}
	id := 1
	for ; id <= len(tasks); id++ {
		if _, ok := tasks[id]; !ok {
			break
		}
	}
	newTask.ID = id
	newTask.Added = time.Now()
	tasks[id] = newTask
	c.IndentedJSON(http.StatusOK, newTask)

}

// func startTaskById(c *gin.Context) {
// 	id := c.Param("id")
// }

// func cancelTaskById(c *gin.Context) {
// 	id := c.Param("id")
// }

// func finishTaskById(c *gin.Context) {
// 	id := c.Param("id")
// }

func getTaskById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
	}
	c.IndentedJSON(http.StatusOK, tasks[id])
}

func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func main() {
	r := gin.Default()
	r.GET("/tasks", getTasks)
	r.POST("/tasks", postTasks)
	r.Run("localhost:8883")
}

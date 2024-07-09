package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

func main() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	app := NewApp(config)

	r := gin.Default()
	r.GET("/tasks", app.getTasks)
	r.POST("/tasks", app.postTasks)
	r.GET("/tasks/:id", app.getTaskById)
	r.PUT("/tasks/:id/start", app.startTaskById)
	r.PUT("/tasks/:id/cancel", app.cancelTaskById)
	r.PUT("/tasks/:id/complete", app.completeTaskById)
	r.Run("localhost:8883")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("POSTGRES")
	viper.BindEnv("HOST")
	viper.BindEnv("USER")
	viper.BindEnv("PASSWORD")
	viper.BindEnv("PORT")
	var config Config
	config.Database.Host = viper.GetString("HOST")
	config.Database.Port = viper.GetString("PORT")
	config.Database.Username = viper.GetString("USER")
	config.Database.Password = viper.GetString("PASSWORD")
	config.Database.DBname = viper.GetString("USER")
	// _, err := toml.DecodeFile("config.toml", &config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	app := NewApp(config)

	r := gin.Default()
	r.GET("/tasks", app.getTasks)
	r.POST("/tasks", app.postTasks)
	r.GET("/tasks/:id", app.getTaskById)
	r.PUT("/tasks/:id/start", app.startTaskById)
	r.PUT("/tasks/:id/cancel", app.cancelTaskById)
	r.PUT("/tasks/:id/complete", app.completeTaskById)
	r.Run()
}

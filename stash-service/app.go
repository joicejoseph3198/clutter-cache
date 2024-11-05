package app

import (
	"github.com/gin-gonic/gin"
	"joicejoseph.dev/clutter-cache/stash-service/config"
	"joicejoseph.dev/clutter-cache/stash-service/pkg/database"
)

type App struct{
	// Add all handlers here
}

// NewApp is responsible for injecting all dependencies required for the application
// - loads all env variables
// - initializes a database connection, makes all neccessary migrations
// - initializes required repositories, services and handlers.
func NewApp() (*App, error){
	config.LoadEnvVariables()

	postgresConn := database.NewPostgresConnection()
	database.SyncDatabase(postgresConn)

	// Initialize repos, services and handlers

	return &App{}, nil
}
// Run sets up all the routes and starts the server
func (a *App) Run(){
	// Here we can write the code which setups all the routes and starts a server
	router := gin.Default()
	//health check url
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := router.Run(":4000") 
	if err != nil {
        panic("[Error] failed to start Gin server due to: " + err.Error())
    }
}
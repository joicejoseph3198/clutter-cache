package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"joicejoseph.dev/clutter-cache/stash-service/config"
	"joicejoseph.dev/clutter-cache/stash-service/internal/handler"
	"joicejoseph.dev/clutter-cache/stash-service/internal/repository"
	"joicejoseph.dev/clutter-cache/stash-service/internal/routes"
	"joicejoseph.dev/clutter-cache/stash-service/internal/service"
	"joicejoseph.dev/clutter-cache/stash-service/pkg/database"
)

type App struct{
	EntryHandler handler.EntryHandler
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
	validator := validator.New()
	entryRepository := repository.NewEntryRepository(postgresConn.DB)
	entryService := service.NewEntryService(entryRepository, validator)
	entryHandler := handler.NewEntryHandler(entryService)

	return &App{EntryHandler: *entryHandler}, nil
}

// Run sets up all the routes and starts the server
func (a *App) Run(){

	router := gin.New()
	router.Use(gin.Logger())
	

	//health check url
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "UP",
		})
	})

	//Defining routes
	entryRoutes :=routes.NewEntryRoutes(a.EntryHandler)
	entryRoutes.EntryRoutesImpl(router)

	err := router.Run(":4000") 
	if err != nil {
        panic("[Error] failed to start Gin server due to: " + err.Error())
    }
}
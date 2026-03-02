package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ArdhiCode/go-auth/db"
	"github.com/ArdhiCode/go-auth/internal/api/controller"
	"github.com/ArdhiCode/go-auth/internal/api/repository"
	"github.com/ArdhiCode/go-auth/internal/api/routes"
	"github.com/ArdhiCode/go-auth/internal/api/service"
	"github.com/ArdhiCode/go-auth/internal/middleware"
	myjwt "github.com/ArdhiCode/go-auth/internal/pkg/jwt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type RestConfig struct {
	server *gin.Engine
}

func NewRest() (RestConfig, error) {
	db := db.New()
	if db == nil {
		return RestConfig{}, fmt.Errorf("database connection failed")
	}

	if mode := os.Getenv("APP_MODE"); mode == "production" || mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression))

	server := NewRouter(app)
	_ = middleware.New(db)

	// Injections
	userRepo := repository.NewUser(db)

	jwtService := myjwt.NewJWT()
	authService := service.NewAuth(userRepo, jwtService)
	authController := controller.NewAuth(authService)

	indexController := controller.NewIndex()

	//m := middleware.New(db)

	// Register all routes
	server.GET("/", indexController.Index)
	routes.Auth(server, authController)

	return RestConfig{
		server: server,
	}, nil
}

func (ap *RestConfig) Start() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	if port == "" {
		port = "8080"
	}

	serve := fmt.Sprintf("%s:%s", host, port)
	if err := ap.server.Run(serve); err != nil {
		log.Panicf("failed to start server: %s", err)
	}
	log.Println("server start on port ", serve)
}

func (ap *RestConfig) GetServer() *gin.Engine {
	return ap.server
}

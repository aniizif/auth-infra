package app

import (
	"github.com/aniizif/stack-mate/auth-service/internal/handlers"
	"github.com/aniizif/stack-mate/auth-service/internal/middleware"
	"github.com/aniizif/stack-mate/auth-service/internal/models"
	"github.com/aniizif/stack-mate/auth-service/internal/repository"
	"github.com/aniizif/stack-mate/auth-service/internal/routes"
	"github.com/aniizif/stack-mate/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type App struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewApp() *App {
	app := &App{}
	app.initDB()
	app.initRouter()

	return app
}

func (app *App) initRouter() {
	userRepo := repository.NewUserRepository(app.db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	r := gin.Default()
	r.Use(middleware.PrometheusMiddleware())
	routes.RegisterAuthRoutes(r, authHandler)
	app.router = r
}

func (app *App) initDB() {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	app.db = db
}

func (app *App) Run() {
	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)

	if err := app.router.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

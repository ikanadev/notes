package restbooksserver

import (
	"books/internal/pkg/configs"
	"books/internal/pkg/db"
	"books/internal/pkg/db/migrations"
	"books/internal/pkg/repo"
	"books/internal/pkg/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type App struct {
	dbConn *gorm.DB
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := db.ProvideDBConn(&appConfig.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	a.dbConn = dbConn

	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, dbConn)
	if err != nil {
		log.Fatal(err)
	}
	migrator.RunMigrations()

	bookRepo := repo.GetNewBookRepository(a.dbConn)
	bookService := service.GetNewBooksService(bookRepo)
	r := ProvideRouter(bookService)

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", appConfig.ServerConfig.Host, appConfig.ServerConfig.Port),
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Starting server")
	log.Fatal(server.ListenAndServe())
}

func (a *App) Shutdown() {
	dbInstance, _ := a.dbConn.DB()
	_ = dbInstance.Close()
}

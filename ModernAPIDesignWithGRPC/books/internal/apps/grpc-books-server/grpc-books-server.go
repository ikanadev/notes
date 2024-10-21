package grpcbooksserver

import (
	"books/internal/pkg/configs"
	"books/internal/pkg/db"
	"books/internal/pkg/db/migrations"
	"books/internal/pkg/middleware"
	"books/internal/pkg/proto"
	"books/internal/pkg/repo"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	proto.UnimplementedBookServiceServer

	dbConn   *gorm.DB
	bookRepo *repo.BookRepository
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
	a.bookRepo = repo.GetNewBookRepository(a.dbConn)

	migrator, err := migrations.ProvideMigrator(appConfig.DBConfig, a.dbConn)
	if err != nil {
		log.Fatal(err)
	}
	migrator.RunMigrations()

	serverAddr := fmt.Sprintf("0.0.0.0:%d", appConfig.ServerConfig.Port)

	fmt.Println("starting books gRPC server at", serverAddr)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.ServerOption{}
	middlewareOpts := middleware.ProvideGrpcMiddlewareServerOpts()
	opts = append(opts, middlewareOpts...)

	s := grpc.NewServer(opts...)
	proto.RegisterBookServiceServer(s, a)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (a *App) Shutdown() {
	db, _ := a.dbConn.DB()
	_ = db.Close()
}

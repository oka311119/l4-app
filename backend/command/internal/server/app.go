package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"

	areaRepo "github.com/oka311119/l4-app/backend/command/internal/area/repository/localstorage"
	"github.com/oka311119/l4-app/backend/command/internal/auth"
	authhttp "github.com/oka311119/l4-app/backend/command/internal/auth/delivery"
	authRepo "github.com/oka311119/l4-app/backend/command/internal/auth/repository/localstorage"
	authusecase "github.com/oka311119/l4-app/backend/command/internal/auth/usecase"
	"github.com/oka311119/l4-app/backend/command/internal/config"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/saltgen"
	"github.com/oka311119/l4-app/backend/command/internal/helpers/uuidgen"
	// userrepo "github.com/oka311119/l4-app/backend/command/internal/auth/repository/dynamodb"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase
}

func NewApp(cfg *config.Config) *App {
	// db := initDB(cfg)

	// userRepo := userrepo.NewUserRepository(db, cfg.AWS.DynamoDB.UserTableName)
	userRepo := authRepo.NewUserLocalStorage()
	areaRepo := areaRepo.NewAreaLocalStorage()

	return &App{
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			areaRepo,
			cfg.Auth.Pepper,
			[]byte(cfg.Auth.SigningKey),
			time.Duration(cfg.Auth.TokenTTL),
			&uuidgen.UUID{},
			&saltgen.Salt{},
		),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	// authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	// api := router.Group("/api", authMiddleware)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB(cfg *config.Config) *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWS.Region),
	})
	if err != nil {
		log.Fatalf("Error occurred while establishing connection to AWS: %s", err)
	}

	db := dynamodb.New(sess)

	// Try to list tables to verify connection
	input := &dynamodb.ListTablesInput{}
	_, err = db.ListTables(input)
	if err != nil {
		log.Fatalf("Error occurred while trying to connect to DynamoDB: %s", err)
	}

	return db
}

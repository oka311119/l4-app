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
	"github.com/spf13/viper"

	"github.com/oka311119/l4-app/backend/command/internal/auth"
	authhttp "github.com/oka311119/l4-app/backend/command/internal/auth/delivery"
	userrepo "github.com/oka311119/l4-app/backend/command/internal/auth/repository/dynamodb"
	authusecase "github.com/oka311119/l4-app/backend/command/internal/auth/usecase"
)

type App struct {
	httpServer *http.Server

	authUC auth.UseCase
}

func NewApp() *App {
	db := initDB()

	userRepo := userrepo.NewUserRepository(db, viper.GetString("dynamodb.user_tablename"))

	return &App{
		authUC: authusecase.NewAuthUseCase(
			userRepo,
			viper.GetString("auth.hash_salt"),
			[]byte(viper.GetString("auth.signing_key")),
			viper.GetDuration("auth.token_ttl"),
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
		Addr: ":" + port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt,os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(viper.GetString("aws.region")),
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


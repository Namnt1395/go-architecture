package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-architecture/api/handle"
	"go-architecture/config"
	"go-architecture/docs"
	"go-architecture/entities"
	util "go-architecture/pkg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Infra            *Infrastructure
	companyCapacityHandler handle.CompanyCapacityHandler
}

func (r App) Start() error {
	// setup routerr
	r.SetupRouters()

	// migration
	r.Migrate()

	// run Gin engine
	util.CheckError(r.Infra.Router.Engine.Run(fmt.Sprintf(":%s", r.Infra.Config.Server.Port)))

	gracefulShutdown(&http.Server{
		Addr:    fmt.Sprintf(":%s", r.Infra.Config.Server.Port),
		Handler: r.Infra.Router.Engine,
	})

	return nil
}

func (r App) Stop() {
	if err := r.Infra.Database.Close(); err != nil {
		panic(err)
	}
}

func (r App) Migrate() {
	_ = r.Infra.Database.DB.AutoMigrate(&entities.CompanyCapacity{})
	_ = r.Infra.Database.DB.AutoMigrate(&entities.CompanyBehavior{})
}

func (r App) SetupRouters() {
	// public api
	groupPublic := r.Infra.Router.Engine.Group("/api/public/v1")
	{
		// login api
		groupPublic.POST("/login")
	}
	// authorized api
	groupV1 := r.Infra.Router.Engine.Group("/api/v1")

	r.companyCapacityHandler.Setup(groupV1.Group("capacities", r.Infra.Router.AuthMiddleware.MiddlewareFunc()))

	// init swagger
	r.InitSwagger(r.Infra.Config)
}

func (r App) InitSwagger(c config.Config) {
	docs.SwaggerInfo.Host = c.Swagger.Url
	r.Infra.Router.InitSwagger(c)
}

// @title Sample Service API By Layer
// @version 1.0
// @description This is Sample Service API By Layer.

// @contact.name Nguyen Thanh Nam
// @contact.email namnguyenthanh024@gmail.com

// @host localhost:8099
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// wire
	app, err := InitApp()
	util.CheckError(err)

	err = app.Start()
	util.CheckError(err)

	defer app.Stop()

	log.Info().Msg("App started")
}

func gracefulShutdown(srv *http.Server) {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("listen")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")
}

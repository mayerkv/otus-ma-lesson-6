package main

import (
	"context"
	"embed"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/jmoiron/sqlx"
	"go.uber.org/dig"
	"log"
	"mayerkv/otus-ma-lesson-6/app/common"
	"mayerkv/otus-ma-lesson-6/app/di"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed infrastructure/persistence/pg/migrations/*.sql
var migrationFS embed.FS

func main() {
	providers := []func(c *dig.Container) error{
		di.ProvideConfig,
		di.ProvideController,
		di.ProvideDatabase,
		di.ProvideRepository,
		di.ProvideRouter,
		di.ProvideService,
		di.ProvideMetrics,
	}
	container := dig.New()
	for _, provide := range providers {
		if err := provide(container); err != nil {
			log.Fatal(err)
		}
	}

	var runMigrationsFlag, runWebServerFlag bool
	flag.BoolVar(&runMigrationsFlag, "migrate", false, "Run database migrations")
	flag.BoolVar(&runWebServerFlag, "web", false, "Run http server")
	flag.Parse()

	if !runWebServerFlag && !runMigrationsFlag {
		flag.Usage()
		return
	}

	if runMigrationsFlag {
		if err := container.Invoke(runMigrations); err != nil {
			log.Fatal(err)
		}
	}

	if runWebServerFlag {
		if err := container.Invoke(runWebServer); err != nil {
			log.Fatal(err)
		}
	}
}

func runWebServer(router *gin.Engine, config *common.Config) {
	srv := &http.Server{
		Addr:         "0.0.0.0:80",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.GracefulDelay)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds.", config.GracefulDelay)
	}
	log.Println("Server exiting")
}

func runMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	fs := http.FS(migrationFS)
	s, err := httpfs.New(fs, "infrastructure/persistence/pg/migrations")
	if err != nil {
		return err
	}

	instance, err := migrate.NewWithInstance("httpfs", s, "pg", driver)
	if err != nil {
		return err
	}

	if err := instance.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}

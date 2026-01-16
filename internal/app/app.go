package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/ultrabor/warehouse-api/internal/config"
	v1 "github.com/ultrabor/warehouse-api/internal/delivery/http/v1"
	"github.com/ultrabor/warehouse-api/internal/repository/postgres"
	"github.com/ultrabor/warehouse-api/internal/usecase"
	"github.com/ultrabor/warehouse-api/pkg/logger"
)

func RunApp() {
	logger := slog.New(logger.NewHandler(nil))
	slog.SetDefault(logger)

	cfg := config.Load(logger)

	db, err := sqlx.Connect("postgres", cfg.DBURL)
	if err != nil {
		logger.Error("Failed to connect to DB", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	logger.Info("Connected to Database")

	if err := config.RunMigrations(cfg.DBURL); err != nil {
		logger.Error("Migration failed", "error", err)
		os.Exit(1)
	}

	repo := postgres.NewProductRepository(db)
	uc := usecase.NewProductUseCase(repo)
	h := v1.NewProductHandler(uc)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(v1.LoggerMiddleware(logger), gin.Recovery())

	h.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		logger.Info("Server starting", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Listen error", "error", err)
		}
	}()

	config.ServerShutdown(srv, logger)
}

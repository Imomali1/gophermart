package gophermart

import (
	"context"
	"errors"
	"github.com/Imomali1/gophermart/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Imomali1/gophermart/internal/api"
	"github.com/Imomali1/gophermart/internal/apps/gophermart/config"
	"github.com/Imomali1/gophermart/internal/pkg/logger"
	"github.com/Imomali1/gophermart/internal/pkg/server"
	"github.com/Imomali1/gophermart/internal/pkg/storage"
)

const _timeout = 5 * time.Second

func Run() {
	var conf config.Config
	config.Parse(&conf)

	log := logger.NewLogger(os.Stdout, conf.LogLevel, conf.ServiceName)

	ctx, cancel := context.WithTimeout(context.Background(), _timeout)
	defer cancel()
	store, err := storage.New(ctx, conf)
	if err != nil {
		log.Logger.Info().Err(err).Msg("failed to initialize storage")
		return
	}

	repo := repository.New(store)
	serviceManager := services.New(repo)
	handler := api.NewRouter(api.Options{
		Logger:         log,
		Conf:           conf,
		ServiceManager: serviceManager,
	})
	srv := server.NewServer(conf.ServerAddress, handler)

	go func() {
		if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Logger.Info().Err(err).Msg("failed to initialize storage")
		}
	}()

	log.Logger.Info().Msg("Server is up and running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit

	if err = store.Close(); err != nil {
		log.Logger.Info().Err(err).Msg("error in closing storage")
	}

	ctx, shutdown := context.WithTimeout(context.Background(), _timeout)
	defer shutdown()
	if err = srv.GracefulShutdown(ctx); err != nil {
		log.Logger.Info().Err(err).Msg("error in shutting down server")
		return
	}
	log.Logger.Info().Msg("server stopped successfully")
}

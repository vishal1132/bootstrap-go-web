package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/vishal1132/bootstrap-go-web/config"
	common_http "github.com/vishal1132/bootstrap-go-web/http"
	"github.com/vishal1132/bootstrap-go-web/log"
	"github.com/vishal1132/bootstrap-go-web/postgres"
	"github.com/vishal1132/bootstrap-go-web/redis"
	"github.com/vishal1132/bootstrap-go-web/status"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf := config.LoadConfig()
	log.SetupLogger(conf.LogConfig)
	router := chi.NewRouter()

	// init dependencies here.
	db := postgres.InitDatabase(conf.DBConfig)
	cache := redis.InitRedis(ctx, conf.RedisConfig)

	// Routes here.
	status.NewStatusController(router, status.NewPGMonitor("maindb", db), status.NewRedisMonitor("maincache", cache))

	srv := common_http.InitServer(router, conf.AppConfig.ServerWriteTimeout, conf.AppConfig.ServerReadTimeout)
	zap.L().Info("Starting server")
	go srv.ListenAndServe()
	waitForShutdownSignal(ctx, conf.AppConfig.GracefulShutDownTimeout, srv, cancel)
}

func waitForShutdownSignal(ctx context.Context, gracefulShutDownTimeout int64, srv *http.Server, cancel context.CancelFunc) {
	var gracefulStop = make(chan os.Signal, 1)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	signal.Notify(gracefulStop, syscall.SIGQUIT)

	select {
	case <-gracefulStop:
		cancel()
		// if stop signal is received, wait for some time so that background workers get time to exit
		<-time.After(time.Duration(gracefulShutDownTimeout) * time.Millisecond)
	case <-ctx.Done():
		// shutdown if context was cancelled by something else before shutdown signal
	}
	srv.Shutdown(ctx)
}

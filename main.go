package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Sn0wo2/QuickNote/config"
	"github.com/Sn0wo2/QuickNote/database/orm"
	"github.com/Sn0wo2/QuickNote/database/table"
	"github.com/Sn0wo2/QuickNote/log"
	"github.com/Sn0wo2/QuickNote/router"
	"github.com/Sn0wo2/QuickNote/setup"
	"go.uber.org/zap"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	defer func() {
		_ = log.Instance.Sync()
	}()

	err := config.Init()
	if err != nil {
		log.Instance.Fatal("Failed to load config",
			zap.Error(err),
		)
	}

	err = orm.Init(config.Instance.Database.Type, config.Instance.Database.URL)
	if err != nil {
		log.Instance.Fatal("Failed to initialize database",
			zap.Error(err),
		)
	}

	err = table.Init()
	if err != nil {
		log.Instance.Fatal("Failed to initialize tables",
			zap.Error(err),
		)
	}

	app := setup.Fiber()
	router.Setup(app)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Instance.Info("Starting server",
			zap.String("address", "0.0.0.0:3000"),
		)

		if err := app.Listen(":3000"); err != nil {
			log.Instance.Fatal("Server failed to start",
				zap.Error(err),
			)
		}
	}()

	sig := <-shutdownChan
	log.Instance.Info("Received shutdown signal",
		zap.String("signal", sig.String()),
	)

	if err = app.Shutdown(); err != nil {
		log.Instance.Error("Server shutdown error",
			zap.Error(err),
		)
	} else {
		log.Instance.Info("Server gracefully stopped")
	}
}

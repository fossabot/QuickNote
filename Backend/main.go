package main

import (
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"

	"github.com/Sn0wo2/QuickNote/Backend/internal/router"
	"github.com/Sn0wo2/QuickNote/Backend/internal/setup"
	"github.com/Sn0wo2/QuickNote/Backend/pkg/config"
	"github.com/Sn0wo2/QuickNote/Backend/pkg/database/orm"
	"github.com/Sn0wo2/QuickNote/Backend/pkg/database/table"
	"github.com/Sn0wo2/QuickNote/Backend/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// stw
	debug.SetGCPercent(50)
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

	if !fiber.IsChild() {
		log.Instance.Info("Starting server",
			zap.String("address", "0.0.0.0:3000"),
			zap.Int("pid", os.Getpid()),
		)
	}

	app := setup.Fiber()
	router.Setup(app)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err = app.Listen(":3000"); err != nil {
			log.Instance.Fatal("Server failed to start",
				zap.Error(err),
			)
		}
	}()

	<-shutdownChan

	if err = app.Shutdown(); err != nil {
		log.Instance.Error("Server shutdown error",
			zap.Error(err),
		)
	}
}

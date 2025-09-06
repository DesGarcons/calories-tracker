package main

import (
	"ccb/config"
	app "ccb/internal/app"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	pathToEnv := flag.String("f", "./.env", "path to .env file")
	runApp := flag.Bool("app", false, "run app")

	flag.Parse()

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	// Configuration
	cfg, err := config.NewConfig(*pathToEnv)
	if err != nil {
		slog.Error(fmt.Sprintf("Config error: %w", err))
		panic(err)
	}

	ae, err := app.InitSE(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("Service entities init failed: %w", err))
		panic(err)
	}

	if *runApp {
		go app.RunApp(ae)
	}

	if !*runApp {
		stop()
	}

	<-ctx.Done()
}

package main

import (
	"time"

	_ "github.com/lib/pq"

	"github.com/Verifieddanny/sendstack-engine/internal/db"
	"github.com/Verifieddanny/sendstack-engine/internal/env"
	"go.uber.org/zap"
)

const currentVersion = "0.0.1"

func main() {
	cfg := config{
		addr:   env.GetEnvAsString("ADDR", ":8080"),
		env:    env.GetEnvAsString("ENV", "development"),
		apiUrl: env.GetEnvAsString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetEnvAsString("DATABASE_URL", ""),
			maxOpenConns: env.GetEnvAsInt("DM_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvAsInt("DM_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetEnvAsString("DM_MAX_IDLE_TIME", "15m"),
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("Database connection pool established")

	app := &application{
		config:          cfg,
		logger:          logger,
		showdownTimeout: 5 * time.Second,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}

package main

import (
	"ManogyaDahal/ecom/internal/env"
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"github.com/jackc/pgx/v5"
)

func main(){ 
	ctx := context.Background()
	err := godotenv.Load(); if err != nil { 
		log.Fatal("Failed to load .env, ",err)
	}

	cfg := config{ 
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil { 
		log.Fatal("Failed to establish db connection ",err)	
	}
	defer conn.Close(ctx)

	logger.Info("Connected to database", "dsn", cfg.db.dsn)

	api := &api{
		config: &cfg,
		db: conn,
	}

	
	// mounting and runing server
	if err := api.run(api.mount()); err != nil { 
		log.Fatal("server has failed to start, err: ",err)
	}
}

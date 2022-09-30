package main

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/kntajus/arraytest/internal/db/store"
	"github.com/kntajus/pgxtra/v4"
	"github.com/rs/zerolog/log"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://arraytest:postgres@arraytest-db:5432/arraytest?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer conn.Close(context.Background())

	err = pgxtra.RegisterEnumArrayTypes(conn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to register custom enum array types")
	}

	db := store.New(conn)
	err = db.AddChoice(context.Background(), []store.Fruit{store.FruitApple})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Info().Msg("Row inserted successfully")
}

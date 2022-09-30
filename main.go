package main

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/kntajus/arraytest/internal/db/store"
	"github.com/rs/zerolog/log"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://arraytest:postgres@arraytest-db:5432/arraytest?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer conn.Close(context.Background())

	var oid uint32
	name := "fruit"
	err = conn.QueryRow(context.Background(), "SELECT oid FROM pg_type WHERE typname = $1 AND typcategory = 'E'", name).Scan(&oid)
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't retrieve custom enum data")
	}
	conn.ConnInfo().RegisterDataType(pgtype.DataType{
		Value: &pgtype.EnumType{},
		Name:  name,
		OID:   oid,
	})

	err = conn.QueryRow(context.Background(), "SELECT oid, typname FROM pg_type WHERE typelem = $1 AND typcategory = 'A'", oid).Scan(&oid, &name)
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't retrieve custom enum array data")
	}
	conn.ConnInfo().RegisterDataType(pgtype.DataType{
		Value: &pgtype.EnumArray{},
		Name:  name,
		OID:   oid,
	})

	db := store.New(conn)
	err = db.AddChoice(context.Background(), []store.Fruit{store.FruitApple})
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	log.Info().Msg("Row inserted successfully")
}

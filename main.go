package main

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/kntajus/arraytest/internal/db/store"
	"github.com/rs/zerolog/log"
)

const retrieveCustomEnumArrayTypes = `SELECT arrtype.oid, arrtype.typname
FROM pg_type arrtype
JOIN pg_type enumtype ON arrtype.typelem = enumtype.oid AND arrtype.typcategory = 'A'
WHERE enumtype.typcategory = 'E'`

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://arraytest:postgres@arraytest-db:5432/arraytest?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	defer conn.Close(context.Background())

	err = registerEnumArrayTypes(conn)
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

func registerEnumArrayTypes(conn *pgx.Conn) error {
	enumTypes, err := fetchEnumArrayTypesInfo(conn)
	if err != nil {
		return err
	}

	for k, v := range enumTypes {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{Value: &pgtype.EnumArray{}, Name: v, OID: k})
	}
	return nil
}

func fetchEnumArrayTypesInfo(conn *pgx.Conn) (map[uint32]string, error) {
	rows, err := conn.Query(context.Background(), retrieveCustomEnumArrayTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	enumTypes := make(map[uint32]string)
	for rows.Next() {
		var oid uint32
		var name string
		rows.Scan(&oid, &name)
		enumTypes[oid] = name
	}
	return enumTypes, nil
}

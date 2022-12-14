// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"database/sql/driver"
	"fmt"
)

type Fruit string

const (
	FruitApple  Fruit = "apple"
	FruitBanana Fruit = "banana"
	FruitKiwi   Fruit = "kiwi"
)

func (e *Fruit) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Fruit(s)
	case string:
		*e = Fruit(s)
	default:
		return fmt.Errorf("unsupported scan type for Fruit: %T", src)
	}
	return nil
}

type NullFruit struct {
	Fruit Fruit
	Valid bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFruit) Scan(value interface{}) error {
	if value == nil {
		ns.Fruit, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Fruit.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFruit) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Fruit, nil
}

type Choice struct {
	ChoiceID int32
	Fruits   []Fruit
}

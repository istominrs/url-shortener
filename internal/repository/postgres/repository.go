package postgres

import "github.com/uptrace/bun"

type PostgresRepo struct {
	con *bun.DB
}

func New(c *bun.DB) *PostgresRepo {
	return &PostgresRepo{
		con: c,
	}
}

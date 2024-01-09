package postgres

import "github.com/uptrace/bun"

type url struct {
	bun.BaseModel `bun:"url"`

	Id    int    `bun:"id"`
	Alias string `bun:"alias"`
	Url   string `bun:"url"`
}

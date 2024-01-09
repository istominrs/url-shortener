package postgres

import (
	"context"
	"fmt"
)

func (p *PostgresRepo) SaveURL(ctx context.Context, urlToSave string, alias string) error {
	const op = "repository.postgres.save_url.SaveURL"

	u := url{
		Alias: alias,
		Url:   urlToSave,
	}

	_, err := p.con.NewInsert().Model(&u).Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

package postgres

import (
	"context"
	"fmt"
)

func (p *PostgresRepo) DeleteURL(ctx context.Context, alias string) error {
	const op = "repository.postgres.delete_url.DeleteURL"

	var u url

	_, err := p.con.NewDelete().Model(&u).Where("alias = ?", alias).Exec(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

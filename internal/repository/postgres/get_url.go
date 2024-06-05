package postgres

import (
	"context"
	"fmt"
)

func (p *PostgresRepo) GetURL(ctx context.Context, alias string) (string, error) {
	const op = "repository.postgres.get_url.GetURL"

	var u url

	err := p.con.NewSelect().Model(&u).Where("alias = ?", alias).Scan(ctx)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return u.Url, nil
}

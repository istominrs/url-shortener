package usecase

import "context"

func (uc *UseCase) DeleteURL(ctx context.Context, alias string) error {
	return uc.urlCase.DeleteURL(ctx, alias)
}

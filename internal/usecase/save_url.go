package usecase

import "context"

func (uc *UseCase) GetURL(ctx context.Context, alias string) (string, error) {
	return uc.urlCase.GetURL(ctx, alias)
}

package usecase

import "context"

type Url interface {
	SaveURL(context.Context, string, string) error
	GetURL(context.Context, string) (string, error)
	DeleteURL(context.Context, string) error
}

type UseCase struct {
	urlCase Url
}

func New(u Url) *UseCase {
	return &UseCase{
		urlCase: u,
	}
}

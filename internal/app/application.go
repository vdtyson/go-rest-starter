package app

import "context"

type Application struct {
	UseCases *UseCases
}

func New(ctx context.Context) (*Application, error) {
	useCases := &UseCases{}
	return &Application{UseCases: useCases}, nil
}

type UseCases struct {
}
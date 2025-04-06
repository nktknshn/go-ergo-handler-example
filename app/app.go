package app

import (
	"context"

	"github.com/nktknshn/go-ergo-handler-example/app/adapters/http_adapter"
	"github.com/nktknshn/go-ergo-handler-example/app/repositories"
	"github.com/nktknshn/go-ergo-handler-example/app/use_cases"
)

type App struct {
	repositories *repositories.Repositories
	useCases     *use_cases.UseCases
	httpAdapter  *http_adapter.HttpAdapter
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init(ctx context.Context) error {
	a.repositories = repositories.New()
	err := a.repositories.Init(ctx)
	if err != nil {
		return err
	}
	a.useCases = use_cases.New()
	err = a.useCases.Init(ctx, a.repositories)
	if err != nil {
		return err
	}
	a.httpAdapter = http_adapter.New()
	err = a.httpAdapter.Init(ctx, a.useCases)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Start(ctx context.Context) error {
	return nil
}

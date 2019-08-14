package controller

import (
	"context"

	"github.com/Khamliuk/testsCI/model"
	"github.com/stretchr/testify/mock"
)

type personProviderMock struct {
	mock.Mock
}

func (p *personProviderMock) Create(ctx context.Context, req model.Person) (*model.Person, error) {
	args := p.Called(ctx, req)
	return args.Get(0).(*model.Person), args.Error(1)
}

func (p *personProviderMock) List(ctx context.Context) ([]model.Person, error) {
	args := p.Called(ctx)
	return args.Get(0).([]model.Person), args.Error(1)
}

func (p *personProviderMock) Update(ctx context.Context, req model.Person) error {
	args := p.Called(ctx, req)
	return args.Error(0)
}

func (p *personProviderMock) Delete(ctx context.Context, id string) error {
	args := p.Called(ctx, id)
	return args.Error(0)
}

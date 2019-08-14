package controller

import (
	"context"
	"fmt"

	"github.com/rockspoon/testsCI/model"
)

type Controller struct {
	db PersonProvider
}

type PersonProvider interface {
	Create(ctx context.Context, req model.Person) (*model.Person, error)
	List(ctx context.Context) ([]model.Person, error)
	Update(ctx context.Context, req model.Person) error
	Delete(ctx context.Context, id string) error
}

func New(db PersonProvider) Controller {
	return Controller{
		db: db,
	}
}

func (t Controller) Create(ctx context.Context, req model.Person) (*model.Person, error) {
	resp, err := t.db.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not create new person: %v", err)
	}
	return resp, nil
}

func (t Controller) List(ctx context.Context) ([]model.Person, error) {
	resp, err := t.db.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not find any person: %v", err)
	}
	return resp, nil
}

func (t Controller) Update(ctx context.Context, req model.Person) error {
	err := t.db.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("could not update person: %v", err)
	}
	return nil
}

func (t Controller) Delete(ctx context.Context, id string) error {
	err := t.db.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("could not delete person: %v", err)
	}
	return nil
}

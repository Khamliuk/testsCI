package controller

import (
	"context"
	"fmt"
	"testing"

	"github.com/Khamliuk/testsCI/model"
	"github.com/stretchr/testify/assert"
)

func TestController_Create(t *testing.T) {
	tt := []struct {
		name     string
		req      model.Person
		resp     *model.Person
		fn       func(p *personProviderMock)
		expError error
	}{
		{
			name: "create error",
			req: model.Person{
				FirstName: "Nikita",
				LastName:  "Khamliuk",
				Age:       22,
			},
			resp: nil,
			fn: func(p *personProviderMock) {
				p.On("Create", context.Background(), model.Person{
					FirstName: "Nikita",
					LastName:  "Khamliuk",
					Age:       22,
				}).Return(&model.Person{}, fmt.Errorf(""))
			},
			expError: fmt.Errorf("could not create new person: %v", fmt.Errorf("")),
		},
		{
			name: "all ok",
			req: model.Person{
				FirstName: "Nikita",
				LastName:  "Khamliuk",
				Age:       22,
			},
			resp: &model.Person{},
			fn: func(p *personProviderMock) {
				p.On("Create", context.Background(), model.Person{
					FirstName: "Nikita",
					LastName:  "Khamliuk",
					Age:       22,
				}).Return(&model.Person{}, nil)
			},
			expError: nil,
		},
	}
	for _, tc := range tt {
		testPersonProvider := new(personProviderMock)
		ss := New(testPersonProvider)
		tc.fn(testPersonProvider)
		t.Run(tc.name, func(t *testing.T) {
			resp, err := ss.Create(context.Background(), tc.req)
			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.resp, resp)
		})
	}
}

func TestController_Update(t *testing.T) {
	tt := []struct {
		name     string
		req      model.Person
		fn       func(p *personProviderMock)
		expError error
	}{
		{
			name: "update error",
			req: model.Person{
				FirstName: "Nikita",
				LastName:  "Khamliuk",
				Age:       22,
			},
			fn: func(p *personProviderMock) {
				p.On("Update", context.Background(), model.Person{
					FirstName: "Nikita",
					LastName:  "Khamliuk",
					Age:       22,
				}).Return(fmt.Errorf(""))
			},
			expError: fmt.Errorf("could not update person: %v", fmt.Errorf("")),
		},
		{
			name: "all ok",
			req: model.Person{
				FirstName: "Nikita",
				LastName:  "Khamliuk",
				Age:       22,
			},
			fn: func(p *personProviderMock) {
				p.On("Update", context.Background(), model.Person{
					FirstName: "Nikita",
					LastName:  "Khamliuk",
					Age:       22,
				}).Return(nil)
			},
			expError: nil,
		},
	}
	for _, tc := range tt {
		testPersonProvider := new(personProviderMock)
		ss := New(testPersonProvider)
		tc.fn(testPersonProvider)
		t.Run(tc.name, func(t *testing.T) {
			err := ss.Update(context.Background(), tc.req)
			assert.Equal(t, tc.expError, err)
		})
	}
}

func TestController_List(t *testing.T) {
	tt := []struct {
		name     string
		resp     []model.Person
		fn       func(p *personProviderMock)
		expError error
	}{
		{
			name: "list error",
			fn: func(p *personProviderMock) {
				p.On("List", context.Background()).Return([]model.Person{}, fmt.Errorf(""))
			},
			expError: fmt.Errorf("could not find any person: %v", fmt.Errorf("")),
		},
		{
			name: "all ok",
			fn: func(p *personProviderMock) {
				p.On("List", context.Background()).Return([]model.Person{}, nil)
			},
			resp:     []model.Person{},
			expError: nil,
		},
	}
	for _, tc := range tt {
		testPersonProvider := new(personProviderMock)
		ss := New(testPersonProvider)
		tc.fn(testPersonProvider)
		t.Run(tc.name, func(t *testing.T) {
			resp, err := ss.List(context.Background())
			assert.Equal(t, tc.expError, err)
			assert.Equal(t, tc.resp, resp)
		})
	}
}

func TestController_Delete(t *testing.T) {
	tt := []struct {
		name     string
		id       string
		fn       func(p *personProviderMock)
		expError error
	}{
		{
			name: "delete error",
			id:   "1",
			fn: func(p *personProviderMock) {
				p.On("Delete", context.Background(), "1").Return(fmt.Errorf(""))
			},
			expError: fmt.Errorf("could not delete person: %v", fmt.Errorf("")),
		},
		{
			name: "all ok",
			id:   "2",
			fn: func(p *personProviderMock) {
				p.On("Delete", context.Background(), "2").Return(nil)
			},
			expError: nil,
		},
	}
	for _, tc := range tt {
		testPersonProvider := new(personProviderMock)
		ss := New(testPersonProvider)
		tc.fn(testPersonProvider)
		t.Run(tc.name, func(t *testing.T) {
			err := ss.Delete(context.Background(), tc.id)
			assert.Equal(t, tc.expError, err)
		})
	}
}

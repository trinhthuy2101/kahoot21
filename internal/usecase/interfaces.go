// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"ecommerce/customer/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Customer interface {
		Get(ctx context.Context, id int64) (*entity.Customer, error)
		Create(context.Context, *entity.Customer) error
		Collection(context.Context) ([]entity.Customer, int64, error)
	}

	CustomerRepo interface {
		CreateOne(context.Context, *entity.Customer) error
		UpdateOne(context.Context, *entity.Customer) error
		GetOne(context.Context, int64) (*entity.Customer, error)
		GetAll(context.Context, *int64) ([]entity.Customer, error)
		DeleteOne(context.Context, int64) error
	}

	CustomerWebAPI interface {
		Translate(*entity.Customer) (*entity.Customer, error)
	}
)

package usecase

import (
	"context"

	"ecommerce/customer/internal/entity"
)

// CustomerUseCase -.
type CustomerUseCase struct {
	repo CustomerRepo
}

func NewCustomer(r CustomerRepo) *CustomerUseCase {
	return &CustomerUseCase{
		repo: r,
	}
}

func (uc *CustomerUseCase) Create(ctx context.Context, customer *entity.Customer) error {
	return uc.repo.CreateOne(ctx, customer)
}

func (uc *CustomerUseCase) Get(ctx context.Context, id int64) (*entity.Customer, error) {
	return uc.repo.GetOne(ctx, id)
}

func (uc *CustomerUseCase) Collection(ctx context.Context) (result []entity.Customer, count int64, err error) {
	result, err = uc.repo.GetAll(ctx, &count)

	return
}

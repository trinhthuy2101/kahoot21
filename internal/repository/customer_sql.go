package repository

import (
	"context"
	"time"

	"gorm.io/gorm"

	"ecommerce/customer/internal/entity"
	"ecommerce/customer/internal/usecase"
)

var _ usecase.CustomerRepo = (*CustomerRepo)(nil)

type CustomerRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db}
}

func (r *CustomerRepo) withContext(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).Model(&entity.Customer{})
}

func (r *CustomerRepo) GetAll(ctx context.Context, count *int64) (customers []entity.Customer, err error) {
	db := r.withContext(ctx).Find(&customers).Limit(10).Scan(&customers) //nolint:gomnd // todo

	err = db.Error
	if err != nil {
		return
	}

	if count != nil {
		db.Count(count)
	}

	return
}

func (r *CustomerRepo) CreateOne(ctx context.Context, input *entity.Customer) error {
	input.CreatedAt = time.Now().Format(time.RFC3339)
	input.UpdatedAt = time.Now().Format(time.RFC3339)

	return r.withContext(ctx).Create(input).Error
}

func (r *CustomerRepo) GetOne(ctx context.Context, id int64) (customer *entity.Customer, err error) {
	err = r.withContext(ctx).First(&customer, id).Error

	return
}

func (r *CustomerRepo) UpdateOne(ctx context.Context, input *entity.Customer) error {
	return r.withContext(ctx).Updates(&input).Error
}

func (r *CustomerRepo) DeleteOne(ctx context.Context, id int64) error {
	return r.withContext(ctx).Delete(&entity.Customer{}).Where("id = ?", id).Error
}

package grpc

import (
	"github.com/uchin-mentorship/ecommerce-go/customer"

	"ecommerce/customer/internal/entity"
)

func transformCustomerCollectionResponseData(items []entity.Customer) []*customer.BaseCustomer {
	var result = make([]*customer.BaseCustomer, 0, len(items))
	for _, value := range items {
		result = append(result, &customer.BaseCustomer{
			Id:        value.ID,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Gender:    value.Gender,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return result
}

func transformCustomerGetResponseData(items *entity.Customer) *customer.BaseCustomer {
	return &customer.BaseCustomer{
		Id:        items.ID,
		FirstName: items.FirstName,
		LastName:  items.LastName,
		Gender:    items.Gender,
		Email:     items.Email,
		CreatedAt: items.CreatedAt,
		UpdatedAt: items.UpdatedAt,
	}
}

func transformCustomerCreateRequestData(request *customer.BaseCustomer) (result *entity.Customer) {
	return &entity.Customer{
		ID:        request.Id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Gender:    request.Gender,
		Email:     request.Email,
	}
}

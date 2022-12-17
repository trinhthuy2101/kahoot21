package grpc

import (
	"context"

	"github.com/uchin-mentorship/ecommerce-go/customer"

	"ecommerce/customer/internal/usecase"
	"ecommerce/customer/pkg/errors"
	"ecommerce/customer/pkg/logger"
)

type customerService struct {
	customer.UnimplementedCustomerServiceServer
	u usecase.Customer
}

var _ customer.CustomerServiceServer = (*customerService)(nil)

func (s *customerService) Create(ctx context.Context, req *customer.BaseCustomer) (result *customer.CreateCustomerResponse, err error) {
	reqEntity := transformCustomerCreateRequestData(req)

	err = s.u.Create(ctx, reqEntity)
	if err != nil {
		logger.Error("failed to create Customer. Error %w", err)

		result = &customer.CreateCustomerResponse{
			Status:  StatusFailure,
			Message: errors.ErrGeneral.Error(),
			TraceId: "",
		}

		return
	}

	result = &customer.CreateCustomerResponse{
		Status:  StatusSuccess,
		Message: MessageSuccess,
		TraceId: "",
	}

	return
}

func (s *customerService) Collection(ctx context.Context, _ *customer.Empty) (result *customer.CollectionResponse, err error) {
	items, count, err := s.u.Collection(ctx)
	if err != nil {
		logger.Error("failed to get collection. Error %w", err)

		result = &customer.CollectionResponse{
			Status:  StatusFailure,
			Message: errors.ErrGeneral.Error(),
			TraceId: "",
		}

		return result, nil
	}

	result = &customer.CollectionResponse{
		Status:  StatusSuccess,
		Message: MessageSuccess,
		TraceId: "",
		Data: &customer.CollectionResponse_Data{
			Results: transformCustomerCollectionResponseData(items),
			Pagination: &customer.Pagination{
				TotalPages: count,
			},
		},
	}

	return
}

func (s *customerService) GetByID(ctx context.Context, req *customer.GetByIDRequest) (result *customer.BaseCustomer, err error) {
	resp, err := s.u.Get(ctx, req.Id)
	if err != nil {
		return
	}

	result = transformCustomerGetResponseData(resp)

	return
}

func NewCustomerService(u usecase.Customer) customer.CustomerServiceServer {
	return &customerService{
		u: u,
	}
}

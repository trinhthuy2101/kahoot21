package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"ecommerce/customer/internal/usecase"
	"ecommerce/customer/pkg/errors"
	"ecommerce/customer/pkg/logger"
	"ecommerce/customer/pkg/response"
)

type customerRoutes struct {
	t usecase.Customer
}

func newCustomerRoutes(handler *gin.RouterGroup, t usecase.Customer) {
	r := &customerRoutes{t}

	h := handler.Group("/customers")
	{
		h.GET("/", response.GinWrap(r.collection))
		h.GET("/:id", response.GinWrap(r.get))
	}
}

// @Summary     Get customers
// @Description Customer collection
// @Tags  	    Customer
// @Accept      json
// @Produce     json
// @Success     200 {object} response.Response{data=response.Collection{results=[]entity.Customer}}
// @Failure     500 {object} response.FailureResponse
// @Router      /v1/customers [get]
func (r *customerRoutes) collection(c *gin.Context) *response.Response {
	customers, count, err := r.t.Collection(c.Request.Context())
	if err != nil {
		logger.Error("failed to get collection. Error %w", err)

		return response.Failure(err)
	}

	return response.SuccessWithCollection(
		customers,
		&response.Pagination{
			TotalRecords: count,
		})
}

// @Summary     Get customer
// @Description Get customer by id
// @Tags  	    Customer
// @Accept      json
// @Produce     json
// @Param    		id path int  true  "Customer ID"
// @Success     200 {object} response.Response{data=response.Data{result=entity.Customer}}
// @Failure     500 {object} response.FailureResponse
// @Router      /v1/customers/{id} [get]
func (r *customerRoutes) get(c *gin.Context) *response.Response {
	id := cast.ToInt64(c.Param("id"))
	if id == 0 {
		err := errors.ErrInvalidCustomerID

		return response.Failure(err)
	}

	result, err := r.t.Get(c, id)
	if err != nil {
		logger.Error(err)
		err = errors.ErrCustomerNotFound

		return response.Failure(err)
	}

	return response.SuccessWithData(result)
}

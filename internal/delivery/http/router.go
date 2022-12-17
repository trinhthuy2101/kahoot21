// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	v1 "ecommerce/customer/internal/delivery/http/v1"
	"ecommerce/customer/internal/usecase"
)

type Router struct {
	handler *gin.Engine
	t       usecase.Customer
}

func (r *Router) Register() {
	// Options
	r.handler.Use(gin.Logger())
	r.handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	r.handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	r.handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	r.handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	v1.RegisterRouter(r.handler.Group("/v1"), r.t)
}

func NewRouter(handler *gin.Engine, t usecase.Customer) *Router {
	return &Router{
		handler: handler,
		t:       t,
	}
}

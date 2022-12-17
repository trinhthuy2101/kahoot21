//go:build wireinject
// +build wireinject

package di

import (
	"examples/kahootee/config"
	grpcDelivery "examples/kahootee/internal/delivery/grpc"
	"examples/kahootee/internal/delivery/http"
	"examples/kahootee/internal/repository"
	"examples/kahootee/internal/usecase"
	"examples/kahootee/pkg/grpcserver"
	"examples/kahootee/pkg/httpserver"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var useCaseSet = wire.NewSet(config.NewConfig, provideGormDB, provideCustomerUseCase, provideCustomerRepo)

func InitializeHttpServer() (*httpserver.Server, func(), error) {
	panic(wire.Build(
		useCaseSet,
		gin.New,
		provideHttpServer,
		http.NewRouter,
	))
}

// func InitializeGRPCServer() (*grpcserver.GRPCServer, func(), error) {
// 	panic(wire.Build(
// 		useCaseSet,
// 		grpc.NewServer,
// 		provideGRPCServerOptions,
// 		provideGRPCServer,
// 		provideGRPCCustomerService,
// 	))
// }

// func provideGRPCCustomerService(u usecase.Customer) customer.CustomerServiceServer {
// 	return grpcDelivery.NewCustomerService(u)
// }

// func provideGRPCServerOptions() []grpc.ServerOption {
// 	return nil
// }

// func provideGRPCServer(cfg *config.Config, server *grpc.Server, delivery customer.CustomerServiceServer) *grpcserver.GRPCServer {
// 	customer.RegisterCustomerServiceServer(server, delivery)
// 	return grpcserver.New(server, cfg.GRPC.Address)
// }

// func provideCustomerRepo(db *gorm.DB) usecase.CustomerRepo {
// 	return repository.New(db)
// }

// func provideCustomerUseCase(r usecase.CustomerRepo) usecase.Customer {
// 	return usecase.NewCustomer(r)
// }

func provideGormDB(cfg *config.Config) (*gorm.DB, func(), error) {
	db, err := gorm.Open(postgres.Open(cfg.PG.URL), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return db, func() {
		conn, err := db.DB()
		if err != nil {
			log.Printf("failed to get db connection, %v", err)
			return
		}
		conn.Close()
	}, nil
}

func provideHttpServer(router *http.Router, handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	router.Register()
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}

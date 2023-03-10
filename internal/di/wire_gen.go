// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"examples/kahootee/config"
	"examples/kahootee/internal/delivery/http"
	"examples/kahootee/internal/repository"
	"examples/kahootee/internal/service/jwthelper"
	"examples/kahootee/internal/usecase"
	"examples/kahootee/pkg/httpserver"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Injectors from wire.go:

func InitializeHttpServer() (*httpserver.Server, func(), error) {
	engine := gin.New()
	specification := provideConfig()
	jwtHelper := provideJWTService(specification)
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := provideGormDB(configConfig)
	if err != nil {
		return nil, nil, err
	}
	authRepo := provideAuthRepo(db)
	authUsecase := provideAuthUseCase(authRepo, jwtHelper)
	kahootRepo := provideKahootRepo(db)
	kahootUsecase := provideKahootUseCase(kahootRepo)
	groupRepo := provideGroupRepo(db)
	groupUsecase := provideGroupUseCase(groupRepo)
	router := http.NewRouter(engine, jwtHelper, authUsecase, kahootUsecase, groupUsecase)
	server := provideHttpServer(router, engine, configConfig)
	return server, func() {
		cleanup()
	}, nil
}

// wire.go:

var useCaseSet = wire.NewSet(config.NewConfig, provideGormDB, provideKahootUseCase, provideKahootRepo, provideGroupUseCase, provideGroupRepo, provideAuthUseCase, provideAuthRepo, provideJWTService, provideConfig)

func provideKahootRepo(db *gorm.DB) usecase.KahootRepo {
	return repo.NewKahootRepo(db)
}

func provideKahootUseCase(k usecase.KahootRepo) usecase.KahootUsecase {
	return usecase.NewKahootUsecase(k)
}

func provideGroupRepo(db *gorm.DB) usecase.GroupRepo {
	return repo.NewGroupRepo(db)
}

func provideGroupUseCase(g usecase.GroupRepo) usecase.GroupUsecase {
	return usecase.NewGroupUsecase(g)
}

func provideAuthRepo(db *gorm.DB) usecase.AuthRepo {
	return repo.NewAuthRepo(db)
}

func provideAuthUseCase(g usecase.AuthRepo, jwtService service.JWTHelper) usecase.AuthUsecase {
	return usecase.NewAuthUsecase(g, jwtService)
}

func provideJWTService(s *config.Specification) service.JWTHelper {
	return service.NewJWTService(s)
}

func provideConfig() *config.Specification {
	return config.LoadEnvConfig()
}

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

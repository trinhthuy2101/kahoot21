package usecase

import repo "examples/kahootee/internal/repository"

type kahootUsecase struct {
	repo repo.KahootRepo
}


func NewKahootUsecase(repo repo.KahootRepo) KahootUsecase {
	return &kahootUsecase{
		repo: repo,
	}

}

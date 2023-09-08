package usecase

import "github.com/halosatrio/bebop/service"

type generalUseCase struct {
	generalRepository service.GeneralRepository
}

func (uc generalUseCase) Welcome() map[string]string {
	return uc.generalRepository.Welcome()
}

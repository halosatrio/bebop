package service

type GeneralUseCase interface {
	Welcome() (result map[string]string)
}

package service

type GeneralRepository interface {
	Welcome() (result map[string]string)
}

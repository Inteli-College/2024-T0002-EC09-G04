package usecase

import (
	"github.com/Inteli-College/2024-T0002-EC09-G04/backend/internal/domain/entity"
)

type TransmitUseCase struct {
	LogRepository entity.LogRepository
}

func NewTransmitUseCase(logRepository entity.LogRepository) *TransmitUseCase {
	return &TransmitUseCase{LogRepository: logRepository}
}

func (t *TransmitUseCase) Execute(log *entity.Log) error {
	err := t.LogRepository.Transmit(log)
	if err != nil {
		return err
	}
	return nil
}
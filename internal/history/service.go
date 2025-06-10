package history

import (
	"event-driven-sample/pkg/entity"
	"event-driven-sample/pkg/mongodb"
	"fmt"
)

type Service struct {
	Repository *mongodb.MongoDB
}

func NewService(repository *mongodb.MongoDB) *Service {
	return &Service{repository}
}

func (s Service) SaveCalculation(calculation entity.Calculation) error {
	if err := s.Repository.Save(calculation); err != nil {
		return fmt.Errorf("failed to save calculation: %w", err)
	}
	return nil
}

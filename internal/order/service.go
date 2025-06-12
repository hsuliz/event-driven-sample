package order

import (
	"encoding/json"
	"event-driven-sample/pkg/entity"
	"event-driven-sample/pkg/hash"
	"event-driven-sample/pkg/kafka"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
)

type Service struct {
	Repository    *Repository
	KafkaProducer *kafka.Producer
}

func NewService(repository *Repository, producer *kafka.Producer) *Service {
	return &Service{repository, producer}
}

func (s Service) SaveCalculation(calculation entity.Calculation) error {
	if err := s.Repository.Save(calculation); err != nil {
		return fmt.Errorf("failed to save calculation: %w", err)
	}
	return nil
}

func (s Service) ProcessMatrix(matrix [][]int) error {
	marshaledMatrix, err := json.Marshal(matrix)
	if err != nil {
		log.Fatal(err)
	}

	// just for hash i dont care
	flattenMatrixInt := slices.Concat(matrix...)
	flattenMatrixIntMarshalled, err := json.Marshal(flattenMatrixInt)
	if err != nil {
		log.Fatalln(err)
	}

	matrixHash := hash.Encode(flattenMatrixIntMarshalled)
	log.Println("saving and sending matrix with hash ", matrixHash)

	if err := s.Repository.Save(entity.Calculation{
		Hash:  matrixHash,
		Done:  false,
		Value: 0,
	}); err != nil {
		return fmt.Errorf("failed to save: %w", err)
	}

	if err := s.KafkaProducer.SendMessage(marshaledMatrix); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (s Service) GenerateMatrix(size int) [][]int {
	matrix := MakeMatrix(size)
	var wg sync.WaitGroup
	for i := range size {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.PopulateRow(&matrix[i])
		}()
	}
	wg.Wait()
	return matrix
}

func MakeMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return matrix
}

func (s Service) PopulateRow(row *[]int) {
	for i := range *row {
		(*row)[i] = rand.Int() % 100
	}
}

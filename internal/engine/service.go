package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"event-driven-sample/pkg/kafka"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"slices"
)

type Service struct {
	Producer *Producer
}

func NewService(producer *Producer) *Service {
	return &Service{producer}
}

func (s Service) Process(hash string, determinant int) error {
	if err := s.Producer.SendMessage(kafka.EngineMsg{
		Hash:  hash,
		Done:  true,
		Value: determinant,
	}); err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

func (s Service) Calculate(matrix [][]int) (string, int, error) {
	flattenMatrixInt := slices.Concat(matrix...)

	// START MONSTROSITY
	flattenMatrixIntMarshalled, err := json.Marshal(flattenMatrixInt)
	if err != nil {
		log.Fatalln(err)
	}

	matrixHashByte := sha256.Sum256(flattenMatrixIntMarshalled)
	matrixHash := hex.EncodeToString(matrixHashByte[:])

	if err := s.Producer.SendMessage(kafka.EngineMsg{
		Hash:  matrixHash,
		Done:  false,
		Value: 0,
	}); err != nil {
		return "", 0, fmt.Errorf("failed to send message: %v", err)
	}

	var flattenMatrixFloat []float64
	if err := json.Unmarshal(flattenMatrixIntMarshalled, &flattenMatrixFloat); err != nil {
		log.Fatalln(err)
	}
	// END MONSTROSITY

	n := len(matrix)
	det := mat.Det(mat.NewDense(n, n, flattenMatrixFloat))
	return matrixHash, int(det), nil
}

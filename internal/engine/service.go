package engine

import (
	"encoding/json"
	"event-driven-sample/pkg/hash"
	"event-driven-sample/pkg/kafka"
	"gonum.org/v1/gonum/mat"
	"log"
	"slices"
)

type Service struct {
	Producer *kafka.Producer
}

func NewService(producer *kafka.Producer) *Service {
	return &Service{producer}
}

func (s Service) Process(matrixHash string, determinant int) error {
	msg := kafka.EngineMsg{
		MatrixHash: matrixHash,
		Done:       true,
		Value:      determinant,
	}

	marshaledMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	log.Println("sending calculated matrix with hash", matrixHash)
	if err := s.Producer.SendMessage(marshaledMsg); err != nil {
		return err
	}
	return nil
}

func (s Service) Calculate(matrix [][]int) (string, int, error) {
	// START MONSTROSITY
	flattenMatrixInt := slices.Concat(matrix...)
	flattenMatrixIntMarshalled, err := json.Marshal(flattenMatrixInt)
	if err != nil {
		log.Fatalln(err)
	}
	matrixHash := hash.Encode(flattenMatrixIntMarshalled)

	var flattenMatrixFloat []float64
	if err := json.Unmarshal(flattenMatrixIntMarshalled, &flattenMatrixFloat); err != nil {
		log.Fatalln(err)
	}
	// END MONSTROSITY

	n := len(matrix)
	det := mat.Det(mat.NewDense(n, n, flattenMatrixFloat))
	return matrixHash, int(det), nil
}

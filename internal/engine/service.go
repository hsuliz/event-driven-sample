package engine

import (
	"encoding/json"
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

func (s Service) Save(determinant int) error {
	log.Println("saved", determinant)
	if err := s.Producer.SendMessage(determinant); err != nil {
		return fmt.Errorf("server failed to send message: %v", err)
	}
	return nil
}

func (s Service) Calculate(matrix [][]int) int {
	flattenMatrixInt := slices.Concat(matrix...)

	// START MONSTROSITY
	flattenMatrixIntMarshalled, err := json.Marshal(flattenMatrixInt)
	if err != nil {
		log.Fatalln(err)
	}
	var flattenMatrixFloat []float64
	if err := json.Unmarshal(flattenMatrixIntMarshalled, &flattenMatrixFloat); err != nil {
		log.Fatalln(err)
	}
	// END MONSTROSITY

	n := len(matrix)
	det := mat.Det(mat.NewDense(n, n, flattenMatrixFloat))
	return int(det)
}

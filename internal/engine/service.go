package engine

import (
	"encoding/json"
	"gonum.org/v1/gonum/mat"
	"log"
	"slices"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) CalculateDeterminant(matrix [][]int) float64 {
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
	return det
}

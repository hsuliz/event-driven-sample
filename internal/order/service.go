package order

import (
	"math/rand"
	"sync"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
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

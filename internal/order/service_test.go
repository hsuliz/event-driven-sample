package order

import (
	"fmt"
	"testing"
)

var service = NewService()

func TestGenerateMatrix(t *testing.T) {
	fmt.Println(service.GenerateMatrix(100))
}

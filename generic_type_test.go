package belajargolanggenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Surya", "Yasin"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag[int](numbers)
}

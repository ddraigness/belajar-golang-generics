package belajargolanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}


func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func TestData(t *testing.T) {
	data := Data[string]{
		First: "Surya",
		Second: "Yasin",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First: "Surya",
		Second: "Yasin",
	}
	assert.Equal(t, "Febry", data.ChangeFirst("Febry"))
	assert.Equal(t, "Hello Surya", data.SayHello("Surya"))
}
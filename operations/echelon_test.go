package operations_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linear-algebra/matrix"
	"github.com/linear-algebra/operations"
)

func Test_EchelonForm_LIMatrix(t *testing.T) {
	m := matrix.New(3, 3, [][]float64{{1, 2, 4}, {1, 1, 3}, {1, 3, 1}})
	reduced := operations.EchelonForm(m)
	log.Println(m.String())
	log.Println(reduced.String())
	assert.Equal(t, []float64{1, 2, 4}, reduced.Row(0))
	assert.Equal(t, []float64{0, -1, -1}, reduced.Row(1))
	assert.Equal(t, []float64{0, 0, -4}, reduced.Row(2))
}

func Test_EchelonForm_LDMatrix(t *testing.T) {
	m := matrix.New(3, 3, [][]float64{{1, 2, 4}, {2, 4, 8}, {1, 3, 3}})
	reduced := operations.EchelonForm(m)
	log.Println("\n" + m.String())
	log.Println(reduced.String())
	assert.Equal(t, []float64{1, 2, 4}, reduced.Row(0))
	assert.Equal(t, []float64{0, 1, -1}, reduced.Row(1))
	assert.Equal(t, []float64{0, 0, 0}, reduced.Row(2))
}

func Test_EchelonForm_OutOfOrderMatrix(t *testing.T) {
	m := matrix.New(4, 3, [][]float64{{1, 2, 3}, {2, 4, 8}, {1, 3, 3}, {3, 4, 2}})
	reduced := operations.EchelonForm(m)
	log.Println("\n" + m.String())
	log.Println(reduced.String())
	assert.Equal(t, []float64{1, 2, 3}, reduced.Row(0))
	assert.Equal(t, []float64{0, 1, 0}, reduced.Row(1))
	assert.Equal(t, []float64{0, 0, 2}, reduced.Row(2))
	assert.Equal(t, []float64{0, 0, 0}, reduced.Row(3))
}

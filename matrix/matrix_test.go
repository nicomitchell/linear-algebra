package matrix_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linear-algebra/matrix"
)

var (
	m = matrix.New(3, 3, [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
)

func Test_At_Works(t *testing.T) {
	v := m.At(1, 1)
	assert.Equal(t, 5.0, v)
}

func Test_String_Works(t *testing.T) {
	s := m.String()
	log.Printf(s)
	assert.Equal(t, "|\t 1.000 2.000 3.000 \t|\n|\t 4.000 5.000 6.000 \t|\n|\t 7.000 8.000 9.000 \t|\n", s)
}

func Test_Row_Works(t *testing.T) {
	r := m.Row(0)
	assert.Equal(t, []float64{1, 2, 3}, r)
}
func Test_Column_Works(t *testing.T) {
	c := m.Column(1)
	assert.Equal(t, []float64{2, 5, 8}, c)
}

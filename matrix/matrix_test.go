package matrix_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linear-algebra/matrix"
)

func Test_String_Works(t *testing.T) {
	m := matrix.New(3, 3, [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	s := m.String()
	log.Printf(s)
	assert.Equal(t, "|\t 1.000 2.000 3.000 \t|\n|\t 4.000 5.000 6.000 \t|\n|\t 7.000 8.000 9.000 \t|\n", s)
}

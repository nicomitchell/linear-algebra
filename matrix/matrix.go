package matrix

import (
	"fmt"
)

//Matrix represents an mxn matrix
type Matrix struct {
	m      int
	n      int
	values [][]float64
}

//New returns a new matrix
func New(m, n int, values [][]float64) *Matrix {
	return &Matrix{
		m:      m,
		n:      n,
		values: values,
	}
}

//At returns the value in matrix
func (m *Matrix) At(row, col int) float64 {
	return m.values[row][col]
}

func (m *Matrix) String() string {
	var out string
	for _, row := range m.values {
		out += "|\t"
		for _, val := range row {
			out = fmt.Sprintf("%s %.3f", out, val)
		}
		out += " \t|\n"
	}
	return out
}

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

//At returns the value at the given position in the matrix
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

//Row returns the row at the given position
func (m *Matrix) Row(row int) []float64 {
	return m.values[row]
}

//Column returns the column at the given position
func (m *Matrix) Column(col int) []float64 {
	var out []float64
	for i := 0; i < m.m; i++ {
		out = append(out, m.values[i][col])
	}
	return out
}

//Width returns the number of columns in the matrix
func (m *Matrix) Width() int {
	return m.n
}

//Height returns the number of rows in the matrix
func (m *Matrix) Height() int {
	return m.m
}

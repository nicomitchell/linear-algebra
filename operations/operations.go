package operations

import (
	"github.com/linear-algebra/matrix"
)

//EchelonForm returns the matrix m in Echelon form
func EchelonForm(m *matrix.Matrix) *matrix.Matrix {
	h := m.Height()
	w := m.Width()
	newRows := make([][]float64, h)
	for i := 0; i < h; i++ {
		newRows[i] = m.Row(i)
	}
	newRows = echelonArithmetic(newRows, h, w)
	newRows = echelonRowSwap(newRows, h, w)
	return matrix.New(h, w, newRows)
}

func subtractRow(r1, r2 []float64) []float64 {
	out := make([]float64, len(r1))
	for i := range r1 {
		out[i] = r1[i] - r2[i]
		//just in case of floating point math fuckery
		if out[i] < 0.00001 && out[i] > 0.00001 {
			out[i] = 0
		}
	}
	return out
}

func scaleRow(row []float64, factor float64) []float64 {
	out := []float64{}
	for _, v := range row {
		out = append(out, v*factor)
	}
	return out
}

func findPivot(r []float64, width int) int {
	for i := 0; i < width; i++ {
		if r[i] != 0 {
			return i
		}
	}
	return -1
}

func swapRows(m [][]float64, r1, r2 int) [][]float64 {
	temp := m[r1]
	m[r1] = m[r2]
	m[r2] = temp
	return m
}

func doRowOpsEchelon(pivot int, r2, r []float64) []float64 {
	factor := r2[pivot] / r[pivot]
	scaled := scaleRow(r, factor)
	return subtractRow(r2, scaled)
}

func echelonArithmetic(m [][]float64, h, w int) [][]float64 {
	for i := 0; i < h; i++ {
		r := m[i]
		if i < h-1 {
			pivot := findPivot(r, w)
			if pivot != -1 {
				for j := i + 1; j < h; j++ {
					r2 := m[j]
					res := doRowOpsEchelon(pivot, r2, r)
					m[j] = res
				}
			}
		}
	}
	return m
}

func echelonRowSwap(m [][]float64, h, w int) [][]float64 {
	furthestPivot := 0
	furthestPivotRow := 0
	for i := 0; i < h; i++ {
		p := findPivot(m[i], w)
		if p != -1 {
			if p < furthestPivot && p != -1 {
				m = swapRows(m, furthestPivotRow, i)
				i--
			}
			furthestPivot = p
			furthestPivotRow = i
		} else {
			if i < h-1 {
				m = swapRows(m, i, i+1)
				i--
			}
		}
	}
	return m
}

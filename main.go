package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/linear-algebra/matrix"
	"github.com/linear-algebra/operations"
)

func main() {
	m := 300
	n := 300
	rows := make([][]float64, m)
	for i := 0; i < m; i++ {
		rows[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			rows[i][j] = float64(rand.Int63() % 300)
		}
	}
	matrix := matrix.New(m, n, rows)
	fmt.Println(matrix.String())
	echelon := operations.EchelonForm(matrix)
	fmt.Println(echelon.String())
	os.Exit(0)
}

package main

import (
	"fmt"
	"math"
)

const E = 0.001
const n = 4

func main() {
	A := [n][n]float64{
		{1.08, -0.04, 0.21, -18.00},
		{0.25, -1.23, 0.22, -0.09},
		{-0.21, 0.22, 0.8, -0.13},
		{0.15, -1.31, 0.06, -1.16},
	}
	var L, U [n][n]float64
	B := [n]float64{-1.24, -1.16, 2.56, 1.08}
	var sum float64

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i <= j {
				L[i][j] = 0
				if i == j {
					L[i][j] = 1
				}
				for k := 0; k < i; k++ {
					sum += L[i][k] * U[k][j]
				}
				U[i][j] = A[i][j] - sum
			} else if i > j {
				U[i][j] = 0
				for k := 0; k < j; k++ {
					sum += L[i][k] * U[k][j]
				}
				L[i][j] = (A[i][j] - sum) / U[j][j]
			}
			sum = 0
		}
	}

	fmt.Println("L:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%.3f  ", L[i][j])
		}
		fmt.Println()
	}

	fmt.Println("\nU:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%.3f  ", U[i][j])
		}
		fmt.Println()
	}

	var x, y [n]float64
	for k := 0; k < n; k++ {
		for j := 0; j < k; j++ {
			sum += L[k][j] * y[j]
		}
		y[k] = B[k] - sum
		sum = 0
	}
	for k := n - 1; k >= 0; k-- {
		for j := k + 1; j < n; j++ {
			sum += U[k][j] * x[j]
		}
		x[k] = (y[k] - sum) / U[k][k]
		sum = 0
	}

	fmt.Println()
	for k := 0; k < n; k++ {
		fmt.Printf("X%d : %.3f  Y%d : %.3f\n", k+1, x[k], k+1, y[k])
	}

	var xIter, xTime, f [n]float64
	var C [n][n]float64
	iter := 0

	for k := 0; k < n; k++ {
		xIter[k] = 0
	}

	fmt.Printf("%10s%20s%12s%12s%10s%25s\n", "N", "x1", "x2", "x3", "x4", "En")
	for {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					C[i][j] = 0
				} else {
					C[i][j] = -A[i][j] / A[i][i]
					sum += C[i][j]
				}
			}
			xTime[i] = xIter[i]
			f[i] = B[i] / A[i][i]
			xIter[i] = f[i] + sum*xIter[i]
			sum = 0
		}

		if math.Abs(xIter[0]-xTime[0]) < E {
			fmt.Println("Сходится")
			fmt.Println()
			for k := 0; k < n; k++ {
				fmt.Printf("X%d : %.3f\n", k+1, xIter[k])
			}
			break
		}
		if iter > 16 {
			fmt.Println("Расходится. Причины: Диагональное преобладание: Матрица коэффициентов системы линейных уравнений должна быть диагонально доминирующей. Это означает, что модуль элемента на главной диагонали матрицы должен быть больше суммы модулей всех остальных элементов в соответствующей строке (или столбце).")
			fmt.Println("Норма матрицы меньше 1: Норма матрицы, обратной к матрице коэффициентов системы, должна быть меньше 1. Это обеспечивает сходимость итерационного процесса.")
			fmt.Println()
			break
		}
		iter++
		fmt.Printf("%10d%25.3f%15.3f%10.3f%10.3f%25.3f\n", iter, xIter[0], xIter[1], xIter[2], xIter[3], xIter[0]-xTime[0])
	}
}



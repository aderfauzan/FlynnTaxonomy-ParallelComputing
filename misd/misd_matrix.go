package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Masukkan ukuran matrix: ")
	fmt.Scan(&n)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = i + j + 1
		}
	}

	fmt.Println("\n=== MISD MATRIX PIPELINE ===")
	start := time.Now()

	fmt.Println("Stage 1: Normalisasi")
	time.Sleep(500 * time.Millisecond)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] /= 2
		}
	}

	fmt.Println("Stage 2: Transformasi (kuadrat)")
	time.Sleep(500 * time.Millisecond)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] *= matrix[i][j]
		}
	}

	fmt.Println("Stage 3: Evaluasi (+5)")
	time.Sleep(500 * time.Millisecond)
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] += 5
		}
	}

	fmt.Println("Selesai.")
	fmt.Println("Waktu:", time.Since(start))
}
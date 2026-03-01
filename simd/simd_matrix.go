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
			matrix[i][j] = i + j
		}
	}

	fmt.Println("\n=== SIMD MATRIX ===")
	fmt.Println("Instruksi: semua elemen dikali 2")

	start := time.Now()

	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("Elemen (%d,%d) diproses\n", i, j)
			matrix[i][j] *= 2
			time.Sleep(150 * time.Millisecond)
		}
	}

	fmt.Println("Selesai.")
	fmt.Println("Waktu:", time.Since(start))
}
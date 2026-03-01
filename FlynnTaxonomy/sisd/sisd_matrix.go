package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Masukkan ukuran matrix: ")
	fmt.Scan(&n)

	A := make([][]int, n)
	B := make([][]int, n)
	C := make([][]int, n)

	for i := 0; i < n; i++ {
		A[i] = make([]int, n)
		B[i] = make([]int, n)
		C[i] = make([]int, n)
		for j := 0; j < n; j++ {
			A[i][j] = i + 1
			B[i][j] = j + 1
		}
	}

	fmt.Println("\n=== SISD MATRIX (Sequentia) ===")

	start := time.Now()

	for i := 0; i < n; i++ {
		fmt.Printf("Memproses baris %d...\n", i)
		for j := 0; j < n; j++ {
			fmt.Printf("  Menghitung elemen (%d,%d)\n", i, j)
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
			time.Sleep(200 * time.Millisecond)
		}
	}

	fmt.Println("Selesai.")
	fmt.Println("Waktu:", time.Since(start))
}
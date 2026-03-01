package main

import (
	"fmt"
	"sync"
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

	fmt.Println("\n=== MIMD MATRIX PARALLEL ===")

	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			fmt.Printf("[START] Baris %d\n", row)
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					C[row][j] += A[row][k] * B[k][j]
				}
				time.Sleep(200 * time.Millisecond)
			}
			fmt.Printf("[DONE] Baris %d\n", row)
		}(i)
	}

	wg.Wait()

	fmt.Println("Selesai.")
	fmt.Println("Waktu:", time.Since(start))
}
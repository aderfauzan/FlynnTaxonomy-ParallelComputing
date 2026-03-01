package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	var increment int

	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(&n)

	fmt.Print("Masukkan nilai yang akan ditambahkan ke semua data: ")
	fmt.Scan(&increment)

	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i + 1
	}

	fmt.Println("\n=== SIMD DEMONSTRATION ===")
	fmt.Println("Data sebelum diproses:", data)
	fmt.Println("Instruksi tunggal    : x +", increment)
	fmt.Println("\nProses berjalan:")

	start := time.Now()

	for i := range data {
		oldValue := data[i]
		newValue := oldValue + increment
		data[i] = newValue

		fmt.Printf("Data %d -> %d + %d = %d\n",
			i+1, oldValue, increment, newValue)

		time.Sleep(400 * time.Millisecond)
	}

	elapsed := time.Since(start)

	fmt.Println("\nData setelah diproses:", data)
	fmt.Println("Waktu eksekusi       :", elapsed)
}
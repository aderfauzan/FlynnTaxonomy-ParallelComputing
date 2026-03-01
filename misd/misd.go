package main

import (
	"fmt"
	"time"
)

func main() {
	var x int

	fmt.Print("Masukkan satu angka: ")
	fmt.Scan(&x)

	fmt.Println("\n=== MISD PIPELINE ===")
	fmt.Println("Data awal:", x)

	start := time.Now()

	// Stage 1 - Validasi
	fmt.Println("\n[Stage 1] Validasi data...")
	time.Sleep(500 * time.Millisecond)
	if x < 0 {
		fmt.Println("Data negatif, diubah ke positif")
		x = -x
	}
	fmt.Println("Output Stage 1:", x)

	// Stage 2 - Normalisasi
	fmt.Println("\n[Stage 2] Normalisasi data...")
	time.Sleep(500 * time.Millisecond)
	x = x / 2
	fmt.Println("Output Stage 2:", x)

	// Stage 3 - Transformasi
	fmt.Println("\n[Stage 3] Transformasi data (kuadrat)...")
	time.Sleep(500 * time.Millisecond)
	x = x * x
	fmt.Println("Output Stage 3:", x)

	// Stage 4 - Evaluasi akhir
	fmt.Println("\n[Stage 4] Evaluasi akhir...")
	time.Sleep(500 * time.Millisecond)
	final := x + 100
	fmt.Println("Final Result:", final)

	fmt.Println("\nPipeline selesai.")
	fmt.Println("Waktu eksekusi:", time.Since(start))
}
package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	var delayTime float64

	fmt.Print("Masukan jumlah angka: ")
	fmt.Scan(&n)

	fmt.Print("Masukan delay per step (detik): ")
	fmt.Scan(&delayTime)
	
	fmt.Println("\n=== SISD DEMONSTRATION ===")
	start := time.Now()
	sum := 0
	
	for i := 1; i <= n; i++ {
		fmt.Printf("Step %d: Add %d\n", i, i)
		sum += 1
		time.Sleep(time.Duration(delayTime * float64(time.Second)))
	}

	fmt.Println("Final Result:", sum)
	fmt.Println("Time:", time.Since(start))
}
package main

import (
	"fmt"
	"sync"
	"time"
)

func processTask(id int, value int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[START] Task %d | Initial Value: %d\n", id, value)

	for step := 1; step <= 2; step++ {
		fmt.Printf("Task %d - Processing Step %d...\n", id, step)
		time.Sleep(200 * time.Millisecond)
	}

	var result int
	
	if id%3 == 0 {
		result = value * value
		fmt.Printf("[DONE] Task %d | Operation: Square | Result: %d\n", id, result)
	} else if id%2 == 0 {
		result = value * 2
		fmt.Printf("[DONE] Task %d | Operation: Multiply by 2 | Result: %d\n", id, result)
	} else {
		result = value + 10
		fmt.Printf("[DONE] Task %d | Operation: Add 10 | Result: %d\n", id, result)
	}
}

func main() {
	var jumlah int

	fmt.Print("Masukkan jumlah task: ")
	fmt.Scan(&jumlah)

	fmt.Println("\n=== MIMD TASK PROCESSING ===")

	var wg sync.WaitGroup
	start := time.Now()

	for i := 1; i <= jumlah; i++ {
		initialValue := i * 5 // nilai awal tiap task
		wg.Add(1)
		go processTask(i, initialValue, &wg)
	}

	wg.Wait()

	fmt.Println("\nSemua task selesai.")
	fmt.Println("Total waktu eksekusi:", time.Since(start))
}
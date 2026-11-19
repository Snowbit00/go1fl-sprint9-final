package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	chunkMaxima := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			// Последняя часть включает остаток от деления.
			end = len(data)
		}

		go func(index int, chunk []int) {
			defer wg.Done()
			chunkMaxima[index] = maximum(chunk)
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение последовательно")
	started := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d горутинах\n", CHUNKS)
	started = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(started).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}

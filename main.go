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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rnd.Int()
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	m := data[0]
	for _, v := range data {
		if v > m {
			m = v
		}
	}
	return m
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	partSize := len(data) / CHUNKS
	remain := len(data) % CHUNKS
	parts := make([][]int, CHUNKS)
	start := 0
	for i := 0; i < CHUNKS; i++ {
		currSize := partSize
		if i < remain {
			currSize++
		}
		end := start + currSize
		parts[i] = data[start:end]
		start = end
	}
	maxInParts := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(part []int) {
			defer wg.Done()
			m := maximum(part)
			maxInParts[i] = m
		}(parts[i])
	}
	wg.Wait()
	return maximum(maxInParts)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

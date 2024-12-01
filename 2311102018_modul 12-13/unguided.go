//2311102018 - Modul 12 Aryo Tegar Sukarno
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numbers []int
	reader := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan bilangan (akhiri dengan bilangan negatif):")
	for {
		reader.Scan()
		input := reader.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid, masukkan angka!")
			continue
		}
		if num < 0 {
			break
		}
		numbers = append(numbers, num)
	}

	insertionSort(numbers)

	fmt.Println("Data setelah diurutkan:", numbers)

	if isConsistentSpacing(numbers) {
		fmt.Println("Data berjarak", numbers[1]-numbers[0])
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func isConsistentSpacing(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}

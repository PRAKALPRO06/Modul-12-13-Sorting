package main

import (
	"fmt"
)

func main() {
	var angka []int
	var input int

	fmt.Print("Masukkan bilangan (0 untuk menghitung median, -5313 untuk keluar):")

	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error membaca input:", err)
			break
		}

		if input == -5313 {
			break
		}

		if input == 0 {
			if len(angka) > 0 {
				median := hitungMedian(angka)
				fmt.Println(median)
			}
			continue
		}

		angka = append(angka, input)
	}
}

func hitungMedian(nums []int) int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	n := len(nums)
	if n%2 == 1 {
		return nums[n/2]
	} else {
		return (nums[n/2-1] + nums[n/2]) / 2
	}
}

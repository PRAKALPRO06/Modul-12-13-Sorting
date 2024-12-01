package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(arr_2311102253 []int) {
	n := len(arr_2311102253)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr_2311102253[j] < arr_2311102253[minIdx] {
				minIdx = j
			}
		}
		arr_2311102253[i], arr_2311102253[minIdx] = arr_2311102253[minIdx], arr_2311102253[i]
	}
}

func calculateMedian(arr_2311102253 []int) int {
	n := len(arr_2311102253)
	if n%2 == 1 {
		return arr_2311102253[n/2]
	}
	return (arr_2311102253[n/2-1] + arr_2311102253[n/2]) / 2
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan data (akhiri dengan -5313):")
	data := []int{}
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		values := strings.Fields(input)

		for _, value := range values {
			num, _ := strconv.Atoi(value)

			if num == -5313 {
				return
			}

			if num == 0 {
				selectionSort(data)
				median := calculateMedian(data)
				fmt.Println(median)
			} else {
				data = append(data, num)
			}
		}
	}
}

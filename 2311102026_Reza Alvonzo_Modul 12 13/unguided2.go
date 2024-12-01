package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func calculateMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
	return float64(arr[n/2])
}
//Reza Alvonzo 2311102026 IF 11 06
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan angka yang dipisahkan dengan spasi (akhiri dengan -5313):")
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, " ")
	var data []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Input tidak valid, harap masukkan angka bulat saja.")
			return
		}

		if num == -5313 {
			break
		} else if num == 0 {
			selectionSort(data)
			median := calculateMedian(data)
			fmt.Printf("%.0f\n", median)
		} else {
			data = append(data, num)
		}
	}
}
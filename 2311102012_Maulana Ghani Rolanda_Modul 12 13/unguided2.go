package main
import (
	"fmt"
)

func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}
func median(arr []int, n int) int {
	if (n % 2) == 0 {
		return int((arr[n/2 - 1] + arr[n/2])/2)
	} else {
		return arr[(n-1)/2]
	}
}

func main() {

	arr := make([]int, 1000000)
	arrMedian := make([]int, 1000000)
	n := 0
	nMed := 0
	datum := 0

	fmt.Println("Masukkan data: ")
	for datum != -5313 && n < 1000000 {
		fmt.Scan(&datum)
		if datum == 0 {
			selectionSort(arr, n)
			arrMedian[nMed] = median(arr, n)
			nMed++
		} else {
			arr[n] = datum
			n++
		}
	}

	fmt.Println("Output: ")
	for i:=0; i<nMed; i++ {
		fmt.Println(arrMedian[i], " ")
	}
}
//Maulana Ghani Rolanda 2311102012 IF-11-06//
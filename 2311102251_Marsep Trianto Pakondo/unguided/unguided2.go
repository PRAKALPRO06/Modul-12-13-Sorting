package main

import (
	"fmt"
)

func findMedian(arrBilBul []int) float64 {
	for i := 1; i < len(arrBilBul); i++ {
		key := arrBilBul[i]
		j := i - 1
		for j >= 0 && arrBilBul[j] > key {
			arrBilBul[j+1] = arrBilBul[j]
			j--
		}
		arrBilBul[j+1] = key
	}

	n := len(arrBilBul)
	if n%2 == 0 {
		return float64((arrBilBul[n/2-1] + arrBilBul[n/2])) / 2.0
	}
	return float64(arrBilBul[n/2])
}

func main() {
	var arrBilBul []int
    var median []float64
    jumlahM := 0
	for {
		var input int
		fmt.Scan(&input)

		if input < 0 {
			break
		}

		if input == 0 {
			if len(arrBilBul) == 0 {
				fmt.Println(0)
			} else {
				median = append(median, findMedian(arrBilBul))
                jumlahM++
			}
			continue
		}

		arrBilBul = append(arrBilBul, input)
	}

    fmt.Println()
    for i := 0; i < jumlahM; i++ {
        fmt.Printf("Median %v : %.2f\n", i+1, median[i])
    }
}

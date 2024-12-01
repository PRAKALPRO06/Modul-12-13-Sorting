package main

import "fmt"

const nMax = 1000000

func selectionSort(arr []int, n int) {
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

func hitungMedian(arr []int, n int) float64 {
    if n%2 == 0 {
        return float64(arr[n/2-1]+arr[n/2]) / 2.0
    }
    return float64(arr[n/2])
}

func main() {
    
    var numbers [nMax]int
    var temp [nMax]int
    var count int
    var num226 int
	fmt.Println("Masukkan rangkaian bilangan bulat (0 untuk median sementara, akhiri dengan -5313): ")
    for {
        fmt.Scan(&num226)
        if num226 == -5313 {
            break
        }

        if num226 != 0 {
            numbers[count] = num226
            count++
        } else {
            if count > 0 {
                for i := 0; i < count; i++ {
                    temp[i] = numbers[i]
                }
                selectionSort(temp[:], count)

                median := hitungMedian(temp[:], count)
                fmt.Printf("%.0f\n", median)
            }
        }
    }
}
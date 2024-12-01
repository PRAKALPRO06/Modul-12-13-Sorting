package main

import "fmt"

func selectionSortGanjil(arr []int) {
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

func selectionSortGenap(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] > arr[minIdx] {
                minIdx = j
            }
        }
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func main() {
    fmt.Print("Masukkan jumlah daerah: ")
    var n226 int
    fmt.Scan(&n226)

    resultSemuaDaerah := make([]string, n226)

    for i := 0; i < n226; i++ {
        fmt.Printf("\nUntuk daerah ke-%d:\n", i+1)
        fmt.Print("Masukkan jumlah rumah di daerah ini: ")

        var m226 int
        fmt.Scan(&m226)
        fmt.Printf("Masukkan %d nomor rumah (dipisahkan dengan spasi): ", m226)
        numbers := make([]int, m226)
        for j := 0; j < m226; j++ {
            fmt.Scan(&numbers[j])
        }

        var odd, even []int
        for _, num := range numbers {
            if num%2 == 1 {
                odd = append(odd, num)
            } else {
                even = append(even, num)
            }
        }

        selectionSortGanjil(odd)
        selectionSortGenap(even)

        var hasil string
        for j := 0; j < len(odd); j++ {
            if j > 0 {
                hasil += " "
            }
            hasil += fmt.Sprintf("%d", odd[j])
        }

        if len(even) > 0 {
            if len(odd) > 0 {
                hasil += " "
            }
            for j := 0; j < len(even); j++ {
                if j > 0 {
                    hasil += " "
                }
                hasil += fmt.Sprintf("%d", even[j])
            }
        }
        resultSemuaDaerah[i] = hasil
    }

    fmt.Println("\nHasil pengurutan untuk setiap daerah:")
    for i, hasil := range resultSemuaDaerah {
        fmt.Printf("Daerah %d: %s\n", i+1, hasil)
    }
}
package main

import "fmt"

func insertionSort(arr []int) {
        for i := 1; i < len(arr); i++ {
                key := arr[i]
                j := i - 1
                for j >= 0 && arr[j] > key {
                        arr[j+1] = arr[j]
                        j = j - 1
                }
                arr[j+1] = key
        }
}

func cekJarak(arr []int) string {
        jarak := arr[1] - arr[0]
        for i := 2; i < len(arr); i++ {
                if arr[i]-arr[i-1] != jarak {
                        return "Data berjarak tidak tetap"
                }
        }
        return fmt.Sprintf("Data berjarak %d", jarak)
}

func main() {
        var angka []int
        var input int

        fmt.Println("Masukkan angka (akhiri dengan angka negatif):")
        for {
                fmt.Scan(&input)
                if input < 0 {
                        break
                }
                angka = append(angka, input)
        }

        insertionSort(angka)
        fmt.Println("Array setelah diurutkan:", angka)
        fmt.Println(cekJarak(angka))
}
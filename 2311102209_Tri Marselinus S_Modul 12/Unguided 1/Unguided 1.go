package main

import "fmt"

func main() {
        var n int
        fmt.Scan(&n)

        for i := 0; i < n; i++ {
                var m int
                fmt.Scan(&m)

                rumah := make([]int, m)
                for j := 0; j < m; j++ {
                        fmt.Scan(&rumah[j])
                }

                ganjil := make([]int, 0)
                genap := make([]int, 0)
                for _, num := range rumah {
                        if num%2 == 1 {
                                ganjil = append(ganjil, num)
                        } else {
                                genap = append(genap, num)
                        }
                }

                for i := 1; i < len(ganjil); i++ {
                        key := ganjil[i]
                        j := i - 1
                        for j >= 0 && ganjil[j] > key {
                                ganjil[j+1] = ganjil[j]
                                j = j - 1
                        }
                        ganjil[j+1] = key
                }

                for i := 1; i < len(genap); i++ {
                        key := genap[i]
                        j := i - 1
                        for j >= 0 && genap[j] < key {
                                genap[j+1] = genap[j]
                                j = j - 1
                        }
                        genap[j+1] = key
                }

                for _, num := range ganjil {
                        fmt.Print(num, " ")
                }
                for i := len(genap) - 1; i >= 0; i-- {
                        fmt.Print(genap[i], " ")
                }
                fmt.Println()
        }
}
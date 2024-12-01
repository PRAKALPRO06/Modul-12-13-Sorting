package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah di daerah %d: ", i)
		fmt.Scanln(&m)

		rumah := make([]int, m)
		fmt.Printf("Masukkan nomor rumah di daerah %d (pisahkan dengan spasi): ", i)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		var ganjil, genap []int
		for _, r := range rumah {
			if r%2 != 0 {
				ganjil = append(ganjil, r)
			} else {
				genap = append(genap, r)
			}
		}

		for j := 0; j < len(ganjil)-1; j++ {
			minIndex := j
			for k := j + 1; k < len(ganjil); k++ {
				if ganjil[k] < ganjil[minIndex] {
					minIndex = k
				}
			}
			ganjil[j], ganjil[minIndex] = ganjil[minIndex], ganjil[j]
		}

		for j := 0; j < len(genap)-1; j++ {
			maxIndex := j
			for k := j + 1; k < len(genap); k++ {
				if genap[k] > genap[maxIndex] {
					maxIndex = k
				}
			}
			genap[j], genap[maxIndex] = genap[maxIndex], genap[j]
		}

		fmt.Printf("Nomor rumah di daerah %d (terurut): ", i)
		for _, r := range ganjil {
			fmt.Print(r, " ")
		}
		for _, r := range genap {
			fmt.Print(r, " ")
		}
		fmt.Println()
	}
}

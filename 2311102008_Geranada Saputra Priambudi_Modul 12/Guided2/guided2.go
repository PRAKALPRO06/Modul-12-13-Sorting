package main

import "fmt"

const (
	NMAX int = 1000
	MMAX int = 1000
)

type tabDaerah [NMAX]tabRumah
type arrRumah [MMAX]int
type tabRumah struct {
	rumah  arrRumah
	nRumah int
}

func inputRumah(daerah *tabDaerah, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&daerah[i].nRumah)
		if daerah[i].nRumah <= MMAX {
			for j := 0; j < daerah[i].nRumah; j++ {
				fmt.Scan(&daerah[i].rumah[j])
			}
		} else {
			fmt.Println("Jumlah Rumah tidak boleh melebihi ", MMAX)
		}
	}
}

func cetakArray(daerah tabDaerah, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < daerah[i].nRumah; j++ {
			if daerah[i].rumah[j]%2 != 0 {
				fmt.Print(daerah[i].rumah[j], " ")
			}
		}
		for j := 0; j < daerah[i].nRumah; j++ {
			if daerah[i].rumah[j]%2 == 0 {
				fmt.Print(daerah[i].rumah[j], " ")
			}
		}
		fmt.Printf("\n")
	}
}

func selectionSort(daerah *tabDaerah, n int) {
	var temp, idx_min int

	for k := 0; k < n; k++ {
		for i := 0; i < daerah[k].nRumah; i++ {
			idx_min = i
			for j := i + 1; j < daerah[k].nRumah; j++ {
				if daerah[k].rumah[idx_min] > daerah[k].rumah[j] {
					idx_min = j
				}
			}
			temp = daerah[k].rumah[idx_min]
			daerah[k].rumah[idx_min] = daerah[k].rumah[i]
			daerah[k].rumah[i] = temp
		}
	}
}

func main() {
	var (
		nDaerah int
		daerah  tabDaerah
	)

	fmt.Scanln(&nDaerah)
	if nDaerah <= NMAX {
		inputRumah(&daerah, nDaerah)
		selectionSort(&daerah, nDaerah)
		cetakArray(daerah, nDaerah)
	} else {
		fmt.Println("Jumlah daerah tidak dapat melebihi", NMAX)
	}

}

package main

import (
	"fmt"
)

const NMAX int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

func main() {
	var n int
	var pustaka [NMAX]Buku

	fmt.Print("MASUKAN JUMLAH BUKU YANG INGIN DI INPUT: ")
	fmt.Scanln(&n)

	if n > NMAX {
		fmt.Println("Jumlah buku tidak boleh lebih dari", NMAX)
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Data Buku ke-%d:\n", i+1)
		fmt.Print("ID Buku: ")
		fmt.Scanln(&pustaka[i].id)
		fmt.Print("Judul: ")
		fmt.Scanln(&pustaka[i].judul)
		fmt.Print("Penulis: ")
		fmt.Scanln(&pustaka[i].penulis)
		fmt.Print("Penerbit: ")
		fmt.Scanln(&pustaka[i].penerbit)
		fmt.Print("Eksemplar: ")
		fmt.Scanln(&pustaka[i].eksemplar)
		fmt.Print("Terbit: ")
		fmt.Scanln(&pustaka[i].tahun)
		fmt.Print("Rating: ")
		fmt.Scanln(&pustaka[i].rating)
	}

	var bukuFav Buku
	bukuFav = pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > bukuFav.rating {
			bukuFav = pustaka[i]
		}
	}

	fmt.Printf("\nBUKU TERFAVORIT BERDASARKAN BUKU : %s oleh %s, penerbit %s, tahun %d dengan rating %d\n",
		bukuFav.judul, bukuFav.penulis, bukuFav.penerbit, bukuFav.tahun, bukuFav.rating)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if pustaka[i].rating < pustaka[j].rating {
				pustaka[i], pustaka[j] = pustaka[j], pustaka[i]
			}
		}
	}
	fmt.Println("\nDATA BUKU BEDASARKAN RATING TERURUT:")
	limit := n
	if n > 5 {
		limit = 5
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
	var rating int
	fmt.Print("\nMASUKAN RATING UNTUK MENCARI BUKU: ")
	fmt.Scanln(&rating)

	found := false
	for i := 0; i < n; i++ {
		if pustaka[i].rating == rating {
			fmt.Printf("\nBuku ditemukan: %s oleh %s, penerbit %s, tahun %d, eksemplar %d, rating %d\n",
				pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].eksemplar, pustaka[i].rating)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

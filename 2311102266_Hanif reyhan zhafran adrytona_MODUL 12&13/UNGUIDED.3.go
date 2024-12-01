// 2311102266_Hanif Reyhan Zhafran Arytona

package main

import (
	"fmt"
	"sort"
)

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah buku harus lebih besar dari 0.")
		return
	}

	var pustaka []Buku
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		pustaka = append(pustaka, buku)
	}

	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}
	fmt.Println("\nBuku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		terfavorit.id, terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.eksemplar, terfavorit.tahun, terfavorit.rating)

	sort.SliceStable(pustaka, func(i, j int) bool {
		return pustaka[i].rating > pustaka[j].rating
	})

	fmt.Println("\nLima buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < len(pustaka); i++ {
		buku := pustaka[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}

	var rating int
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	ditemukan := false
	for _, buku := range pustaka {
		if buku.rating == rating {
			ditemukan = true
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

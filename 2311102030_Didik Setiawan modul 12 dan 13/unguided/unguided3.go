package main

import (
	"fmt"
	"sort"
)

type Buku struct {
	id        int
	eksemplar int
	tahun     int
	rating    int
	judul     string
	penerbit  string
	penulis   string
}

func DaftarkanBuku(pustaka *[]Buku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

func CetakFavorit(pustaka []Buku) {
	if len(pustaka) == 0 {
		fmt.Println("Pustaka kosong.")
		return
	}
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}
	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		terfavorit.id, terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.eksemplar, terfavorit.tahun, terfavorit.rating)
}

func UrutkanBuku(pustaka *[]Buku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].rating > (*pustaka)[j].rating
	})
}

func Cetak5Terbaik(pustaka []Buku) {
	fmt.Println("Lima buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < len(pustaka); i++ {
		buku := pustaka[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func CariBuku(pustaka []Buku, rating int) {
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

func main() {
	var n int
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&n)

	if n <= 0 || n > 7919 {
		fmt.Println("Jumlah buku harus antara 1 hingga 7919.")
		return
	}

	var pustaka []Buku

	DaftarkanBuku(&pustaka, n)
	CetakFavorit(pustaka)
	UrutkanBuku(&pustaka)
	Cetak5Terbaik(pustaka)

	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}

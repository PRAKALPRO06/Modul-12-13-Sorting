// 2311102037_BRIAN FARREL EVANDHIKA_IF 11_06
package main

import (
	"fmt"
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

func DaftarkanBuku(perpustakaan *[]Buku, jumlah int) {
	for i := 0; i < jumlah; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		*perpustakaan = append(*perpustakaan, buku)
	}
}

func CetakBukuRatingTertinggi(perpustakaan []Buku) {
	if len(perpustakaan) == 0 {
		fmt.Println("Perpustakaan kosong.")
		return
	}
	bukuFavorit := perpustakaan[0]
	for _, buku := range perpustakaan {
		if buku.rating > bukuFavorit.rating {
			bukuFavorit = buku
		}
	}
	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		bukuFavorit.id, bukuFavorit.judul, bukuFavorit.penulis, bukuFavorit.penerbit, bukuFavorit.eksemplar, bukuFavorit.tahun, bukuFavorit.rating)
}

func UrutkanBukuBerdasarkanRating(perpustakaan *[]Buku) {
	for i := 1; i < len(*perpustakaan); i++ {
		kunci := (*perpustakaan)[i]
		j := i - 1
		for j >= 0 && (*perpustakaan)[j].rating < kunci.rating {
			(*perpustakaan)[j+1] = (*perpustakaan)[j]
			j--
		}
		(*perpustakaan)[j+1] = kunci
	}
}

func Cetak5BukuTerbaik(perpustakaan []Buku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < len(perpustakaan); i++ {
		buku := perpustakaan[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func CariBukuBerdasarkanRating(perpustakaan []Buku, targetRating int) {
	ditemukan := false
	for _, buku := range perpustakaan {
		if buku.rating == targetRating {
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
	var jumlahBuku int
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&jumlahBuku)

	if jumlahBuku <= 0 || jumlahBuku > 7919 {
		fmt.Println("Jumlah buku harus antara 1 hingga 7919.")
		return
	}

	var perpustakaan []Buku

	DaftarkanBuku(&perpustakaan, jumlahBuku)
	CetakBukuRatingTertinggi(perpustakaan)
	UrutkanBukuBerdasarkanRating(&perpustakaan)
	Cetak5BukuTerbaik(perpustakaan)

	var ratingDicari int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&ratingDicari)
	CariBukuBerdasarkanRating(perpustakaan, ratingDicari)
}

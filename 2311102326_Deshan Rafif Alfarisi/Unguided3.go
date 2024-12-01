package main

import (
	"fmt"
)

// Definisi struct untuk Buku
type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

// Daftar Buku
type DaftarBuku []Buku

// Fungsi untuk mengurutkan buku berdasarkan rating secara descending menggunakan insertion sort
func UrutBuku(pustaka DaftarBuku) DaftarBuku {
	n := len(pustaka)
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		// Pindahkan elemen yang lebih kecil dari key ke posisi selanjutnya
		for j >= 0 && pustaka[j].Rating < key.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
	return pustaka
}

// Fungsi untuk mencetak buku terfavorit (rating tertinggi)
func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) > 0 {
		buku := pustaka[0]
		fmt.Printf("Buku Terfavorit:\nJudul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	} else {
		fmt.Println("Tidak ada buku dalam daftar.")
	}
}

// Fungsi untuk mencetak 5 buku terbaru dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < len(pustaka); i++ {
		buku := pustaka[i]
		fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating (binary search)
func CariBuku(pustaka DaftarBuku, rating int) {
	left, right := 0, len(pustaka)-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].Rating == rating {
			buku := pustaka[mid]
			fmt.Printf("Buku ditemukan:\nJudul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Eksemplar, buku.Rating)
			return
		} else if pustaka[mid].Rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	// Input data buku
	pustaka := DaftarBuku{
		{1, "Buku A", "Penulis A", "Penerbit A", 10, 2020, 5},
		{2, "Buku B", "Penulis B", "Penerbit B", 5, 2018, 4},
		{3, "Buku C", "Penulis C", "Penerbit C", 3, 2019, 3},
		{4, "Buku D", "Penulis D", "Penerbit D", 2, 2021, 5},
		{5, "Buku E", "Penulis E", "Penerbit E", 7, 2022, 2},
		{6, "Buku F", "Penulis F", "Penerbit F", 1, 2023, 4},
	}

	// Mengurutkan buku berdasarkan rating
	pustaka = UrutBuku(pustaka)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Cetak 5 buku terbaru
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	fmt.Println("Mencari buku dengan rating 4:")
	CariBuku(pustaka, 4)

	fmt.Println("Mencari buku dengan rating 6:")
	CariBuku(pustaka, 6)
}

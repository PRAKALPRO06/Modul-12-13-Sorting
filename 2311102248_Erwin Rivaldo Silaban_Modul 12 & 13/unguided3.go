package main

import (
	"fmt"
)

// Erwin Rivaldo Silaban/2311102248
const nMax = 7919

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku struct {
	Pustaka  [nMax]Buku
	nPustaka int
}

// Fungsi untuk menambahkan buku ke daftar
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Masukkan data buku (id, judul, penulis, penerbit, eksemplar, tahun, rating):")
		fmt.Scan(&pustaka.Pustaka[i].id, &pustaka.Pustaka[i].judul, &pustaka.Pustaka[i].penulis, &pustaka.Pustaka[i].penerbit, &pustaka.Pustaka[i].eksemplar, &pustaka.Pustaka[i].tahun, &pustaka.Pustaka[i].rating)
	}
	pustaka.nPustaka = n
}

// Fungsi untuk mencetak buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}

	terfavorit := pustaka.Pustaka[0]
	for i := 1; i < pustaka.nPustaka; i++ {
		if pustaka.Pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka.Pustaka[i]
		}
	}

	fmt.Println("Buku terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating (Insertion Sort)
func UrutBuku(pustaka *DaftarBuku) {
	for i := 1; i < pustaka.nPustaka; i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].rating < key.rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

// Fungsi untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating (Binary Search)
func CariBuku(pustaka DaftarBuku, rating int) {
	low, high := 0, pustaka.nPustaka-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka[mid].rating == rating {
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Buku ditemukan: Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
			return
		} else if pustaka.Pustaka[mid].rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka)
	UrutBuku(&pustaka)
	Cetak5Terbaru(pustaka)
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}

package main

import (
	"fmt"
	"sort"
	// NAMA : Egi Umar Ferdhika
	// NIM : 2311102277
)

type DaftarBuku struct {
	nBuku int
	buku  []DataBuku
}

type DataBuku struct {
	id        int
	judul     string
	pengarang string
	tahun     int
	rating    int
}

// DaftarkanBuku initializes a new DaftarBuku with specified size
// I.S: sejumlah n data buku telah siap para piranti masukan
// F.S: n berisi sebuah nilai, dan pustaka berisi sejumlah n data buku
func DaftarkanBuku(n int) *DaftarBuku {
	return &DaftarBuku{
		nBuku: 0,
		buku:  make([]DataBuku, n),
	}
}

// CetakTerfavorit prints book data sorted by rating
// I.S: array pustaka berisi n buah data buku dan belum terurut
// F.S: Tampilan data buku (judul, penulis, penerbit, tahun) terfavorit
func CetakTerfavorit(pustaka *DaftarBuku, n int) {
	if n == 0 || pustaka.nBuku == 0 {
		return
	}

	// Create temporary slice for sorting
	temp := make([]DataBuku, pustaka.nBuku)
	copy(temp, pustaka.buku[:pustaka.nBuku])

	// Sort by rating (highest first)
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].rating > temp[j].rating
	})

	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul: %s, Pengarang: %s, Tahun: %d, Rating: %d\n",
		temp[0].judul, temp[0].pengarang, temp[0].tahun, temp[0].rating)
}

// UrutBuku sorts books by rating using insertion sort
// I.S: Array pustaka berisi n data buku
// F.S: Array pustaka terurut menurun/mengecil terhadap rating
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < pustaka.nBuku; i++ {
		key := pustaka.buku[i]
		j := i - 1

		for j >= 0 && pustaka.buku[j].rating < key.rating {
			pustaka.buku[j+1] = pustaka.buku[j]
			j--
		}
		pustaka.buku[j+1] = key
	}
}

// Cetak5Terbaru prints the 5 books with highest rating
// I.S: pustaka berisi n data buku yang sudah terurut menurut rating
// F.S: Laporan 5 judul buku dengan rating tertinggi
func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
	fmt.Println("5 Buku Rating Tertinggi:")
	limit := 5
	if pustaka.nBuku < 5 {
		limit = pustaka.nBuku
	}

	for i := 0; i < limit; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %d\n",
			i+1, pustaka.buku[i].judul, pustaka.buku[i].rating)
	}
}

// CariBuku searches for books with specific rating using binary search
// I.S: pustaka berisi n data buku yang sudah terurut menurut rating
// F.S: Laporan salah satu buku dengan rating yang diberikan
func CariBuku(pustaka *DaftarBuku, n int, r int) {
	left := 0
	right := pustaka.nBuku - 1
	found := false

	for left <= right {
		mid := (left + right) / 2
		if pustaka.buku[mid].rating == r {
			fmt.Printf("Buku ditemukan:\nJudul: %s\nPengarang: %s\nTahun: %d\nRating: %d\n",
				pustaka.buku[mid].judul, pustaka.buku[mid].pengarang,
				pustaka.buku[mid].tahun, pustaka.buku[mid].rating)
			found = true
			break
		}
		if pustaka.buku[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if !found {
		fmt.Printf("Tidak ada buku dengan rating %d\n", r)
	}
}

func main() {
	perpustakaan := DaftarkanBuku(10)

	perpustakaan.buku[0] = DataBuku{1, "Go Programming", "John Doe", 2020, 5}
	perpustakaan.buku[1] = DataBuku{2, "Database Design", "Jane Smith", 2019, 4}
	perpustakaan.buku[2] = DataBuku{3, "Web Development", "John Doe", 2021, 3}
	perpustakaan.buku[3] = DataBuku{4, "AI Basics", "Alice Johnson", 2022, 5}
	perpustakaan.nBuku = 4

	fmt.Println("Original books:")
	for i := 0; i < perpustakaan.nBuku; i++ {
		fmt.Printf("%+v\n", perpustakaan.buku[i])
	}

	fmt.Println("\nTesting CetakTerfavorit:")
	CetakTerfavorit(perpustakaan, perpustakaan.nBuku)

	fmt.Println("\nTesting UrutBuku:")
	UrutBuku(perpustakaan, perpustakaan.nBuku)

	fmt.Println("\nTesting Cetak5Terbaru:")
	Cetak5Terbaru(perpustakaan, perpustakaan.nBuku)

	fmt.Println("\nTesting CariBuku with rating 4:")
	CariBuku(perpustakaan, perpustakaan.nBuku, 4)
}

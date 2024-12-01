package main

import (
	"fmt"
	"sort"
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

func DaftarkanBuku(n int) *DaftarBuku {
	return &DaftarBuku{
		nBuku: 0,
		buku:  make([]DataBuku, n),
	}
}

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

	perpustakaan.buku[0] = DataBuku{1, "Frontend Developer", "Umar", 2005, 5}
	perpustakaan.buku[1] = DataBuku{2, "Backend Developer", "Slamet", 2006, 2}
	perpustakaan.buku[2] = DataBuku{3, "Fullstack Developer", "Miskan", 2007, 4}
	perpustakaan.buku[3] = DataBuku{4, "Mobile Developer", "Yanto", 2008, 3}
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
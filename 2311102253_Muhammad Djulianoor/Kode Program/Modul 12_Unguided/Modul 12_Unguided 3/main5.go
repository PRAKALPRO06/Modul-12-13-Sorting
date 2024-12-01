package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id_2311102253, judul, penulis, penerbit string
	eksemplar, tahun, rating                int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	nPustaka = n
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&pustaka[i].id_2311102253, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := -1
	var favorit Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			favorit = pustaka[i]
		}
	}
	fmt.Printf("Buku Terfavorit: %s oleh %s, diterbitkan oleh %s pada tahun %d\n", favorit.judul, favorit.penulis, favorit.penerbit, favorit.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j = j - 1
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. %s (%d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n, rating int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == rating {
			fmt.Printf("Ditemukan buku dengan rating %d: %s oleh %s, diterbitkan oleh %s pada tahun %d, Eksemplar: %d\n",
				rating, pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar)
			return
		} else if pustaka[mid].rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, n)
	UrutBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)
	Cetak5Terbaru(pustaka, n)

	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}

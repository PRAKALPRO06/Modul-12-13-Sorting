package main

import "fmt"

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

const nMax = 7919

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n) // Membaca jumlah buku
	for i := 0; i < *n; i++ {
		var buku Buku
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		pustaka[i] = buku
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := -1
	var bukuFavorit Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			bukuFavorit = pustaka[i]
		}
	}
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
		bukuFavorit.judul, bukuFavorit.penulis, bukuFavorit.penerbit, bukuFavorit.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	m := 5
	if n < 5 {
		m = n
	}
	for i := 0; i < m; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			return
		} else if pustaka[mid].rating < r {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var pustaka DaftarBuku
	var n int
	DaftarkanBuku(&pustaka, &n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	var rating int
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}

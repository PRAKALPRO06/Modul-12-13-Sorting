package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&n)

	// Validasi jumlah buku
	if n <= 0 || n > nMax {
		fmt.Println("Jumlah buku tidak valid!")
		return
	}

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	// Input rating untuk pencarian
	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scanln(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}

// Procedure DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scanln(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Procedure CetakTerfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := -1
	var bukuFavorit Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			bukuFavorit = pustaka[i]
		}
	}

	fmt.Println("Buku terfavorit:")
	fmt.Println(bukuFavorit.judul, bukuFavorit.penulis, bukuFavorit.penerbit, bukuFavorit.tahun)
}

// Procedure UrutBuku
func UrutBuku(pustaka *DaftarBuku, n int) {
	// Insertion Sort untuk mengurutkan berdasarkan tahun (descending)
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].tahun < key.tahun {
			pustaka[j+1] = pustaka[j]
			j = j - 1
		}
		pustaka[j+1] = key
	}
}

// Procedure Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku terbaru:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

// Procedure CariBuku
func CariBuku(pustaka DaftarBuku, n int, r int) {
	// Urutkan berdasarkan rating (ascending) sebelum melakukan binary search
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating > key.rating {
			pustaka[j+1] = pustaka[j]
			j = j - 1
		}
		pustaka[j+1] = key
	}

	// Binary Search
	left := 0
	right := n - 1
	found := false

	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			fmt.Println("Buku ditemukan:")
			fmt.Println(pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			found = true
			break
		} else if pustaka[mid].rating < r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

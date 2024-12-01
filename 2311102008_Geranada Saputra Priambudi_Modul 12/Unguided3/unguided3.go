package main

import "fmt"

const NMAX int = 7919

type buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]buku

func DaftarkanBuku(pustaka *DaftarBuku, N *int) {
	fmt.Print("Banyak Buku: ")
	fmt.Scanln(N)

	for i := 0; i < *N; i++ {
		fmt.Println("Data buku ke-", i+1)
		fmt.Scanln(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var buku_fav buku

	buku_fav = pustaka[0]
	for i := 1; i < n; i++ {
		if buku_fav.rating < pustaka[i].rating {
			buku_fav = pustaka[i]
		}
	}

	fmt.Printf("Buku terfavorit adalah %s oleh %s, penerbit %s, tahun %d dengan rating %d\n", buku_fav.judul, buku_fav.penulis, buku_fav.penerbit, buku_fav.tahun, buku_fav.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp buku
	var j int

	for i := 1; i < n; i++ {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("Top Buku Terbaru:")
	if n < 5 {
		for i := 0; i < n; i++ {
			fmt.Println(i+1, ". ", pustaka[i].judul, "Rating:", pustaka[i].rating)
		}
	} else {
		for i := 0; i < 5; i++ {
			fmt.Println(i+1, ". ", pustaka[i].judul, "Rating:", pustaka[i].rating)
		}
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	found := false
	var idx int
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			found = true
			idx = mid
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if found {
		fmt.Println("Buku ditemukan: ")
		fmt.Println(pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun, pustaka[idx].eksemplar, pustaka[idx].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var Pustaka DaftarBuku
	var nPustaka int
	var rating int

	DaftarkanBuku(&Pustaka, &nPustaka)
	CetakTerfavorit(Pustaka, nPustaka)
	UrutBuku(&Pustaka, nPustaka)
	Cetak5Terbaru(Pustaka, nPustaka)
	fmt.Print("Buku dengan rating yang dicari: ")
	fmt.Scan(&rating)
	CariBuku(Pustaka, nPustaka, rating)
}

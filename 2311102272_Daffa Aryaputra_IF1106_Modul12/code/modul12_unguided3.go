package main

import "fmt"

const NMAX int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, N *int) {
	fmt.Print("Berapa banyaknya Buku: ")
	fmt.Scanln(N)

	for i := 0; i < *N; i++ {
		fmt.Printf("Data buku ke-%d:\n", i+1)
		fmt.Print("ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating (pisahkan dengan spasi): ")
		fmt.Scanln(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	bukuFav := pustaka[0]
	for i := 1; i < n; i++ {
		if bukuFav.rating < pustaka[i].rating {
			bukuFav = pustaka[i]
		}
	}
	fmt.Printf("Buku terfavorit adalah %s oleh %s, penerbit %s, tahun %d dengan rating %d\n",
		bukuFav.judul, bukuFav.penulis, bukuFav.penerbit, bukuFav.tahun, bukuFav.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i
		for j > 0 && pustaka[j-1].rating < temp.rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("Top 5 Buku Terbaru:")
	limit := n
	if n > 5 {
		limit = 5
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
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
		} else if pustaka[mid].rating > r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if found {
		fmt.Println("Buku ditemukan:")
		fmt.Printf("%s oleh %s, penerbit %s, tahun %d, eksemplar %d, rating %d\n",
			pustaka[idx].judul, pustaka[idx].penulis, pustaka[idx].penerbit, pustaka[idx].tahun, pustaka[idx].eksemplar, pustaka[idx].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var rating int

	DaftarkanBuku(&pustaka, &nPustaka)
	CetakTerfavorit(pustaka, nPustaka)
	UrutBuku(&pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)
	fmt.Print("Masukkan rating buku yang dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, nPustaka, rating)
}

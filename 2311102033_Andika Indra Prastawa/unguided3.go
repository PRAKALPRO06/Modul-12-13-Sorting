package main

import (
	"fmt"
)

const NMAX int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]Buku

func DaftarBuku_2311102033(pustaka *DaftarBuku, N *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(N)

	if *N > NMAX {
		fmt.Printf("Jumlah buku tidak boleh lebih dari %d.\n", NMAX)
		*N = 0
		return
	}

	for i := 0; i < *N; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID Buku: ")
		fmt.Scanln(&pustaka[i].id)
		fmt.Print("Judul Buku: ")
		fmt.Scanln(&pustaka[i].judul)
		fmt.Print("Penulis: ")
		fmt.Scanln(&pustaka[i].penulis)
		fmt.Print("Penerbit: ")
		fmt.Scanln(&pustaka[i].penerbit)
		fmt.Print("Jumlah Eksemplar: ")
		fmt.Scanln(&pustaka[i].eksemplar)
		fmt.Print("Tahun Terbit: ")
		fmt.Scanln(&pustaka[i].tahun)
		fmt.Print("Rating Buku: ")
		fmt.Scanln(&pustaka[i].rating)
	}
}

func CetakTerfavorit_2311102033(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku yang terdaftar.")
		return
	}

	bukuFav := pustaka[0]
	for i := 1; i < n; i++ {
		if bukuFav.rating < pustaka[i].rating {
			bukuFav = pustaka[i]
		}
	}
	fmt.Printf("\nBuku terfavorit: %s oleh %s, penerbit %s, tahun %d dengan rating %d\n",
		bukuFav.judul, bukuFav.penulis, bukuFav.penerbit, bukuFav.tahun, bukuFav.rating)
}

func UrutanBuku_2311102033(pustaka *DaftarBuku, n int) {
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

func CetakTerbaru_2311102033(pustaka DaftarBuku, n int) {
	fmt.Println("\nTop 5 Buku dengan Rating Tertinggi:")
	limit := n
	if n > 5 {
		limit = 5
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku_2311102033(pustaka DaftarBuku, n int, r int) {
	if n == 0 {
		fmt.Println("Tidak ada buku yang terdaftar.")
		return
	}

	found := false
	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			fmt.Printf("\nBuku ditemukan: %s oleh %s, penerbit %s, tahun %d, eksemplar %d, rating %d\n",
				pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].eksemplar, pustaka[i].rating)
			break
		}
	}
	if !found {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int
	var rating int

	DaftarBuku_2311102033(&pustaka, &nPustaka)

	CetakTerfavorit_2311102033(pustaka, nPustaka)

	UrutanBuku_2311102033(&pustaka, nPustaka)
	CetakTerbaru_2311102033(pustaka, nPustaka)

	fmt.Print("\nMasukkan rating untuk mencari buku: ")
	fmt.Scanln(&rating)
	CariBuku_2311102033(pustaka, nPustaka, rating)
}

package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating      int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(daftar_buku *DaftarBuku, n int) {
	fmt.Println("Masukkan data buku dengan format (id judul penulis penerbit eksemplar tahun rating)")
	for i := 0; i < n; i++ {
		fmt.Scan(&daftar_buku[i].id, &daftar_buku[i].judul, &daftar_buku[i].penulis, &daftar_buku[i].penerbit, &daftar_buku[i].eksemplar, &daftar_buku[i].tahun, &daftar_buku[i].rating)
	}
}

func CetakTerfavorit(daftar_buku DaftarBuku, n int) {
	var buku_terfavorit Buku = daftar_buku[0]
	for i := 1; i < n; i++ {
		if buku_terfavorit.rating < daftar_buku[i].rating {
			buku_terfavorit = daftar_buku[i]
		}
	}
	fmt.Println("Buku terfavorit adalah:\"",buku_terfavorit.judul, "\"")
}

func UrutBukuBerdasarkanRating(daftar_buku *DaftarBuku, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if daftar_buku[j].rating > daftar_buku[idxMax].rating {
				idxMax = j
			}
		}
		daftar_buku[i], daftar_buku[idxMax] = daftar_buku[idxMax], daftar_buku[i]
	}
}

func Cetak5Terbaru(daftar_buku DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi: ")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(i+1, ".",daftar_buku[i].judul, ", rating: ",daftar_buku[i].rating)
	}
	fmt.Print("\n")
}

func CariBuku(daftar_buku DaftarBuku, n, rating int) {
	ditemukan := false
	for i := 0; i < n; i++ {
		if daftar_buku[i].rating == rating {
			ditemukan = true
			fmt.Print("Buku dengan rating ", rating, " ditemukan: ")
			fmt.Println(daftar_buku[i].judul)
		}
	}

	if !ditemukan {
		fmt.Println("Buku dengan rating", rating, "tidak ditemukan")
	}
}

func main() {
	var daftar_buku DaftarBuku
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	DaftarkanBuku(&daftar_buku, n)
	UrutBukuBerdasarkanRating(&daftar_buku, n)

	fmt.Print("\n")
	CetakTerfavorit(daftar_buku, n)
	fmt.Print("\n")
	Cetak5Terbaru(daftar_buku, n)

	var rating int
	fmt.Print("Masukkan rating buku yang dicari: ")
	fmt.Scan(&rating)
	CariBuku(daftar_buku, n, rating)
}
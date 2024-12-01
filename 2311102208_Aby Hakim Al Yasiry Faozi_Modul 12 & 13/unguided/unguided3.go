package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	Judul, Penulis, Penerbit     string
	Eksemplar, Tahun, Rating, ID int
}

type DaftarBuku struct {
	Data     [nMax]Buku
	nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	pustaka.nPustaka = n
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)
		fmt.Scan(&pustaka.Data[i].ID, &pustaka.Data[i].Judul, &pustaka.Data[i].Penulis, &pustaka.Data[i].Penerbit,
			&pustaka.Data[i].Eksemplar, &pustaka.Data[i].Tahun, &pustaka.Data[i].Rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku dalam daftar.")
		return
	}

	favorit := pustaka.Data[0]
	for i := 1; i < pustaka.nPustaka; i++ {
		if pustaka.Data[i].Rating > favorit.Rating {
			favorit = pustaka.Data[i]
		}
	}

	fmt.Println("Buku Terfavorit:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
		favorit.ID, favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun)
}

func UrutBuku(pustaka *DaftarBuku) {
	for i := 0; i < pustaka.nPustaka-1; i++ {
		for j := 0; j < pustaka.nPustaka-i-1; j++ {
			if pustaka.Data[j].Rating < pustaka.Data[j+1].Rating {
				temp := pustaka.Data[j]
				pustaka.Data[j] = pustaka.Data[j+1]
				pustaka.Data[j+1] = temp
			}
		}
	}
}

func Cetak5Buku(pustaka DaftarBuku) {
	fmt.Println("5 Buku Dengan Rating Tertinggi:")
	for i := 0; i < pustaka.nPustaka && i < 5; i++ {
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Rating: %d\n",
			pustaka.Data[i].ID, pustaka.Data[i].Judul, pustaka.Data[i].Penulis,
			pustaka.Data[i].Penerbit, pustaka.Data[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku, target int) {
	found := false
	for i := 0; i < pustaka.nPustaka; i++ {
		if pustaka.Data[i].Rating == target {
			fmt.Println("Buku ditemukan:")
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
				pustaka.Data[i].ID, pustaka.Data[i].Judul, pustaka.Data[i].Penulis,
				pustaka.Data[i].Penerbit, pustaka.Data[i].Tahun, pustaka.Data[i].Rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var jumlahBuku, cariRating int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&jumlahBuku)

	DaftarkanBuku(&pustaka, jumlahBuku)

	fmt.Println("\nBuku Terfavorit")
	CetakTerfavorit(pustaka)

	UrutBuku(&pustaka)

	fmt.Println("\n5 Buku Dengan Rating Tertinggi")
	Cetak5Buku(pustaka)

	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scan(&cariRating)

	CariBuku(pustaka, cariRating)
}

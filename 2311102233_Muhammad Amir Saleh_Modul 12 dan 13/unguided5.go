package main

import (
	"fmt"
)

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&pustaka[i].id)
		fmt.Print("Judul: ")
		fmt.Scan(&pustaka[i].judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&pustaka[i].penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Print("Eksemplar: ")
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Print("Tahun: ")
		fmt.Scan(&pustaka[i].tahun)
		fmt.Print("Rating: ")
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}
	maxIndex := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[maxIndex].rating {
			maxIndex = i
		}
	}
	fmt.Println("Buku terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
		pustaka[maxIndex].judul, pustaka[maxIndex].penulis,
		pustaka[maxIndex].penerbit, pustaka[maxIndex].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	count := 5
	if n < 5 {
		count = n
	}
	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == r {
			fmt.Println("Buku ditemukan:")
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit,
				pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			return
		} else if pustaka[mid].rating > r {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}

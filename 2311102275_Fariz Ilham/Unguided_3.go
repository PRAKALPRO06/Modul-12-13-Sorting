package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		var buku Buku
		fmt.Printf("\nData buku ke-%d\n", i+1)

		fmt.Print("ID: ")
		fmt.Scan(&buku.id)

		fmt.Print("Judul: ")
		fmt.Scan(&buku.judul)

		fmt.Print("Penulis: ")
		fmt.Scan(&buku.penulis)

		fmt.Print("Penerbit: ")
		fmt.Scan(&buku.penerbit)

		fmt.Print("Eksemplar: ")
		fmt.Scan(&buku.eksemplar)

		fmt.Print("Tahun: ")
		fmt.Scan(&buku.tahun)

		fmt.Print("Rating: ")
		fmt.Scan(&buku.rating)

		(*pustaka)[i] = buku
	}
}

func CetakTertaforit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}
	terbaik := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terbaik.rating {
			terbaik = pustaka[i]
		}
	}
	fmt.Println("\nBuku favorit dengan rating tertinggi:")
	fmt.Printf("%s | %s | %s | %s | %d | %d | %d\n", terbaik.id, terbaik.judul, terbaik.penulis, terbaik.penerbit, terbaik.eksemplar, terbaik.tahun, terbaik.rating)
}

func UrutkanBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n5 Buku dengan rating tertinggi:")
	count := 0
	for i := 0; i < n && count < 5; i++ {
		buku := pustaka[i]
		fmt.Printf("%s | %s | %s | %s | %d | %d | %d\n", buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
		count++
	}
	if count < 5 {
		fmt.Println("Jumlah buku kurang dari 5.")
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Println("\nBuku ditemukan:")
			fmt.Printf("%s | %s | %s | %s | %d | %d | %d\n", buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("\nTidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	DaftarkanBuku(&pustaka, &n)

	CetakTertaforit(pustaka, n)

	UrutkanBuku(&pustaka, n)

	Cetak5Terbaru(pustaka, n)

	var rating int
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}

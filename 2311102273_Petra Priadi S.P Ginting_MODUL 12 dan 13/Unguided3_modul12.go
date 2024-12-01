package main

import "fmt"

const nMax = 7919

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID,Judul, Penulis, Penerbit,Eksemplar, Tahun, Rating):\n", i+1)
		fmt.Scan(&pustaka[i].ID, &pustaka[i].Judul, &pustaka[i].Penulis, &pustaka[i].Penerbit, &pustaka[i].Eksemplar, &pustaka[i].Tahun, &pustaka[i].Rating)
	}
}

func CetakTerfavorit(pustaka *DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}
	terbaik := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].Rating > terbaik.Rating {
			terbaik = pustaka[i]
		}
	}
	fmt.Printf("Buku Terfavorit: %s oleh %s (%s, %d)\n", terbaik.Judul, terbaik.Penulis, terbaik.Penerbit, terbaik.Tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].Rating < key.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
	fmt.Println("5 Buku dengan Rating Tertinggi: ")
	for i := 0; i < n && i < 5; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].Judul, pustaka[i].Rating)
	}
}

func CariBuku(pustaka *DaftarBuku, n, rating int) {
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].Rating == rating {
			b := pustaka[mid]
			fmt.Printf("Buku ditemukan: %s oleh %s (%s, %d) [%d eksemplar, Rating %d]\n",
				b.Judul, b.Penulis, b.Penerbit, b.Tahun, b.Eksemplar, b.Rating)
			return
		} else if pustaka[mid].Rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)

	CetakTerfavorit(&pustaka, n)

	UrutBuku(&pustaka, n)
	Cetak5Terbaru(&pustaka, n)

	fmt.Print("Masukkan rating yang dicari: ")
	fmt.Scan(&rating)
	CariBuku(&pustaka, n, rating)
}

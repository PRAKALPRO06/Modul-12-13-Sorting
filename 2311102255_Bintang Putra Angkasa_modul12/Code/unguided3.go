package main

import "fmt"

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

const nMax = 7919
type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan jumlah buku:")
	fmt.Scanln(n)

	for i := 0; i < *n; i++ {
		fmt.Println("Masukkan data buku ke-", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&pustaka[i].id)
		fmt.Print("Judul: ")
		fmt.Scan(&pustaka[i].judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&pustaka[i].penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Print("Tahun: ")
		fmt.Scan(&pustaka[i].tahun)
		fmt.Print("Eksemplar: ")
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Print("Rating: ")
		fmt.Scan(&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", 
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		// Mengurutkan berdasarkan rating secara menurun
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("Lima Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %d\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low := 0
	high := n - 1
	found := false

	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			fmt.Printf("Buku dengan rating %d ditemukan:\n", r)
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n", 
				pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			found = true
			break
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka int

	DaftarkanBuku(&pustaka, &nPustaka)
	UrutBuku(& pustaka, nPustaka)
	CetakTerfavorit(pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)

	var ratingCari int
	fmt.Print("Masukkan rating yang akan dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, nPustaka, ratingCari)
}
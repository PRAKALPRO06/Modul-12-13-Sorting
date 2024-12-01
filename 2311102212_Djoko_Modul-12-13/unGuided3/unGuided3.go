package main

import "fmt"

type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Rating int
}

type Perpustakaan struct {
	Pustaka  []Buku
	NPustaka int
}

// Fungsi untuk membaca data buku
func DaftarkanBuku(p *Perpustakaan, n int) {
	p.NPustaka = n
	p.Pustaka = make([]Buku, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)
		fmt.Scan(&p.Pustaka[i].ID, &p.Pustaka[i].Judul, &p.Pustaka[i].Penulis, &p.Pustaka[i].Penerbit, &p.Pustaka[i].Eksemplar, &p.Pustaka[i].Tahun, &p.Pustaka[i].Rating)
	}
}

// Fungsi untuk mencetak buku favorit
func CetakTerfavorit(p Perpustakaan) {
	maxRating := -1
	var favorit Buku
	for _, buku := range p.Pustaka {
		if buku.Rating > maxRating {
			maxRating = buku.Rating
			favorit = buku
		}
	}
	fmt.Printf("Buku Terfavorit: %s - %s (%d), Rating: %d\n", favorit.Judul, favorit.Penerbit, favorit.Tahun, favorit.Rating)
}

// Fungsi pengurutan menggunakan metode insertion sort
func insertionSort(p *Perpustakaan) {
	for i := 1; i < p.NPustaka; i++ {
		temp := p.Pustaka[i]
		j := i
		// Geser elemen-elemen yang lebih kecil dari temp ke kanan
		for j > 0 && temp.Rating > p.Pustaka[j-1].Rating {
			p.Pustaka[j] = p.Pustaka[j-1]
			j = j - 1
		}
		p.Pustaka[j] = temp
	}
}

// Fungsi untuk mencetak 5 buku dengan rating tertinggi
func Cetak5Terbaru(p Perpustakaan) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < p.NPustaka && i < 5; i++ {
		fmt.Printf("%s - %s (%d), Rating: %d\n", p.Pustaka[i].Judul, p.Pustaka[i].Penerbit, p.Pustaka[i].Tahun, p.Pustaka[i].Rating)
	}
}

// Fungsi untuk mencari buku berdasarkan rating menggunakan pencarian biner
func CariBuku(p Perpustakaan, rating int) {
	left, right := 0, p.NPustaka-1
	for left <= right {
		mid := (left + right) / 2
		if p.Pustaka[mid].Rating == rating {
			fmt.Printf("Ditemukan buku: %s - %s (%d), Rating: %d\n", p.Pustaka[mid].Judul, p.Pustaka[mid].Penerbit, p.Pustaka[mid].Tahun, p.Pustaka[mid].Rating)
			return
		} else if p.Pustaka[mid].Rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var perpustakaan2311102212 Perpustakaan
	var n, rating int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Input data buku
	DaftarkanBuku(&perpustakaan2311102212, n)

	// Urutkan buku berdasarkan rating
	insertionSort(&perpustakaan2311102212)

	// Cetak buku favorit
	CetakTerfavorit(perpustakaan2311102212)

	// Cetak 5 buku terbaru
	Cetak5Terbaru(perpustakaan2311102212)

	// Cari buku berdasarkan rating
	fmt.Print("Masukkan rating yang dicari: ")
	fmt.Scan(&rating)
	CariBuku(perpustakaan2311102212, rating)
}

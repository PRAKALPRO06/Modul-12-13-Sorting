package main

import (
	"fmt"
)

// Struktur data Buku
type Buku struct {
	id      int
	judul   string
	penulis string
	penerbit string
	eksemplar int
	rating   int
}

// Fungsi untuk melakukan Insertion Sort berdasarkan rating buku
func insertionSortBuku(arr []Buku) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Geser elemen-elemen yang lebih besar dari key berdasarkan rating
		for j >= 0 && arr[j].rating > key.rating {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	var N int
	var pustaka []Buku

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku yang ada di dalam perpustakaan: ")
	fmt.Scan(&N)

	// Input data buku
	for i := 0; i < N; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)

		fmt.Print("ID Buku: ")
		fmt.Scan(&buku.id)

		fmt.Print("Judul Buku: ")
		fmt.Scan(&buku.judul)

		fmt.Print("Penulis Buku: ")
		fmt.Scan(&buku.penulis)

		fmt.Print("Penerbit Buku: ")
		fmt.Scan(&buku.penerbit)

		fmt.Print("Jumlah Eksemplar: ")
		fmt.Scan(&buku.eksemplar)

		fmt.Print("Rating Buku (1-5): ")
		fmt.Scan(&buku.rating)

		pustaka = append(pustaka, buku)
	}

	// Urutkan data buku berdasarkan rating
	insertionSortBuku(pustaka)

	// Output data buku setelah diurutkan berdasarkan rating
	fmt.Println("Data Buku di Perpustakaan (diurutkan berdasarkan rating):")
	for _, buku := range pustaka {
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.rating)
	}
}

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    float64
}

// Struct untuk daftar buku
type DaftarBuku struct {
	Pustaka []Buku
}

func main() {
	var daftarBuku DaftarBuku
	var n int
	fmt.Println("Masukkan jumlah buku:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d:\n", i+1)

		fmt.Print("ID: ")
		fmt.Scan(&buku.ID)

		fmt.Print("Judul: ")
		buku.Judul = inputString()

		fmt.Print("Penulis: ")
		buku.Penulis = inputString()

		fmt.Print("Penerbit: ")
		buku.Penerbit = inputString()

		fmt.Print("Eksemplar: ")
		fmt.Scan(&buku.Eksemplar)

		fmt.Print("Tahun: ")
		fmt.Scan(&buku.Tahun)

		fmt.Print("Rating: ")
		fmt.Scan(&buku.Rating)

		daftarBuku.Pustaka = append(daftarBuku.Pustaka, buku)
	}

	tampilkanBukuTertinggi(&daftarBuku)
	urutkanBukuBerdasarkanTahun(&daftarBuku)
	fmt.Println("\nDaftar buku setelah diurutkan berdasarkan tahun:")
	tampilkanDaftarBuku(&daftarBuku)
}

func inputString() string {
	var input string
	fmt.Scan(&input) // Menggunakan fmt.Scan untuk membaca string
	return strings.TrimSpace(input)
}

func tampilkanBukuTertinggi(daftarBuku *DaftarBuku) {
	if len(daftarBuku.Pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	sort.SliceStable(daftarBuku.Pustaka, func(i, j int) bool {
		return daftarBuku.Pustaka[i].Rating > daftarBuku.Pustaka[j].Rating
	})

	bukuTertinggi := daftarBuku.Pustaka[0]
	fmt.Println("\nBuku dengan rating tertinggi:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Rating: %.2f\n",
		bukuTertinggi.Judul, bukuTertinggi.Penulis, bukuTertinggi.Penerbit, bukuTertinggi.Rating)
}

func urutkanBukuBerdasarkanTahun(daftarBuku *DaftarBuku) {
	sort.SliceStable(daftarBuku.Pustaka, func(i, j int) bool {
		return daftarBuku.Pustaka[i].Tahun < daftarBuku.Pustaka[j].Tahun
	})
}

func tampilkanDaftarBuku(daftarBuku *DaftarBuku) {
	for _, buku := range daftarBuku.Pustaka {
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %.2f\n",
			buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
	}
}
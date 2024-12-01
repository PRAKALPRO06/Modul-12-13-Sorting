// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import "fmt"

const nMax = 10000

type Buku struct {
	id                       string
	judul                    string
	penulis                  string
	penerbit                 string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit,
			&pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

// Prosedur untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("\nTidak ada data buku.")
		return
	}

	terfavorit := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("\n=============================================")
	fmt.Println("Buku Terfavorit:")
	fmt.Println("=============================================")
	fmt.Printf("Judul    : %s\n", terfavorit.judul)
	fmt.Printf("Penulis  : %s\n", terfavorit.penulis)
	fmt.Printf("Penerbit : %s\n", terfavorit.penerbit)
	fmt.Printf("Tahun    : %d\n", terfavorit.tahun)
	fmt.Printf("Rating   : %d\n", terfavorit.rating)
	fmt.Println("=============================================")
}

// Prosedur untuk mengurutkan buku berdasarkan rating secara descending
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

// Prosedur untuk mencetak 5 buku dengan rating tertinggi
func CetakTerbaru(pustaka DaftarBuku, n int) {
	fmt.Println("\n=============================================")
	fmt.Println("5 Buku Dengan Rating Tertinggi:")
	fmt.Println("=============================================")
	count := 5
	if n < 5 {
		count = n
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
	fmt.Println("=============================================")
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku(pustaka DaftarBuku, n, r int) {
	found := false
	fmt.Printf("\nMencari Buku dengan Rating %d:\n", r)

	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			fmt.Printf("\nJudul    : %s\n", pustaka[i].judul)
			fmt.Printf("Penulis  : %s\n", pustaka[i].penulis)
			fmt.Printf("Penerbit : %s\n", pustaka[i].penerbit)
			fmt.Printf("Tahun    : %d\n", pustaka[i].tahun)
			fmt.Printf("Eksemplar: %d\n", pustaka[i].eksemplar)
			fmt.Printf("Rating   : %d\n", pustaka[i].rating)
			fmt.Println("=============================================")
		}
	}

	if !found {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var ratingCari int

	// Untuk menjalankan Prosedur Daftarkan Buku
	DaftarkanBuku(&pustaka, &n)

	// Untuk menjalankan Prosedur Cetak Buku Terfavorit
	CetakTerfavorit(pustaka, n)

	// Untuk menjalankan Prosedur Urutkan Buku berdasarkan Rating
	UrutBuku(&pustaka, n)

	// Untuk menjalankan Prosedur Cetak 5 buku dengan rating tertinggi
	CetakTerbaru(pustaka, n)

	// Untuk mencari buku berdasarkan rating
	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}

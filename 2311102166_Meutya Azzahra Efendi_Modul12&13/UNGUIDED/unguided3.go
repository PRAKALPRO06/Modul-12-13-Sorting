// Meutya Azzahra Efendi
// 2311102166
// IF-11-06

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const nMax = 10000

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

// Fungsi untuk membaca input dari pengguna
func bacaInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Prosedur untuk mendaftarkan buku ke perpustakaan
func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	jumlah, _ := strconv.Atoi(bacaInput("Masukkan jumlah buku yang ingin didaftarkan: "))
	*n = jumlah

	for i := 0; i < *n; i++ {
		fmt.Printf("\nMasukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		input := bacaInput("")
		data := strings.Split(input, " ")
		if len(data) != 7 {
			fmt.Println("Input tidak valid. Harap masukkan 7 data sesuai format.")
			i-- // Ulangi iterasi untuk input yang salah
			continue
		}
		pustaka[i] = Buku{
			id:        data[0],
			judul:     data[1],
			penulis:   data[2],
			penerbit:  data[3],
			eksemplar: atoi(data[4]),
			tahun:     atoi(data[5]),
			rating:    atoi(data[6]),
		}
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

	fmt.Println("\nBuku Terfavorit:")
	tampilkanBuku(terfavorit)
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
	fmt.Println("\n5 Buku Dengan Rating Tertinggi:")
	count := min(5, n)
	for i := 0; i < count; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// Prosedur untuk mencari buku berdasarkan ratingnya
func CariBuku(pustaka DaftarBuku, n, r int) {
	fmt.Printf("\nMencari Buku dengan Rating %d:\n", r)
	found := false
	for i := 0; i < n; i++ {
		if pustaka[i].rating == r {
			found = true
			tampilkanBuku(pustaka[i])
		}
	}
	if !found {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

// Fungsi untuk mencetak data buku
func tampilkanBuku(buku Buku) {
	fmt.Printf("Judul    : %s\n", buku.judul)
	fmt.Printf("Penulis  : %s\n", buku.penulis)
	fmt.Printf("Penerbit : %s\n", buku.penerbit)
	fmt.Printf("Tahun    : %d\n", buku.tahun)
	fmt.Printf("Eksemplar: %d\n", buku.eksemplar)
	fmt.Printf("Rating   : %d\n", buku.rating)
}

// Fungsi pembantu untuk konversi string ke int
func atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

// Fungsi pembantu untuk mendapatkan nilai minimum
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var pustaka DaftarBuku
	var n int

	DaftarkanBuku(&pustaka, &n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	CetakTerbaru(pustaka, n)

	ratingCari := atoi(bacaInput("\nMasukkan rating yang ingin dicari: "))
	CariBuku(pustaka, n, ratingCari)
}

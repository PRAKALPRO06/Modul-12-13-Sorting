package main 

import (
	"fmt"
	"math"
)

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}
type DaftarBuku [nMax]Buku

func DaftarkanBuku(daftar_buku *DaftarBuku, n int) {
	fmt.Println("Masukkan data buku di setiap baris (id judul penulis penerbit eksemplar tahun rating)")
	for i:=0; i<n; i++ {
		fmt.Scan(&daftar_buku[i].id, &daftar_buku[i].judul, &daftar_buku[i].penulis, &daftar_buku[i].penerbit, &daftar_buku[i].eksemplar, &daftar_buku[i].tahun, &daftar_buku[i].rating)
	}
}
func CetakTerfavorit(daftar_buku DaftarBuku, n int) {
	var buku_terfavorit Buku = daftar_buku[0]
	// // Jika buku belum diurutkan maka uncomment bagian ini
	// for i:=1; i<n; i++ {
	// 	if buku_terfavorit.rating < daftar_buku[i].rating {
	// 		buku_terfavorit = daftar_buku[i]
	// 	}
	// }
	fmt.Println("Buku terfavorit adalah buku dengan judul: \"", buku_terfavorit.judul, "\"")
}
func UrutBukuBerdasarkanRating(daftar_buku *DaftarBuku, n int) {
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if daftar_buku[j].rating > daftar_buku[idxMax].rating {
				idxMax = j
			}
		}
		// Tukar elemen terkecil dengan elemen di posisi i
		daftar_buku[i], daftar_buku[idxMax] = daftar_buku[idxMax], daftar_buku[i]
	}
}
func Cetak5Terbaru(daftar_buku DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi: ")
	for i:=0; i<5; i++ {
		fmt.Println(i+1, ". ", daftar_buku[i].judul, ", rating: ", daftar_buku[i].rating)
	}
	fmt.Print("\n")
}
func CariBuku(daftar_buku DaftarBuku, n, rating int) {
	ditemukan := false
	top := n-1
	bottom := 0
	middle := int((top + bottom)/2)
	// Mengecek index pertama dan terakhir daftar_buku terlebih dahulu sebelum lanjut ke proses Binary Search
	if daftar_buku[bottom].rating == rating && rating != 0 {
		ditemukan = true
		fmt.Println(daftar_buku[bottom].judul)
		// Mulai menelusuri buku-buku yang berada di dekat buku daftar_buku[0] dan mencetak judulnya jika ratingnya sama
		for i:=bottom+1; i<n; i++ {
			if(daftar_buku[i].rating != rating) { break }
			fmt.Println(daftar_buku[i].judul)
		}
		return
	}
	if daftar_buku[top].rating == rating && rating != 0 {
		ditemukan = true
		fmt.Println(daftar_buku[top].judul)
		// Mulai menelusuri buku-buku yang berada di dekat buku daftar_buku[0] dan mencetak judulnya jika ratingnya sama
		for i:=top+1; i>=0; i-- {
			if(daftar_buku[i].rating != rating) { break }
			fmt.Println(daftar_buku[i].judul)
		}
		return
	}
	for middle > 0 && middle < n {
		if daftar_buku[middle].rating == rating && rating != 0 {
			ditemukan = true
			fmt.Println(daftar_buku[middle].judul)
			// Mulai menelusuri buku-buku yang berada di dekat buku daftar_buku[middle] dan mencetak judulnya jika ratingnya sama
			for i:=middle-1; i>=0; i-- {
				if(daftar_buku[i].rating != rating) { break }
				fmt.Println(daftar_buku[i].judul)
			}
			for i:=middle+1; i<n; i++ {
				if(daftar_buku[i].rating != rating) { break }
				fmt.Println(daftar_buku[i].judul)
			}
			return
		} else if daftar_buku[middle].rating < rating {
			bottom = middle
			middle = int(math.Ceil(float64(top + bottom)/2))
		} else {
			top = middle
			middle = int(math.Floor(float64(top + bottom)/2))
		}
		fmt.Println(middle)
	}
	if ! ditemukan {
		fmt.Println("Buku dengan rating ", rating, " tidak ditemukan")
	}
}

func main() {

	var daftar_buku DaftarBuku
	var n int
	fmt.Print("Masukkan jumlah buku (n): ")
	fmt.Scan(&n)

	DaftarkanBuku(&daftar_buku, n)
	UrutBukuBerdasarkanRating(&daftar_buku, n)

	fmt.Print("\n")
	CetakTerfavorit(daftar_buku, n)
	fmt.Print("\n")
	Cetak5Terbaru(daftar_buku, n)

	var rating int
	fmt.Print("Masukkan rating buku yang dicari: ")
	fmt.Scan(&rating)
	CariBuku(daftar_buku, n, rating)
	
}

// Contoh input:
// 1 Buku_A Jhon Penerbit_A 200 2024 4
// 2 Buku_B Alex Penerbit_A 150 2022 3
// 3 Buku_C Jhon Penerbit_B 200 2020 4
// 4 Buku_D Steve Penerbit_A 150 2024 4
// 5 Buku_E Jhon Penerbit_A 300 2020 5
// 6 Buku_F Jhon Penerbit_B 400 2024 4
// 7 Buku_G Alex Penerbit_A 200 2024 3
// 8 Buku_H Jhon Penerbit_A 150 2022 3
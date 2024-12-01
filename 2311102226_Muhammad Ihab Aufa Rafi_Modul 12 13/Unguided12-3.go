package main

import "fmt"

const maxBuku = 7919

type Buku struct {
    id, judul, penulis, penerbit    string
    eksemplar, tahun, rating        int
}

type DaftarBuku [maxBuku]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
    fmt.Print("Masukkan ID Buku: ")
    fmt.Scanln(&pustaka[*n].id)
    fmt.Print("Masukkan Judul Buku: ")
    fmt.Scanln(&pustaka[*n].judul)
    fmt.Print("Masukkan Nama Penulis: ")
    fmt.Scanln(&pustaka[*n].penulis)
    fmt.Print("Masukkan Nama Penerbit: ")
    fmt.Scanln(&pustaka[*n].penerbit)
    fmt.Print("Masukkan Jumlah Eksemplar: ")
    fmt.Scanln(&pustaka[*n].eksemplar)
    fmt.Print("Masukkan Tahun Terbit: ")
    fmt.Scanln(&pustaka[*n].tahun)
    fmt.Print("Masukkan Rating Buku (1-10): ")
    fmt.Scanln(&pustaka[*n].rating)
    *n++
}

func CetakFavorit(pustaka DaftarBuku, n int) {
    fmt.Println("\nBuku Terfavorit:")
    maxRating := 0
    for i := 0; i < n; i++ {
        if pustaka[i].rating > maxRating {
            maxRating = pustaka[i].rating
        }
    }
    
    for i := 0; i < n; i++ {
        if pustaka[i].rating == maxRating {
            fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", 
            pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].rating)
        }
    }
}

func UrutBuku(pustaka *DaftarBuku, n int) {
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
	fmt.Println("5 Judul Buku dengan Rating Tertinggi:")
	for i := 0; i < n && i < 5; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
    fmt.Printf("\nBuku dengan Rating %d:\n", r)
    found := false
    for i := 0; i < n; i++ {
        if pustaka[i].rating == r {
            fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", 
            pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].tahun, pustaka[i].rating)
            found = true
        }
    }
    if !found {
        fmt.Println("Tidak ada buku dengan rating seperti itu")
    }
}

func main() {
    var pustaka DaftarBuku
    var n226, cariRating int
    var count int

    fmt.Print("Masukkan jumlah buku: ")
    fmt.Scanln(&n226)

    fmt.Println("\nMasukkan data untuk setiap buku:")
    for i := 0; i < n226; i++ {
        fmt.Printf("\nData Buku ke-%d:\n", i+1)
        DaftarkanBuku(&pustaka, &count)
    }

    CetakFavorit(pustaka, n226)
    
    fmt.Println()
    UrutBuku(&pustaka, n226)
    Cetak5Terbaru(pustaka, n226)

    fmt.Print("\nMasukkan rating yang dicari: ")
    fmt.Scanln(&cariRating)
    CariBuku(pustaka, n226, cariRating)
}
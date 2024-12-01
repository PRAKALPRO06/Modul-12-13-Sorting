package main
import (
	"fmt"
)

type Buku struct {
	id       int
	judul    string
	penulis  string
	penerbit string
	tahun    int
	rating   int
}

type DaftarBuku []Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Println("Masukkan data buku (id, judul, penulis, penerbit, tahun, rating):")
	for i := 0; i < *n; i++ {
		var id, tahun, rating int
		var judul, penulis, penerbit string
		fmt.Scanln(&id, &judul, &penulis, &penerbit, &tahun, &rating)
		*pustaka = append(*pustaka, Buku{id, judul, penulis, penerbit, tahun, rating})
	}
	fmt.Printf("=================================================================== \n")
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := pustaka[0].rating
	var favorit Buku
	for _, buku := range pustaka {
		if buku.rating > maxRating {
			maxRating = buku.rating
			favorit = buku
		}
	}
	fmt.Printf("Buku Terfavorit: %s oleh %s (%s, %d)\n", favorit.judul, favorit.penulis, favorit.penerbit, favorit.tahun)
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
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		buku := pustaka[i]
		fmt.Printf("%s oleh %s (%s, %d) - Rating: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			buku := pustaka[mid]
			fmt.Printf("Ditemukan: %s oleh %s (%s, %d) - Rating: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Print("Input Jumlah Buku: ")
	fmt.Scanln(&n)

	DaftarkanBuku(&pustaka, &n)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	var r int
	fmt.Print("Masukkan rating yang ingin dicari: ")
	fmt.Scanln(&r)
	CariBuku(pustaka, n, r)
}
//WISNU RANANTA RADITYA PUTRA (2311102013) IF-11-06
package main
import (
	"fmt"
)
type Buku struct {
	id_2311102013 int
	eksemplar int
	tahun int
	rating int
	judul, penerbit, penulis string
}

func DaftarkanBuku(pustaka *[]Buku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (ID, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id_2311102013, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

func CetakFavorit(pustaka []Buku, n int) {
	if len(pustaka) == 0 {
		fmt.Println("Pustaka kosong.")
		return
	}
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}
	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		terfavorit.id_2311102013, terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.eksemplar, terfavorit.tahun, terfavorit.rating)
}

func UrutkanBuku(pustaka *[]Buku, n int) {
	for i := 1; i < len(*pustaka); i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

func Cetak5Terbaik(pustaka []Buku, n int) {
	fmt.Println("Lima buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < len(pustaka); i++ {
		buku := pustaka[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id_2311102013, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func CariBuku(pustaka []Buku, n int, r int) {
	ditemukan := false
	for _, buku := range pustaka {
		if buku.rating == r {
			ditemukan = true
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.id_2311102013, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}
func main() {
	var n int
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&n)

	if n <= 0 || n > 7919 {
		fmt.Println("Jumlah buku harus antara 1 hingga 7919.")
		return
	}

	var pustaka []Buku

	DaftarkanBuku(&pustaka, n)
	CetakFavorit(pustaka, n)
	UrutkanBuku(&pustaka, n)

	Cetak5Terbaik(pustaka, n)
	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}
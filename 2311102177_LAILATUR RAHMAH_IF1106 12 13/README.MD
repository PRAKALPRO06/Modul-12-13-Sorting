<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XII & XIII</strong></h2>
<h2 align="center"><strong> PENGURUTAN DATA </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  LAILATUR RAHMAH / 2311102177<br>
  S1 IF 11 06
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Abednego Dwi Septiadi
  
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Features

- [DASAR TEORI](#dasar-teori)
- [GUIDED](#guided)
- [UNGUIDED](#unguided)

## DASAR TEORI

**Definisi**
Pengurutan data (sorting) adalah proses penyusunan elemen-elemen data dalam urutan tertentu, seperti urutan menaik (ascending) atau menurun (descending). Proses pengurutan ini digunakan dalam banyak algoritma komputer, baik itu untuk pencarian, perbandingan, maupun pengolahan data yang lebih efisien. Dalam pemrograman bahasa Go, pengurutan data dapat dilakukan menggunakan berbagai algoritma yang masing-masing memiliki karakteristik yang berbeda dalam hal kompleksitas waktu dan ruang.

### Prinsip algoritma Selection Sort
Selection Sort adalah salah satu algoritma pengurutan sederhana yang bekerja dengan cara memilih elemen terkecil (atau terbesar, tergantung urutan yang diinginkan) dalam suatu array dan menempatkannya di posisi yang benar, kemudian mengulangi proses ini untuk elemen-elemen yang tersisa.

#### Prinsip Kerja Selection Sort
- Iterasi 1: Cari elemen terkecil dalam array dan tukar dengan elemen pertama.
- Iterasi 2: Cari elemen terkecil di bagian array yang tersisa dan tukar dengan elemen kedua.
- Iterasi 3: Cari elemen terkecil di bagian array yang tersisa dan tukar dengan elemen ketiga.
- Proses ini terus berulang hingga seluruh array terurut.

#### Langkah-Langkah Selection Sort
- Mulai dari elemen pertama dalam array, tentukan indeks elemen terkecil dalam array (dari indeks yang sedang diproses hingga akhir array).
- Tukar elemen terkecil tersebut dengan elemen di indeks yang sedang diproses.
- Pindah ke elemen berikutnya dan ulangi proses ini hingga array terurut.

### Prinsip algoritma Insertion Sort
Insertion Sort adalah algoritma pengurutan yang lebih efisien dibandingkan Selection Sort untuk array kecil atau array yang sudah sebagian terurut. Prinsip dasar dari Insertion Sort adalah membangun array yang terurut satu elemen pada satu waktu dengan cara memasukkan elemen baru ke posisi yang benar dalam urutan yang sudah terurut sebelumnya.

#### Prinsip Kerja Insertion Sort
- Mulai dari elemen kedua (karena elemen pertama sudah dianggap terurut).
- Ambil elemen yang akan dimasukkan ke dalam urutan yang sudah terurut, dan bandingkan dengan elemen-elemen sebelumnya.
- Geser elemen-elemen yang lebih besar satu posisi ke kanan untuk memberi ruang bagi elemen yang baru.
- Tempatkan elemen tersebut pada posisi yang tepat.
- Ulangi proses ini untuk elemen-elemen berikutnya sampai array terurut.

#### Langkah-Langkah Insertion Sort
- Mulai dari elemen kedua (karena elemen pertama dianggap sudah terurut).
- Ambil elemen tersebut dan bandingkan dengan elemen-elemen yang lebih kecil di sebelah kirinya.
- Geser elemen-elemen yang lebih besar ke kanan.
- Tempatkan elemen tersebut pada posisi yang benar di array.
- Ulangi proses ini untuk setiap elemen dari kiri ke kanan hingga array terurut.


## GUIDED

### Guided 1

#### Study Case

Mengurutkan Nomor Rumah Kerabat di Berbagai Daerah

#### Source Code

```go
package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int, n int) {
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			// Cari elemen terkecil
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		// Tukar elemen terkecil dengan elemen di posisi i
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses tiap daerah
	for daerah := 1; daerah <= n; daerah++ {
		var m int
		fmt.Printf("\nMasukkan jumlah nomor rumah kerabat untuk daerah %d: ", daerah)
		fmt.Scan(&m)

		// Membaca nomor rumah untuk daerah ini
		arr := make([]int, m)
		fmt.Printf("Masukkan %d nomor rumah kerabat: ", m)
		for i := 0; i < m; i++ {
			fmt.Scan(&arr[i])
		}

		// Urutkan array dari terkecil ke terbesar
		selectionSort(arr, m)

		// Tampilkan hasil
		fmt.Printf("Nomor rumah terurut untuk daerah %d: ", daerah)
		for _, num := range arr {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/1de33c7a-afa2-4fac-918a-a5713ff01648)

#### Deskripsi Program

Program ini dirancang untuk membantu pengguna mengurutkan nomor rumah kerabat di berbagai daerah menggunakan algoritma Selection Sort. Pengguna dapat memasukkan jumlah daerah, jumlah nomor rumah di setiap daerah, dan daftar nomor rumah untuk setiap daerah tersebut. Program akan mengurutkan nomor rumah dari terkecil ke terbesar untuk setiap daerah secara terpisah dan menampilkan hasilnya.

### Guided 2

#### Study Case

Mengelola Data Usia Peserta Kursus

#### Source Code

```go
package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan Insertion Sort
func insertionSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah selisih elemen array tetap
func isConstantDifference(arr []int, n int) (bool, int) {
	if n < 2 {
		return true, 0
	}

	difference := arr[1] - arr[0]
	for i := 1; i < n-1; i++ {
		if arr[i+1]-arr[i] != difference {
			return false, 0
		}
	}
	return true, difference
}

func main() {
	var arr []int
	var num int

	// Input data hingga bilangan negatif ditemukan
	fmt.Println("Masukkan data integer (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}

	n := len(arr)

	// Urutkan array menggunakan Insertion Sort
	insertionSort(arr, n)

	// Periksa apakah selisih elemen tetap
	isConstant, difference := isConstantDifference(arr, n)

	// Tampilkan hasil pengurutan
	fmt.Println("Array setelah diurutkan:")
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// Tampilkan status jarak
	if isConstant {
		fmt.Printf("Data berjarak %d\n", difference)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/e6c66e0c-d90e-4d7e-aca5-13c174169ba8)

#### Deskripsi Program

Program ini dirancang untuk mengelola data integer, seperti usia peserta kursus, dengan dua fungsi utama: mengurutkan data menggunakan algoritma Insertion Sort dan memeriksa apakah selisih antar elemen memiliki jarak konstan. Pengguna dapat memasukkan data integer secara berulang hingga memasukkan angka negatif untuk menghentikan input. Setelah data diurutkan, program menganalisis perbedaan antar elemen array untuk menentukan apakah jaraknya tetap, dan jika ya, menampilkan nilai jaraknya. Program ini cocok digunakan untuk analisis sederhana dalam berbagai aplikasi, seperti pengelompokan data usia atau identifikasi pola jarak data.

## Unguided

### Unguided 1

#### Study Case

Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

Format **Masukan** masih persis sama seperti sebelumnya.

**Keluaran** terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

| No | Masukan        | Keluaran       |
|---|----------------|----------------|
| 1 | 3              | 1 13 12 8 2   |
|   | 5 2 1 7 9 13  | 15 27 39 75 133 189 |
|   | 6 189 15 27 39 75 133 | 8 4 2        |
|   | 3 4 9 1        |                |

**Keterangan:**
Terdapat 3 daerah dalam contoh masukan. Baris kedua berisi campuran bilangan ganjil dan genap. Baris berikutnya hanya berisi bilangan ganjil, dan baris terakhir hanya berisi bilangan genap.

**Petunjuk:**
- Waktu pembacaan data, bilangan ganjil dan genap dipisahkan ke dalam dua array yang berbeda, untuk kemudian masing-masing diurutkan tersendiri.
- Atau, tetap disimpan dalam satu array, diurutkan secara keseluruhan. Tetapi pada waktu pencetakan, mulai dengan mencetak semua nilai ganjil lebih dulu, kemudian setelah selesai cetaklah semua nilai genapnya.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import "fmt"

// Fungsi untuk mengurutkan array dengan selection sort
func selectionSort(arr []int, asc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && arr[j] < arr[idx]) || (!asc && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

// Fungsi untuk mencetak garis batas panjang
func printSeparator() {
	fmt.Println("\n=================================================")
}

func printDashedSeparator() {
	fmt.Println("-------------------------------------------------")
}

// Fungsi untuk menampilkan nomor rumah berdasarkan daerah
func processDaerah(i, m int) ([]int, []int) {
	var arr = make([]int, m)
	// Input nomor rumah kerabat untuk setiap daerah
	fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

	// Memisahkan nomor ganjil dan genap
	var odd, even []int
	for _, num := range arr {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return odd, even
}

// Fungsi untuk menampilkan hasil terurut
func printSortedResults(i int, odd, even []int) {
	// Mengurutkan nomor rumah
	selectionSort(odd, true)   // urutkan ganjil secara menaik
	selectionSort(even, false) // urutkan genap secara menurun

	// Output nomor rumah yang terurut
	fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)

	// Tampilkan nomor ganjil
	for _, num := range odd {
		fmt.Printf("%d ", num)
	}

	// Tampilkan nomor genap
	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Fungsi utama program
func main() {
	var n int
	// Input jumlah daerah kerabat
	printSeparator()
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	// Proses untuk setiap daerah
	for i := 0; i < n; i++ {
		var m int
		// Input jumlah rumah kerabat untuk setiap daerah
		printDashedSeparator()
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		// Ambil nomor rumah untuk daerah tersebut
		odd, even := processDaerah(i, m)

		// Menampilkan hasil terurut untuk daerah tersebut
		printSortedResults(i, odd, even)

		// Cetak garis pemisah untuk setiap daerah
		printDashedSeparator()
	}
}

```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/7ff83c85-0fad-4ddc-956b-fcaa1235bf54)

#### Deskripsi Program

Program ini dirancang untuk mengelola daftar nomor rumah kerabat di beberapa daerah. Pengguna diminta memasukkan jumlah daerah yang ingin diproses dan jumlah nomor rumah untuk setiap daerah. Program memisahkan nomor rumah menjadi dua kategori: ganjil dan genap, kemudian mengurutkan nomor ganjil secara menaik dan nomor genap secara menurun menggunakan algoritma Selection Sort. Hasil pengurutan nomor rumah untuk setiap daerah ditampilkan dalam format yang rapi. Program ini memberikan solusi terstruktur untuk mengelompokkan dan mengurutkan data numerik berdasarkan atribut tertentu (ganjil/genap) dengan pendekatan sederhana namun efektif.

**Alur Kerja Program:**

**1. Memasukkan Jumlah Daerah**
- Program meminta pengguna untuk memasukkan jumlah daerah (`n`) yang akan dikelola.

**2. Memproses Setiap Daerah**
- Untuk setiap daerah, pengguna diminta memasukkan:
  - Jumlah rumah kerabat di daerah tersebut (`m`).
  - Daftar nomor rumah sebanyak `m`.

**3. Memisahkan Nomor Rumah**
- Nomor rumah dipisahkan menjadi dua kategori:
  - **Ganjil**: Nomor yang tidak habis dibagi 2.
  - **Genap**: Nomor yang habis dibagi 2.

**4. Mengurutkan Nomor Rumah**
- Nomor ganjil diurutkan secara **menaik** (ascending).
- Nomor genap diurutkan secara **menurun** (descending).

**5. Menampilkan Hasil**
- Untuk setiap daerah, program menampilkan nomor rumah yang telah diurutkan:
  - Nomor ganjil (menaik).
  - Nomor genap (menurun).

**6. Mengulang untuk Semua Daerah**
- Langkah 2–5 diulangi untuk semua daerah yang dimasukkan oleh pengguna.

**Struktur Output**
- Setiap hasil untuk daerah tertentu dipisahkan dengan garis batas untuk memberikan kejelasan.

### Unguided 2

#### Study Case

Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

"Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah."

Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

**Masukan:** berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

**Keluaran:** adalah median yang diminta, satu data per baris.

| No | Masukan                                 | Keluaran |
|---|-------------------------------------------|----------|
| 1 | 7 23 11 8 5 19 2 29 3 13 17 0 -5313 | 11       |
|   |                                          | 12       |

**Keterangan:**
Sampai bilangan 0 yang pertama, data terbaca adalah 7 23 11, setelah tersusun: 7 11 23, maka median saat itu adalah 11.

Sampai bilangan 0 yang kedua, data adalah 7 23 11 5 19 2 29 3 13 17, setelah tersusun diperoleh: 2 3 5 7 11 13 17 19 23 29. Karena ada 10 data, genap, maka median adalah (11+13)/2 = 12.

**Petunjuk:**
Untuk setiap data bukan 0 (dan bukan marker -5313541) simpan ke dalam array. Dan setiap kali menemukan bilangan 0, urutkanlah data yang sudah tersimpan dengan menggunakan metode insertion sort dan ambil mediannya.

#### Source Code

```go
// LAILATUR RAHMAH
// 2311102177
// IF 11 06

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk melakukan selection sort
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Cari elemen terkecil di bagian yang belum terurut
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Tukar elemen terkecil yang ditemukan dengan elemen pertama
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi untuk menghitung median dari array yang sudah terurut
func calculateMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		// Jika jumlah elemen genap, median adalah rata-rata dari dua nilai tengah
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	}
	// Jika jumlah elemen ganjil, median adalah elemen tengah
	return float64(arr[n/2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca seluruh input sebagai satu baris
	fmt.Println("Masukkan angka yang dipisahkan dengan spasi (akhiri dengan -5313):")
	scanner.Scan()
	line := scanner.Text()

	// Memisahkan input ke dalam bagian-bagian
	parts := strings.Split(line, " ")
	var data []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Input tidak valid, harap masukkan angka bulat saja.")
			return
		}

		if num == -5313 {
			// Akhir dari input
			break
		} else if num == 0 {
			// Urutkan data dan hitung median
			selectionSort(data)
			median := calculateMedian(data)
			fmt.Printf("%.0f\n", median)
		} else {
			// Tambahkan angka ke array data
			data = append(data, num)
		}
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/757ddf8e-c696-4cd6-8251-200dcbc2b9c4)

#### Deskripsi Program

Program ini membaca serangkaian angka dari pengguna, mengurutkannya menggunakan algoritma Selection Sort, dan menghitung nilai median dari angka yang sudah diurutkan. Pengguna dapat memasukkan angka yang dipisahkan dengan spasi, dan setiap kali angka 0 dimasukkan, program akan mengurutkan data yang telah dimasukkan, menghitung median, dan menampilkan hasilnya. Program berhenti ketika pengguna memasukkan angka -5313. Median dihitung berdasarkan jumlah elemen: jika jumlah elemen ganjil, median adalah elemen tengah; jika genap, median adalah rata-rata dua elemen tengah. Program ini memastikan fleksibilitas dan memberikan hasil secara interaktif.

**Cara Kerja Program:**

**1. Input Angka**
- Program membaca satu baris input angka yang dipisahkan dengan spasi menggunakan `bufio.Scanner`.
- Input terdiri dari angka-angka bulat yang diakhiri dengan **`-5313`** sebagai tanda akhir input.

**2. Proses Input**
- Angka-angka diproses sebagai berikut:
  - Jika angka **`0`** dimasukkan, program akan:
    - Mengurutkan semua angka yang telah dimasukkan menggunakan **Selection Sort**.
    - Menghitung nilai **median** dari angka yang sudah diurutkan.
    - Menampilkan hasil median.
  - Angka lain (selain `0` dan `-5313`) akan ditambahkan ke dalam array `data`.
  - Jika angka **`-5313`** ditemukan, program berhenti memproses input.

**3. Pengurutan Data**
- Angka-angka yang dimasukkan diurutkan menggunakan algoritma **Selection Sort**:
  - Cari elemen terkecil di array dan tukar dengan elemen pertama.
  - Ulangi langkah tersebut untuk elemen berikutnya hingga seluruh array terurut.

**4. Perhitungan Median**
- Jika jumlah elemen **ganjil**, median adalah elemen tengah.
- Jika jumlah elemen **genap**, median adalah rata-rata dari dua elemen tengah.

### Unguided 3

#### Study Case

Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:

```
const nMax : integer = 7919;
type Buku = record
  id, judul, penulis, penerbit : string;
  eksemplar, tahun, rating : integer;
end;

type DaftarBuku = array [1..nMax] of Buku;

var
  Pustaka : DaftarBuku;
  nPustaka : integer;
```
**Masukan:**
Terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

**Keluaran:**
Terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### Source Code

```go
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
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/30764820-ab00-4383-8d40-6109abc467d7)
![image](https://github.com/user-attachments/assets/7a989312-8f57-43b9-82e6-43212eeffe64)

#### Deskripsi Program

Program di atas merupakan aplikasi manajemen data buku perpustakaan. Program memungkinkan pengguna untuk mendaftarkan sejumlah buku dengan memasukkan data seperti ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating. Setelah data dimasukkan, program dapat menampilkan buku dengan rating tertinggi (buku terfavorit), mengurutkan semua buku berdasarkan rating secara menurun (descending), dan mencetak hingga lima buku dengan rating tertinggi. Selain itu, program juga menyediakan fitur pencarian buku berdasarkan rating yang dimasukkan pengguna. Jika buku dengan rating tersebut ditemukan, program akan menampilkan detail buku; jika tidak, akan memberi pesan bahwa buku tidak ditemukan.

**Cara Kerja Program Manajemen Perpustakaan**

Program ini merupakan aplikasi untuk mengelola data buku di perpustakaan. Fitur-fiturnya mencakup pendaftaran buku, pengurutan berdasarkan rating, pencarian buku, dan menampilkan buku dengan rating tertinggi. Berikut adalah cara kerja program secara rinci:

**1. Pendaftaran Buku**
- Program meminta input jumlah buku yang akan didaftarkan.
- Untuk setiap buku, program meminta data berikut:
  - **ID buku**
  - **Judul buku**
  - **Penulis buku**
  - **Penerbit buku**
  - **Jumlah eksemplar**
  - **Tahun terbit**
  - **Rating buku**
- Data yang dimasukkan disimpan dalam array `pustaka`.

**2. Menampilkan Buku Terfavorit**
- Program mencari buku dengan **rating tertinggi**.
- Buku dengan rating tertinggi ditampilkan dengan informasi lengkap:
  - Judul buku
  - Penulis buku
  - Penerbit buku
  - Tahun terbit
  - Rating

**3. Pengurutan Buku Berdasarkan Rating**
- Buku diurutkan berdasarkan **rating** secara **menurun (descending)**.
- Algoritma yang digunakan adalah **Insertion Sort**.

**4. Mencetak 5 Buku Dengan Rating Tertinggi**
- Setelah pengurutan, program menampilkan **5 buku** dengan rating tertinggi.
- Jika jumlah buku kurang dari 5, program menampilkan semua buku yang ada.

**5. Mencari Buku Berdasarkan Rating**
- Pengguna memasukkan **rating** tertentu untuk mencari buku.
- Program mencari buku dengan rating tersebut dan menampilkan detail:
  - Judul buku
  - Penulis buku
  - Penerbit buku
  - Tahun terbit
  - Jumlah eksemplar
  - Rating
- Jika tidak ada buku dengan rating tersebut, program memberikan pesan bahwa buku tidak ditemukan.

**Contoh Alur Program**
1. Pengguna memasukkan jumlah buku yang akan didaftarkan, misalnya **3**.
2. Program meminta data buku:
   - **Buku 1**: ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating.
   - **Buku 2**: ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating.
   - **Buku 3**: ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating.
3. Program menampilkan buku dengan rating tertinggi.
4. Program mengurutkan buku berdasarkan rating dan menampilkan 5 buku teratas.
5. Pengguna memasukkan rating tertentu, dan program mencari serta menampilkan buku dengan rating tersebut.

**Struktur Data yang Digunakan**
- Program menggunakan array `DaftarBuku` untuk menyimpan data buku.
- Setiap buku direpresentasikan dalam struct `Buku` dengan atribut:
  - `id` (string)
  - `judul` (string)
  - `penulis` (string)
  - `penerbit` (string)
  - `eksemplar` (int)
  - `tahun` (int)
  - `rating` (int)

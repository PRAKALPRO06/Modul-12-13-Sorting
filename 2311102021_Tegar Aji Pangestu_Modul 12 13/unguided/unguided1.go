package main
//Tegar Aji Pangestu 2311102021 IF 06//
import "fmt"
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

func printSeparator() {
	fmt.Println("\n=================================================")
}

func printDashedSeparator() {
	fmt.Println("-------------------------------------------------")
}

func processDaerah(i, m int) ([]int, []int) {
	var arr = make([]int, m)
	fmt.Printf("Masukkan %d nomor rumah kerabat untuk daerah ke-%d: ", m, i+1)
	for j := 0; j < m; j++ {
		fmt.Scan(&arr[j])
	}

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

func printSortedResults(i int, odd, even []int) {

	selectionSort(odd, true) 
	selectionSort(even, false) 

	fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", i+1)

	for _, num := range odd {
		fmt.Printf("%d ", num)
	}

	for _, num := range even {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
func main() {
	var n int
	printSeparator()
	fmt.Print("Masukkan jumlah daerah kerabat (n): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		printDashedSeparator()
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d (m): ", i+1)
		fmt.Scan(&m)

		odd, even := processDaerah(i, m)

		printSortedResults(i, odd, even)
		printDashedSeparator()
	}
}
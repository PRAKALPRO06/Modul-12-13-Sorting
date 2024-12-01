package main

import "fmt"

func sortNumbers(arr []int, isAsc bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		extremeIdx := i
		for j := i + 1; j < n; j++ {
			if (isAsc && arr[j] < arr[extremeIdx]) || (!isAsc && arr[j] > arr[extremeIdx]) {
				extremeIdx = j
			}
		}
		arr[i], arr[extremeIdx] = arr[extremeIdx], arr[i]
	}
}

func collectHouseNumbers(areaIndex, numHouses int) ([]int, []int) {
	var houseNumbers = make([]int, numHouses)
	fmt.Printf("Masukkan %d nomor rumah untuk daerah ke-%d: ", numHouses, areaIndex+1)
	for i := 0; i < numHouses; i++ {
		fmt.Scan(&houseNumbers[i])
	}

	oddNumbers, evenNumbers := categorizeNumbers(houseNumbers)
	return oddNumbers, evenNumbers
}

func categorizeNumbers(houseNumbers []int) (odd []int, even []int) {
	for _, num := range houseNumbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return odd, even
}

func displaySortedResults(areaIndex int, odd, even []int) {
	sortNumbers(odd, true)
	sortNumbers(even, false)

	fmt.Printf("\nNomor rumah terurut untuk daerah ke-%d:\n", areaIndex+1)
	printArray(odd)
	printArray(even)
}

func printArray(arr []int) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	var totalAreas int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scan(&totalAreas)

	for areaIndex := 0; areaIndex < totalAreas; areaIndex++ {
		var numHouses int
		fmt.Printf("Masukkan jumlah rumah kerabat di daerah ke-%d: ", areaIndex+1)
		fmt.Scan(&numHouses)

		odd, even := collectHouseNumbers(areaIndex, numHouses)

		displaySortedResults(areaIndex, odd, even)
	}
}

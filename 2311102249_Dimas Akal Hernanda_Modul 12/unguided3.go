package main

import (
	"fmt"
)

const MaxBooks int = 7919

type Book struct {
	ID, Title, Author, Publisher string
	Copies, Year, Rating         int
}

type BookList [MaxBooks]Book

func RegisterBooks(library *BookList, count *int) {
	fmt.Print("Enter the number of books: ")
	fmt.Scanln(count)

	for i := 0; i < *count; i++ {
		fmt.Printf("Enter details for book %d:\n", i+1)
		fmt.Print("ID, Title, Author, Publisher, Copies, Year, Rating (space-separated): ")
		fmt.Scanln(&library[i].ID, &library[i].Title, &library[i].Author, &library[i].Publisher, &library[i].Copies, &library[i].Year, &library[i].Rating)
	}
}

func PrintFavoriteBook(library BookList, count int) {
	favorite := library[0]
	for i := 1; i < count; i++ {
		if library[i].Rating > favorite.Rating {
			favorite = library[i]
		}
	}
	fmt.Printf("Favorite Book: %s by %s, published by %s in %d with a rating of %d\n",
		favorite.Title, favorite.Author, favorite.Publisher, favorite.Year, favorite.Rating)
}

func SortBooksByRating(library *BookList, count int) {
	for i := 1; i < count; i++ {
		temp := library[i]
		j := i - 1
		for j >= 0 && library[j].Rating < temp.Rating {
			library[j+1] = library[j]
			j--
		}
		library[j+1] = temp
	}
}

func PrintTop5Books(library BookList, count int) {
	fmt.Println("Top 5 Books by Rating:")
	limit := 5
	if count < 5 {
		limit = count
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, library[i].Title, library[i].Rating)
	}
}

func SearchBookByRating(library BookList, count int, rating int) {
	low, high := 0, count-1
	found := false
	var idx int

	for low <= high {
		mid := (low + high) / 2
		if library[mid].Rating == rating {
			found = true
			idx = mid
			break
		} else if library[mid].Rating < rating {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		fmt.Printf("Book found: %s by %s, published by %s in %d, %d copies, Rating: %d\n",
			library[idx].Title, library[idx].Author, library[idx].Publisher, library[idx].Year, library[idx].Copies, library[idx].Rating)
	} else {
		fmt.Println("No book found with the given rating.")
	}
}

func main() {
	var library BookList
	var bookCount, searchRating int

	RegisterBooks(&library, &bookCount)
	PrintFavoriteBook(library, bookCount)
	SortBooksByRating(&library, bookCount)
	PrintTop5Books(library, bookCount)

	fmt.Print("Enter the rating of the book to search for: ")
	fmt.Scanln(&searchRating)
	SearchBookByRating(library, bookCount, searchRating)
}

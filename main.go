package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	inventory := NewInMemoryInventory()
	library := NewLibrary(inventory)

	bookReq1 := AddBookRequest{
		Title:           "Oliver Twist",
		Authors:         []string{"Charles Dickens"},
		PublicationYear: "1996",
		Genre:           "social novel",
		Price:           "INR 200",
	}

	bookReq2 := AddBookRequest{
		Title:           "System Design Vol 1",
		Authors:         []string{"Alex Xu"},
		PublicationYear: "2004",
		Genre:           "tech",
		Price:           "INR 899",
	}

	bookReq3 := AddBookRequest{
		Title:           "Designing Data Intensive Applications",
		Authors:         []string{"Martin Kleppmann"},
		PublicationYear: "2004",
		Genre:           "tech",
		Price:           "INR 1699",
	}

	bookReq4 := AddBookRequest{
		Title:           "System Design Vol 2",
		Authors:         []string{"Alex Xu", "Sahn Lam"},
		PublicationYear: "2004",
		Genre:           "tech",
		Price:           "INR 899",
	}

	req := []AddBookRequest{bookReq1, bookReq2, bookReq3, bookReq4}

	fmt.Println("Adding books to inventory:-")

	if err := library.Add(req); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Books successfully added")
	}

	allBooks := library.GetAllBooks()
	fmt.Println("Displaying all books:-")

	// Create a tab writer with padding and alignment
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// Print header
	fmt.Fprintln(w, "ID\tTitle\tAuthor\tPublication Year\tGenre\tPrice")

	// Print books
	for _, book := range allBooks {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", book.ID, book.Title, book.Authors, book.PublicationYear, book.Genre, book.Price)
	}

	// Flush the buffer
	w.Flush()

	booksByAuthor, err := library.SearchByAuthor("Alex Xu")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Printing books by author Alex Xu:")
		fmt.Fprintln(w, "ID\tTitle\tAuthor\tPublication Year\tGenre\tPrice")
		defer w.Flush()
		for _, book := range booksByAuthor {
			authors := ""
			for _, author := range book.Authors {
				authors += author + ", "
			}
			authors = authors[:len(authors)-2] // remove trailing comma and space
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", book.ID, book.Title, authors, book.PublicationYear, book.Genre, book.Price)
		}
	}

	//if books, err := library.SearchByGenre("tech"); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(books)
	//}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	inventory := NewInMemoryInventory()
	library := NewLibrary(inventory)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Library Management System!")

	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Add a book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Search books by author")
		fmt.Println("4. Search books by title")
		fmt.Println("5. Search books by genre")
		fmt.Println("6. Display all books")
		fmt.Println("7. Exit")

		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addBook(reader, library)
		case "2":
			removeBook(reader, library)
		case "3":
			searchByAuthor(reader, library)
		case "4":
			searchByTitle(reader, library)
		case "5":
			searchByGenre(reader, library)
		case "6":
			displayAllBooks(library)
		case "7":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addBook(reader *bufio.Reader, library *Library) {
	fmt.Println("Adding a new book:")
	fmt.Print("Enter title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter author: ")
	names, _ := reader.ReadString('\n')
	authors := strings.Split(names, ",")
	for i := 0; i < len(authors); i++ {
		authors[i] = strings.TrimSpace(authors[i])
	}

	fmt.Print("Enter publication year: ")
	year, _ := reader.ReadString('\n')
	year = strings.TrimSpace(year)

	fmt.Print("Enter genre: ")
	genre, _ := reader.ReadString('\n')
	genre = strings.TrimSpace(genre)

	bookReq := AddBookRequest{
		Title:           title,
		Authors:         authors,
		PublicationYear: year,
		Genre:           genre,
	}
	if err := library.Add([]AddBookRequest{bookReq}); err != nil {
		fmt.Println("Error adding book:", err)
	} else {
		fmt.Println("Book added successfully")
	}
}

func removeBook(reader *bufio.Reader, library *Library) {
	fmt.Println("Removing a book:")
	fmt.Print("Enter title of the book to remove: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	if err := library.Remove(title); err != nil {
		fmt.Println("Error removing book:", err)
	} else {
		fmt.Println("Book removed successfully")
	}
}

func searchByAuthor(reader *bufio.Reader, library *Library) {
	fmt.Println("Searching books by author:")
	fmt.Print("Enter author name: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	books, err := library.SearchByAuthor(author)
	if err != nil {
		fmt.Println("Error searching for books:", err)
		return
	}

	displayBooks(books)
}

func searchByTitle(reader *bufio.Reader, library *Library) {
	fmt.Println("Searching books by title:")
	fmt.Print("Enter title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	books, err := library.SearchByTitle(title)
	if err != nil {
		fmt.Println("Error searching for books:", err)
		return
	}

	displayBooks(books)
}

func searchByGenre(reader *bufio.Reader, library *Library) {
	fmt.Println("Searching books by genre:")
	fmt.Print("Enter genre: ")
	genre, _ := reader.ReadString('\n')
	genre = strings.TrimSpace(genre)

	books, err := library.SearchByGenre(genre)
	if err != nil {
		fmt.Println("Error searching for books:", err)
		return
	}

	displayBooks(books)
}

func displayAllBooks(library *Library) {
	fmt.Println("Displaying all books:")

	allBooks := library.GetAllBooks()
	if len(allBooks) > 0 {
		displayBooks(allBooks)
	} else {
		fmt.Println("No books in the library at the moment!")
	}
}

func displayBooks(books []Book) {
	// Create a tab writer with padding and alignment
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	// Print header
	fmt.Fprintln(w, "ID\tTitle\tAuthor\tPublication Year\tGenre\tPrice")

	// Print books
	for _, book := range books {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", book.ID, book.Title, strings.Join(book.Authors, ", "), book.PublicationYear, book.Genre, book.Price)
	}

	// Flush the buffer
	w.Flush()
}

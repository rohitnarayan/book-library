package main

import "fmt"

type Library struct {
	inventory BookInventory
}

func NewLibrary(Inventory BookInventory) *Library {
	return &Library{
		inventory: Inventory,
	}
}

func (l *Library) Add(bookReq []AddBookRequest) error {
	_, err := l.inventory.Add(bookReq)
	if err != nil {
		return fmt.Errorf("failed to add books to library, err: %+v", err)
	}

	return nil
}

func (l *Library) Remove(title string) error {
	books, err := l.inventory.Search("title", title)
	if err != nil {
		return fmt.Errorf("failed to find book by title: %s, err: %+v", title, err)
	}

	if err := l.inventory.Remove(books[0].ID); err != nil {
		return fmt.Errorf("failed to remove bookID: %s from library", books[0].ID)
	}

	return nil
}

func (l *Library) SearchByAuthor(name string) ([]Book, error) {
	books, err := l.inventory.Search("author", name)
	if err != nil {
		return []Book{}, fmt.Errorf("failed to search books by author: %s", name)
	}

	return books, nil
}

func (l *Library) SearchByGenre(name string) ([]Book, error) {
	books, err := l.inventory.Search("genre", name)
	if err != nil {
		return []Book{}, fmt.Errorf("failed to search books by genre: %s", name)
	}

	return books, nil
}

func (l *Library) SearchByTitle(name string) (Book, error) {
	books, err := l.inventory.Search("title", name)
	if err != nil {
		return Book{}, fmt.Errorf("failed to search books by title: %s", name)
	}

	return books[0], nil
}

func (l *Library) GetAllBooks() []Book {
	return l.inventory.AllBooks()
}

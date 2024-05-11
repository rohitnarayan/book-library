package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	ErrBookAlreadyPresent   = fmt.Errorf("book already present in inventory")
	ErrBookNotFound         = fmt.Errorf("book not found")
	ErrRequiredParamMissing = fmt.Errorf("required parameters missing")
	ErrInvalidAttribute     = fmt.Errorf("invalid attribute provided")
)

type BookInventory interface {
	Add(request []AddBookRequest) ([]Book, error)
	Remove(ID string) error
	Search(attribute, attributeVal string) ([]Book, error)
	AllBooks() []Book
}

type inMemoryDB struct {
	bookIDCounter int
	booksMap      map[string]Book
	authorIndex   map[string][]string
	genreIndex    map[string][]string
	titleIndex    map[string]string
	lock          sync.Mutex
}

func NewInMemoryInventory() BookInventory {
	return &inMemoryDB{
		bookIDCounter: 0,
		booksMap:      make(map[string]Book),
		authorIndex:   make(map[string][]string),
		genreIndex:    make(map[string][]string),
		titleIndex:    make(map[string]string),
	}
}

func (db *inMemoryDB) Add(requests []AddBookRequest) ([]Book, error) {
	var books []Book
	for _, req := range requests {
		book, err := db.addBook(req)
		if err != nil && err != ErrBookAlreadyPresent {
			return []Book{}, fmt.Errorf("error while adding books to inventory, err: %+v", err)
		}

		books = append(books, book)
		db.updateIndexes("add", &book)
	}

	return books, nil
}

func (db *inMemoryDB) Remove(ID string) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	book, ok := db.booksMap[ID]
	if !ok {
		return ErrBookNotFound
	}

	delete(db.booksMap, book.ID)
	db.updateIndexes("remove", &book)

	return nil
}

func (db *inMemoryDB) Search(attribute, attributeVal string) ([]Book, error) {
	if len(attribute) == 0 || len(attributeVal) == 0 {
		return []Book{}, ErrRequiredParamMissing
	}

	switch attribute {
	case "genre":
		if bookIds, ok := db.genreIndex[attributeVal]; ok {
			return db.getBookByIDs(bookIds), nil
		}
	case "title":
		if bookId, ok := db.titleIndex[attributeVal]; ok {
			return db.getBookByIDs([]string{bookId}), nil
		}
	case "author":
		if booksIds, ok := db.authorIndex[attributeVal]; ok {
			return db.getBookByIDs(booksIds), nil
		}
	case "id":
		if book, ok := db.booksMap[attributeVal]; ok {
			return []Book{book}, nil
		}
	}

	return []Book{}, ErrInvalidAttribute
}

func (db *inMemoryDB) AllBooks() []Book {
	var books []Book
	for _, book := range db.booksMap {
		books = append(books, book)
	}

	return books
}

func (db *inMemoryDB) getBookByIDs(ids []string) []Book {
	var books []Book
	for _, id := range ids {
		if book, ok := db.booksMap[id]; ok {
			books = append(books, book)
		}
	}

	return books
}

func (db *inMemoryDB) addBook(req AddBookRequest) (Book, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	if _, ok := db.booksMap[req.Title]; ok {
		return Book{}, ErrBookAlreadyPresent
	}

	db.bookIDCounter++

	book := Book{
		ID:              strconv.Itoa(db.bookIDCounter),
		Title:           req.Title,
		Genre:           req.Genre,
		Authors:         req.Authors,
		Price:           req.Price,
		PublicationYear: req.PublicationYear,
		AddedAt:         time.Now().UTC(),
	}

	db.booksMap[book.ID] = book

	return book, nil
}

func (db *inMemoryDB) updateIndexes(operation string, book *Book) {
	db.lock.Lock()
	defer db.lock.Unlock()

	switch operation {
	case "add":
		db.addToIndex(book)
	case "remove":
		db.removeFromIndex(book)
	}
}

func (db *inMemoryDB) addToIndex(book *Book) {
	// update genre index
	if books, ok := db.genreIndex[book.Genre]; !ok {
		db.genreIndex[book.Genre] = []string{book.ID}
	} else if ok {
		books = append(books, book.ID)
		db.genreIndex[book.Genre] = books
	}

	// update title index
	if _, ok := db.titleIndex[book.Title]; !ok {
		db.titleIndex[book.Title] = book.ID
	}

	// update author index
	for _, author := range book.Authors {
		if books, ok := db.authorIndex[author]; !ok {
			db.authorIndex[author] = []string{book.ID}
		} else if ok {
			books = append(books, book.ID)
			db.authorIndex[author] = books
		}
	}
}

func (db *inMemoryDB) removeFromIndex(book *Book) {
	// update only the title index as there can be
	// other books of the same author and
	// other books with the same genre

	if _, ok := db.titleIndex[book.Title]; ok {
		delete(db.titleIndex, book.Title)
	}
}

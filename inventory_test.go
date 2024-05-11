package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBooks(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	books, err := db.Add(requests)
	assert.NoError(t, err)

	expectedBooks := []Book{
		{ID: "1", Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{ID: "2", Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	assert.Equal(t, expectedBooks[0].Authors, books[0].Authors)
	assert.Equal(t, expectedBooks[0].Title, books[0].Title)
}

func TestRemoveBooks(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	books, err := db.Add(requests)
	assert.NoError(t, err)

	err = db.Remove(books[0].ID)
	assert.NoError(t, err)
}

func TestSearchByGenre(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	_, err := db.Add(requests)
	assert.NoError(t, err)

	res, err := db.Search("genre", "Fiction")
	assert.NoError(t, err)

	assert.Equal(t, "Book1", res[0].Title)
}

func TestSearchByAuthor(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	_, err := db.Add(requests)
	assert.NoError(t, err)

	res, err := db.Search("author", "Author2")
	assert.NoError(t, err)

	assert.Equal(t, "Book1", res[0].Title)
	assert.Equal(t, "Book2", res[1].Title)
}

func TestSearchByTitle(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	_, err := db.Add(requests)
	assert.NoError(t, err)

	res, err := db.Search("title", "Book1")
	assert.NoError(t, err)

	assert.Equal(t, "Book1", res[0].Title)
}

func TestSearchByID(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	_, err := db.Add(requests)
	assert.NoError(t, err)

	res, err := db.Search("id", "1")
	assert.NoError(t, err)

	assert.Equal(t, "Book1", res[0].Title)
}

func TestGetAllBooks(t *testing.T) {
	db := NewInMemoryInventory()

	// Test adding multiple books
	requests := []AddBookRequest{
		{Title: "Book1", Genre: "Fiction", Authors: []string{"Author1", "Author2"}, Price: "101", PublicationYear: "2021"},
		{Title: "Book2", Genre: "Non-Fiction", Authors: []string{"Author2"}, Price: "1299", PublicationYear: "2020"},
	}

	_, err := db.Add(requests)
	assert.NoError(t, err)

	books := db.AllBooks()
	assert.Equal(t, "Book1", books[0].Title)
	assert.Equal(t, "Book2", books[1].Title)
}

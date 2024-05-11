package main

import "time"

type AddBookRequest struct {
	Title           string
	Authors         []string
	Price           string
	PublicationYear string
	Genre           string
}

type Book struct {
	ID              string
	Title           string
	Authors         []string
	Price           string
	PublicationYear string
	Genre           string
	Available       bool
	AddedAt         time.Time
}

package service

import (
	"books/internal/pkg/model"
	"books/internal/pkg/repo"
	"errors"
	"fmt"
)

type BooksService struct {
	booksRepo *repo.BookRepository
}

func GetNewBooksService(booksRepo *repo.BookRepository) BooksService {
	return BooksService{booksRepo: booksRepo}
}

func (bs *BooksService) AddBook(book *model.DBBook) {
	bs.booksRepo.AddBook(book)
}

func (bs *BooksService) GetBook(isbn int) (*model.DBBook, error) {
	book := bs.booksRepo.GetBook(isbn)
	if book != nil {
		return book, nil
	}
	return nil, errors.New(fmt.Sprintf("book with isbn %d was not found", isbn))
}

func (bs *BooksService) GetAllBooks() ([]*model.DBBook, error) {
	books, err := bs.booksRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, errors.New("no books found")
	}
	return books, nil
}

func (bs *BooksService) RemoveBook(isbn int) {
	bs.booksRepo.RemoveBook(isbn)
}

func (bs *BooksService) UpdateBook(book *model.DBBook) {
	bs.booksRepo.UpdateBook(book)
}

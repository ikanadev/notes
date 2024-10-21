package grpcbooksserver

import (
	"books/internal/pkg/model"
	"books/internal/pkg/proto"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// AddBook implements proto.BookServiceServer.
func (a *App) AddBook(ctx context.Context, req *proto.Book) (*proto.AddBookResponse, error) {
	log.Println("adding book")

	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}

	a.bookRepo.AddBook(book)
	return &proto.AddBookResponse{
		Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher(%s) added successfully", book.Isbn, book.Name, book.Publisher),
	}, nil
}

// GetBook implements proto.BookServiceServer.
func (a *App) GetBook(ctx context.Context, req *proto.GetBookRequest) (*proto.Book, error) {
	log.Println("fetching book")

	book := a.bookRepo.GetBook(int(req.Isbn))

	return &proto.Book{
		Isbn:      int32(book.Isbn),
		Name:      book.Name,
		Publisher: book.Publisher,
	}, nil
}

// ListBooks implements proto.BookServiceServer.
func (a *App) ListBooks(ctx context.Context, req *proto.Empty) (*proto.ListBooksResponse, error) {
	log.Println("listing book")

	books, err := a.bookRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(books)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal books: %s\n", err.Error())
	}

	pbBooks := []*proto.Book{}
	err = json.Unmarshal(b, &pbBooks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal books: %s\n", err.Error())
	}
	return &proto.ListBooksResponse{Books: pbBooks}, nil
}

// RemoveBook implements proto.BookServiceServer.
func (a *App) RemoveBook(ctx context.Context, req *proto.RemoveBookRequest) (*proto.RemoveBookResponse, error) {
	log.Println("removing book")
	a.bookRepo.RemoveBook(int(req.Isbn))
	return &proto.RemoveBookResponse{
		Status: fmt.Sprintf("book with isbn(%d) removed successfully", req.Isbn),
	}, nil
}

// UpdateBook implements proto.BookServiceServer.
func (a *App) UpdateBook(ctx context.Context, req *proto.Book) (*proto.UpdateBookResponse, error) {
	log.Println("updating book book")
	book := &model.DBBook{
		Isbn:      int(req.Isbn),
		Name:      req.Name,
		Publisher: req.Publisher,
	}
	a.bookRepo.UpdateBook(book)
	return &proto.UpdateBookResponse{
		Status: fmt.Sprintf("book with isbn(%d), name(%s), publisher(%s) updated successfully", book.Isbn, book.Name, book.Publisher),
	}, nil
}

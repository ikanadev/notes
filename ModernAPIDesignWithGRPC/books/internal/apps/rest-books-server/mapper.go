package restbooksserver

import "books/internal/pkg/model"

func DBBook(book *model.Book) *model.DBBook {
	return &model.DBBook{
		Isbn:      book.Isbn,
		Name:      book.Name,
		Publisher: book.Publisher,
	}
}

func Book(dbbook *model.DBBook) *model.Book {
	return &model.Book{
		Isbn:      dbbook.Isbn,
		Name:      dbbook.Name,
		Publisher: dbbook.Publisher,
	}
}

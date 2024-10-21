package grpcbooksclient

import (
	"books/internal/pkg/configs"
	"books/internal/pkg/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	serverConn *grpc.ClientConn
	client     proto.BookServiceClient
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	appConfig, err := configs.ProvideAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	serverAddr := appConfig.ClientConfig.ServerAddress

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())

	a.serverConn, err = grpc.NewClient(serverAddr, opts)
	if err != nil {
		log.Fatal(err)
	}

	a.client = proto.NewBookServiceClient(a.serverConn)
	a.performClientOperations()
}

func (a *App) Shutdown() {
	a.serverConn.Close()
}

func (a *App) performClientOperations() {
	book := &proto.Book{
		Isbn:      12348,
		Name:      "Atomic Habits",
		Publisher: "random house business books",
	}
	a.AddBook(book)

	found := false
	books := a.ListBooks()

	for i, b := range books {
		fmt.Printf("[%d.] Nmae(%s), Isbn(%d), Publisher(%s)\n", i, b.Name, b.Isbn, b.Publisher)
		if b.Isbn == book.Isbn && b.Name == book.Name && b.Publisher == book.Publisher {
			found = true
		}
	}

	if found {
		fmt.Println("Book sent using 'AddBook' request was verified to have been added while listing books.")
	}

	fetchedBook := a.FetchBook(12345)
	if fetchedBook.Isbn == 12345 {
		fmt.Println("Book added via migrations was successfully fetched.")
	}

	book.Name = "atomic habits vol-2"
	a.UpdateBook(book)
	a.RemoveBook(int(book.Isbn))
}

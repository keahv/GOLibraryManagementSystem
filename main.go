package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

func (b Book) Initialize(title, author, isbn string, available bool) Book {
	return Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: available,
	}
}

func (b Book) DisplayDetails() string {
	status := "Unavailable"
	if b.Available {
		status = "Available"
	}
	return fmt.Sprintf("Title: %s, Author: %s, ISBN: %s, Status: %s", b.Title, b.Author, b.ISBN, status)
}

type EBook struct {
	Book
	FileSize int
}

func (e EBook) DisplayDetails() string {
	return fmt.Sprintf("%s, File Size: %d MB", e.Book.DisplayDetails(), e.FileSize)
}

type BookInterface interface {
	DisplayDetails() string
}

type Library struct {
	Collection []BookInterface
}

func (l Library) AddBook(book BookInterface) Library {
	l.Collection = append(l.Collection, book)
	return l
}

func (l Library) RemoveBook(isbn string) (Library, error) {
	for i, item := range l.Collection {
		if book, ok := item.(Book); ok && book.ISBN == isbn {
			l.Collection = append(l.Collection[:i], l.Collection[i+1:]...)
			return l, nil
		} else if ebook, ok := item.(EBook); ok && ebook.ISBN == isbn {
			l.Collection = append(l.Collection[:i], l.Collection[i+1:]...)
			return l, nil
		}
	}
	return l, fmt.Errorf("book with ISBN %s not found", isbn)
}

func (l Library) SearchBookByTitle(title string) ([]BookInterface, error) {
	var results []BookInterface
	for _, item := range l.Collection {
		switch book := item.(type) {
		case Book:
			if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
				results = append(results, book)
			}
		case EBook:
			if strings.Contains(strings.ToLower(book.Title), strings.ToLower(title)) {
				results = append(results, book)
			}
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no books found with title containing '%s'", title)
	}
	return results, nil
}

func (l Library) ListBooks() {
	if len(l.Collection) == 0 {
		fmt.Println("No books in the library.")
		return
	}
	for _, item := range l.Collection {
		fmt.Println(item.DisplayDetails())
	}
}

func main() {
	library := Library{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Add an EBook")
		fmt.Println("3. Remove a Book/EBook")
		fmt.Println("4. Search for Books by Title")
		fmt.Println("5. List all Books/EBooks")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			fmt.Print("Enter ISBN: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)
			fmt.Print("Is it available (true/false): ")
			availableInput, _ := reader.ReadString('\n')
			availableInput = strings.TrimSpace(availableInput)
			available := strings.ToLower(availableInput) == "true"
			book := Book{}.Initialize(title, author, isbn, available)
			library = library.AddBook(book)
			fmt.Println("Book added successfully!")

		case "2":
			fmt.Print("Enter Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			fmt.Print("Enter Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			fmt.Print("Enter ISBN: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)
			fmt.Print("Enter File Size (MB): ")
			fileSizeInput, _ := reader.ReadString('\n')
			fileSizeInput = strings.TrimSpace(fileSizeInput)
			fileSize := 0
			fmt.Sscanf(fileSizeInput, "%d", &fileSize)
			fmt.Print("Is it available (true/false): ")
			availableInput, _ := reader.ReadString('\n')
			availableInput = strings.TrimSpace(availableInput)
			available := strings.ToLower(availableInput) == "true"
			ebook := EBook{Book: Book{}.Initialize(title, author, isbn, available), FileSize: fileSize}
			library = library.AddBook(ebook)
			fmt.Println("EBook added successfully!")

		case "3":
			fmt.Print("Enter ISBN to remove: ")
			isbn, _ := reader.ReadString('\n')
			isbn = strings.TrimSpace(isbn)
			var err error
			library, err = library.RemoveBook(isbn)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Book removed successfully!")
			}

		case "4":
			fmt.Print("Enter title to search: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			results, err := library.SearchBookByTitle(title)
			if err != nil {
				fmt.Println(err)
			} else {
				for _, item := range results {
					fmt.Println(item.DisplayDetails())
				}
			}

		case "5":
			library.ListBooks()

		case "6":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

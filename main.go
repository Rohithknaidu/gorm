package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"fmt"
)

// Define Author and Book structs
type Author struct {
	gorm.Model
	Name  string
	Books []Book // One-to-Many relationship
}

type Book struct {
	gorm.Model
	Title     string
	AuthorID  uint // Foreign key
	Author    Author `gorm:"foreignKey:AuthorID"` // Belongs To relationship
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Author{}, &Book{})

	// Create an author and some books
	author := Author{Name: "John Doe"}
	book1 := Book{Title: "Golang Basics", Author: author}
	book2 := Book{Title: "Golang Advanced", Author: author}

	// Save author and books to the database
	db.Create(&author)
	db.Create(&book1)
	db.Create(&book2)

	// Query the author along with their books
	var queriedAuthor Author
	db.Preload("Books").First(&queriedAuthor, "name = ?", "John Doe")

	// Print the author and their books
	fmt.Printf("Author: %s\n", queriedAuthor.Name)
	fmt.Println("Books:")
	for _, book := range queriedAuthor.Books {
		fmt.Printf("- %s\n", book.Title)
	}
}

package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "For the Love of Go",
		Author: "Alan A.A. Donovan",
		Copies: 99,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "Alan A.A. Donovan",
		Copies: 2,
	}

	var want int = 1
	result, err := bookstore.Buy(book)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if got != want {
		t.Errorf("Copies = %d; want %d", got, want)
	}
}

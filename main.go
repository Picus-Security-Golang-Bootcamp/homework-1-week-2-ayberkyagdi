package main

import (
	"fmt"
	"os"
	"strings"
)

type Books struct {
	Items []BookItem
}

type BookItem struct {
	Name, Author, Date string
}

func (b *Books) AddItem(item BookItem) []BookItem {
	b.Items = append(b.Items, item)
	return b.Items
}

func (b *Books) FindItem(s string) BookItem {
	item := BookItem{}
	for i := range b.Items {
		if strings.ToLower(b.Items[i].Name) == s {
			item = b.Items[i]
			return item
		}
	}
	return item
}

func (b *Books) ListItems() {
	column_length := 50
	empty_string := strings.Repeat("\t", 6)
	fmt.Printf("\tNAME%sAUTHOR%sDATE\n", empty_string, empty_string)
	for index, book := range b.Items {
		fmt.Printf("%d %s", index+1, book.Name)
		fmt.Printf("%s", strings.Repeat(" ", column_length-len(book.Name)))
		fmt.Printf("%s", book.Author)
		fmt.Printf("%s", strings.Repeat(" ", column_length-len(book.Author)))
		fmt.Printf("%s\n", book.Date)

	}
}
func main() {
	arg := os.Args[1:]

	to_string := strings.Join(arg, " ")
	to_string = strings.ToLower(to_string)

	book_list := map[string][]string{
		"The History of Tom Jones": {"Henry Fielding", "28 February 1749"},
		"Pride and Prejudice":      {"Jane Austen", "28 January 1813"},
		"The Red and the Black":    {"Stendhal", "22 November 1830"},
		"Le Pere Goriot":           {"Honore de Balzac", "8 March 1835 "},
		"David Copperfield":        {"Charles Dickens", "6 November 1850"},
		"Madame Bovary":            {"Gustave Flaubert", "5 April 1857"},
		"Moby-Dick":                {"Herman Melville", "18 October 1851"},
		"Wuthering Heights":        {"Emily Bronte", "3 December 1847"},
		"The Brothers Karamazov":   {"Dostoevsky", "10 January 1879"},
		"War and Peace":            {"Tolstoy", "19 June 1869"},
	}
	book_struct := new(Books)
	for k, v := range book_list {
		book := new(BookItem)
		book.Name = k
		book.Author = v[0]
		book.Date = v[1]
		book_struct.AddItem(*book)

	}
	if to_string == "list" {
		book_struct.ListItems()
	} else if strings.HasPrefix(to_string, "search ") {
		exp_book := to_string[len("search "):]
		result := book_struct.FindItem(exp_book)
		if !(result == BookItem{}) {
			fmt.Printf("\nThe book has found in the library !\n")
			fmt.Printf("Name:%s\nAuthor:%s\nDate:%s", result.Name, result.Author, result.Date)
		} else {
			fmt.Printf("The book was not found in the library !")
		}
	} else {
		fmt.Printf("Invalid input! You must enter 'list' or 'search <bookName>'!")
	}
}

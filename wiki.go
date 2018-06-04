package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

// Save method
/*
"This is a method named save that takes as its receiver p, a pointer to Page.
It takes no parameters, and returns a value of type error." It saves the page's
body to a file with the page's title.
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Load Method
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// main
func main() {
	p1 := &Page{Title: "Test Page", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("Test Page")
	fmt.Println(string(p2.Body))
}

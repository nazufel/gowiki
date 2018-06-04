package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

// HANDLERS //

// View Handler
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

}

// main
func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))

}

/*
	p1 := &Page{Title: "Test Page", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("Test Page")
	fmt.Println(string(p2.Body))
*/

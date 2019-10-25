package structs

import (
	io "io/ioutil"
)


//Page struct that is a webpage
type Page struct {
	Title string
	Body  []byte
}

// Save a page to the disk
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return io.WriteFile(filename, p.Body, 0600)
}

// LoadPage from disk
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := io.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	page := Page{
		Title: title,
		Body: body}
	return &page, nil
}
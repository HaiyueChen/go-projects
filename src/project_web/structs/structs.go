package structs

import (
	io "io/ioutil"
	"project_web/utils"
)

//Page struct that is a webpage
type Page struct {
	Title string
	Body  []byte
}

// Save a page to the disk
func (p *Page) Save(filename string, path string) error {

	return io.WriteFile(path, p.Body, 0600)
}

// LoadPage from disk
func LoadPage(path string) (*Page, error) {

	body, err := io.ReadFile(path)
	if err != nil {
		return nil, err
	}
	filename := utils.ParseFilenameFromPath(path)
	page := Page{
		Title: filename,
		Body:  body}
	return &page, nil
}

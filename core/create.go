package core

import (
	"google.golang.org/api/docs/v1"
)

// create blank document
func CreateBlankDoc(srv *docs.Service, title string) (*docs.Document, error) {
	data := srv.Documents.Create(&docs.Document{
		Title: title,
	})
	return data.Do()
}

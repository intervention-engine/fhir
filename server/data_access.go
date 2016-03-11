package server

import (
	"net/url"

	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
)

type DataAccessLayer interface {
	Get(id, resourceType string) (interface{}, error)
	Post(resource interface{}) (id string, err error)
	PostWithId(id string, resource interface{}) error
	Put(id string, resource interface{}) error
	Delete(id, resourceType string) error
	Search(baseURL url.URL, searchQuery search.Query) (*models.Bundle, error)
}

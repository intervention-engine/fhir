package server

import (
	"errors"
	"net/url"

	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
)

// DataAccessLayer is an interface for the various interactions that can occur on a FHIR data store.
type DataAccessLayer interface {
	Get(id, resourceType string) (result interface{}, err error)
	Post(resource interface{}) (id string, err error)
	PostWithID(id string, resource interface{}) error
	Put(id string, resource interface{}) (createdNew bool, err error)
	ConditionalPut(query search.Query, resource interface{}) (id string, createdNew bool, err error)
	Delete(id, resourceType string) error
	ConditionalDelete(query search.Query) (count int, err error)
	Search(baseURL url.URL, searchQuery search.Query) (result *models.Bundle, err error)
}

// ErrNotFound indicates an error
var ErrNotFound = errors.New("Resource Not Found")

// ErrMultipleMatches indicates that the conditional update query returned multiple matches
var ErrMultipleMatches = errors.New("Multiple Matches")

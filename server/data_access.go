package server

import (
	"errors"
	"net/url"

	"github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
)

// DataAccessLayer is an interface for the various interactions that can occur on a FHIR data store.
type DataAccessLayer interface {
	// Get retrieves a single resource instance identified by its resource type and ID
	Get(id, resourceType string) (result interface{}, err error)
	// Post creates a resource instance, returning its new ID.
	Post(resource interface{}) (id string, err error)
	// PostWithID creates a resource instance with the given ID.
	PostWithID(id string, resource interface{}) error
	// Put creates or updates a resource instance with the given ID.
	Put(id string, resource interface{}) (createdNew bool, err error)
	// ConditionalPut creates or updates a resource based on search criteria.  If the criteria results in zero matches,
	// the resource is created.  If the criteria results in one match, it is updated.  Otherwise, a ErrMultipleMatches
	// error is returned.
	ConditionalPut(query search.Query, resource interface{}) (id string, createdNew bool, err error)
	// ConditionalPutPreflight performs the ConditionalPut logic but does not modify the resource instance or the
	// database.  It is used to determine what would happen if a ConditionalPut were executed.  Note that when
	// ConditionalPutPreflight indicates that a resource instance will be createdNew, the returned ID is a
	// randomly created ID and will not match the actual new ID if a ConditionalPut is performed.
	ConditionalPutPreflight(query search.Query, resource interface{}) (id string, createdNew bool, err error)
	// Delete removes the resource instance with the given ID.  This operation cannot be undone.
	Delete(id, resourceType string) error
	// ConditionalDelete removes zero or more resources matching the passed in search criteria.  This operation cannot
	// be undone.
	ConditionalDelete(query search.Query) (count int, err error)
	// Search executes a search given the baseURL and searchQuery.
	Search(baseURL url.URL, searchQuery search.Query) (result *models.Bundle, err error)
}

// ErrNotFound indicates an error
var ErrNotFound = errors.New("Resource Not Found")

// ErrMultipleMatches indicates that the conditional update query returned multiple matches
var ErrMultipleMatches = errors.New("Multiple Matches")

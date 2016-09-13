package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

// ConfigureIndexes ensures that all indexes listed in the provided indexes.conf file
// are part of the Mongodb fhir database. If an index does not exist yet ConfigureIndexes
// creates a new index in the background using mgo.collection.EnsureIndex(). Depending
// on the size of the collection it may take some time before the index is created.
// This will block the current thread until the indexing completes, but will not block
// other connections to the mongo database.
func ConfigureIndexes(session *mgo.Session, config Config) {
	var err error

	db := session.DB(config.DatabaseName)

	// Read the config file
	idxConfig, err := ioutil.ReadFile(config.IndexConfigPath)
	if err != nil {
		log.Fatal("Could not find indexes configuration file")
	}

	// parse the config file
	indexMap, err := parseIndexes(string(idxConfig), config.DatabaseName)

	if err != nil {
		log.Fatal(err.Error())
	}

	// ensure all indexes in the config file
	for k := range indexMap {
		collection := db.C(k)

		for _, index := range indexMap[k] {
			log.Printf("Ensuring index: %s.%s: %s\n", config.DatabaseName, k, index.Key[0])
			err = collection.EnsureIndex(index)

			if err != nil {
				log.Printf("[WARNING] Could not ensure index: %s.%s: %s\n", config.DatabaseName, k, index.Key[0])
			}
		}
	}
}

func parseIndexes(fileContents string, configuredDBName string) (map[string][]mgo.Index, error) {
	var indexMap = make(map[string][]mgo.Index)
	lines := strings.Split(fileContents, "\n")
	for _, line := range lines {

		if line == "" {
			continue
		}

		if string(line[0]) == "#" {
			continue
		}

		// create a new index from the next line in the config file
		// format: <db_name>.<collection_name>.<key>_<(-)1>
		config := strings.SplitN(line, ".", 3)

		if len(config) < 3 {
			// malformed index name
			return nil, newParseIndexError(line, "Not of format <db_name>.<collection_name>.<key>_<(-)1>")
		}

		dbName := config[0]

		if dbName != configuredDBName {
			// malformed index name
			return nil, newParseIndexError(line, "DB name does not match server config")
		}

		collectionName := config[1]
		rest := strings.Split(config[2], "_")

		if len(rest) != 2 {
			// malformed index name
			return nil, newParseIndexError(line, "Missing field name or index order (asc/desc)")
		}

		key := rest[0]

		// direction determines ascending or descending order
		direction := ""
		if rest[1] == "-1" {
			direction = "-"
		}

		indexMap[collectionName] = append(indexMap[collectionName], mgo.Index{
			Key:        []string{fmt.Sprintf("%s%s", direction, key)},
			Background: true, // build the index in the background; do not block other connections
		})
	}
	return indexMap, nil
}

func newParseIndexError(indexName, reason string) error {
	return fmt.Errorf("Index name '%s' is invalid: %s", indexName, reason)
}

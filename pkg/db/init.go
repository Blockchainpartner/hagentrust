
package db

// Init initialize connection to the DB, create collection objects and pushes to the B first needed scheduled jobs if
// not found
func Init() {
	InitClient()
	InitCollections()
}
package db

// Table defines a database and collection pair.
type Table struct {
	Database   string
	Collection string
}

// NewTable returns a new table.
func NewTable(database, collection string) Table {
	return Table{
		Database:   database,
		Collection: collection,
	}
}

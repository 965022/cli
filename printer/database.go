package printer

import (
	"time"

	ps "github.com/planetscale/planetscale-go"
)

// Database returns a table-serializable database model.
type Database struct {
	Name      string `header:"name"`
	Notes     string `header:"notes"`
	CreatedAt int64  `header:"created_at,timestamp(ms|utc|human)"`
	UpdatedAt int64  `header:"updated_at,timestamp(ms|utc|human)"`
}

// NewDatabasePrinter returns a struct that prints out the various fields of a
// database model.
func NewDatabasePrinter(db *ps.Database) *Database {
	return &Database{
		Name:      db.Name,
		Notes:     db.Notes,
		CreatedAt: db.CreatedAt.UTC().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
		UpdatedAt: db.UpdatedAt.UTC().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond)),
	}
}

// NewDatabaseSlicePrinter returns a slice of printable databases.
func NewDatabaseSlicePrinter(databases []*ps.Database) []*Database {
	dbs := make([]*Database, 0, len(databases))

	for _, db := range databases {
		dbs = append(dbs, NewDatabasePrinter(db))
	}

	return dbs
}
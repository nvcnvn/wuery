package translator

import (
	"database/sql"

	"github.com/lib/pq"
)

var _ = pq.BoolArray{}

// CockRoachTranslate implement bla bla bla
type CockRoachTranslate struct {
	db *sql.DB
}

// Translate recieves bla bla bla
func (t *CockRoachTranslate) Translate(rows *sql.Rows) []byte {
	return nil
}

package translator

import (
	"database/sql"

	"github.com/Jeffail/gabs"
	"github.com/lib/pq"
)

var _ = pq.BoolArray{}

// CockRoachTranslate implement bla bla bla
type CockRoachTranslate struct {
	db *sql.DB
}

// Translate recieves bla bla bla
func (t *CockRoachTranslate) Translate(rows *sql.Rows) []byte {
	jsonResult := gabs.New()
	jsonResult.Array("results")

	for i := 0; i < 3; i++ {
		jsonObj := gabs.New()
		jsonObj.Set(i, "value")
		jsonObj.Set("bla bla", "bla")
		jsonResult.ArrayAppend(jsonObj, "results")
	}
	return jsonResult.Bytes()
}

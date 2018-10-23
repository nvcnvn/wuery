package wuery

import (
	"context"
	"database/sql"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
	"github.com/cockroachdb/cockroach/pkg/sql/pgwire/pgerror"
	"github.com/cockroachdb/cockroach/pkg/sql/sem/tree"
)

// Wuery can validate the SQL statement
type Wuery struct {
	db     dbQueryInterface
	parser *parser.Parser
}

// NewWuery retuns new Wuery
func NewWuery(db *sql.DB) *Wuery {
	return &Wuery{
		db:     db,
		parser: &parser.Parser{},
	}
}

type dbQueryInterface interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
}

func (w *Wuery) validate(sql string) error {
	stmts, err := w.parser.Parse(sql)
	if err != nil {
		return err
	}

	if len(stmts) != 1 {
		return pgerror.NewAssertionErrorf("expected 1 statement, but found %d", len(stmts))
	}

	if stmts[0].StatementType() != tree.Rows {
		return pgerror.NewAssertionErrorf("expected the statement returns the affected rows")
	}

	println(stmts[0].String())
	return nil
}

// Query actually send the query to DB
func (w *Wuery) Query(ctx context.Context, query string) ([]byte, error) {
	err := w.validate(query)
	if err != nil {
		return nil, err
	}

	_, err = w.db.QueryContext(ctx, query)
	return nil, err
}

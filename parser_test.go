package sipdb

import (
	"testing"
	"github.com/lfittl/pg_query_go"
)

func TestSelect(t *testing.T) {
	tree, err := pg_query.Parse("select 1")
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)
}

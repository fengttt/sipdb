package sipdb_tests

import (
	"fmt"
	"io/ioutil"
	"testing"
	"github.com/lfittl/pg_query_go"
)

func parseQ(fn string) (string, error) {
	b, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", err
	}
	qs := string(b)
	return pg_query.ParseToJSON(qs)
}

func TestParse(t *testing.T) {
	for i:=1; i<=22; i++ {
		fn := fmt.Sprintf("./tpch/q%d.sql", i)
		tree, err := parseQ(fn)
		if err != nil {
			t.Error(err)
		}

		t.Log(" --- TPCH Query ", i, " --- ")
		t.Log(tree)
	}
}


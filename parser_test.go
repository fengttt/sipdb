package sipdb

import (
	"testing"
	"github.com/lfittl/pg_query_go"
)

func TestSelect(t *testing.T) {
	tree, err := pg_query.Parse("select * from t where i = $d")
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)

	js, err := pg_query.ParseToJSON("select t.id, x.id from t, t.p as p, p.x as x")
	if err != nil {
		t.Error(err)
	}
	t.Log(js);

}

func TestCreateTable(t *testing.T) {
	tree, err := pg_query.Parse("create table t (i int, j int) with (url = 'LocateMe!')")
	if err != nil {
		t.Error(err)
	}
	t.Log(tree)

	js, err := pg_query.ParseToJSON("create table t (i int, j int) with (url = 'LocateMe!')")
	if err != nil {
		t.Error(err)
	}
	t.Log(js);
}



package conn

import (
	"sql"

	"sipdb/rowset"
	"sipdb/errors"
)

type Conn struct {
	BindRels map[string]rowset.Rowset
	BindFunc map[string]interface{}
}

func (c *Conn) BindRowset(name string, rs rowset.Rowset) {
	_, ok = c.BindRels[name]
	errors.PanicIf(ok, errors.AlreadyExists, "Relation %s has already been bound", name)

	c.BindRels[name] = rs
}

func (c *Conn) UnbindRowset(name string) {
	_, ok = c.BindRels[name]
	errors.PanicIf(!ok, errors.NotFound, "Relation %s is not bound", name)
	delete(c.BindRels, name)
}

func (c *Conn) RebindRowset(name string) {
	rs, ok = c.BindRels[name]
	errors.PanicIf(!ok, errors.NotFound, "Relation %s is not bound", name)
	rs.Rebind()
}

func (c *Conn) BindFunc(name string, f interface{}) {
	c.BindFunc[name] = f
}

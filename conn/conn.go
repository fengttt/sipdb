package conn

import (
	"github.com/fengttt/sipdb/errors"
	"github.com/fengttt/sipdb/rowset"
)

type Conn struct {
	bindRels map[string]rowset.Rowset
	bindFunc map[string]interface{}
}

func (c *Conn) BindRowset(name string, rs rowset.Rowset) {
	_, ok := c.bindRels[name]
	errors.PanicIf(ok, errors.AlreadyExists, "Relation %s has already been bound", name)

	c.bindRels[name] = rs
}

func (c *Conn) UnbindRowset(name string) {
	_, ok := c.bindRels[name]
	errors.PanicIf(!ok, errors.NotFound, "Relation %s is not bound", name)
	delete(c.bindRels, name)
}

func (c *Conn) RebindRowset(name string) {
	rs, ok := c.bindRels[name]
	errors.PanicIf(!ok, errors.NotFound, "Relation %s is not bound", name)
	rs.Rebind()
}

func (c *Conn) BindFunct(name string, f interface{}) {
	c.bindFunc[name] = f
}

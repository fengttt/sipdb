package rowset

import (
	"fmt"

	"github.com/fengttt/sipdb/types"
)

type sliceRowset struct {
	s types.Schema
	cols []ColVec
}

type sliceRowsetCursor struct {
	rs *sliceRowset
	batch *Batch
}

func (c *sliceRowsetCursor) Rewind() {
	c.batch = &Batch { c.rs.cols }
}

func (c *sliceRowsetCursor) Next() (b *Batch) {
	b = c.batch
	c.batch = nil
	return
}

func (rs *sliceRowset) Schema() *types.Schema {
	return &rs.s
}

func (rs *sliceRowset) Rebind() {
}

func (rs *sliceRowset) Cursor() RowsetCursor {
	c := &sliceRowsetCursor { rs, nil }
	c.Rewind()
	return c
}

func FromSlice(ss []interface{}) Rowset {
	var rs sliceRowset
	for i, s := range ss {
		rs.s.AddCol(types.TypeOfSlice(s), fmt.Sprintf("_%d", i))
		rs.cols = append(rs.cols, ColVec{nil, s})
	}
	return &rs
}

func FromOneSlice(s interface{}) Rowset {
	var rs sliceRowset
	rs.s.AddCol(types.TypeOfSlice(s), "_0")
	rs.cols = append(rs.cols, ColVec{nil, s})
	return &rs
}

package rowset

import (
	"reflect"
	"github.com/fengttt/sipdb/types"
)

type ColVec struct {
	Nulls []types.NullInfo
	i interface{}
}

func (c *ColVec) B() []bool {
	return c.i.([]bool)
}

func (c *ColVec) I8() []int8 {
	return c.i.([]int8)
}

func (c *ColVec) I16() []int16 {
	return c.i.([]int16)
}

func (c *ColVec) I32() []int32 {
	return c.i.([]int32)
}

func (c *ColVec) I64() []int64 {
	return c.i.([]int64)
}

func (c *ColVec) F32() []float32 {
	return c.i.([]float32)
}

func (c *ColVec) F64() []float64 {
	return c.i.([]float64)
}

func (c *ColVec) Str() []string {
	return c.i.([]string)
}

func (c *ColVec) Bytes() [][]byte {
	return c.i.([][]byte)
}

func (c *ColVec) Object() []interface{} {
	return c.i.([]interface{})
}

func (c *ColVec) Sz() int {
	v := reflect.ValueOf(c.i)
	return v.Len()
}

type Batch struct {
	Cols []ColVec
}

// cols should not be jagged.   We pick the first col.
func (b *Batch) Sz() int {
	return b.Cols[0].Sz()
}

type RowsetCursor interface {
	Next() *Batch
	Rewind()
}

type Rowset interface {
	Schema() *types.Schema
	Rebind()
	Cursor() RowsetCursor
}


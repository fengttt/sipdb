package rowset

import (
	"sipdb/types"
)

type ColVec struct {
	Nulls []types.NullInfo
	B []bool
	I8 []int8
	I16 []int16
	I32 []int32
	I64 []int64
	F32 []float32
	F64 []float64
	S []string
	Bytes [][]byte
	Obj []interface{}
}

type Batch struct {
	Sz int
	Cols []ColVec
}

type Rowset interface {
	Schema() types.Schema
	Next() *Batch
}





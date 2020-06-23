package types

type Type int16

const (
	ANY		Type	= iota
	BOOL
	INT8
	INT16
	INT32
	INT64
	FLOAT32
	FLOAT64
	DATE
	TIME
	TIMESTAMP
	STRING
	BYTES
	JSON
	OBJECT
	ARRAY
	BAG
)

func (t Type) String() string {
	switch t {
	case ANY:
		return "any"
	case BOOL:
		return "bool"
	case INT8:
		return "int8"
	case INT16:
		return "int16"
	case INT32:
		return "int32"
	case INT64:
		return "int64"
	case FLOAT32:
		return "float32"
	case FLOAT64:
		return "float64"
	case DATE:
		return "date"
	case TIME:
		return "time"
	case TIMESTAMP:
		return "timestamp"
	case STRING:
		return "string"
	case BYTES:
		return "bytes"
	case JSON:
		return "json"
	case OBJECT:
		return "object"
	case ARRAY:
		return "array"
	case BAG:
		return "bag"
	}
	return "unknown type"
}

func TypeOfSlice(i interface{}) Type {
	if _, ok := i.([]bool); ok {
		return BOOL
	}
	if _, ok := i.([]int8); ok {
		return INT8
	}
	if _, ok := i.([]int16); ok {
		return INT16
	}
	if _, ok := i.([]int32); ok {
		return INT32
	}
	if _, ok := i.([]int64); ok {
		return INT64
	}
	if _, ok := i.([]float32); ok {
		return FLOAT32 
	}
	if _, ok := i.([]float64); ok {
		return FLOAT64
	}
	if _, ok := i.([]string); ok {
		return STRING
	}
	if _, ok := i.([][]byte); ok {
		return BYTES 
	}
	if _, ok := i.([]interface{}); ok {
		return OBJECT
	}
	return ANY
}

type NullInfo int8

const (
	DATA	NullInfo = 1 << iota
	NULL
	MISSING
)

type ColInfo struct {
	Typ Type
	TypMod int
	NullInfo NullInfo
	Name string
}

type Schema struct {
	Cols []ColInfo
}

func (s *Schema) AddCol(t Type, n string) *Schema {
	var ci ColInfo
	ci.Typ = t
	ci.Name = n
	s.Cols = append(s.Cols, ci)
	return s
}



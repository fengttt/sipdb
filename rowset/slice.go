package rowset

type SliceRowset struct {
	s Schema
	c ColVec
}

func FromSlice(s interface{}) {
	var rs SliceRowset
	rs.s.AddCol(types.OBJECT, "_1")

	if t, ok := s.([]bool); ok {
		rs.s.Cols[0].Typ = types.BOOL
		rs.c.B = t
		return
	}

	if t, ok := s.([]int8); ok {
		rs.s.Cols[0].Typ = types.INT8
		rs.c.I8 = t
		return
	}
}

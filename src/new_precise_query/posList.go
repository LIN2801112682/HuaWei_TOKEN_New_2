package new_precise_query

import (
	"build_VToken_index"
)

type PosList struct {
	Sid      build_VToken_index.SeriesId
	PosArray []int
}

func NewPosList(sid build_VToken_index.SeriesId, posArray []int) PosList {
	return PosList{
		Sid:      sid,
		PosArray: posArray,
	}
}

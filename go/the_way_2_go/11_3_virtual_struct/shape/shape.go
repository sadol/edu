package shape

import "../shaper"

//pure virtual class analogon from C++, can be now embedded in the  whole range of
//structs
type Shape struct {
	shaper.Shaper
}

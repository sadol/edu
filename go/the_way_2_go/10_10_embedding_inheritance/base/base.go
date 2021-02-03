package base

type Base struct {
	id int  "some named & PRIVATE field (getters and setters nedded)"
}

// getter
func (b Base) Id() int {
	return b.id
}

// setter
func (b *Base) SetId(newId int) {
	b.id = newId
}

// `constructor'
func NewBase(i int) *Base {
    return &Base{i} // no need to use setter here
}

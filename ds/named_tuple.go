package ds

type TupleName struct {
	namesToIndex map[string]int
	indexToName  map[int]string
}

func (tn *TupleName) Names() []string {
	return MapConvertKeyToSlice(tn.namesToIndex)
}

func (tn *TupleName) Set(name string, index int) {
	tn.namesToIndex[name] = index
	tn.indexToName[index] = name
}

func (tn *TupleName) Index(name string) int {
	if v, ok := tn.namesToIndex[name]; ok {
		return v
	}
	return -1
}

func (tn *TupleName) Name(index int) string {
	if v, ok := tn.indexToName[index]; ok {
		return v
	}
	return ""
}

type NamedTuple2E[A any, B any] struct {
	TupleName
	E1 A
	E2 B
}

func (nt *NamedTuple2E[A, B]) Get(name string) interface{} {
	index := nt.TupleName.Index(name)
	if index < 0 {
		var empty interface{}
		return empty
	}
	switch index {
	case 0:
		return nt.E1
	case 1:
		return nt.E2
	default:
		var empty interface{}
		return empty
	}
}

func (nt *NamedTuple2E[A, B]) Unpack() (A, B) {
	return nt.E1, nt.E2
}

type NamedTuple3E[A any, B any, C any] struct {
	NamedTuple2E[A, B]
	E3 C
}

func (nt *NamedTuple3E[A, B, C]) Get(name string) interface{} {
	index := nt.TupleName.Index(name)
	if index < 0 {
		var empty interface{}
		return empty
	}
	switch index {
	case 0:
		return nt.E1
	case 1:
		return nt.E2
	case 2:
		return nt.E3
	default:
		var empty interface{}
		return empty
	}
}
func (nt *NamedTuple3E[A, B, C]) Unpack() (A, B, C) {
	return nt.E1, nt.E2, nt.E3
}

type NamedTuple4E[A any, B any, C any, D any] struct {
	NamedTuple3E[A, B, C]
	E4 D
}

func (nt *NamedTuple4E[A, B, C, D]) Get(name string) interface{} {
	index := nt.TupleName.Index(name)
	if index < 0 {
		var empty interface{}
		return empty
	}
	switch index {
	case 0:
		return nt.E1
	case 1:
		return nt.E2
	case 2:
		return nt.E3
	case 3:
		return nt.E4
	default:
		var empty interface{}
		return empty
	}
}
func (nt *NamedTuple4E[A, B, C, D]) Unpack() (A, B, C, D) {
	return nt.E1, nt.E2, nt.E3, nt.E4
}

type NamedTuple5E[A any, B any, C any, D any, E any] struct {
	NamedTuple4E[A, B, C, D]
	E5 E
}

func (nt *NamedTuple5E[A, B, C, D, E]) Get(name string) interface{} {
	index := nt.TupleName.Index(name)
	if index < 0 {
		var empty interface{}
		return empty
	}
	switch index {
	case 0:
		return nt.E1
	case 1:
		return nt.E2
	case 2:
		return nt.E3
	case 3:
		return nt.E4
	case 4:
		return nt.E5
	default:
		var empty interface{}
		return empty
	}
}
func (nt *NamedTuple5E[A, B, C, D, E]) Unpack() (A, B, C, D, E) {
	return nt.E1, nt.E2, nt.E3, nt.E4, nt.E5
}

func NewNamedTuple2E[A any, B any](na string, a A, nb string, b B) *NamedTuple2E[A, B] {
	return &NamedTuple2E[A, B]{
		TupleName: TupleName{
			namesToIndex: map[string]int{
				na: 0,
				nb: 1,
			},
			indexToName: map[int]string{
				0: na,
				1: nb,
			},
		},
		E1: a,
		E2: b,
	}
}

func NewNamedTuple3E[A any, B any, C any](na string, a A, nb string, b B, nc string, c C) *NamedTuple3E[A, B, C] {
	tuple2e := NewNamedTuple2E(na, a, nb, b)
	tuple3e := &NamedTuple3E[A, B, C]{
		NamedTuple2E: *tuple2e,
		E3:           c,
	}
	tuple3e.TupleName.Set(nc, 2)
	return tuple3e
}

func NewNamedTuple4E[A any, B any, C any, D any](na string, a A, nb string, b B, nc string, c C, nd string, d D) *NamedTuple4E[A, B, C, D] {
	tuple3e := NewNamedTuple3E(na, a, nb, b, nc, c)
	tuple4e := &NamedTuple4E[A, B, C, D]{
		NamedTuple3E: *tuple3e,
		E4:           d,
	}
	tuple4e.TupleName.Set(nd, 3)
	return tuple4e
}

func NewNamedTuple5E[A any, B any, C any, D any, E any](na string, a A, nb string, b B, nc string, c C, nd string, d D, ne string, e E) *NamedTuple5E[A, B, C, D, E] {
	tuple4e := NewNamedTuple4E(na, a, nb, b, nc, c, nd, d)
	tuple5e := &NamedTuple5E[A, B, C, D, E]{
		NamedTuple4E: *tuple4e,
		E5:           e,
	}
	tuple5e.TupleName.Set(ne, 4)
	return tuple5e
}

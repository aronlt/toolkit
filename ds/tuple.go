package ds

// Tuple2E 自定义两个元素的Tuple
type Tuple2E[A any, B any] struct {
	E1 A
	E2 B
}

func (t *Tuple2E[A, B]) Unpack() (A, B) {
	return t.E1, t.E2
}

type Tuple3E[A any, B any, C any] struct {
	Tuple2E[A, B]
	E3 C
}

func (t *Tuple3E[A, B, C]) Unpack() (A, B, C) {
	return t.E1, t.E2, t.E3
}

type Tuple4E[A any, B any, C any, D any] struct {
	Tuple3E[A, B, C]
	E4 D
}

func (t *Tuple4E[A, B, C, D]) Unpack() (A, B, C, D) {
	return t.E1, t.E2, t.E3, t.E4
}

type Tuple5E[A any, B any, C any, D any, E any] struct {
	Tuple4E[A, B, C, D]
	E5 E
}

func (t *Tuple5E[A, B, C, D, E]) Unpack() (A, B, C, D, E) {
	return t.E1, t.E2, t.E3, t.E4, t.E5
}

func NewTuple2E[A any, B any](a A, b B) *Tuple2E[A, B] {
	return &Tuple2E[A, B]{E1: a, E2: b}
}

func NewTuple3E[A any, B any, C any](a A, b B, c C) *Tuple3E[A, B, C] {
	return &Tuple3E[A, B, C]{Tuple2E: *NewTuple2E(a, b), E3: c}
}

func NewTuple4E[A any, B any, C any, D any](a A, b B, c C, d D) *Tuple4E[A, B, C, D] {
	return &Tuple4E[A, B, C, D]{Tuple3E: *NewTuple3E(a, b, c), E4: d}
}

func NewTuple5E[A any, B any, C any, D any, E any](a A, b B, c C, d D, e E) *Tuple5E[A, B, C, D, E] {
	return &Tuple5E[A, B, C, D, E]{Tuple4E: *NewTuple4E(a, b, c, d), E5: e}
}

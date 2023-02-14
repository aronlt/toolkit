package ds

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

type Tuple6E[A any, B any, C any, D any, E any, F any] struct {
	Tuple5E[A, B, C, D, E]
	E6 F
}

func (t *Tuple6E[A, B, C, D, E, F]) Unpack() (A, B, C, D, E, F) {
	return t.E1, t.E2, t.E3, t.E4, t.E5, t.E6
}

type Tuple7E[A any, B any, C any, D any, E any, F any, G any] struct {
	Tuple6E[A, B, C, D, E, F]
	E7 G
}

func (t *Tuple7E[A, B, C, D, E, F, G]) Unpack() (A, B, C, D, E, F, G) {
	return t.E1, t.E2, t.E3, t.E4, t.E5, t.E6, t.E7
}

type Tuple8E[A any, B any, C any, D any, E any, F any, G any, H any] struct {
	Tuple7E[A, B, C, D, E, F, G]
	E8 H
}

func (t *Tuple8E[A, B, C, D, E, F, G, H]) Unpack() (A, B, C, D, E, F, G, H) {
	return t.E1, t.E2, t.E3, t.E4, t.E5, t.E6, t.E7, t.E8
}

type Tuple9E[A any, B any, C any, D any, E any, F any, G any, H any, I any] struct {
	Tuple8E[A, B, C, D, E, F, G, H]
	E9 I
}

func (t *Tuple9E[A, B, C, D, E, F, G, H, I]) Unpack() (A, B, C, D, E, F, G, H, I) {
	return t.E1, t.E2, t.E3, t.E4, t.E5, t.E6, t.E7, t.E8, t.E9
}

type Tuple10E[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] struct {
	Tuple9E[A, B, C, D, E, F, G, H, I]
	E10 J
}

func (t *Tuple10E[A, B, C, D, E, F, G, H, I, J]) Unpack() (A, B, C, D, E, F, G, H, I, J) {
	return t.E1, t.E2, t.E3, t.E4, t.E5, t.E6, t.E7, t.E8, t.E9, t.E10
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

func NewTuple6E[A any, B any, C any, D any, E any, F any](a A, b B, c C, d D, e E, f F) *Tuple6E[A, B, C, D, E, F] {
	return &Tuple6E[A, B, C, D, E, F]{Tuple5E: *NewTuple5E(a, b, c, d, e), E6: f}
}

func NewTuple7E[A any, B any, C any, D any, E any, F any, G any](a A, b B, c C, d D, e E, f F, g G) *Tuple7E[A, B, C, D, E, F, G] {
	return &Tuple7E[A, B, C, D, E, F, G]{Tuple6E: *NewTuple6E(a, b, c, d, e, f), E7: g}
}

func NewTuple8E[A any, B any, C any, D any, E any, F any, G any, H any](a A, b B, c C, d D, e E, f F, g G, h H) *Tuple8E[A, B, C, D, E, F, G, H] {
	return &Tuple8E[A, B, C, D, E, F, G, H]{Tuple7E: *NewTuple7E(a, b, c, d, e, f, g), E8: h}
}

func NewTuple9E[A any, B any, C any, D any, E any, F any, G any, H any, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) *Tuple9E[A, B, C, D, E, F, G, H, I] {
	return &Tuple9E[A, B, C, D, E, F, G, H, I]{Tuple8E: *NewTuple8E(a, b, c, d, e, f, g, h), E9: i}
}

func NewTuple10E[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) *Tuple10E[A, B, C, D, E, F, G, H, I, J] {
	return &Tuple10E[A, B, C, D, E, F, G, H, I, J]{Tuple9E: *NewTuple9E(a, b, c, d, e, f, g, h, i), E10: j}
}

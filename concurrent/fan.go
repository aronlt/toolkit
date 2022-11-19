package concurrent

import "github.com/aronlt/toolkit/ttypes"

func FanOut[T any](data ttypes.IFanOut[T], number int) []T {
	in := data.Produce()

	// FAN-OUT
	outChs := make([]chan T, number)
	for i := 0; i < number; i++ {
		outChs[i] = data.Compute(in)
	}

	// consumer
	out := make([]T, 0)
	for ret := range data.Merge(outChs...) {
		out = append(out, ret)
	}
	return out
}

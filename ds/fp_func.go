package ds

// copy from https://github.com/samber/lo

// FpPartial returns new function that, when called, has its first argument set to the provided value.
func FpPartial[T1, T2, R any](f func(a T1, b T2) R, arg1 T1) func(T2) R {
	return func(t2 T2) R {
		return f(arg1, t2)
	}
}

// FpPartial1 returns new function that, when called, has its first argument set to the provided value.
func FpPartial1[T1, T2, R any](f func(T1, T2) R, arg1 T1) func(T2) R {
	return FpPartial(f, arg1)
}

// FpPartial2 returns new function that, when called, has its first argument set to the provided value.
func FpPartial2[T1, T2, T3, R any](f func(T1, T2, T3) R, arg1 T1) func(T2, T3) R {
	return func(t2 T2, t3 T3) R {
		return f(arg1, t2, t3)
	}
}

// FpPartial3 returns new function that, when called, has its first argument set to the provided value.
func FpPartial3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) R, arg1 T1) func(T2, T3, T4) R {
	return func(t2 T2, t3 T3, t4 T4) R {
		return f(arg1, t2, t3, t4)
	}
}

// FpPartial4 returns new function that, when called, has its first argument set to the provided value.
func FpPartial4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) R, arg1 T1) func(T2, T3, T4, T5) R {
	return func(t2 T2, t3 T3, t4 T4, t5 T5) R {
		return f(arg1, t2, t3, t4, t5)
	}
}

// FpPartial5 returns new function that, when called, has its first argument set to the provided value
func FpPartial5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) R, arg1 T1) func(T2, T3, T4, T5, T6) R {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) R {
		return f(arg1, t2, t3, t4, t5, t6)
	}
}

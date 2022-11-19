package ttypes

import "context"

type FetchHandler[T any] func(context.Context, chan T)

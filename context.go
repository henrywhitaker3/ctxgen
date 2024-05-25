package ctxgen

import "context"

type keyName any

func WithValue[T any](parent context.Context, key any, val T) context.Context {
	return context.WithValue(parent, keyName(key), val)
}

func ValueOk[T any](ctx context.Context, key any) (T, bool) {
	var out T
	rv, ok := ctx.Value(key).(T)
	if ok {
		return rv, true
	}
	return out, false
}

func Value[T any](ctx context.Context, key any) T {
	out, _ := ValueOk[T](ctx, key)
	return out
}

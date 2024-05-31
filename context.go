package ctxgen

import "context"

type keyName any

func WithValue[T any](parent context.Context, key any, val T) context.Context {
	if parent == nil {
		parent = context.Background()
	}
	return context.WithValue(parent, keyName(key), val)
}

func ValueOk[T any](ctx context.Context, key any) (T, bool) {
	var out T
	if ctx == nil {
		return out, false
	}
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

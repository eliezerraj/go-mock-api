package utils

import (
	"context"
)

func AddToContext(parent context.Context, key, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}

func FindToContext(ctx context.Context, key interface{}) interface{} {
	return ctx.Value(key)
}

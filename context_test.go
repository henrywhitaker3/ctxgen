package ctxgen_test

import (
	"context"
	"testing"

	"github.com/henrywhitaker3/ctxgen"
)

func TestItDoesNotReturnOkayWhenNotSet(t *testing.T) {
	ctx := context.Background()
	ctx = ctxgen.WithValue(ctx, "bongo", 5)

	_, ok := ctxgen.ValueOk[int](ctx, "bingo")
	if ok {
		t.Error("key was found in context when it shouldn't have been")
	}
}

func TestItRetrievesAValueFromContext(t *testing.T) {
	ctx := context.Background()
	ctx = ctxgen.WithValue(ctx, "bongo", 5)

	val, ok := ctxgen.ValueOk[int](ctx, "bongo")
	if !ok {
		t.Error("key was not found in context when it should have been")
	}
	if val != 5 {
		t.Error("value retrieved from context is not 5")
	}
}

func TestValueReturnsFalseWhenCtxIsNil(t *testing.T) {
	val, ok := ctxgen.ValueOk[int](nil, "bongo")
	if ok {
		t.Fail()
	}
	if val != 0 {
		t.Fail()
	}
}

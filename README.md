# Ctxgen

A package to store/retreive typed values from a context

```go
ctx := context.Background()
ctx = ctxgen.WithValue(ctx, "bongo", 5)

...

value, ok := ctx.ValueOk[int](ctx, "bongo")
fmt.Println(value, ok) // prints 5 true
```

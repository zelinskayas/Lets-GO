package main

import (
	"context"
	"fmt"
)

type ctxKey1 string
type ctxKey2 string

func main() {
	ctx := context.Background()

	var key1 ctxKey1 = "some key1"
	var key2 ctxKey2 = "some key2"

	ctx = context.WithValue(ctx, key1, "some value1")
	ctx = context.WithValue(ctx, key2, "some value2")

	valcontex(ctx)
}
func valcontex(ctx context.Context) {
	var key1 ctxKey1 = "some key1"
	var key2 ctxKey2 = "some key2"

	fmt.Printf("%v: %v\n", key1, ctx.Value(key1))
	fmt.Printf("%v: %v\n", key2, ctx.Value(key2))
}

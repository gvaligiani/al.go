package set

import (
	"github.com/gvaligiani/al.go/dict"
	"github.com/gvaligiani/al.go/util"
)

func NoneOf[T comparable, S ~map[T]struct{}](items S, predicate util.Predicate[T]) bool {
	return dict.NoKeyOf(items, util.TestOnFirstArg[T, struct{}](predicate))
}

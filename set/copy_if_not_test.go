package set_test

import (
	"testing"

	"go.uber.org/zap"

	"github.com/gvaligiani/al.go/set"
	"github.com/gvaligiani/al.go/test"
	"github.com/gvaligiani/al.go/util"
)

func TestCopyIfNotInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[int64]
		predicate util.Predicate[int64]
		wantItems set.Set[int64]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyInt64Set,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: EmptyInt64Set,
		},
		"remove-none": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return false },
			wantItems: DefaultInt64Set,
		},
		"remove-odd": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i%2 == 0 },
			wantItems: set.New[int64](
				21,
				87,
			),
		},
		"remove-even": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return i%2 == 1 },
			wantItems: set.New[int64](
				12,
				34,
				52,
			),
		},
		"remove-all": {
			items:     DefaultInt64Set,
			predicate: func(i int64) bool { return true },
			wantItems: EmptyInt64Set,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.CopyIfNot(testCase.items, testCase.predicate)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfNotStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[Item]
		predicate func(Item) bool
		wantItems set.Set[Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemSet,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemSet,
		},
		"remove-none": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return false },
			wantItems: DefaultItemSet,
		},
		"remove-odd": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value%2 == 0 },
			wantItems: set.New(
				Item{Value: 21},
				Item{Value: 87},
			),
		},
		"remove-even": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return item.Value%2 == 1 },
			wantItems: set.New(
				Item{Value: 12},
				Item{Value: 34},
				Item{Value: 52},
			),
		},
		"remove-all": {
			items:     DefaultItemSet,
			predicate: func(item Item) bool { return true },
			wantItems: EmptyItemSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.CopyIfNot(testCase.items, testCase.predicate)

		// assert
		assertEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

func TestCopyIfNotStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     set.Set[*Item]
		predicate func(*Item) bool
		wantItems set.Set[*Item]
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: nil,
		},
		"empty": {
			items:     EmptyItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: EmptyItemPointerSet,
		},
		"remove-none": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return false },
			wantItems: DefaultItemPointerSet,
		},
		"remove-odd": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%2 == 0 },
			wantItems: set.New(
				&Item{Value: 21},
				&Item{Value: 87},
			),
		},
		"remove-even": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return item.Value%2 == 1 },
			wantItems: set.New(
				&Item{Value: 12},
				&Item{Value: 34},
				&Item{Value: 52},
			),
		},
		"remove-all": {
			items:     DefaultItemPointerSet,
			predicate: func(item *Item) bool { return true },
			wantItems: EmptyItemPointerSet,
		},
	}

	//
	// run
	//

	test.RunTestCases(t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItems := set.CopyIfNot(testCase.items, testCase.predicate)

		// assert
		assertDeepEqual(t, testCase.wantItems, gotItems, "wrong items!")
	})
}

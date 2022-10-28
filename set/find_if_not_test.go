package set_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/set"
	"github.com/gvaligiani/algo/test"
)

func TestFindIfNotInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int64]struct{}
		predicate func(int64) bool
		wantItem  int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"empty": {
			items:     test.EmptyInt64Set,
			predicate: func(i int64) bool { return i%10 == 3 },
			wantFound: false,
			wantItem:  0,
		},
		"no-match": {
			items:     test.Int64Set,
			predicate: func(i int64) bool { return i < 100 },
			wantFound: false,
			wantItem:  0,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.Int64Set,
		// 	predicate: func(i int64) bool { return i%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  21,
		// },
		"one-match": {
			items:     test.Int64Set,
			predicate: func(i int64) bool { return i < 60 },
			wantFound: true,
			wantItem:  87,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot[int64](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[test.Item]struct{}
		predicate func(test.Item) bool
		wantItem  test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  test.Item{},
		},
		"empty": {
			items:     test.EmptyItemSet,
			predicate: func(item test.Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  test.Item{},
		},
		"no-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  test.Item{},
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.ItemSet,
		// 	predicate: func(item test.Item) bool { return item.Value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  test.Item{Value: 21},
		// },
		"one-match": {
			items:     test.ItemSet,
			predicate: func(item test.Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  test.Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot[test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

func TestFindIfNotStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[*test.Item]struct{}
		predicate func(*test.Item) bool
		wantItem  *test.Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"empty": {
			items:     test.EmptyItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value%10 == 3 },
			wantFound: false,
			wantItem:  nil,
		},
		"no-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value < 100 },
			wantFound: false,
			wantItem:  nil,
		},
		// NOTE order could not be guarantee >>> expected 21, 12, 87, 52
		// "some-match": {
		// 	items:     test.ItemPointerSet,
		// 	predicate: func(item *test.Item) bool { return item.Value%10 == 4 },
		// 	wantFound: true,
		// 	wantItem:  &test.Item{Value: 21},
		// },
		"one-match": {
			items:     test.ItemPointerSet,
			predicate: func(item *test.Item) bool { return item.Value < 60 },
			wantFound: true,
			wantItem:  &test.Item{Value: 87},
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotItem, gotFound := set.FindIfNot[*test.Item](testCase.items, testCase.predicate)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
		require.Equalf(t, testCase.wantItem, gotItem, "wrong item!")
	})
}

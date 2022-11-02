package dict_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/gvaligiani/algo/dict"
	"github.com/gvaligiani/algo/test"
)

func TestFindValueInt64(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]int64
		value     int64
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     34,
			wantFound: false,
		},
		"empty": {
			items:     EmptyInt64Dict,
			value:     34,
			wantFound: false,
		},
		"found": {
			items:     DefaultInt64Dict,
			value:     34,
			wantFound: true,
		},
		"not-found": {
			items:     DefaultInt64Dict,
			value:     99,
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindValue[int,int64](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindValueStruct(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]Item
		value     Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemDict,
			value:     Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     DefaultItemDict,
			value:     Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemDict,
			value:     Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindValue[int,Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

func TestFindValueStructPointer(t *testing.T) {

	//
	// test cases
	//

	type TestCase struct {
		items     map[int]*Item
		value     *Item
		wantFound bool
	}

	testCases := map[string]TestCase{
		"nil": {
			items:     nil,
			value:     &Item{Value: 34},
			wantFound: false,
		},
		"empty": {
			items:     EmptyItemPointerDict,
			value:     &Item{Value: 34},
			wantFound: false,
		},
		"found": {
			items:     DefaultItemPointerDict,
			value:     &Item{Value: 34},
			wantFound: true,
		},
		"not-found": {
			items:     DefaultItemPointerDict,
			value:     &Item{Value: 99},
			wantFound: false,
		},
	}

	//
	// run
	//

	test.RunTestCases[TestCase](t, testCases, func(t *testing.T, logger *zap.Logger, testCase TestCase) {

		// execute
		gotFound := dict.FindValue[int, *Item](testCase.items, testCase.value)

		// assert
		require.Equalf(t, testCase.wantFound, gotFound, "wrong found!")
	})
}

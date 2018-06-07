package common

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		s   interface{}
		w   *Array
		err bool
	}{
		{
			s:   1,
			err: true,
		},
		{
			s: []int{1},
			w: &Array{
				1,
			},
		},
		{
			s: [1]int{1},
			w: &Array{
				1,
			},
		},
		{
			s: []int{1, 2},
			w: &Array{
				1, 2,
			},
		},
		{
			s: [2]int{1, 2},
			w: &Array{
				1, 2,
			},
		},
		{
			s: []interface{}{1, 2, "hoge"},
			w: &Array{
				1, 2, "hoge",
			},
		},
	}

	for _, test := range tests {
		a, err := NewArray(test.s)
		if test.err && err == nil {
			t.Error("error is empty")
		}

		if !test.err && err != nil {
			t.Error("error is not empty")
		}

		if test.err {
			continue
		}

		if a == nil {
			t.Error("array is empty")
		}

		if !reflect.DeepEqual(*a, *test.w) {
			t.Errorf("miss match result: %v, %v", *a, test.w)
		}
	}
}

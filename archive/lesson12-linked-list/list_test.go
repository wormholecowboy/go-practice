package main

import "testing"

func TestPushHead(t *testing.T) {
	var tests = []struct {
		val  []int
		want []int
	}{
		{val: []int{15}, want: []int{15}},
			{val: []int{1,2,3}, want: []int{3,2,1}},
			{val: []int{}, want: []int{}},
	}

	for _, test := range tests {
		var list = List[int]{}

		for _, v := range test.val {
		list.PushHead(v)
		}

		got := list.ToSlice()
		if len(got) != len(test.want) {
			t.Errorf("PushHead(%v): got %d len for list, expected: %d", test.val, len(got), len(test.val))
			continue
		}

		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("PushHead(%v): got %v, wanted %v", test.val, got, test.want)
				break
			}
		}

	}

}

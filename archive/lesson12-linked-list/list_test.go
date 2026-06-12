package main

import "testing"

func TestPushHead(t *testing.T) {
	var tests = []struct {
		val  int
		want int
	}{
		{val: 15, want: 15},
	}

	for _, test := range tests {
		var list = List[int]{}
		list.PushHead(test.val)

		got, ok := list.PeekHead()
		if !ok {
			t.Fatalf("PeekHead returned no value after PushHead(%d)", test.val)
		}

		if got != test.want {
			t.Errorf("PushHead(%d) does not equal wanted: %d", test.val, test.want)
		}

	}

}

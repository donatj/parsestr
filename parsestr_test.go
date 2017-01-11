package parsestr

import (
	"reflect"
	"testing"
)

func TestGetSub(t *testing.T) {
	var tests = []struct {
		input  string
		key    []string
		output []Values
	}{
		{
			"foo[]=bar&foo[]=baz",
			[]string{"foo"},
			[]Values{
				{"": []string{"bar", "baz"}},
			},
		},
		{
			"foo[bar]=baz&foo[qux]=quux&dog[]=cat",
			[]string{"foo", "dog"},
			[]Values{
				{"bar": []string{"baz"}, "qux": []string{"quux"}},
				{"": []string{"cat"}},
			},
		},
		{
			"dog=cat",
			[]string{"dog"},
			[]Values{{}},
		},
	}

	for _, test := range tests {
		v, e := ParseQuery(test.input)
		if e != nil {
			t.Error(e)
		}

		for i, key := range test.key {
			output := test.output[i]
			actual := v.GetSub(key)

			if !reflect.DeepEqual(actual, output) {
				t.Errorf(`v = ParseQuery(%q); v.GetSub(%q) = %#v; want %#v`, test.input, key, actual, output)
			}
		}

	}
}

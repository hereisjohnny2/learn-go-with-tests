package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{} 
		ExpectedCalls []string
	} {
		{
			"struct with one string field",
			struct { Name string }{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two fields",
			struct { 
				Name string
				City string
			}{"Chris", "RJ"},
			[]string{"Chris", "RJ"},
		},
		{
			"struct with non string field",
			struct { 
				Name string
				City string
				Age int 
			}{"Chris", "RJ", 28},
			[]string{"Chris", "RJ"},
		},
		{
			"nested fields",
			Person {
				"Chris",
				Profile{33, "RJ"},
			},
			[]string{"Chris", "RJ"},
		},
		{
			"pointers to things",
			&Person {
				"Chris",
				Profile{33, "RJ"},
			},
			[]string{"Chris", "RJ"},
		},
		{
			"slices",
			[]Profile{{33, "RJ"}, {35, "SP"}},
			[]string{"RJ", "SP"},
		},
		{
			"arrays",
			[2]Profile{{33, "RJ"}, {35, "SP"}},
			[]string{"RJ", "SP"},
		},
		{
			"maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)				
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}


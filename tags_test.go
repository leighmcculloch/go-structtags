package structtags

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		String string
		Tag    Tag
	}{
		{
			"",
			Tag{},
		},
		{
			"omitempty",
			Tag{"omitempty", TagOptions{}},
		},
		{
			"omitempty,abc",
			Tag{"omitempty", TagOptions{"abc"}},
		},
		{
			"omitempty,abc,def",
			Tag{"omitempty", TagOptions{"abc", "def"}},
		},
	}

	for _, test := range tests {
		tag := Parse(test.String)
		if !reflect.DeepEqual(tag, test.Tag) {
			t.Errorf("Parse tag %v got %v, want %v", test.String, tag, test.Tag)
		}
	}
}

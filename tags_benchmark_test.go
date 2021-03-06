package structtags_test

import (
	"testing"

	"4d63.com/structtags"
)

func BenchmarkParse(b *testing.B) {
	s := "name,abc,def,efg"
	for i := 0; i < b.N; i++ {
		structtags.Parse(s)
	}
}

func BenchmarkTagOptionsContains(b *testing.B) {
	s := "name,abc,def,efg"
	t := structtags.Parse(s)
	o := t.Options
	for i := 0; i < b.N; i++ {
		o.Contains("efg")
	}
}

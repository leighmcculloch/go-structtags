package structtags_test

import (
	"fmt"
	"reflect"

	"4d63.com/structtags"
)

func ExampleParse() {
	s := struct {
		A string `myTag:"a,omitempty,flatten"`
		B string `myTag:",omitempty"`
		C int
	}{}

	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if tagStr, ok := f.Tag.Lookup("myTag"); ok {
			tag := structtags.Parse(tagStr)

			fmt.Printf("Tag: %#v\n", tag.Value)
			fmt.Println("  Omit Empty:", tag.Options.Contains("omitempty"))
			fmt.Println("  Flatten:", tag.Options.Contains("flatten"))
		}
	}

	// Output:
	// Tag: "a"
	//   Omit Empty: true
	//   Flatten: true
	// Tag: ""
	//   Omit Empty: true
	//   Flatten: false
}

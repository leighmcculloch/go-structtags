// Package structtags parses struct tags based on the pattern used in the encoding packages of the go stdlib.
package structtags

import "strings"

// Parse parses a string tag into a Tag struct.
// A tag defined as `myTag:"value,abc,def"` will
// be parsed as:
// Tag {
//   Value: "value",
//   Options: ["abc", "def"],
// }
func Parse(tag string) Tag {
	if len(tag) == 0 {
		return Tag{}
	}
	tagParts := strings.Split(tag, ",")
	return Tag{
		Value:   tagParts[0],
		Options: tagParts[1:],
	}
}

// Tag contains a tag value and options.
type Tag struct {
	Value   string
	Options TagOptions
}

// TagOptions contains the options for a tag.
type TagOptions []string

// Contains checks if the option is contained in the
// TagOptions.
func (o TagOptions) Contains(option string) bool {
	for _, v := range o {
		if v == option {
			return true
		}
	}
	return false
}

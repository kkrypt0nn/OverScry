import sys

TEMPLATE = """package tags

import (
	"reflect"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// {{CAPITALIZED_TAG}} is used to tag CHANGEME.
// https://wiki.openstreetmap.org/wiki/CHANGEME
type {{CAPITALIZED_TAG}} struct {
	// ChangeMe represents CHANGEME.
	ChangeMe *Feature `yaml:"changeme,omitempty"`
}

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	temp := &struct {
		ChangeMe *Feature `yaml:"changeme,omitempty"`
	}{}
	if err := unmarshal(temp); err != nil {
		return err
	}

	val := reflect.ValueOf({{FIRST_CHAR}}).Elem()
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)

		tempField := reflect.ValueOf(temp).Elem().FieldByName(field.Name)
		if tempField.IsValid() && !tempField.IsNil() {
			fieldVal.Set(tempField)
			feature := fieldVal.Interface().(*Feature)
			feature.Tag = strings.ToLower(field.Name)
			feature.Prefix = helpers.Ptr({{FIRST_CHAR}}.GetTag())
		}
	}

	return nil
}

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) GetTag() string {
	return "{{TAG}}"
}

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) ToOQL() string {
	var queryParts []string

	val := reflect.ValueOf({{FIRST_CHAR}}).Elem()
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldVal := val.Field(i)
		if fieldVal.IsValid() && !fieldVal.IsNil() {
			queryParts = append(queryParts, fieldVal.Interface().(*Feature).ToOQL())
		}
	}

	return strings.Join(queryParts, "")
}

"""

try:
    _tag = sys.argv[1]
except:
    print("Error: Please provide a tag as argument when executing the tool.")
    sys.exit(1337)

_first_char = _tag[0].lower()
_capitalized_tag = _tag.capitalize()

replace_with = {
    "{{FIRST_CHAR}}": _first_char,
    "{{TAG}}": _tag,
    "{{CAPITALIZED_TAG}}": _capitalized_tag,
}

result = TEMPLATE

for key in replace_with.keys():
    result = result.replace(key, replace_with[key])

print(result)

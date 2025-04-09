import sys

TEMPLATE = """package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// {{TYPE}} represents the valid {{TAG}} types.
// https://wiki.openstreetmap.org/wiki/Key:{{TAG}}
type {{TYPE}} string

const (
	// {{CAPITALIZED_TAG}}ChangeMe represents CHANGEME.
	{{CAPITALIZED_TAG}}ChangeMe {{TYPE}} = "changeme"
)

var {{VALID_TYPES}} = map[string]{{TYPE}}{
	"changeme": {{CAPITALIZED_TAG}}ChangeMe,
}

// {{CAPITALIZED_TAG}} is used to tag CHANGEME.
// https://wiki.openstreetmap.org/wiki/CHANGEME
type {{CAPITALIZED_TAG}} Feature

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := {{VALID_TYPES}}[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len({{VALID_TYPES}}))
			for k := range {{VALID_TYPES}} {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", {{FIRST_CHAR}}.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = {{FIRST_CHAR}}.GetTag()

	*{{FIRST_CHAR}} = {{CAPITALIZED_TAG}}(f)
	return nil
}

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) GetTag() string {
	return "{{TAG}}"
}

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) ToOQL() string {
	return Feature(*{{FIRST_CHAR}}).ToOQL()
}

"""


try:
    _tag = sys.argv[1]
except:
    print("Error: Please provide a tag as argument when executing the tool.")
    sys.exit(1337)


_first_char = _tag[0].lower()
_capitalized_tag = _tag.capitalize()
_type = f"{_capitalized_tag}Type"
_valid_types = f"valid{_capitalized_tag}Types"


replace_with = {
    "{{FIRST_CHAR}}": _first_char,
    "{{TAG}}": _tag,
    "{{CAPITALIZED_TAG}}": _capitalized_tag,
    "{{TYPE}}": _type,
    "{{VALID_TYPES}}": _valid_types,
}


result = TEMPLATE


for key in replace_with.keys():
    result = result.replace(key, replace_with[key])


print(result)

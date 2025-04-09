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
type {{CAPITALIZED_TAG}} struct {
	Value  *{{TYPE}}   `yaml:"value,omitempty"`
	Match  MatchMethod `yaml:"match,omitempty"`
	Inline *Feature    `yaml:"-"`

	// ChangeMe represents CHANGEME.
	ChangeMe *Feature `yaml:"changeme,omitempty"`
}

func ({{FIRST_CHAR}} *{{CAPITALIZED_TAG}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	temp := &struct {
		Value *string     `yaml:"value,omitempty"`
		Match MatchMethod `yaml:"match,omitempty"`

		ChangeMe *Feature `yaml:"changeme,omitempty"`
	}{}
	if err := unmarshal(temp); err != nil {
		return err
	}

	if temp.Value != nil {
		val := strings.ToLower(*temp.Value)
		if _, ok := {{VALID_TYPES}}[val]; ok {
			{{FIRST_CHAR}}.Value = (*{{TYPE}})(helpers.Ptr(val))
		} else {
			validKeys := make([]string, 0, len({{VALID_TYPES}}))
			for k := range {{VALID_TYPES}} {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", {{FIRST_CHAR}}.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	{{FIRST_CHAR}}.Inline = &Feature{
		Match: temp.Match,
		Value: helpers.Ptr(*temp.Value),
		Tag:   {{FIRST_CHAR}}.GetTag(),
	}

	tempVal := reflect.ValueOf(temp).Elem()
	targetVal := reflect.ValueOf({{FIRST_CHAR}}).Elem()
	for i := 0; i < tempVal.NumField(); i++ {
		field := tempVal.Type().Field(i)
		if field.Name == "Value" || field.Name == "Match" {
			continue
		}

		tempField := tempVal.FieldByName(field.Name)
		if tempField.IsValid() && !tempField.IsNil() {
			feature := tempField.Interface().(*Feature)
			tempFeature := &Feature{
				Value:  feature.Value,
				Match:  feature.Match,
				Tag:    strings.ToLower(field.Name),
				Prefix: helpers.Ptr({{FIRST_CHAR}}.GetTag()),
			}
			if validator, ok := validators.{{CAPITALIZED_TAG}}Validators[tempFeature.Tag]; ok {
				tempFeature.Validator = validator
				if err := tempFeature.Validate(); err != nil {
					return err
				}
			}
			targetField := targetVal.FieldByName(field.Name)
			targetField.Set(reflect.ValueOf(tempFeature))
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
		field := t.Field(i)
		fieldVal := val.Field(i)

		if field.Name == "Value" || field.Name == "Match" {
			continue
		}

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

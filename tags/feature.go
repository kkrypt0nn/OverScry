package tags

import (
	"fmt"
	"strings"

	"github.com/kkrypt0nn/overscry/tags/validators"
)

type Feature struct {
	Prefix *string     `yaml:"-"`
	Tag    string      `yaml:"-"`
	Value  *string     `yaml:"value,omitempty"`
	Match  MatchMethod `yaml:"match,omitempty"`

	Validator *validators.FeatureValidator `yaml:"-"`
}

func (f *Feature) Validate() error {
	if f.Validator != nil && f.Value != nil && !f.Validator.Validate(*f.Value) {
		return fmt.Errorf("invalid %s value: %q (allowed: %s)",
			f.Tag, *f.Value, strings.Join(f.Validator.AllowedValues, ", "))
	}
	return nil
}

func (f Feature) ToOQL() string {
	if f.Value == nil {
		return ""
	}

	prefix := ""
	if f.Prefix != nil {
		prefix = fmt.Sprintf("%s:", *f.Prefix)
	}
	return fmt.Sprintf(`["%s%s"%s"%s"]`, prefix, f.Tag, f.Match.ToOperator(), *f.Value)
}

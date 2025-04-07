package tags

import "fmt"

type Feature struct {
	Prefix *string     `yaml:"-"`
	Tag    string      `yaml:"-"`
	Value  *string     `yaml:"value,omitempty"`
	Match  MatchMethod `yaml:"match,omitempty"`
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

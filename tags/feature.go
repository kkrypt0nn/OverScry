package tags

import "fmt"

type Feature struct {
	Prefix string      `yaml:"-"`
	Tag    string      `yaml:"-"`
	Value  *string     `yaml:"value,omitempty"`
	Match  MatchMethod `yaml:"match,omitempty"`
}

func (f Feature) ToOQL() string {
	if f.Value == nil {
		return ""
	}
	return fmt.Sprintf(`["%s:%s"%s"%s"]`, f.Prefix, f.Tag, f.Match.ToOperator(), *f.Value)
}

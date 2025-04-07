package tags

type Tag interface {
	ToOQL() string
}

type MatchMethod string

const (
	Equals    MatchMethod = "equals"
	NotEquals MatchMethod = "not_equals"
	Regex     MatchMethod = "regex"
)

func (m MatchMethod) ToOperator() string {
	switch m {
	case Equals:
		return "="
	case NotEquals:
		return "!="
	case Regex:
		return "~"
	default:
		return "="
	}
}

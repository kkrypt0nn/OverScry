package tags

type MatchMethod string

const (
	Equals    MatchMethod = "="
	NotEquals MatchMethod = "!="
	Regex     MatchMethod = "~"
)

func (m MatchMethod) ToOperator() MatchMethod {
	switch m {
	case "equals":
		return Equals
	case "not_equals":
		return NotEquals
	case "regex":
		return Regex
	default:
		return Equals
	}
}

package models

import (
	"fmt"
	"strings"

	"github.com/kkrypt0nn/overscry/tags"
)

type Node struct {
	Addr tags.Addr `yaml:"addr"`
}

func (n Node) ToOQL() string {
	queryParts := []string{}

	if addrQuery := n.Addr.ToOQL(); addrQuery != "" {
		queryParts = append(queryParts, addrQuery)
	}

	if len(queryParts) == 0 {
		return ""
	}

	return fmt.Sprintf("node%s;", strings.Join(queryParts, ""))
}

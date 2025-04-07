package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kkrypt0nn/overscry/tags"
)

type Node struct {
	Addr      *tags.Addr      `yaml:"addr,omitempty"`
	Aerialway *tags.Aerialway `yaml:"aerialway,omitempty"`
	Aeroway   *tags.Aeroway   `yaml:"aeroway,omitempty"`
}

func (n *Node) ToOQL() string {
	var queryParts []string

	val := reflect.ValueOf(n).Elem()
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldVal := val.Field(i)
		if fieldVal.IsValid() && !fieldVal.IsNil() {
			queryParts = append(queryParts, fieldVal.Interface().(tags.Tag).ToOQL())
		}
	}

	return fmt.Sprintf("node%s;", strings.Join(queryParts, ""))
}

package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// BoundaryType represents the valid boundary types.
// https://wiki.openstreetmap.org/wiki/Key:boundary
type BoundaryType string

const (
	// BoundaryMarker represents a boundary marker, border marker, boundary stone, or border stone is a robust physical marker that identifies the start of a land boundary or the change in a boundary, especially a change in direction of a boundary.
	BoundaryMarker BoundaryType = "marker"
)

var validBoundaryTypes = map[string]BoundaryType{
	"marker": BoundaryMarker,
}

// Boundary is used to tag administrative and other boundaries.
// https://wiki.openstreetmap.org/wiki/Key:boundary
type Boundary Feature

func (b *Boundary) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validBoundaryTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validBoundaryTypes))
			for k := range validBoundaryTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", b.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = b.GetTag()

	*b = Boundary(f)
	return nil
}

func (b *Boundary) GetTag() string {
	return "boundary"
}

func (b *Boundary) ToOQL() string {
	return Feature(*b).ToOQL()
}

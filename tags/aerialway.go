package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// AerialwayType represents the valid aerialway types.
// https://wiki.openstreetmap.org/wiki/Key:aerialway
type AerialwayType string

const (
	// AerialwayPylon represents a pylon supporting the aerialway cable.
	AerialwayPylon AerialwayType = "pylon"
	// AerialwayStation represents station where passengers and/or goods can enter and/or leave the aerialway
	AerialwayStation AerialwayType = "station"
)

var validAerialwayTypes = map[string]AerialwayType{
	"pylon":   AerialwayPylon,
	"station": AerialwayStation,
}

// Aerialway is used to tag different forms of transportation for people or goods by using aerial wires.
// https://wiki.openstreetmap.org/wiki/Map_features#Aerialway
type Aerialway Feature

func (a *Aerialway) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validAerialwayTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validAerialwayTypes))
			for k := range validAerialwayTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", a.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = a.GetTag()

	*a = Aerialway(f)
	return nil
}

func (a *Aerialway) GetTag() string {
	return "aerialway"
}

func (a *Aerialway) ToOQL() string {
	return Feature(*a).ToOQL()
}

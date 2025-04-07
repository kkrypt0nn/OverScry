package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// AerowayType represents the valid aeroway types.
// https://wiki.openstreetmap.org/wiki/Key:aeroway
type AerowayType string

const (
	// AerowayAerodrome represents an aerodrome, airport or airfield
	AerowayAerodrome AerowayType = "aerodrome"
	// AerowayAircraftCrossing represents a point where the flow of traffic is impacted by crossing aircraft.
	AerowayAircraftCrossing AerowayType = "aircraft_crossing"
	// AerowayGate represents the bounded space, inside the airport terminal, where passengers wait before boarding their flight
	AerowayGate AerowayType = "gate"
	// AerowayHelipad represents a landing area or platform designed for helicopters.
	AerowayHelipad AerowayType = "helipad"
	// AerowayHeliport represents a special aerodrome built for helicopters.
	AerowayHeliport AerowayType = "heliport"
	// AerowayNavigationaid represents a facility that supports visual navigation for aircraft.
	AerowayNavigationaid AerowayType = "navigationaid"
	// AerowaySpaceport represents a spaceport or cosmodrome: a site for launching or receiving spacecraft.
	AerowaySpaceport AerowayType = "spaceport"
	// AerowayTerminal represents an airport passenger building.
	AerowayTerminal AerowayType = "terminal"
	// AerowayWindsock represents an object that shows wind direction and speed.
	AerowayWindsock AerowayType = "windsock"
)

var validAerowayTypes = map[string]AerowayType{
	"aerodrome":         AerowayAerodrome,
	"aircraft_crossing": AerowayAircraftCrossing,
	"gate":              AerowayGate,
	"helipad":           AerowayHelipad,
	"heliport":          AerowayHeliport,
	"navigationaid":     AerowayNavigationaid,
	"spaceport":         AerowaySpaceport,
	"terminal":          AerowayTerminal,
	"windsock":          AerowayWindsock,
}

// Aeroway is used to tag aerodromes, airfields other ground facilities that support the operation of airplanes and helicopters.
// https://wiki.openstreetmap.org/wiki/Aeroways
type Aeroway Feature

func (a *Aeroway) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validAerowayTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validAerowayTypes))
			for k := range validAerowayTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid aeroway value: %q (allowed: %s)", val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = "aeroway"

	*a = Aeroway(f)
	return nil
}

func (a Aeroway) ToOQL() string {
	return Feature(a).ToOQL()
}

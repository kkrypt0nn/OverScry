package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// PublicTransportType represents the valid public_transport types.
// https://wiki.openstreetmap.org/wiki/Key:public_transport
type PublicTransportType string

const (
	// PublicTransportStopPosition represents the position on the street or rails where a public transport vehicle stops.
	PublicTransportStopPosition PublicTransportType = "stop_position"
	// PublicTransportPlatform represents the place where passengers are waiting for the public transport vehicles.
	PublicTransportPlatform PublicTransportType = "platform"
	// PublicTransportStation represents a station is an area designed to access public transport.
	PublicTransportStation PublicTransportType = "station"
)

var validPublicTransportTypes = map[string]PublicTransportType{
	"stop_position": PublicTransportStopPosition,
	"platform":      PublicTransportPlatform,
	"station":       PublicTransportStation,
}

// PublicTransport is used to tag features related to public transport. For example: railway stations, bus stops and services.
// https://wiki.openstreetmap.org/wiki/Public_transport
type PublicTransport Feature

func (p *PublicTransport) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validPublicTransportTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validPublicTransportTypes))
			for k := range validPublicTransportTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", p.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = p.GetTag()

	*p = PublicTransport(f)
	return nil
}

func (p *PublicTransport) GetTag() string {
	return "public_transport"
}

func (p *PublicTransport) ToOQL() string {
	return Feature(*p).ToOQL()
}

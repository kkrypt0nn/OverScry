package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// PowerType represents the valid power types.
// https://wiki.openstreetmap.org/wiki/Key:power
type PowerType string

const (
	// PowerCatenaryMast represents a pole supporting the overhead wires used to supply electricity to vehicles equipped with a pantograph such as trams and trains.
	PowerCatenaryMast PowerType = "catenary_mast"
	// PowerCompensator represents a static power device used to ensure power quality and electrical network resilience.
	PowerCompensator PowerType = "compensator"
	// PowerConnection represents a free-standing electrical connection between two or more power lines or cables.
	PowerConnection PowerType = "connection"
	// PowerConverter represents a device to convert power between alternating and direct current electrical power: often, but not only, over high voltage networks.
	PowerConverter PowerType = "converter"
	// PowerGenerator represents a device which converts one form of energy to another, for example, an electrical generator.
	PowerGenerator PowerType = "generator"
	// PowerHeliostat represents a mirror of a heliostat device.
	PowerHeliostat PowerType = "heliostat"
	// PowerInsulator represents an electrical insulator which connects a power line to a support structure and prevents grounding.
	PowerInsulator PowerType = "insulator"
	// PowerInverter represents a device to convert power from direct current to alternating current.
	PowerInverter PowerType = "inverter"
	// PowerPole represents a single pole supporting power lines, often a wood, steel, or concrete mast designed to carry minor power lines.
	PowerPole PowerType = "pole"
	// PowerPortal represents a supporting structure for power lines, composed of vertical legs with cables between them attached to a horizontal crossarm.
	PowerPortal PowerType = "portal"
	// PowerSubstation represents a facility which controls the flow of electricity in a power network with transformers, switchgear or compensators.
	PowerSubstation PowerType = "substation"
	// PowerSwitch represents a device which allows electrical network operators to power up & down lines and transformers in substations or along the power grid.
	PowerSwitch PowerType = "switch"
	// PowerTerminal represents a point of connection where an overhead power line ends on a building or wall; for example, when connecting it to an indoor substation.
	PowerTerminal PowerType = "terminal"
	// PowerTower represents a tower or pylon carrying high voltage electricity cables. Often constructed from steel latticework but tubular or solid pylons are also used.
	PowerTower PowerType = "tower"
	// PowerTransformer represents a device for stepping up or down electric voltage. Large power transformers are typically located inside substations.
	PowerTransformer PowerType = "transformer"
)

var validPowerTypes = map[string]PowerType{
	"catenary_mast": PowerCatenaryMast,
	"compensator":   PowerCompensator,
	"connection":    PowerConnection,
	"converter":     PowerConverter,
	"generator":     PowerGenerator,
	"heliostat":     PowerHeliostat,
	"insulator":     PowerInsulator,
	"inverter":      PowerInverter,
	"pole":          PowerPole,
	"portal":        PowerPortal,
	"substation":    PowerSubstation,
	"switch":        PowerSwitch,
	"terminal":      PowerTerminal,
	"tower":         PowerTower,
	"transformer":   PowerTransformer,
}

// Power is used to tag electrical power generation and distributions systems.
// https://wiki.openstreetmap.org/wiki/Power
type Power Feature

func (p *Power) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validPowerTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validPowerTypes))
			for k := range validPowerTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", p.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = p.GetTag()

	*p = Power(f)
	return nil
}

func (p *Power) GetTag() string {
	return "power"
}

func (p *Power) ToOQL() string {
	return Feature(*p).ToOQL()
}

package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// BarrierType represents the valid barrier types.
// https://wiki.openstreetmap.org/wiki/Key:barrier
type BarrierType string

const (
	// Linear barriers

	// BarrierKerb represents a stone edging to a pavement or raised path. The right side is considered the bottom, and the left side is the top.
	BarrierKerb BarrierType = "kerb"

	// Access control on highways

	// BarrierBlock represents one or more large immobile block(s) usually barring free access along a way.
	BarrierBlock BarrierType = "block"
	// BarrierBollard represents one or more solid (usually concrete or metal) pillar(s) used to control traffic.
	BarrierBollard BarrierType = "bollard"
	// BarrierBorderControl represents a control point at an international border between two countries.
	BarrierBorderControl BarrierType = "border_control"
	// BarrierBumpGate represents a drive-through gate used in rural areas to provide a barrier to livestock that does not require the driver to exit the vehicle.
	BarrierBumpGate BarrierType = "bump_gate"
	// BarrierBusTrap represents a short section of the roadway where there is a deep dip in the middle to prevent passage by some traffic.
	BarrierBusTrap BarrierType = "bus_trap"
	// BarrierCattleGrid represents a hole in the road surface covered in a series of bars that allow wheeled vehicles but not animals to cross.
	BarrierCattleGrid BarrierType = "cattle_grid"
	// BarrierChain represents a chain used to prevent motorised vehicles.
	BarrierChain BarrierType = "chain"
	// BarrierCycleBarrier represents a barrier along a path that slows or prevents access for bicycle users.
	BarrierCycleBarrier BarrierType = "cycle_barrier"
	// BarrierDebris represents debris blocking a road.
	BarrierDebris BarrierType = "debris"
	// BarrierEntrance represents an opening or gap in a barrier.
	BarrierEntrance BarrierType = "entrance"
	// BarrierFullHeightTurnstile represents a full-height turnstile.
	BarrierFullHeightTurnstile BarrierType = "full-height_turnstile"
	// BarrierGate represents a section in a wall or fence which can be opened to allow access.
	BarrierGate BarrierType = "gate"
	// BarrierHampshireGate represents a section of wire fence which can be removed temporarily.
	BarrierHampshireGate BarrierType = "hampshire_gate"
	// BarrierHeightRestrictor represents a height restrictor which prevents access of vehicles higher than a set limit.
	BarrierHeightRestrictor BarrierType = "height_restrictor"
	// BarrierHorseStile represents a horse stile allows pedestrians and horses to cross a gap through a fence.
	BarrierHorseStile BarrierType = "horse_stile"
	// BarrierJerseyBarrier represents a barrier made of heavy prefabricated blocks.
	BarrierJerseyBarrier BarrierType = "jersey_barrier"
	// BarrierKissingGate represents a type of gate where you have to go into an enclosure and open a gate to get through.
	BarrierKissingGate BarrierType = "kissing_gate"
	// BarrierLiftGate represents a bar or pole pivoted (rotates upwards to open) in such a way as to allow the boom to block vehicular access through a controlled point.
	BarrierLiftGate BarrierType = "lift_gate"
	// BarrierLog represents a log blocking access.
	BarrierLog BarrierType = "log"
	// BarrierMotorcycleBarrier represents a barrier along a path preventing access by motorcycles.
	BarrierMotorcycleBarrier BarrierType = "motorcycle_barrier"
	// BarrierRope represents a flexible barrier made of rope.
	BarrierRope BarrierType = "rope"
	// BarrierSallyPort represents a secure, controlled entryway to a fortification or prison.
	BarrierSallyPort BarrierType = "sally_port"
	// BarrierSpikes represents spikes on the ground preventing unauthorized access.
	BarrierSpikes BarrierType = "spikes"
	// BarrierStile represents a structure which provides people a passage through or over a boundary via steps, ladders or narrow gaps.
	BarrierStile BarrierType = "stile"
	// BarrierSumpBuster represents a barrier to stop cars (two tracked vehicles with less than a certain ground clearance and width between tracks).
	BarrierSumpBuster BarrierType = "sump_buster"
	// BarrierSwingGate represents a gate consisting of a bar or pole pivoted (rotates sidewards to open) in such a way as to allow the boom to block vehicular access through a controlled point.
	BarrierSwingGate BarrierType = "swing_gate"
	// BarrierTollBooth represents a place where a road usage toll or fee is collected.
	BarrierTollBooth BarrierType = "toll_booth"
	// BarrierTurnstile represents a turnstile, a passage on foot designed to allow one person at a time to pass.
	BarrierTurnstile BarrierType = "turnstile"
	// BarrierYes represents an unspecified barrier. If possible, use a more specific value.
	BarrierYes BarrierType = "yes"
)

var validBarrierTypes = map[string]BarrierType{
	// Linear barriers
	"kerb": BarrierKerb,

	// Access control on highways
	"block":                 BarrierBlock,
	"bollard":               BarrierBollard,
	"border_control":        BarrierBorderControl,
	"bump_gate":             BarrierBumpGate,
	"bus_trap":              BarrierBusTrap,
	"cattle_grid":           BarrierCattleGrid,
	"chain":                 BarrierChain,
	"cycle_barrier":         BarrierCycleBarrier,
	"debris":                BarrierDebris,
	"entrance":              BarrierEntrance,
	"full-height_turnstile": BarrierFullHeightTurnstile,
	"gate":                  BarrierGate,
	"hampshire_gate":        BarrierHampshireGate,
	"height_restrictor":     BarrierHeightRestrictor,
	"horse_stile":           BarrierHorseStile,
	"jersey_barrier":        BarrierJerseyBarrier,
	"kissing_gate":          BarrierKissingGate,
	"lift_gate":             BarrierLiftGate,
	"log":                   BarrierLog,
	"motorcycle_barrier":    BarrierMotorcycleBarrier,
	"rope":                  BarrierRope,
	"sally_port":            BarrierSallyPort,
	"spikes":                BarrierSpikes,
	"stile":                 BarrierStile,
	"sump_buster":           BarrierSumpBuster,
	"swing_gate":            BarrierSwingGate,
	"toll_booth":            BarrierTollBooth,
	"turnstile":             BarrierTurnstile,
	"yes":                   BarrierYes,
}

// Barrier is used to tag barriers and obstacles that are usually involved by traveling.
// https://wiki.openstreetmap.org/wiki/Barriers
type Barrier Feature

func (b *Barrier) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validBarrierTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validBarrierTypes))
			for k := range validBarrierTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", b.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = b.GetTag()

	*b = Barrier(f)
	return nil
}

func (b *Barrier) GetTag() string {
	return "barrier"
}

func (b *Barrier) ToOQL() string {
	return Feature(*b).ToOQL()
}

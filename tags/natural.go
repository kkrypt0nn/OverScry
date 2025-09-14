package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// NaturalType represents the valid natural types.
// https://wiki.openstreetmap.org/wiki/Key:natural
type NaturalType string

const (
	// Vegetation

	// NaturalFell represents habitat above tree line covered with grass, dwarf shrubs and mosses.
	NaturalFell NaturalType = "fell"
	// NaturalScrub represents uncultivated land covered with shrubs, bushes or stunted trees.
	NaturalScrub NaturalType = "scrub"
	// NaturalTree represents a single tree.
	NaturalTree NaturalType = "tree"
	// NaturalWood represents tree-covered area (a 'forest' or 'wood').
	NaturalWood NaturalType = "wood"

	// Water related

	// NaturalBay represents named area of water mostly surrounded by land but with level connection to the ocean or a lake.
	NaturalBay NaturalType = "bay"
	// NaturalBeach represents landform along a body of water which consists of sand, shingle or other loose material.
	NaturalBeach NaturalType = "beach"
	// NaturalBlowhole represents an opening to a sea cave which has grown landwards resulting in blasts of water from the opening due to the wave action.
	NaturalBlowhole NaturalType = "blowhole"
	// NaturalCape represents a piece of elevated land sticking out into the sea or large lake.
	NaturalCape NaturalType = "cape"
	// NaturalGeyser represents a spring characterized by intermittent discharge of water ejected turbulently and accompanied by steam.
	NaturalGeyser NaturalType = "geyser"
	// NaturalGlacier represents a permanent body of ice formed naturally from snow that is moving under its own weight.
	NaturalGlacier NaturalType = "glacier"
	// NaturalHotSpring represents a spring of geothermally heated groundwater.
	NaturalHotSpring NaturalType = "hot_spring"
	// NaturalIsthmus represents a narrow strip of land, bordered by water on both sides and connecting two larger land masses.
	NaturalIsthmus NaturalType = "isthmus"
	// NaturalPeninsula represents a piece of land projecting into water from a larger land mass, nearly surrounded by water.
	NaturalPeninsula NaturalType = "peninsula"
	// NaturalReef represents a feature (rock, sandbar, coral, etc) lying permanently beneath the surface of the water.
	NaturalReef NaturalType = "reef"
	// NaturalSpring represents a place where ground water flows naturally from the ground.
	NaturalSpring NaturalType = "spring"
	// NaturalStrait represents a narrow area of water surrounded by land on two sides and by water on two other sides.
	NaturalStrait NaturalType = "strait"
	// NaturalWater represents any inland body of water, from natural such as a lake or pond to artificial like a moat or canal.
	NaturalWater NaturalType = "water"

	// Geology related

	// NaturalArch represents a rock arch naturally formed by erosion, with an opening underneath.
	NaturalArch NaturalType = "arch"
	// NaturalCaveEntrance represents an entrance to a cave: a natural underground space large enough for a human to enter.
	NaturalCaveEntrance NaturalType = "cave_entrance"
	// NaturalCliff represents a vertical or almost vertical natural drop in terrain, usually with a bare rock surface. The bottom of the cliff is on the right side of the way.
	NaturalCliff NaturalType = "cliff"
	// NaturalDune represents a hill of sand formed by wind, covered with no or very little vegetation.
	NaturalDune NaturalType = "dune"
	// NaturalFumarole represents a fumarole is an opening in a planet's crust, which emits steam and gases.
	NaturalFumarole NaturalType = "fumarole"
	// NaturalHill represents a hill is a landform that is elevated above the surrounding terrain, but is smaller than a mountain.
	NaturalHill NaturalType = "hill"
	// NaturalPeak represents the top (summit) of a hill or mountain.
	NaturalPeak NaturalType = "peak"
	// NaturalRock represents a notable rock or group of rocks attached to the underlying bedrock.
	NaturalRock NaturalType = "rock"
	// NaturalSaddle represents the lowest point along a ridge or between two mountain tops.
	NaturalSaddle NaturalType = "saddle"
	// NaturalSinkhole represents a natural depression or hole in the surface topography.
	NaturalSinkhole NaturalType = "sinkhole"
	// NaturalStone represents a single notable freestanding rock, which may differ from the composition of the terrain it lies in.
	NaturalStone NaturalType = "stone"
	// NaturalValley represents a natural depression flanked by ridges or ranges of mountains or hills.
	NaturalValley NaturalType = "valley"
	// NaturalVolcano represents an opening exposed on the earth's surface where volcanic material is emitted.
	NaturalVolcano NaturalType = "volcano"
)

var validNaturalTypes = map[string]NaturalType{
	// Vegetation
	"fell":  NaturalFell,
	"scrub": NaturalScrub,
	"tree":  NaturalTree,
	"wood":  NaturalWood,

	// Water related
	"bay":        NaturalBay,
	"beach":      NaturalBeach,
	"blowhole":   NaturalBlowhole,
	"cape":       NaturalCape,
	"geyser":     NaturalGeyser,
	"glacier":    NaturalGlacier,
	"hot_spring": NaturalHotSpring,
	"isthmus":    NaturalIsthmus,
	"peninsula":  NaturalPeninsula,
	"reef":       NaturalReef,
	"spring":     NaturalSpring,
	"strait":     NaturalStrait,
	"water":      NaturalWater,

	// Geology related
	"arch":          NaturalArch,
	"cave_entrance": NaturalCaveEntrance,
	"cliff":         NaturalCliff,
	"dune":          NaturalDune,
	"fumarole":      NaturalFumarole,
	"hill":          NaturalHill,
	"peak":          NaturalPeak,
	"rock":          NaturalRock,
	"saddle":        NaturalSaddle,
	"sinkhole":      NaturalSinkhole,
	"stone":         NaturalStone,
	"valley":        NaturalValley,
	"volcano":       NaturalVolcano,
}

// Natural is used to describe natural and physical land features.
// https://wiki.openstreetmap.org/wiki/Key:natural
type Natural Feature

func (n *Natural) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validNaturalTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validNaturalTypes))
			for k := range validNaturalTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", n.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = n.GetTag()

	*n = Natural(f)
	return nil
}

func (n *Natural) GetTag() string {
	return "natural"
}

func (n *Natural) ToOQL() string {
	return Feature(*n).ToOQL()
}

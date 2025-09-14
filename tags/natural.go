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
	// NaturalGrassland represents an area where the vegetation is dominated by grasses and other herbaceous (non-woody) plants.
	NaturalGrassland NaturalType = "grassland"
	// NaturalHeath represents a dwarf-shrub habitat, characterised by open, low growing woody vegetation, often dominated by plants of the Ericaceae.
	NaturalHeath NaturalType = "heath"
	// NaturalMoor represents don't use, see wikipage.
	NaturalMoor NaturalType = "moor"
	// NaturalScrub represents uncultivated land covered with shrubs, bushes or stunted trees.
	NaturalScrub NaturalType = "scrub"
	// NaturalShrubbery represents an area of woody shrubbery that is actively maintained or pruned by humans. A slightly wilder look is also possible.
	NaturalShrubbery NaturalType = "shrubbery"
	// NaturalTree represents a single tree.
	NaturalTree NaturalType = "tree"
	// NaturalTreeRow represents a line of trees.
	NaturalTreeRow NaturalType = "tree_row"
	// NaturalTundra represents treeless cold climate habitat with open, low growing sedges, grasses, mosses and lichens.
	NaturalTundra NaturalType = "tundra"
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
	// NaturalCoastline represents the line between sea and land indicating mean high water (springs): the water lies on the right side of the way.
	NaturalCoastline NaturalType = "coastline"
	// NaturalCrevasse represents a large crack in a glacier.
	NaturalCrevasse NaturalType = "crevasse"
	// NaturalGeyser represents a spring characterized by intermittent discharge of water ejected turbulently and accompanied by steam.
	NaturalGeyser NaturalType = "geyser"
	// NaturalGlacier represents a permanent body of ice formed naturally from snow that is moving under its own weight.
	NaturalGlacier NaturalType = "glacier"
	// NaturalHotSpring represents a spring of geothermally heated groundwater.
	NaturalHotSpring NaturalType = "hot_spring"
	// NaturalIsthmus represents a narrow strip of land, bordered by water on both sides and connecting two larger land masses.
	NaturalIsthmus NaturalType = "isthmus"
	// NaturalMud represents area covered with mud: water saturated fine grained soil without significant plant growth.
	NaturalMud NaturalType = "mud"
	// NaturalPeninsula represents a piece of land projecting into water from a larger land mass, nearly surrounded by water.
	NaturalPeninsula NaturalType = "peninsula"
	// NaturalReef represents a feature (rock, sandbar, coral, etc) lying permanently beneath the surface of the water.
	NaturalReef NaturalType = "reef"
	// NaturalShingle represents an accumulation of rounded rock fragments on a beach or riverbed.
	NaturalShingle NaturalType = "shingle"
	// NaturalShoal represents an area of the water floor which nears the water surface and is exposed at low tide or when a river/lake is not full of water.
	NaturalShoal NaturalType = "shoal"
	// NaturalSpring represents a place where ground water flows naturally from the ground.
	NaturalSpring NaturalType = "spring"
	// NaturalStrait represents a narrow area of water surrounded by land on two sides and by water on two other sides.
	NaturalStrait NaturalType = "strait"
	// NaturalWater represents any inland body of water, from natural such as a lake or pond to artificial like a moat or canal.
	NaturalWater NaturalType = "water"
	// NaturalWetland represents a natural area subject to inundation or with waterlogged ground.
	NaturalWetland NaturalType = "wetland"

	// Geology related

	// NaturalArch represents a rock arch naturally formed by erosion, with an opening underneath.
	NaturalArch NaturalType = "arch"
	// NaturalArete represents a thin, almost knife-like, ridge of rock which is typically formed when two glaciers erode parallel U-shaped valleys.
	NaturalArete NaturalType = "arete"
	// NaturalBareRock represents an area with sparse or no vegetation, so that the bedrock becomes visible.
	NaturalBareRock NaturalType = "bare_rock"
	// NaturalBlockfield represents an area covered by boulder- or block-sized rocks.
	NaturalBlockfield NaturalType = "blockfield"
	// NaturalCaveEntrance represents an entrance to a cave: a natural underground space large enough for a human to enter.
	NaturalCaveEntrance NaturalType = "cave_entrance"
	// NaturalCliff represents a vertical or almost vertical natural drop in terrain, usually with a bare rock surface. The bottom of the cliff is on the right side of the way.
	NaturalCliff NaturalType = "cliff"
	// NaturalDune represents a hill of sand formed by wind, covered with no or very little vegetation.
	NaturalDune NaturalType = "dune"
	// NaturalEarthBank represents large erosion gully or steep earth bank.
	NaturalEarthBank NaturalType = "earth_bank"
	// NaturalFumarole represents a fumarole is an opening in a planet's crust, which emits steam and gases.
	NaturalFumarole NaturalType = "fumarole"
	// NaturalHill represents a hill is a landform that is elevated above the surrounding terrain, but is smaller than a mountain.
	NaturalHill NaturalType = "hill"
	// NaturalPeak represents the top (summit) of a hill or mountain.
	NaturalPeak NaturalType = "peak"
	// NaturalRidge represents a mountain landform with a continuous elevated crest or linear hill.
	NaturalRidge NaturalType = "ridge"
	// NaturalRock represents a notable rock or group of rocks attached to the underlying bedrock.
	NaturalRock NaturalType = "rock"
	// NaturalSaddle represents the lowest point along a ridge or between two mountain tops.
	NaturalSaddle NaturalType = "saddle"
	// NaturalSand represents an area covered by sand with no or very little vegetation.
	NaturalSand NaturalType = "sand"
	// NaturalScree represents unconsolidated angular stones formed by rockfall and weathering from adjacent rock faces.
	NaturalScree NaturalType = "scree"
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
	"fell":      NaturalFell,
	"grassland": NaturalGrassland,
	"heath":     NaturalHeath,
	"moor":      NaturalMoor,
	"scrub":     NaturalScrub,
	"shrubbery": NaturalShrubbery,
	"tree":      NaturalTree,
	"tree_row":  NaturalTreeRow,
	"tundra":    NaturalTundra,
	"wood":      NaturalWood,

	// Water related
	"bay":        NaturalBay,
	"beach":      NaturalBeach,
	"blowhole":   NaturalBlowhole,
	"cape":       NaturalCape,
	"coastline":  NaturalCoastline,
	"crevasse":   NaturalCrevasse,
	"geyser":     NaturalGeyser,
	"glacier":    NaturalGlacier,
	"hot_spring": NaturalHotSpring,
	"isthmus":    NaturalIsthmus,
	"mud":        NaturalMud,
	"peninsula":  NaturalPeninsula,
	"reef":       NaturalReef,
	"shingle":    NaturalShingle,
	"shoal":      NaturalShoal,
	"spring":     NaturalSpring,
	"strait":     NaturalStrait,
	"water":      NaturalWater,
	"wetland":    NaturalWetland,

	// Geology related
	"arch":          NaturalArch,
	"arete":         NaturalArete,
	"bare_rock":     NaturalBareRock,
	"blockfield":    NaturalBlockfield,
	"cave_entrance": NaturalCaveEntrance,
	"cliff":         NaturalCliff,
	"dune":          NaturalDune,
	"earth_bank":    NaturalEarthBank,
	"fumarole":      NaturalFumarole,
	"hill":          NaturalHill,
	"peak":          NaturalPeak,
	"ridge":         NaturalRidge,
	"rock":          NaturalRock,
	"saddle":        NaturalSaddle,
	"sand":          NaturalSand,
	"scree":         NaturalScree,
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

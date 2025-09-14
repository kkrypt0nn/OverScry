package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// GeologicalType represents the valid geological types.
// https://wiki.openstreetmap.org/wiki/Key:geological
type GeologicalType string

const (
	// GeologicalMoraine represents any accumulation of unconsolidated rock debris previously carried by a glacier.
	GeologicalMoraine GeologicalType = "moraine"
	// GeologicalOutcrop represents a place where the bedrock or superficial deposits previously covered under the soil have become locally exposed.
	GeologicalOutcrop GeologicalType = "outcrop"
	// GeologicalFold represents planar surfaces, such as sedimentary strata, that are bent or curved ("folded").
	GeologicalFold GeologicalType = "fold"
	// GeologicalPalaeontologicalSite represents a place with fossils.
	GeologicalPalaeontologicalSite GeologicalType = "palaeontological_site"
	// GeologicalVolcanicLavaField represents an area with volcanic lava on the ground.
	GeologicalVolcanicLavaField GeologicalType = "volcanic_lava_field"
	// GeologicalVolcanicVent represents a hole through which the lava erupts.
	GeologicalVolcanicVent GeologicalType = "volcanic_vent"
	// GeologicalGlacialErratic represents a boulder deposited by a glacier.
	GeologicalGlacialErratic GeologicalType = "glacial_erratic"
	// GeologicalRockGlacier represents rock glaciers are mixtures of rock and ice that move slowly downhill when active.
	GeologicalRockGlacier GeologicalType = "rock_glacier"
	// GeologicalGiantsKettle represents a regular hole in a rock created by the rotation of stones in the bed of a stream.
	GeologicalGiantsKettle GeologicalType = "giants_kettle"
	// GeologicalMeteorCrater represents a crater formed by the impact of a meteor.
	GeologicalMeteorCrater GeologicalType = "meteor_crater"
	// GeologicalHoodoo represents a column of rock with a hat.
	GeologicalHoodoo GeologicalType = "hoodoo"
	// GeologicalColumnarJointing represents several hexagonal columns of rock.
	GeologicalColumnarJointing GeologicalType = "columnar_jointing"
	// GeologicalDyke represents a sheet of rock that fills a fracture of a pre-existing rock body or when standing alone looking like a man made wall.
	GeologicalDyke GeologicalType = "dyke"
	// GeologicalMonocline represents step-like fold in rock strata consisting of a zone of a dip comprised between ~65° and 90°.
	GeologicalMonocline GeologicalType = "monocline"
	// GeologicalTor represents a large, free-standing rock outcrop that rises abruptly.
	GeologicalTor GeologicalType = "tor"
	// GeologicalUnconformity represents a buried erosional surface separating two rock masses or strata of different ages.
	GeologicalUnconformity GeologicalType = "unconformity"
	// GeologicalCone represents a landform with a distinctly conical shape.
	GeologicalCone GeologicalType = "cone"
	// GeologicalSinkhole represents a depression or hole in the ground caused by some form of collapse of the surface.
	GeologicalSinkhole GeologicalType = "sinkhole"
	// GeologicalPingo represents pingos are intrapermafrost ice-cored hills.
	GeologicalPingo GeologicalType = "pingo"
	// GeologicalInselberg represents inselberg are isolated rock hill, knob, ridge, or small mountain.
	GeologicalInselberg GeologicalType = "inselberg"
	// GeologicalLimestonePavement represents a natural karst landform consisting of a flat surface of limestone.
	GeologicalLimestonePavement GeologicalType = "limestone_pavement"
)

var validGeologicalTypes = map[string]GeologicalType{
	"moraine":               GeologicalMoraine,
	"outcrop":               GeologicalOutcrop,
	"fold":                  GeologicalFold,
	"palaeontological_site": GeologicalPalaeontologicalSite,
	"volcanic_lava_field":   GeologicalVolcanicLavaField,
	"volcanic_vent":         GeologicalVolcanicVent,
	"glacial_erratic":       GeologicalGlacialErratic,
	"rock_glacier":          GeologicalRockGlacier,
	"giants_kettle":         GeologicalGiantsKettle,
	"meteor_crater":         GeologicalMeteorCrater,
	"hoodoo":                GeologicalHoodoo,
	"columnar_jointing":     GeologicalColumnarJointing,
	"dyke":                  GeologicalDyke,
	"monocline":             GeologicalMonocline,
	"tor":                   GeologicalTor,
	"unconformity":          GeologicalUnconformity,
	"cone":                  GeologicalCone,
	"sinkhole":              GeologicalSinkhole,
	"pingo":                 GeologicalPingo,
	"inselberg":             GeologicalInselberg,
	"limestone_pavement":    GeologicalLimestonePavement,
}

// Geological is used to tag the geological makeup of an area.
// https://wiki.openstreetmap.org/wiki/Key:geological
type Geological Feature

func (g *Geological) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validGeologicalTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validGeologicalTypes))
			for k := range validGeologicalTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", g.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = g.GetTag()

	*g = Geological(f)
	return nil
}

func (g *Geological) GetTag() string {
	return "geological"
}

func (g *Geological) ToOQL() string {
	return Feature(*g).ToOQL()
}

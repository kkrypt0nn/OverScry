package tags

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
	"github.com/kkrypt0nn/overscry/tags/validators"
)

// PlaceType represents the valid place types.
// https://wiki.openstreetmap.org/wiki/Key:place
type PlaceType string

const (
	// Administratively declared places

	// PlaceCountry represents a nation state or other high-level national political/administrative area.
	PlaceCountry PlaceType = "country"
	// PlaceState represents a large sub-national political/administrative area.
	PlaceState PlaceType = "state"
	// PlaceRegion represents used both as a broad tag for geographic or historical areas with no clear boundary and for distinct administration areas (with specific boundaries) in some countries.
	PlaceRegion PlaceType = "region"
	// PlaceProvince represents a subdivision of a country similar to a state.
	PlaceProvince PlaceType = "province"
	// PlaceDistrict represents a district – a type of administrative division that, in some countries, is managed by local government.
	PlaceDistrict PlaceType = "district"
	// PlaceCounty represents a county - a geographical region of a country.
	PlaceCounty PlaceType = "county"
	// PlaceSubdistrict represents a subdistrict - a subdivision of a district used for administrative or other purposes.
	PlaceSubdistrict PlaceType = "subdistrict"
	// PlaceMunicipality represents a municipality - single urban administrative division having corporate status.
	PlaceMunicipality PlaceType = "municipality"

	// Populated settlements, urban

	// PlaceCity represents the largest urban settlement or settlements within the territory.
	PlaceCity PlaceType = "city"
	// PlaceBorough represents a part in a larger city grouped into administrative unit.
	PlaceBorough PlaceType = "borough"
	// PlaceSuburb represents a part of a town or city with a well-known name and often a distinct identity.
	PlaceSuburb PlaceType = "suburb"
	// PlaceQuarter represents a quarter is a named, geographically localised place within a suburb of a larger city or within a town, which is bigger than a neighbourhood.
	PlaceQuarter PlaceType = "quarter"
	// PlaceNeighbourhood represents a neighbourhood is a smaller named, geographically localised place within a suburb of a larger city or within a town or village.
	PlaceNeighbourhood PlaceType = "neighbourhood"
	// PlaceCityBlock represents a named city block, usually surrounded by streets.
	PlaceCityBlock PlaceType = "city_block"
	// PlacePlot represents a named plot is a tract or parcel of land owned or meant to be owned by some owner.
	PlacePlot PlaceType = "plot"

	// Populated settlements, urban and rural

	// PlaceTown represents an important urban centre, between a village and a city in size.
	PlaceTown PlaceType = "town"
	// PlaceVillage represents a smaller distinct settlement, smaller than a town with few facilities available with people traveling to nearby towns to access these.
	PlaceVillage PlaceType = "village"
	// PlaceHamlet represents a smaller rural community, typically with fewer than 100-200 inhabitants, and little infrastructure.
	PlaceHamlet PlaceType = "hamlet"
	// PlaceIsolatedDwelling represents the smallest kind of settlement (1-2 households).
	PlaceIsolatedDwelling PlaceType = "isolated_dwelling"
	// PlaceFarm represents an individually named farm.
	PlaceFarm PlaceType = "farm"
	// PlaceAllotments represents a separate settlement, which is located outside an officially inhabited locality and has its own addressing.
	PlaceAllotments PlaceType = "allotments"

	// Other places

	// PlaceContinent represents one of the seven continents: Africa, Antarctica, Asia, Europe, North America, Oceania, South America.
	PlaceContinent PlaceType = "continent"
	// PlaceIsland represents any piece of land that is completely surrounded by water and isolated from other significant landmasses (bigger than 1 km²).
	PlaceIsland PlaceType = "island"
	// PlaceIslet represents a very small island (smaller than 1 km²).
	PlaceIslet PlaceType = "islet"
	// PlaceSquare represents a town or village square: a (typically) paved open space, generally of architectural significance, which is surrounded by buildings in a built-up area such as a city, town or village.
	PlaceSquare PlaceType = "square"
	// PlaceLocality represents a named place that has no population.
	PlaceLocality PlaceType = "locality"
	// PlacePolder represents a land area that forms an artificial hydrological entity enclosed by embankments and usually is under sea level.
	PlacePolder PlaceType = "polder"
	// PlaceSea represents a large body of salt water part of, or connected to, an ocean.
	PlaceSea PlaceType = "sea"
	// PlaceOcean represents the world's five main major oceanic divisions.
	PlaceOcean PlaceType = "ocean"
)

var validPlaceTypes = map[string]PlaceType{
	// Administratively declared places
	"country":      PlaceCountry,
	"state":        PlaceState,
	"region":       PlaceRegion,
	"province":     PlaceProvince,
	"district":     PlaceDistrict,
	"county":       PlaceCounty,
	"subdistrict":  PlaceSubdistrict,
	"municipality": PlaceMunicipality,

	// Populated settlements, urban
	"city":          PlaceCity,
	"borough":       PlaceBorough,
	"suburb":        PlaceSuburb,
	"quarter":       PlaceQuarter,
	"neighbourhood": PlaceNeighbourhood,
	"city_block":    PlaceCityBlock,
	"plot":          PlacePlot,

	// Populated settlements, urban and rural
	"town":              PlaceTown,
	"village":           PlaceVillage,
	"hamlet":            PlaceHamlet,
	"isolated_dwelling": PlaceIsolatedDwelling,
	"farm":              PlaceFarm,
	"allotments":        PlaceAllotments,

	// Other places
	"continent": PlaceContinent,
	"island":    PlaceIsland,
	"islet":     PlaceIslet,
	"square":    PlaceSquare,
	"locality":  PlaceLocality,
	"polder":    PlacePolder,
	"sea":       PlaceSea,
	"ocean":     PlaceOcean,
}

// Place is used to mainly to give details about settlements.
// https://wiki.openstreetmap.org/wiki/Places
type Place struct {
	Value  *PlaceType  `yaml:"value,omitempty"`
	Match  MatchMethod `yaml:"match,omitempty"`
	Inline *Feature    `yaml:"-"`

	// Population indicates a rough number of citizens in a given place.
	Population *Feature `yaml:"population,omitempty"`
	// IsIn is used to index where a place or feature is.
	IsIn *Feature `yaml:"is_in,omitempty"`
}

func (p *Place) UnmarshalYAML(unmarshal func(any) error) error {
	temp := &struct {
		Value *string     `yaml:"value,omitempty"`
		Match MatchMethod `yaml:"match,omitempty"`

		Population *Feature `yaml:"population,omitempty"`
		IsIn       *Feature `yaml:"is_in,omitempty"`
	}{}
	if err := unmarshal(temp); err != nil {
		return err
	}

	if temp.Value != nil {
		val := strings.ToLower(*temp.Value)
		if _, ok := validPlaceTypes[val]; ok {
			p.Value = (*PlaceType)(helpers.Ptr(val))
		} else {
			validKeys := make([]string, 0, len(validPlaceTypes))
			for k := range validPlaceTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", p.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	p.Inline = &Feature{
		Match: temp.Match,
		Value: helpers.Ptr(*temp.Value),
		Tag:   p.GetTag(),
	}

	tempVal := reflect.ValueOf(temp).Elem()
	targetVal := reflect.ValueOf(p).Elem()
	for i := 0; i < tempVal.NumField(); i++ {
		field := tempVal.Type().Field(i)
		if field.Name == "Value" || field.Name == "Match" {
			continue
		}

		tempField := tempVal.FieldByName(field.Name)
		if tempField.IsValid() && !tempField.IsNil() {
			feature := tempField.Interface().(*Feature)
			tempFeature := &Feature{
				Value:  feature.Value,
				Match:  feature.Match,
				Tag:    strings.ToLower(field.Name),
				Prefix: helpers.Ptr(p.GetTag()),
			}
			if validator, ok := validators.PlaceValidators[tempFeature.Tag]; ok {
				tempFeature.Validator = validator
				if err := tempFeature.Validate(); err != nil {
					return err
				}
			}
			targetField := targetVal.FieldByName(field.Name)
			targetField.Set(reflect.ValueOf(tempFeature))
		}
	}

	return nil
}

func (p *Place) GetTag() string {
	return "place"
}

func (p *Place) ToOQL() string {
	var queryParts []string

	val := reflect.ValueOf(p).Elem()
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)

		if field.Name == "Value" || field.Name == "Match" {
			continue
		}

		if fieldVal.IsValid() && !fieldVal.IsNil() {
			queryParts = append(queryParts, fieldVal.Interface().(*Feature).ToOQL())
		}
	}

	return strings.Join(queryParts, "")
}

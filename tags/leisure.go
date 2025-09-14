package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// LeisureType represents the valid leisure types.
// https://wiki.openstreetmap.org/wiki/Key:leisure
type LeisureType string

const (
	// LeisureAdultGamingCentre represents a venue with gambling machines, usually with a minimum age requirement.
	LeisureAdultGamingCentre LeisureType = "adult_gaming_centre"
	// LeisureAmusementArcade represents a venue with pay-to-play games.
	LeisureAmusementArcade LeisureType = "amusement_arcade"
	// LeisureBeachResort represents a managed beach, including within the boundary any associated facilities. Entry may also require payment of a fee.
	LeisureBeachResort LeisureType = "beach_resort"
	// LeisureBandstand represents a bandstand is an open structure where musical bands can perform concerts.
	LeisureBandstand LeisureType = "bandstand"
	// LeisureBirdHide represents a place that is used to observe wildlife, especially birds.
	LeisureBirdHide LeisureType = "bird_hide"
	// LeisureCommon represents identify land over which the public has general rights of use for certain leisure activities.
	LeisureCommon LeisureType = "common"
	// LeisureDance represents a dance venue or dance hall.
	LeisureDance LeisureType = "dance"
	// LeisureDiscGolfCourse represents a place to play disc golf.
	LeisureDiscGolfCourse LeisureType = "disc_golf_course"
	// LeisureDogPark represents a designated area, with or without a fenced boundary, where dog-owners are permitted to exercise their pets unrestrained.
	LeisureDogPark LeisureType = "dog_park"
	// LeisureEscapeGame represents a physical adventure game in which players solve a series of puzzles using clues, hints and strategy to complete the objectives at hand.
	LeisureEscapeGame LeisureType = "escape_game"
	// LeisureFirepit represents a fire ring or fire pit, often at a campsite or picnic site.
	LeisureFirepit LeisureType = "firepit"
	// LeisureFishing represents a public or private place for fishing.
	LeisureFishing LeisureType = "fishing"
	// LeisureFitnessCentre represents fitness centre, health club or gym with exercise machines, fitness classes or both, for exercise.
	LeisureFitnessCentre LeisureType = "fitness_centre"
	// LeisureFitnessStation represents an outdoor facility where people can practise typical fitness exercises.
	LeisureFitnessStation LeisureType = "fitness_station"
	// LeisureGarden represents a place where flowers and other plants are grown in a decorative and structured manner or for scientific purposes.
	LeisureGarden LeisureType = "garden"
	// LeisureHackerspace represents a place where people with common interests (science, technology, ...) meet.
	LeisureHackerspace LeisureType = "hackerspace"
	// LeisureHorseRiding represents a facility where people practise horse riding, usually in their spare time, e.g. a riding centre.
	LeisureHorseRiding LeisureType = "horse_riding"
	// LeisureIceRink represents a place where you can skate and play bandy or ice hockey.
	LeisureIceRink LeisureType = "ice_rink"
	// LeisureMarina represents a facility for mooring leisure yachts and motor boats.
	LeisureMarina LeisureType = "marina"
	// LeisureMiniatureGolf represents a place or area where you can play miniature golf.
	LeisureMiniatureGolf LeisureType = "miniature_golf"
	// LeisureNatureReserve represents a protected area of importance for wildlife, flora, fauna or features of geological or other special interest.
	LeisureNatureReserve LeisureType = "nature_reserve"
	// LeisurePark represents a park, usually in an urban (municipal) setting, created for recreation and relaxation.
	LeisurePark LeisureType = "park"
	// LeisurePicnicTable represents a table with benches for food and rest.
	LeisurePicnicTable LeisureType = "picnic_table"
	// LeisurePitch represents an area designed for practising a particular sport, normally designated with appropriate markings.
	LeisurePitch LeisureType = "pitch"
	// LeisurePlayground represents a playground: an area designed for children to play.
	LeisurePlayground LeisureType = "playground"
	// LeisureSlipway represents a slipway: a ramp for launching a boat into water.
	LeisureSlipway LeisureType = "slipway"
	// LeisureSportsCentre represents a sports centre is a distinct facility where sports take place within an enclosed area.
	LeisureSportsCentre LeisureType = "sports_centre"
	// LeisureStadium represents a major sports facility with substantial tiered seating.
	LeisureStadium LeisureType = "stadium"
	// LeisureSummerCamp represents a place for supervised camps for children or teenagers conducted during the summer months.
	LeisureSummerCamp LeisureType = "summer_camp"
	// LeisureSwimmingArea represents an area for swimming within a larger body of water (such as a river, lake or the sea) that is marked by a rope, buoys or similar.
	LeisureSwimmingArea LeisureType = "swimming_area"
	// LeisureSwimmingPool represents a swimming pool (water area only).
	LeisureSwimmingPool LeisureType = "swimming_pool"
	// LeisureTrack represents a track for running, cycling and other non-motorised racing such as horses, greyhounds.
	LeisureTrack LeisureType = "track"
	// LeisureWaterPark represents an amusement park with features like water slides, recreational pools (e.g. wave pools) or lazy rivers.
	LeisureWaterPark LeisureType = "water_park"
)

var validLeisureTypes = map[string]LeisureType{
	"adult_gaming_centre": LeisureAdultGamingCentre,
	"amusement_arcade":    LeisureAmusementArcade,
	"beach_resort":        LeisureBeachResort,
	"bandstand":           LeisureBandstand,
	"bird_hide":           LeisureBirdHide,
	"common":              LeisureCommon,
	"dance":               LeisureDance,
	"disc_golf_course":    LeisureDiscGolfCourse,
	"dog_park":            LeisureDogPark,
	"escape_game":         LeisureEscapeGame,
	"firepit":             LeisureFirepit,
	"fishing":             LeisureFishing,
	"fitness_centre":      LeisureFitnessCentre,
	"fitness_station":     LeisureFitnessStation,
	"garden":              LeisureGarden,
	"hackerspace":         LeisureHackerspace,
	"horse_riding":        LeisureHorseRiding,
	"ice_rink":            LeisureIceRink,
	"marina":              LeisureMarina,
	"miniature_golf":      LeisureMiniatureGolf,
	"nature_reserve":      LeisureNatureReserve,
	"park":                LeisurePark,
	"picnic_table":        LeisurePicnicTable,
	"pitch":               LeisurePitch,
	"playground":          LeisurePlayground,
	"slipway":             LeisureSlipway,
	"sports_centre":       LeisureSportsCentre,
	"stadium":             LeisureStadium,
	"summer_camp":         LeisureSummerCamp,
	"swimming_area":       LeisureSwimmingArea,
	"swimming_pool":       LeisureSwimmingPool,
	"track":               LeisureTrack,
	"water_park":          LeisureWaterPark,
}

// Leisure is used to tag leisure and sports facilities.
// https://wiki.openstreetmap.org/wiki/Leisure
type Leisure Feature

func (l *Leisure) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validLeisureTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validLeisureTypes))
			for k := range validLeisureTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", l.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = l.GetTag()

	*l = Leisure(f)
	return nil
}

func (l *Leisure) GetTag() string {
	return "leisure"
}

func (l *Leisure) ToOQL() string {
	return Feature(*l).ToOQL()
}

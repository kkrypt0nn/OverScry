package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// LanduseType represents the valid landuse types.
// https://wiki.openstreetmap.org/wiki/Key:landuse
type LanduseType string

const (
	// Developed land

	// LanduseCommercial represents predominantly commercial businesses and their offices. Commercial businesses which sell goods should be categorised as [LanduseRetail]. Commercial businesses can sell services on site and may include private Doctor's Surgeries, and those non-government services for mental and physical health, such as a Counselor's or Physiotherapist's practice or Veterinary. Commercial businesses can also include office buildings and business parks which have limited interface with the public and sell their services either on site, or externally. Commercial businesses have low amounts of public foot traffic.
	LanduseCommercial LanduseType = "commercial"
	// LanduseConstruction represents a site which is under active development and construction of a building or structure, including any purposeful alteration to the land or vegetation upon it. Abandoned construction projects and sites should not use this tag.
	LanduseConstruction LanduseType = "construction"
	// LanduseEducation represents an area predominately used for educational purposes/facilities.
	LanduseEducation LanduseType = "education"
	// LanduseFairground represents a site where a [fair] takes place.
	//
	// [fair]: https://en.wikipedia.org/wiki/Fair
	LanduseFairground LanduseType = "fairground"
	// LanduseIndustrial represents predominantly industrial landuses such as workshops, factories, or warehouses.
	LanduseIndustrial LanduseType = "industrial"
	// LanduseResidential represents land where people reside; predominantly residential detached (single houses, grouped dwellings), or attached (apartments, flats, units) dwellings. For "Mixed-Use" areas where more than half of the land is residential, tag as residential.
	LanduseResidential LanduseType = "residential"
	// LanduseRetail represents predominantly retail businesses such as shops. Retail businesses sell physical goods such as food (prepared or grocery), clothing, medicine, stationary, appliances, tools, or other similar physical items. Retail businesses have high amounts of public foot traffic. Retail businesses do not exclusively provide or sell their services. For businesses which sell services see [LanduseCommercial].
	LanduseRetail LanduseType = "retail"
	// LanduseInstitutional represents land used for institutional purposes, see [Institution (disambiguation)].
	//
	// [Institution (disambiguation)]: https://en.wikipedia.org/wiki/Institution_(disambiguation)
	LanduseInstitutional LanduseType = "institutional"

	// Rural and agricultural land

	// LanduseAquaculture represents a [piece of land] for the farming of freshwater and saltwater organisms such as finfish, molluscs, crustaceans and aquatic plants.
	//
	// [piece of land]: https://en.wikipedia.org/wiki/Aquaculture
	LanduseAquaculture LanduseType = "aquaculture"
	// LanduseAllotments represents a piece of land given over to local residents for growing vegetables and flowers.
	LanduseAllotments LanduseType = "allotments"
	// LanduseForest represents managed forest or woodland plantation. Some use this to map an area of trees rather than the use of the land.
	LanduseForest LanduseType = "forest"
	// LanduseMeadow represents a meadow or pasture: land primarily vegetated by grass and non-woody plants, mainly used for hay or grazing.
	LanduseMeadow LanduseType = "meadow"
	// LanduseOrchard represents intentional planting of trees or shrubs maintained for food production.
	LanduseOrchard LanduseType = "orchard"
	// LanduseVineyard represents a piece of land where grapes are grown.
	LanduseVineyard LanduseType = "vineyard"

	// Waterbody

	// LanduseBasin represents an area artificially graded to hold water.
	LanduseBasin LanduseType = "basin"
	// LanduseReservoir represents a [reservoir] on Wikipedia.
	//
	// [reservoir]: https://en.wikipedia.org/wiki/Reservoir
	LanduseReservoir LanduseType = "reservoir"

	// Other

	// LanduseBrownfield represents describes land scheduled for new development where old buildings have been demolished and cleared.
	LanduseBrownfield LanduseType = "brownfield"
	// LanduseCemetery represents place for burials. Smaller places (e.g. with a church nearby) may use [AmenityGraveYard] instead.
	LanduseCemetery LanduseType = "cemetery"
	// LanduseGrass represents an area of mown and managed grass not otherwise covered by a more specific tag.
	LanduseGrass LanduseType = "grass"
	// LanduseGreenfield represents describes land scheduled for new development where there have been no buildings before. A greenfield is scheduled to turn into a construction site.
	LanduseGreenfield LanduseType = "greenfield"
	// LanduseLandfill represents place where waste is dumped.
	LanduseLandfill LanduseType = "landfill"
	// LanduseMilitary represents for land areas owned/used by the military for whatever purpose.
	LanduseMilitary LanduseType = "military"
	// LanduseQuarry represents surface mineral extraction.
	LanduseQuarry LanduseType = "quarry"
	// LanduseRecreationGround represents an open green space for general recreation, which may include pitches, nets and so on, usually municipal but possibly also private to colleges or companies.
	LanduseRecreationGround LanduseType = "recreation_ground"
	// LanduseReligious represents an area used for religious purposes.
	LanduseReligious LanduseType = "religious"
	// LanduseVillageGreen represents a village green is a distinctive area of grassy public land in a village centre. Not a generic tag for urban greenery. It is a typical English term – defined separately from 'common land' under the Commons Registration Act 1965 and the Commons Act 2006.
	LanduseVillageGreen LanduseType = "village_green"
	// LanduseWinterSports represents an area dedicated to winter sports (e.g. skiing).
	LanduseWinterSports LanduseType = "winter_sports"
)

var validLanduseTypes = map[string]LanduseType{
	// Developed land
	"commercial":    LanduseCommercial,
	"construction":  LanduseConstruction,
	"education":     LanduseEducation,
	"fairground":    LanduseFairground,
	"industrial":    LanduseIndustrial,
	"residential":   LanduseResidential,
	"retail":        LanduseRetail,
	"institutional": LanduseInstitutional,

	// Rural and agricultural land
	"aquaculture": LanduseAquaculture,
	"allotments":  LanduseAllotments,
	"forest":      LanduseForest,
	"meadow":      LanduseMeadow,
	"orchard":     LanduseOrchard,
	"vineyard":    LanduseVineyard,

	// Waterbody
	"basin":     LanduseBasin,
	"reservoir": LanduseReservoir,

	// Other
	"brownfield":        LanduseBrownfield,
	"cemetery":          LanduseCemetery,
	"grass":             LanduseGrass,
	"greenfield":        LanduseGreenfield,
	"landfill":          LanduseLandfill,
	"military":          LanduseMilitary,
	"quarry":            LanduseQuarry,
	"recreation_ground": LanduseRecreationGround,
	"religious":         LanduseReligious,
	"village_green":     LanduseVillageGreen,
	"winter_sports":     LanduseWinterSports,
}

// Landuse is used to tag the purpose for which an area of land is being used.
// https://wiki.openstreetmap.org/wiki/Landuse
type Landuse Feature

func (l *Landuse) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validLanduseTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validLanduseTypes))
			for k := range validLanduseTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", l.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = l.GetTag()

	*l = Landuse(f)
	return nil
}

func (l *Landuse) GetTag() string {
	return "landuse"
}

func (l *Landuse) ToOQL() string {
	return Feature(*l).ToOQL()
}

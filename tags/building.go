package tags

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
	"github.com/kkrypt0nn/overscry/tags/validators"
)

// BuildingType represents the valid building types.
// https://wiki.openstreetmap.org/wiki/Key:building
type BuildingType string

const (
	// Accomodation

	// BuildingApartments represents a building arranged into individual dwellings, often on separate floors. May also have retail outlets on the ground floor.
	BuildingApartments BuildingType = "apartments"
	// BuildingBarracks represents buildings built to house military personnel or laborers.
	BuildingBarracks BuildingType = "barracks"
	// BuildingBungalow represents a single-storey detached small house, Dacha.
	BuildingBungalow BuildingType = "bungalow"
	// BuildingCabin represents a small, roughly built house usually with a wood exterior and typically found in rural areas. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Cabin
	BuildingCabin BuildingType = "cabin"
	// BuildingDetached represents a detached house, a free-standing residential building usually housing a single family.
	BuildingDetached BuildingType = "detached"
	// BuildingAnnexe represents a small self-contained apartment, cottage, or small residential building on the same property as the main residential unit.
	BuildingAnnexe BuildingType = "annexe"
	// BuildingDormitory represents a shared building intended for college/university students.
	BuildingDormitory BuildingType = "dormitory"
	// BuildingFarm represents a residential building on a farm (farmhouse).
	BuildingFarm BuildingType = "farm"
	// BuildingGer represents a permanent or seasonal round yurt or ger.
	BuildingGer BuildingType = "ger"
	// BuildingHotel represents a building designed with separate rooms available for overnight accommodation.
	BuildingHotel BuildingType = "hotel"
	// BuildingHouse represents a dwelling unit inhabited by a single household.
	BuildingHouse BuildingType = "house"
	// BuildingHouseboat represents a boat used primarily as a home.
	BuildingHouseboat BuildingType = "houseboat"
	// BuildingResidential represents a general tag for a building used primarily for residential purposes.
	BuildingResidential BuildingType = "residential"
	// BuildingSemidetachedHouse represents a residential house that shares a common wall with another on one side.
	BuildingSemidetachedHouse BuildingType = "semidetached_house"
	// BuildingStaticCaravan represents a mobile home (semi)permanently left on a single site.
	BuildingStaticCaravan BuildingType = "static_caravan"
	// BuildingStiltHouse represents a building raised on piles over the surface of the soil or a body of water.
	BuildingStiltHouse BuildingType = "stilt_house"
	// BuildingTerrace represents a linear row of residential dwellings, each of which normally has its own entrance.
	BuildingTerrace BuildingType = "terrace"
	// BuildingTreeHouse represents a small hut or room built on tree posts or on a natural tree.
	BuildingTreeHouse BuildingType = "tree_house"
	// BuildingTrullo represents a stone hut with a conical roof.
	BuildingTrullo BuildingType = "trullo"

	// Commercial

	// BuildingCommercial represents a building for non-specific commercial activities, not necessarily an office building.
	BuildingCommercial BuildingType = "commercial"
	// BuildingIndustrial represents a building for industrial purposes. Use [BuildingWarehouse] if the purpose is known to be primarily for storage/distribution.
	BuildingIndustrial BuildingType = "industrial"
	// BuildingKiosk represents a small one-room retail building.
	BuildingKiosk BuildingType = "kiosk"
	// BuildingOffice represents an office building.
	BuildingOffice BuildingType = "office"
	// BuildingRetail represents a building primarily used for selling goods that are sold to the public.
	BuildingRetail BuildingType = "retail"
	// BuildingSupermarket represents a building constructed to house a self-service large-area store.
	BuildingSupermarket BuildingType = "supermarket"
	// BuildingWarehouse represents a building primarily intended for the storage or goods or as part of a distribution system.
	BuildingWarehouse BuildingType = "warehouse"

	// Religious

	// BuildingReligious represents a building that is related to religion, but not specific to any particular structure type.
	BuildingReligious BuildingType = "religious"
	// BuildingCathedral represents a building that was built as a cathedral.
	BuildingCathedral BuildingType = "cathedral"
	// BuildingChapel represents a building that was built as a chapel.
	BuildingChapel BuildingType = "chapel"
	// BuildingChurch represents a building that was built as a church.
	BuildingChurch BuildingType = "church"
	// BuildingKingdomHall represents a building that was built as a [Kingdom Hall] for Jehovah's Witnesses.
	//
	// [Kingdom Hall]: https://en.wikipedia.org/wiki/Kingdom_Hall
	BuildingKingdomHall BuildingType = "kingdom_hall"
	// BuildingMonastery represents a building constructed as a [monastery], often with multiple structures.
	//
	// [monastery]: https://en.wikipedia.org/wiki/Monastery
	BuildingMonastery BuildingType = "monastery"
	// BuildingMosque represents a building that was erected as a mosque.
	BuildingMosque BuildingType = "mosque"
	// BuildingPresbytery represents a building where priests live and work.
	BuildingPresbytery BuildingType = "presbytery"
	// BuildingShrine represents a building that was built as a shrine.
	BuildingShrine BuildingType = "shrine"
	// BuildingSynagogue represents a building that was built as a synagogue.
	BuildingSynagogue BuildingType = "synagogue"
	// BuildingTemple represents a building that was built as a temple.
	BuildingTemple BuildingType = "temple"

	// Civic/amenity

	// BuildingBakehouse represents a building built for baking bread, often used with amenity=baking_oven and oven=wood_fired.
	BuildingBakehouse BuildingType = "bakehouse"
	// BuildingBridge represents a building used as a bridge, often referred to as a skyway.
	BuildingBridge BuildingType = "bridge"
	// BuildingCivic represents a building created to house civic amenities such as community centres, libraries, and town halls.
	BuildingCivic BuildingType = "civic"
	// BuildingCollege represents a building designed for a college.
	BuildingCollege BuildingType = "college"
	// BuildingFireStation represents a building constructed to house firefighting equipment and personnel.
	BuildingFireStation BuildingType = "fire_station"
	// BuildingGovernment represents a government building, including municipal and regional offices.
	BuildingGovernment BuildingType = "government"
	// BuildingGatehouse represents an entry control point building, typically located at the entrance to a city or compound.
	BuildingGatehouse BuildingType = "gatehouse"
	// BuildingHospital represents a building erected for use as a hospital.
	BuildingHospital BuildingType = "hospital"
	// BuildingKindergarten represents a building used as a kindergarten.
	BuildingKindergarten BuildingType = "kindergarten"
	// BuildingMuseum represents a building designed as a museum.
	BuildingMuseum BuildingType = "museum"
	// BuildingPublic represents a building constructed to be accessible to the general public, such as town halls or police stations.
	BuildingPublic BuildingType = "public"
	// BuildingSchool represents a building constructed as a school.
	BuildingSchool BuildingType = "school"
	// BuildingToilets represents a toilet block building.
	BuildingToilets BuildingType = "toilets"
	// BuildingTrainStation represents a building constructed to be a train station, even if now repurposed.
	BuildingTrainStation BuildingType = "train_station"
	// BuildingTransportation represents a building related to public transport, often tagged with transport-related tags.
	BuildingTransportation BuildingType = "transportation"
	// BuildingUniversity represents a building designed for a university.
	BuildingUniversity BuildingType = "university"

	// Agricultural/plant production

	// BuildingBarn represents an agricultural building used for storage and as a covered workplace.
	BuildingBarn BuildingType = "barn"
	// BuildingConservatory represents a building or room with glass or tarpaulin roofing, used as an indoor garden or sunroom.
	BuildingConservatory BuildingType = "conservatory"
	// BuildingCowshed represents a building for housing cows, commonly found on farms.
	BuildingCowshed BuildingType = "cowshed"
	// BuildingFarmAuxiliary represents a building on a farm that is not a dwelling, such as barns or storage structures.
	BuildingFarmAuxiliary BuildingType = "farm_auxiliary"
	// BuildingGreenhouse represents a [greenhouse], a structure used for growing plants, typically covered with glass or plastic.
	//
	// [greenhouse]: https://en.wikipedia.org/wiki/Greenhouse
	BuildingGreenhouse BuildingType = "greenhouse"
	// BuildingSlurryTank represents a circular building used to hold slurry (liquid animal excreta).
	BuildingSlurryTank BuildingType = "slurry_tank"
	// BuildingStable represents a building constructed as a stable for housing horses.
	BuildingStable BuildingType = "stable"
	// BuildingSty represents a building for raising domestic pigs, often found on farms. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Sty
	BuildingSty BuildingType = "sty"
	// BuildingLivestock represents a building for housing or raising other livestock (apart from cows, horses, or pigs).
	BuildingLivestock BuildingType = "livestock"

	// Sports

	// BuildingGrandstand represents the main stand at racecourses or sports grounds, often roofed and commanding the best view for spectators.
	BuildingGrandstand BuildingType = "grandstand"
	// BuildingPavilion represents a sports [pavilion], typically with changing rooms, storage areas, and possibly a space for functions & events.
	//
	// [pavilion]: https://en.wikipedia.org/wiki/Pavilion
	BuildingPavilion BuildingType = "pavilion"
	// BuildingRidingHall represents a building constructed specifically as a riding hall.
	BuildingRidingHall BuildingType = "riding_hall"
	// BuildingSportsHall represents a building constructed specifically as a sports hall.
	BuildingSportsHall BuildingType = "sports_hall"
	// BuildingSportsCentre represents a building constructed as a sports centre, often with various sports facilities.
	BuildingSportsCentre BuildingType = "sports_centre"
	// BuildingStadium represents a building constructed as a stadium, including those that may have been abandoned and repurposed.
	BuildingStadium BuildingType = "stadium"

	// Storage

	// BuildingAllotmentHouse represents a small outbuilding for short visits in an allotment garden.
	BuildingAllotmentHouse BuildingType = "allotment_house"
	// BuildingBoathouse represents a building used for the storage of boats.
	BuildingBoathouse BuildingType = "boathouse"
	// BuildingHangar represents a building used for the storage of airplanes, helicopters, or spacecraft. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Hangar
	BuildingHangar BuildingType = "hangar"
	// BuildingHut represents a small and crude shelter, possibly used as a shed or residential building of low quality. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Hut
	BuildingHut BuildingType = "hut"
	// BuildingShed represents a simple, single-storey structure often used for storage, hobbies, or as a workshop. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Shed
	BuildingShed BuildingType = "shed"

	// Cars

	// BuildingCarport represents a covered structure used to offer limited protection to vehicles, usually without four walls. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Carport
	BuildingCarport BuildingType = "carport"
	// BuildingGarage represents a building suitable for the storage of one or more motor vehicles. [Learn more]
	//
	// [Learn more]: https://en.wikipedia.org/wiki/Garage_(residential)
	BuildingGarage BuildingType = "garage"
	// BuildingGarages represents a building consisting of multiple discrete storage spaces for different owners or tenants.
	BuildingGarages BuildingType = "garages"
	// BuildingParking represents a structure purpose-built for parking cars.
	BuildingParking BuildingType = "parking"

	// Power/technical buildings

	// BuildingDigester represents a bioreactor used for the production of biogas from biomass.
	BuildingDigester BuildingType = "digester"
	// BuildingService represents a small unmanned building with machinery like pumps or transformers.
	BuildingService BuildingType = "service"
	// BuildingTechCab represents a small prefabricated cabin used for the air-conditioned accommodation of technology.
	BuildingTechCab BuildingType = "tech_cab"
	// BuildingTransformerTower represents a tall building comprising a distribution transformer, often connected to a power line.
	BuildingTransformerTower BuildingType = "transformer_tower"
	// BuildingWaterTower represents a water tower used for storing water at height for distribution.
	BuildingWaterTower BuildingType = "water_tower"
	// BuildingStorageTank represents a container that holds liquids, such as a storage tank.
	BuildingStorageTank BuildingType = "storage_tank"
	// BuildingSilo represents a building for storing bulk materials, such as grain or other agricultural products.
	BuildingSilo BuildingType = "silo"

	// Other

	// BuildingBeachHut represents a small, usually wooden cabin or shelter on popular bathing beaches.
	BuildingBeachHut BuildingType = "beach_hut"
	// BuildingBunker represents a hardened military building.
	BuildingBunker BuildingType = "bunker"
	// BuildingCastle represents a building constructed as a castle.
	BuildingCastle BuildingType = "castle"
	// BuildingConstruction represents a building under construction.
	BuildingConstruction BuildingType = "construction"
	// BuildingContainer represents a permanent building made from a container.
	BuildingContainer BuildingType = "container"
	// BuildingGuardhouse represents a post or small building used as a guardhouse.
	BuildingGuardhouse BuildingType = "guardhouse"
	// BuildingMilitary represents a military building.
	BuildingMilitary BuildingType = "military"
	// BuildingOutbuilding represents a less important building on the same land as a larger building.
	BuildingOutbuilding BuildingType = "outbuilding"
	// BuildingPagoda represents a building constructed as a pagoda.
	BuildingPagoda BuildingType = "pagoda"
	// BuildingQuonsetHut represents a lightweight prefabricated structure in the shape of a semicircle.
	BuildingQuonsetHut BuildingType = "quonset_hut"
	// BuildingRoof represents a structure consisting of a roof with open sides, such as a rain shelter.
	BuildingRoof BuildingType = "roof"
	// BuildingRuins represents a building that is abandoned and in poor repair.
	BuildingRuins BuildingType = "ruins"
	// BuildingShip represents a decommissioned ship or submarine that stays in one place.
	BuildingShip BuildingType = "ship"
	// BuildingTent represents a permanently placed tent.
	BuildingTent BuildingType = "tent"
	// BuildingTower represents a tower building.
	BuildingTower BuildingType = "tower"
	// BuildingTriumphalArch represents a free-standing monumental structure in the shape of an archway.
	BuildingTriumphalArch BuildingType = "triumphal_arch"
	// BuildingWindmill represents a traditional windmill used to mill grain.
	BuildingWindmill BuildingType = "windmill"
	// BuildingYes represents a building type that can't be more specifically determined.
	BuildingYes BuildingType = "yes"
)

var validBuildingTypes = map[string]BuildingType{
	// Accomodation
	"apartments":         BuildingApartments,
	"barracks":           BuildingBarracks,
	"bungalow":           BuildingBungalow,
	"cabin":              BuildingCabin,
	"detached":           BuildingDetached,
	"annexe":             BuildingAnnexe,
	"dormitory":          BuildingDormitory,
	"farm":               BuildingFarm,
	"ger":                BuildingGer,
	"hotel":              BuildingHotel,
	"house":              BuildingHouse,
	"houseboat":          BuildingHouseboat,
	"residential":        BuildingResidential,
	"semidetached_house": BuildingSemidetachedHouse,
	"static_caravan":     BuildingStaticCaravan,
	"stilt_house":        BuildingStiltHouse,
	"terrace":            BuildingTerrace,
	"tree_house":         BuildingTreeHouse,
	"trullo":             BuildingTrullo,

	// Commercial
	"commercial":  BuildingCommercial,
	"industrial":  BuildingIndustrial,
	"kiosk":       BuildingKiosk,
	"office":      BuildingOffice,
	"retail":      BuildingRetail,
	"supermarket": BuildingSupermarket,
	"warehouse":   BuildingWarehouse,

	// Religious
	"religious":    BuildingReligious,
	"cathedral":    BuildingCathedral,
	"chapel":       BuildingChapel,
	"church":       BuildingChurch,
	"kingdom_hall": BuildingKingdomHall,
	"monastery":    BuildingMonastery,
	"mosque":       BuildingMosque,
	"presbytery":   BuildingPresbytery,
	"shrine":       BuildingShrine,
	"synagogue":    BuildingSynagogue,
	"temple":       BuildingTemple,

	// Civic/amenity
	"bakehouse":      BuildingBakehouse,
	"bridge":         BuildingBridge,
	"civic":          BuildingCivic,
	"college":        BuildingCollege,
	"fire_station":   BuildingFireStation,
	"government":     BuildingGovernment,
	"gatehouse":      BuildingGatehouse,
	"hospital":       BuildingHospital,
	"kindergarten":   BuildingKindergarten,
	"museum":         BuildingMuseum,
	"public":         BuildingPublic,
	"school":         BuildingSchool,
	"toilets":        BuildingToilets,
	"train_station":  BuildingTrainStation,
	"transportation": BuildingTransportation,
	"university":     BuildingUniversity,

	// Agricultural/plant production
	"barn":           BuildingBarn,
	"conservatory":   BuildingConservatory,
	"cowshed":        BuildingCowshed,
	"farm_auxiliary": BuildingFarmAuxiliary,
	"greenhouse":     BuildingGreenhouse,
	"slurry_tank":    BuildingSlurryTank,
	"stable":         BuildingStable,
	"sty":            BuildingSty,
	"livestock":      BuildingLivestock,

	// Sports
	"grandstand":    BuildingGrandstand,
	"pavilion":      BuildingPavilion,
	"riding_hall":   BuildingRidingHall,
	"sports_hall":   BuildingSportsHall,
	"sports_centre": BuildingSportsCentre,
	"stadium":       BuildingStadium,

	// Storage
	"allotment_house": BuildingAllotmentHouse,
	"boathouse":       BuildingBoathouse,
	"hangar":          BuildingHangar,
	"hut":             BuildingHut,
	"shed":            BuildingShed,

	// Cars
	"carport": BuildingCarport,
	"garage":  BuildingGarage,
	"garages": BuildingGarages,
	"parking": BuildingParking,

	// Power/technical buildings
	"digester":          BuildingDigester,
	"service":           BuildingService,
	"tech_cab":          BuildingTechCab,
	"transformer_tower": BuildingTransformerTower,
	"water_tower":       BuildingWaterTower,
	"storage_tank":      BuildingStorageTank,
	"silo":              BuildingSilo,

	// Other
	"beach_hut":      BuildingBeachHut,
	"bunker":         BuildingBunker,
	"castle":         BuildingCastle,
	"construction":   BuildingConstruction,
	"container":      BuildingContainer,
	"guardhouse":     BuildingGuardhouse,
	"military":       BuildingMilitary,
	"outbuilding":    BuildingOutbuilding,
	"pagoda":         BuildingPagoda,
	"quonset_hut":    BuildingQuonsetHut,
	"roof":           BuildingRoof,
	"ruins":          BuildingRuins,
	"ship":           BuildingShip,
	"tent":           BuildingTent,
	"tower":          BuildingTower,
	"triumphal_arch": BuildingTriumphalArch,
	"windmill":       BuildingWindmill,
	"yes":            BuildingYes,
}

// Building is used to tag individual buildings or groups of connected buildings.
// https://wiki.openstreetmap.org/wiki/Buildings
type Building struct {
	Value  *BuildingType `yaml:"value,omitempty"`
	Match  MatchMethod   `yaml:"match,omitempty"`
	Inline *Feature      `yaml:"-"`

	// Architecture represents the architectural style of a building.
	Architecture *Feature `yaml:"architecture,omitempty"`
	// Fireproof represents the fire-resistance information.
	Fireproof *Feature `yaml:"fireproof,omitempty"`
	// Flats represents the number of residential units (flats, apartments) in an apartment building ([BuildingApartments]), residential building ([BuildingResidential]), house ([BuildingHouse]), detached house ([BuildingDetached]) or similar building.
	Flats *Feature `yaml:"flats,omitempty"`
	// ArchLevelsitecture represents the number of visible levels (floors) in the building as used in the [Simple 3D buildings] scheme.
	//
	// [Simple 3D buildings]: https://wiki.openstreetmap.org/wiki/Simple_3D_Buildings
	Levels *Feature `yaml:"levels,omitempty"`
	// SoftStorey represents a building where any one level is significantly more flexible (less stiff) than those above and below it.
	SoftStorey *Feature `yaml:"soft_storey,omitempty"`
}

func (b *Building) UnmarshalYAML(unmarshal func(interface{}) error) error {
	temp := &struct {
		Value *string     `yaml:"value,omitempty"`
		Match MatchMethod `yaml:"match,omitempty"`

		Architecture *Feature `yaml:"architecture,omitempty"`
		Fireproof    *Feature `yaml:"fireproof,omitempty"`
		Flats        *Feature `yaml:"flats,omitempty"`
		Levels       *Feature `yaml:"levels,omitempty"`
		SoftStorey   *Feature `yaml:"soft_storey,omitempty"`
	}{}
	if err := unmarshal(temp); err != nil {
		return err
	}

	if temp.Value != nil {
		val := strings.ToLower(*temp.Value)
		if _, ok := validBuildingTypes[val]; ok {
			b.Value = (*BuildingType)(helpers.Ptr(val))
		} else {
			validKeys := make([]string, 0, len(validBuildingTypes))
			for k := range validBuildingTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", b.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	b.Inline = &Feature{
		Match: temp.Match,
		Value: helpers.Ptr(*temp.Value),
		Tag:   b.GetTag(),
	}

	tempVal := reflect.ValueOf(temp).Elem()
	targetVal := reflect.ValueOf(b).Elem()
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
				Prefix: helpers.Ptr(b.GetTag()),
			}
			if validator, ok := validators.BuildingValidators[tempFeature.Tag]; ok {
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

func (b *Building) GetTag() string {
	return "building"
}

func (b *Building) ToOQL() string {
	var queryParts []string

	val := reflect.ValueOf(b).Elem()
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

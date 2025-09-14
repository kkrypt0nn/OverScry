package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// HistoricType represents the valid historic types.
// https://wiki.openstreetmap.org/wiki/Key:historic
type HistoricType string

const (
	// HistoricAircraft represents a decommissioned aircraft which generally remains in one place.
	HistoricAircraft HistoricType = "aircraft"
	// HistoricAnchor represents a historic/retired anchor. Usually found in historic maritime areas.
	HistoricAnchor HistoricType = "anchor"
	// HistoricArchaeologicalSite represents a place in which evidence of past activity is preserved.
	HistoricArchaeologicalSite HistoricType = "archaeological_site"
	// HistoricBattlefield represents the site of a battle or military skirmish in the past.
	HistoricBattlefield HistoricType = "battlefield"
	// HistoricBombCrater represents a bomb crater.
	HistoricBombCrater HistoricType = "bomb_crater"
	// HistoricBoundaryStone represents a historic physical marker that identifies a boundary.
	HistoricBoundaryStone HistoricType = "boundary_stone"
	// HistoricBuilding represents unspecified historic building.
	HistoricBuilding HistoricType = "building"
	// HistoricBullaunStone represents a stone with one or more depressions.
	HistoricBullaunStone HistoricType = "bullaun_stone"
	// HistoricCannon represents a historic/retired cannon. Usually found on forts or battlefields.
	HistoricCannon HistoricType = "cannon"
	// HistoricCastle represents used for various kinds of castles, palaces, fortresses, manors, stately homes, kremlins, shiros and other.
	HistoricCastle HistoricType = "castle"
	// HistoricCharcoalPile represents a historic site of a charcoal pile. Often still in good condition in hilly forest areas.
	HistoricCharcoalPile HistoricType = "charcoal_pile"
	// HistoricChurch represents a church with a historical value.
	HistoricChurch HistoricType = "church"
	// HistoricCityGate represents a city gate within a city wall.
	HistoricCityGate HistoricType = "city_gate"
	// HistoricCreamery represents a creamery is an industrial building where butter and sometimes cheese or ice-cream were made from milk. For rural communities, it also served as a social gathering point.
	HistoricCreamery HistoricType = "creamery"
	// HistoricDistrict represents a designated historic district.
	HistoricDistrict HistoricType = "district"
	// HistoricEpigraph represents a historic inscription on an object.
	HistoricEpigraph HistoricType = "epigraph"
	// HistoricFarm represents a historical farm, kept in its original state.
	HistoricFarm HistoricType = "farm"
	// HistoricFort represents a military fort, a stand-alone defensive structure which differs from a castle in that there is no permanent residence.
	HistoricFort HistoricType = "fort"
	// HistoricGallows represents remains of a gallows.
	HistoricGallows HistoricType = "gallows"
	// HistoricHouse represents a historic house.
	HistoricHouse HistoricType = "house"
	// HistoricHighCross represents an early medieaval standing cross, richly decorated, often with a ring surrounding the crossing point.
	HistoricHighCross HistoricType = "high_cross"
	// HistoricHighwaterMark represents a marker indicating a past flood or high water.
	HistoricHighwaterMark HistoricType = "highwater_mark"
	// HistoricLavoir represents disused lavoir mapped for its historical value.
	HistoricLavoir HistoricType = "lavoir"
	// HistoricLimeKiln represents built structure which was used in the past to produce quicklime from limestone.
	HistoricLimeKiln HistoricType = "lime_kiln"
	// HistoricLocomotive represents a decommissioned locomotive which generally remains in one place.
	HistoricLocomotive HistoricType = "locomotive"
	// HistoricMachine represents a historic machine.
	HistoricMachine HistoricType = "machine"
	// HistoricManor represents historic manors/mansions having different use today.
	HistoricManor HistoricType = "manor"
	// HistoricMemorial represents small memorials, usually remembering special persons, people who lost their lives in the wars, past events or missing places.
	HistoricMemorial HistoricType = "memorial"
	// HistoricMilestone represents a historic marker that shows the distance to important destinations.
	HistoricMilestone HistoricType = "milestone"
	// HistoricMillstone represents a large round stone used for grinding grain or other materials.
	HistoricMillstone HistoricType = "millstone"
	// HistoricMine represents location of historic underground mine workings for minerals such as coal or lead.
	HistoricMine HistoricType = "mine"
	// HistoricMinecart represents a cart used to transport coal or ore from a mine.
	HistoricMinecart HistoricType = "minecart"
	// HistoricMonastery represents building/place that is a historically significant monastery.
	HistoricMonastery HistoricType = "monastery"
	// HistoricMonument represents a memorial object, which is especially large, built to remember, show respect to a person or group of people or to commemorate an event.
	HistoricMonument HistoricType = "monument"
	// HistoricMosque represents a mosque with a historical and archaeological value.
	HistoricMosque HistoricType = "mosque"
	// HistoricOghamStone represents a stone with an Ogham script on it. They are most commonly found in Ireland as free standing stones, lying on the ground, recycled in buildings such as churches or as artefacts in museums.
	HistoricOghamStone HistoricType = "ogham_stone"
	// HistoricOpticalTelegraph represents semaphore system.
	HistoricOpticalTelegraph HistoricType = "optical_telegraph"
	// HistoricPa represents a New Zealand Maori Pā.
	HistoricPa HistoricType = "pa"
	// HistoricPillory represents a pillory.
	HistoricPillory HistoricType = "pillory"
	// HistoricPound represents a former amenity to hold stray or seized animals, often walled.
	HistoricPound HistoricType = "pound"
	// HistoricRailwayCar represents a decommissioned railway car which generally remains in one place.
	HistoricRailwayCar HistoricType = "railway_car"
	// HistoricRoundTower represents slim, conical tower built as a bell tower and watchtower in Ireland.
	HistoricRoundTower HistoricType = "round_tower"
	// HistoricRuins represents remains of structures that were once complete, but have fallen into partial or complete disrepair.
	HistoricRuins HistoricType = "ruins"
	// HistoricRuneStone represents stones, boulders or bedrock with historical runic inscriptions.
	HistoricRuneStone HistoricType = "rune_stone"
	// HistoricShieling represents an abandoned mountain pasture.
	HistoricShieling HistoricType = "shieling"
	// HistoricShip represents a decommissioned ship/submarine which generally remains in one place.
	HistoricShip HistoricType = "ship"
	// HistoricStećak represents specifies megalithic gravestone from medieval Bosnia.
	HistoricStećak HistoricType = "stećak"
	// HistoricStone represents a stone shaped or placed by man with historical value.
	HistoricStone HistoricType = "stone"
	// HistoricTank represents a decommissioned tank which generally remains in one place.
	HistoricTank HistoricType = "tank"
	// HistoricTemple represents an ancient pagan temple in_situ and in various degree of preservation, such as Ziggurat, Egyptian temple, Mithraeum, Erechtheion, Buddhist, Meso-American and other ancient religious and rituals' temples, with a historical and archaeological value.
	HistoricTemple HistoricType = "temple"
	// HistoricTomb represents a structure where somebody has been buried.
	HistoricTomb HistoricType = "tomb"
	// HistoricTower represents this property distinguishes a tower as historic.
	HistoricTower HistoricType = "tower"
	// HistoricVehicle represents a decommissioned vehicle which generally remains in one place.
	HistoricVehicle HistoricType = "vehicle"
	// HistoricWaysideCross represents a historical cross along a way, symbol of Christian faith.
	HistoricWaysideCross HistoricType = "wayside_cross"
	// HistoricWaysideShrine represents a shrine often showing a religious depiction. Tag is used also for modern shrines.
	HistoricWaysideShrine HistoricType = "wayside_shrine"
	// HistoricWreck represents a nautical craft that has been sunk or destroyed.
	HistoricWreck HistoricType = "wreck"
	// HistoricYes represents used to add the historic significance of the objects described by other tags.
	HistoricYes HistoricType = "yes"
)

var validHistoricTypes = map[string]HistoricType{
	"aircraft":            HistoricAircraft,
	"anchor":              HistoricAnchor,
	"archaeological_site": HistoricArchaeologicalSite,
	"battlefield":         HistoricBattlefield,
	"bomb_crater":         HistoricBombCrater,
	"boundary_stone":      HistoricBoundaryStone,
	"building":            HistoricBuilding,
	"bullaun_stone":       HistoricBullaunStone,
	"cannon":              HistoricCannon,
	"castle":              HistoricCastle,
	"charcoal_pile":       HistoricCharcoalPile,
	"church":              HistoricChurch,
	"city_gate":           HistoricCityGate,
	"creamery":            HistoricCreamery,
	"district":            HistoricDistrict,
	"epigraph":            HistoricEpigraph,
	"farm":                HistoricFarm,
	"fort":                HistoricFort,
	"gallows":             HistoricGallows,
	"house":               HistoricHouse,
	"high_cross":          HistoricHighCross,
	"highwater_mark":      HistoricHighwaterMark,
	"lavoir":              HistoricLavoir,
	"lime_kiln":           HistoricLimeKiln,
	"locomotive":          HistoricLocomotive,
	"machine":             HistoricMachine,
	"manor":               HistoricManor,
	"memorial":            HistoricMemorial,
	"milestone":           HistoricMilestone,
	"millstone":           HistoricMillstone,
	"mine":                HistoricMine,
	"minecart":            HistoricMinecart,
	"monastery":           HistoricMonastery,
	"monument":            HistoricMonument,
	"mosque":              HistoricMosque,
	"ogham_stone":         HistoricOghamStone,
	"optical_telegraph":   HistoricOpticalTelegraph,
	"pa":                  HistoricPa,
	"pillory":             HistoricPillory,
	"pound":               HistoricPound,
	"railway_car":         HistoricRailwayCar,
	"round_tower":         HistoricRoundTower,
	"ruins":               HistoricRuins,
	"rune_stone":          HistoricRuneStone,
	"shieling":            HistoricShieling,
	"ship":                HistoricShip,
	"stećak":              HistoricStećak,
	"stone":               HistoricStone,
	"tank":                HistoricTank,
	"temple":              HistoricTemple,
	"tomb":                HistoricTomb,
	"tower":               HistoricTower,
	"vehicle":             HistoricVehicle,
	"wayside_cross":       HistoricWaysideCross,
	"wayside_shrine":      HistoricWaysideShrine,
	"wreck":               HistoricWreck,
	"yes":                 HistoricYes,
}

// Historic is used to tag various historic places.
// https://wiki.openstreetmap.org/wiki/Historic
type Historic Feature

func (h *Historic) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validHistoricTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validHistoricTypes))
			for k := range validHistoricTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", h.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = h.GetTag()

	*h = Historic(f)
	return nil
}

func (h *Historic) GetTag() string {
	return "historic"
}

func (h *Historic) ToOQL() string {
	return Feature(*h).ToOQL()
}

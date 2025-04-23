package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// ManMadeType represents the valid ManMade types.
// https://wiki.openstreetmap.org/wiki/Key:man_made
type ManMadeType string

const (
	// ManMadeAdit represents a type of entrance to an underground mine which is horizontal or nearly horizontal.
	ManMadeAdit ManMadeType = "adit"
	// ManMadeBeacon represents a structure for signalling on land and sea.
	ManMadeBeacon ManMadeType = "beacon"
	// ManMadeBunkerSilo represents an open-sided structure without a roof that can be used with vehicles to fill and empty them.
	ManMadeBunkerSilo ManMadeType = "bunker_silo"
	// ManMadeCarpetHanger represents a construction to hang carpets for cleaning with the help of carpet beaters.
	ManMadeCarpetHanger ManMadeType = "carpet_hanger"
	// ManMadeChimney represents a tall distinctive vertical conduit for venting hot gases or smoke, normally found near power stations or large factories.
	ManMadeChimney ManMadeType = "chimney"
	// ManMadeColumn represents a column used to support a structure or for decoration.
	ManMadeColumn ManMadeType = "column"
	// ManMadeCommunicationsTower represents a huge tower for transmitting radio applications.
	ManMadeCommunicationsTower ManMadeType = "communications_tower"
	// ManMadeCrane represents a stationary, permanent crane.
	ManMadeCrane ManMadeType = "crane"
	// ManMadeCross represents cross, especially one with little historical value.
	ManMadeCross ManMadeType = "cross"
	// ManMadeDovecote represents a place where doves are farmed or stored.
	ManMadeDovecote ManMadeType = "dovecote"
	// ManMadeFlagpole represents a long pole built to hold a flag.
	ManMadeFlagpole ManMadeType = "flagpole"
	// ManMadeGasometer represents a large container in which natural gas or town gas is stored near atmospheric pressure at ambient temperatures.
	ManMadeGasometer ManMadeType = "gasometer"
	// ManMadeGuardStone represents a guard stone: a stone built onto or into the corner of a building or wall to prevent carriages from damaging the structure, often found on either side of an entrance to a laneway, or alongside a wall to protect it.
	ManMadeGuardStone ManMadeType = "guard_stone"
	// ManMadeKiln represents a thermally insulated chamber used for processes such as burning, hardening, drying, or smelting.
	ManMadeKiln ManMadeType = "kiln"
	// ManMadeLighthouse represents tower that emits light to serve as a navigational aid at sea or on inland waterway.
	ManMadeLighthouse ManMadeType = "lighthouse"
	// ManMadeMast represents a mast is a vertical structure built to hold, for example, antennas.
	ManMadeMast ManMadeType = "mast"
	// ManMadeMineshaft represents a vertical tunnel into a mine where minerals are extracted.
	ManMadeMineshaft ManMadeType = "mineshaft"
	// ManMadeMonitoringStation represents a station that monitors something.
	ManMadeMonitoringStation ManMadeType = "monitoring_station"
	// ManMadeObelisk represents tall, narrow, four-sided, tapered monument which usually ends in a pyramid-like shape at the top.
	ManMadeObelisk ManMadeType = "obelisk"
	// ManMadeObservatory represents observatory: a location used for observing terrestrial or celestial events.
	ManMadeObservatory ManMadeType = "observatory"
	// ManMadeOffshorePlatform represents offshore platform, oil platform or offshore drilling rig.
	ManMadeOffshorePlatform ManMadeType = "offshore_platform"
	// ManMadePetroleumWell represents a hole bored in the earth, designed to bring petroleum oil or gas to the surface.
	ManMadePetroleumWell ManMadeType = "petroleum_well"
	// ManMadePump represents a device in charge of moving or raising the level of liquids.
	ManMadePump ManMadeType = "pump"
	// ManMadePumpingStation represents pumping station: a facility including pumps and equipment for pumping fluids from one place to another.
	ManMadePumpingStation ManMadeType = "pumping_station"
	// ManMadeReservoirCovered represents a covered reservoir is a large man-made tank for holding fresh water.
	ManMadeReservoirCovered ManMadeType = "reservoir_covered"
	// ManMadeSewerVent represents used to map sewer vent pipes.
	ManMadeSewerVent ManMadeType = "sewer_vent"
	// ManMadeSilo represents a storage container for bulk material, often grains such as corn or wheat.
	ManMadeSilo ManMadeType = "silo"
	// ManMadeSnowNet represents a netting fence built across steep slopes to reduce risk and severity of (snow) avalanches.
	ManMadeSnowNet ManMadeType = "snow_net"
	// ManMadeStorageTank represents a container that holds liquids or compressed gases.
	ManMadeStorageTank ManMadeType = "storage_tank"
	// ManMadeStreetCabinet represents a cabinet located in the street and hosting technical equipment to operate facilities such as electricity or street lights.
	ManMadeStreetCabinet ManMadeType = "street_cabinet"
	// ManMadeStupa represents a Buddhist dome-shaped structure with a spire on top.
	ManMadeStupa ManMadeType = "stupa"
	// ManMadeSurveillance represents a surveillance camera or other type of surveillance equipment.
	ManMadeSurveillance ManMadeType = "surveillance"
	// ManMadeSurveyPoint represents a triangulation pillar, geodetic vertex, or other piece of fixed equipment used by topographers.
	ManMadeSurveyPoint ManMadeType = "survey_point"
	// ManMadeTelescope represents telescope: an instrument that aids in the observation of remote objects by collecting light or radio waves.
	ManMadeTelescope ManMadeType = "telescope"
	// ManMadeTower represents a tower is a free-standing structure which is higher than it is wide.
	ManMadeTower ManMadeType = "tower"
	// ManMadeVideoWall represents a digital screen, typically constructed out of smaller LED panels.
	ManMadeVideoWall ManMadeType = "video_wall"
	// ManMadeWastewaterPlant represents a wastewater plant is a facility used to treat wastewater.
	ManMadeWastewaterPlant ManMadeType = "wastewater_plant"
	// ManMadeWatermill represents a mill driven by water power.
	ManMadeWatermill ManMadeType = "watermill"
	// ManMadeWaterTower represents structure with a water tank at an altitude to increase pressure in water network.
	ManMadeWaterTower ManMadeType = "water_tower"
	// ManMadeWaterWell represents a structural facility to access ground water, created by digging or drilling.
	ManMadeWaterWell ManMadeType = "water_well"
	// ManMadeWaterTap represents publicly usable water tap.
	ManMadeWaterTap ManMadeType = "water_tap"
	// ManMadeWaterWorks represents a facility where water is treated to make it suitable for human consumption.
	ManMadeWaterWorks ManMadeType = "water_works"
	// ManMadeWindmill represents a traditional windmill, historically used to mill grain with wind power.
	ManMadeWindmill ManMadeType = "windmill"
	// ManMadeWorks represents a factory or industrial production plant.
	ManMadeWorks ManMadeType = "works"
)

var validManMadeTypes = map[string]ManMadeType{
	"adit":                 ManMadeAdit,
	"beacon":               ManMadeBeacon,
	"bunker_silo":          ManMadeBunkerSilo,
	"carpet_hanger":        ManMadeCarpetHanger,
	"chimney":              ManMadeChimney,
	"column":               ManMadeColumn,
	"communications_tower": ManMadeCommunicationsTower,
	"crane":                ManMadeCrane,
	"cross":                ManMadeCross,
	"dovecote":             ManMadeDovecote,
	"flagpole":             ManMadeFlagpole,
	"gasometer":            ManMadeGasometer,
	"guard_stone":          ManMadeGuardStone,
	"kiln":                 ManMadeKiln,
	"lighthouse":           ManMadeLighthouse,
	"mast":                 ManMadeMast,
	"mineshaft":            ManMadeMineshaft,
	"monitoring_station":   ManMadeMonitoringStation,
	"obelisk":              ManMadeObelisk,
	"observatory":          ManMadeObservatory,
	"offshore_platform":    ManMadeOffshorePlatform,
	"petroleum_well":       ManMadePetroleumWell,
	"pump":                 ManMadePump,
	"pumping_station":      ManMadePumpingStation,
	"reservoir_covered":    ManMadeReservoirCovered,
	"sewer_vent":           ManMadeSewerVent,
	"silo":                 ManMadeSilo,
	"snow_net":             ManMadeSnowNet,
	"storage_tank":         ManMadeStorageTank,
	"street_cabinet":       ManMadeStreetCabinet,
	"stupa":                ManMadeStupa,
	"surveillance":         ManMadeSurveillance,
	"survey_point":         ManMadeSurveyPoint,
	"telescope":            ManMadeTelescope,
	"tower":                ManMadeTower,
	"video_wall":           ManMadeVideoWall,
	"wastewater_plant":     ManMadeWastewaterPlant,
	"watermill":            ManMadeWatermill,
	"water_tower":          ManMadeWaterTower,
	"water_well":           ManMadeWaterWell,
	"water_tap":            ManMadeWaterTap,
	"water_works":          ManMadeWaterWorks,
	"windmill":             ManMadeWindmill,
	"works":                ManMadeWorks,
}

// ManMade is used to tag man made (artificial) structures that are added to the landscape.
// https://wiki.openstreetmap.org/wiki/Key:man_made
type ManMade Feature

func (m *ManMade) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validManMadeTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validManMadeTypes))
			for k := range validManMadeTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", m.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = m.GetTag()

	*m = ManMade(f)
	return nil
}

func (m *ManMade) GetTag() string {
	return "man_made"
}

func (m *ManMade) ToOQL() string {
	return Feature(*m).ToOQL()
}

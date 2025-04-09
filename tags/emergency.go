package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// EmergencyType represents the valid emergency types.
// https://wiki.openstreetmap.org/wiki/Key:emergency
type EmergencyType string

const (
	// Medical rescue

	// EmergencyAmbulanceStation represents an ambulance station is a structure or other area set aside for storage of ambulance vehicles, medical equipment, personal protective equipment, and other medical supplies.
	EmergencyAmbulanceStation EmergencyType = "ambulance_station"
	// EmergencyDefibrillator represents a defibrillator (AED), an external and portable electronic device that diagnoses and can correct arrhythmia of the heart automatically.
	EmergencyDefibrillator EmergencyType = "defibrillator"
	// EmergencyLandingSite represents a preselected flat area for a helicopter to land in an emergency situation.
	EmergencyLandingSite EmergencyType = "landing_site"
	// EmergencyEmergencyWardEntrance represents the entrance to an emergency ward.
	EmergencyEmergencyWardEntrance EmergencyType = "emergency_ward_entrance"

	// Firefighters

	// EmergencyFireServiceInlet represents an inlet that allows the fire brigade pump water into a building.
	EmergencyFireServiceInlet EmergencyType = "fire_service_inlet"
	// EmergencyFireAlarmBox represents a device used for notifying a fire department of a fire.
	EmergencyFireAlarmBox EmergencyType = "fire_alarm_box"
	// EmergencyFireExtinguisher represents an active fire protection device used to extinguish or control small fires, often in emergency situations.
	EmergencyFireExtinguisher EmergencyType = "fire_extinguisher"
	// EmergencyFireHose represents a high-pressure hose used to carry water or other fire retardant (such as foam) to a fire to extinguish it.
	EmergencyFireHose EmergencyType = "fire_hose"
	// EmergencyFireHydrant represents an active fire protection measure, and a source of water provided in most urban, suburban, and rural areas with municipal water service to enable fire fighters to tap into the municipal water supply to assist in extinguishing a fire.
	EmergencyFireHydrant EmergencyType = "fire_hydrant"
	// EmergencyWaterTank represents a large water basin or tank for a fire department to take water.
	EmergencyWaterTank EmergencyType = "water_tank"
	// EmergencySuctionPoint represents a preferred point to pump water off a river or other waters for a fire department.
	EmergencySuctionPoint EmergencyType = "suction_point"

	// Lifeguards

	// EmergencyLifeguard represents a place where a lifeguard is on duty.
	EmergencyLifeguard EmergencyType = "lifeguard"
	// EmergencyLifeRing represents a floating ring to throw out to someone who is struggling in water.
	EmergencyLifeRing EmergencyType = "life_ring"

	// Assembly point

	// EmergencyAssemblyPoint represents a designated (safe) place where people can gather or must report to during an emergency or a fire drill.
	EmergencyAssemblyPoint EmergencyType = "assembly_point"

	// Other structure

	// EmergencyPhone represents an emergency telephone.
	EmergencyPhone EmergencyType = "phone"
	// EmergencySiren represents a loud noise maker, such as an air raid siren or a tornado siren.
	EmergencySiren EmergencyType = "siren"
	// EmergencySiren represents a facility that provides drinking water in emergency situations.
	EmergencyDrinkingWater EmergencyType = "drinking_water"
)

var validEmergencyTypes = map[string]EmergencyType{
	// Medical rescue
	"ambulance_station":       EmergencyAmbulanceStation,
	"defibrillator":           EmergencyDefibrillator,
	"landing_site":            EmergencyLandingSite,
	"emergency_ward_entrance": EmergencyEmergencyWardEntrance,

	// Firefighters
	"fire_service_inlet": EmergencyFireServiceInlet,
	"fire_alarm_box":     EmergencyFireAlarmBox,
	"fire_extinguisher":  EmergencyFireExtinguisher,
	"fire_hose":          EmergencyFireHose,
	"fire_hydrant":       EmergencyFireHydrant,
	"water_tank":         EmergencyWaterTank,
	"suction_point":      EmergencySuctionPoint,

	// Lifeguards
	"lifeguard": EmergencyLifeguard,
	"life_ring": EmergencyLifeRing,

	// Assembly point
	"assembly_point": EmergencyAssemblyPoint,

	// Other structure
	"phone":          EmergencyPhone,
	"siren":          EmergencySiren,
	"drinking_water": EmergencyDrinkingWater,
}

// Emergency is used to tag emergency facilities and equipment.
// https://wiki.openstreetmap.org/wiki/Emergency_facilities_and_amenities
type Emergency Feature

func (e *Emergency) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validEmergencyTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validEmergencyTypes))
			for k := range validEmergencyTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", e.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = e.GetTag()

	*e = Emergency(f)
	return nil
}

func (e *Emergency) GetTag() string {
	return "emergency"
}

func (e *Emergency) ToOQL() string {
	return Feature(*e).ToOQL()
}

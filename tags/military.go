package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// MilitaryType represents the valid military types.
// https://wiki.openstreetmap.org/wiki/Key:military
type MilitaryType string

const (
	// MilitaryAcademy represents a training establishment for military service members.
	MilitaryAcademy MilitaryType = "academy"
	// MilitaryAirfield represents a place where military planes take off and land.
	MilitaryAirfield MilitaryType = "airfield"
	// MilitaryBase represents a facility where military personnel and equipment are based.
	MilitaryBase MilitaryType = "base"
	// MilitaryBunker represents a building reinforced to withstand attack.
	MilitaryBunker MilitaryType = "bunker"
	// MilitaryBarracks represents buildings where military personnel live and sleep.
	MilitaryBarracks MilitaryType = "barracks"
	// MilitaryCheckpoint represents place where civilian visitors and vehicles will be checked by a military authority.
	MilitaryCheckpoint MilitaryType = "checkpoint"
	// MilitaryDangerArea represents a danger area is a restricted area posing a threat to life or property.
	MilitaryDangerArea MilitaryType = "danger_area"
	// MilitaryNuclearExplosionSite represents nuclear weapons test site.
	MilitaryNuclearExplosionSite MilitaryType = "nuclear_explosion_site"
	// MilitaryOffice represents military offices, e.g. general staff office, military recruitment office etc.
	MilitaryOffice MilitaryType = "office"
	// MilitaryRange represents where military personnel practice with their weapons (firing, bombing, artillery, ...).
	MilitaryRange MilitaryType = "range"
	// MilitarySchool represents a school for children run as part of a country's military.
	MilitarySchool MilitaryType = "school"
	// MilitaryTrench represents a military trench is an excavation in the ground that is generally deeper than it is wide, dug into the ground as a barrier for military purposes (e.g. trench warfare).
	MilitaryTrench MilitaryType = "trench"
)

var validMilitaryTypes = map[string]MilitaryType{
	"academy":                MilitaryAcademy,
	"airfield":               MilitaryAirfield,
	"base":                   MilitaryBase,
	"bunker":                 MilitaryBunker,
	"barracks":               MilitaryBarracks,
	"checkpoint":             MilitaryCheckpoint,
	"danger_area":            MilitaryDangerArea,
	"nuclear_explosion_site": MilitaryNuclearExplosionSite,
	"office":                 MilitaryOffice,
	"range":                  MilitaryRange,
	"school":                 MilitarySchool,
	"trench":                 MilitaryTrench,
}

// Military is used to tag facilities and on land used by the military. These may include the Navy, Army, Air Force and Marines.
// https://wiki.openstreetmap.org/wiki/Military
type Military Feature

func (m *Military) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validMilitaryTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validMilitaryTypes))
			for k := range validMilitaryTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", m.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = m.GetTag()

	*m = Military(f)
	return nil
}

func (m *Military) GetTag() string {
	return "military"
}

func (m *Military) ToOQL() string {
	return Feature(*m).ToOQL()
}

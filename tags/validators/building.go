package validators

import (
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

var architectureValues = []string{"islamic", "mamluk", "romanesque", "gothic", "renaissance", "mannerism", "ottoman", "baroque", "rococo", "empire", "moorish revival", "neoclassicism", "georgian", "victorian", "historicism", "neo-romanesque", "neo-byzantine", "neo-gothic", "neo-renaissance", "neo-baroque", "art_nouveau", "eclectic", "functionalism", "cubism", "new_objectivity", "art_deco", "modern", "amsterdam_school", "international_style", "constructivism", "stalinist_neoclassicism", "brutalist", "postmodern", "contemporary"}
var yesNoValues = []string{"yes", "no"}
var yesNoReinforcedValues = []string{"yes", "no", "reinforced"}

var BuildingValidators = map[string]*FeatureValidator{
	"architecture": {
		Callback: func(val string) bool {
			return helpers.Contains(architectureValues, strings.ToLower(val))
		},
		AllowedValues: architectureValues,
	},
	"fireproof": {
		Callback: func(val string) bool {
			return helpers.Contains(yesNoValues, strings.ToLower(val))
		},
		AllowedValues: yesNoValues,
	},
	"soft_storey": {
		Callback: func(val string) bool {
			return helpers.Contains(yesNoReinforcedValues, strings.ToLower(val))
		},
		AllowedValues: yesNoReinforcedValues,
	},
}

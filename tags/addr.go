package tags

import "strings"

type Addr struct {
	HouseNumber *Feature `yaml:"housenumber,omitempty"`
}

func (a *Addr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	temp := &struct {
		HouseNumber *Feature `yaml:"housenumber"`
	}{}
	if err := unmarshal(temp); err != nil {
		return err
	}

	if temp.HouseNumber != nil {
		a.HouseNumber = temp.HouseNumber
		a.HouseNumber.Tag = "housenumber"
		a.HouseNumber.Prefix = "addr"
	}

	return nil
}

func (a Addr) ToOQL() string {
	var queryParts []string

	if a.HouseNumber != nil {
		queryParts = append(queryParts, a.HouseNumber.ToOQL())
	}

	return strings.Join(queryParts, "")
}

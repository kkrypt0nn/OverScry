package tags

import (
	"reflect"
	"strings"
)

// Addr is used to provide address information for a building, facility or any other object that has an address.
// https://wiki.openstreetmap.org/wiki/Addresses
type Addr struct {
	// HouseNumber represents the house number (may contain letters, dashes or other characters).
	HouseNumber *Feature `yaml:"housenumber,omitempty"`
	// HouseName represents the name of a house.
	HouseName *Feature `yaml:"housename,omitempty"`
	// Flats represents the unit numbers (a range or a list) of the flats or apartments located behind a single entrance door.
	Flats *Feature `yaml:"flats,omitempty"`
	// ConscriptionNumber represents a special kind of housenumber relates to a settlement instead of a street.
	// Conscription numbers were introduced in the Austro-Hungarian Empire and are still in use in some parts of Europe, sometimes together with street-related housenumbers which are also called orientation numbers.
	ConscriptionNumber *Feature `yaml:"conscriptionnumber,omitempty"`
	// Street represents the name of the respective street. If the street name is very long or nonexistent, the ref of the respective street.
	Street *Feature `yaml:"street,omitempty"`
	// Place represents a part of an address which refers to the name of some territorial zone (usually an island, square or very small village) instead of a street. Should not be used together with [Addr.Street].
	Place *Feature `yaml:"place,omitempty"`
	// PostCode represents the postal code of the building/area.
	PostCode *Feature `yaml:"postcode,omitempty"`
	// City represents the name of the city as given in postal addresses of the building/area.
	City *Feature `yaml:"city,omitempty"`
	// Country represents the ISO 3166-1 alpha-2 two letter country code in upper case.
	// Example: "DE" for Germany, "CH" for Switzerland, "AT" for Austria, "FR" for France, "IT" for Italy.
	Country *Feature `yaml:"country,omitempty"`
	// PostBox represents the postal service [Post Office Box] as alternative to addressing using street names.
	// Example: "PO Box 34"
	//
	// [Post Office Box]: https://en.wikipedia.org/wiki/Post_office_box
	PostBox *Feature `yaml:"postbox,omitempty"`
	// Full represents a full-text, often multi-line, address if you find the structured address fields unsuitable for denoting the address of this particular location. Examples: "Fifth house on the left after the village oak, Smalltown, Smallcountry", or addresses using special delivery names or codes (possibly via an unrelated city name and post code), or PO Boxes.
	Full *Feature `yaml:"full,omitempty"`

	// For countries using hamlet, subdistrict, district, province, state, county

	// Hamlet represents the hamlet of the object. In France, some addresses use hamlets instead of street names, use the generic [Addr.Place] instead.
	Hamlet *Feature `yaml:"hamlet,omitempty"`
	// Suburb represents the name of the settlement. If an address exists several times in a city you have to add it.
	// See Australian definition of [suburb].
	//
	// [suburb]: https://en.wikipedia.org/wiki/Suburb
	Suburb *Feature `yaml:"suburb,omitempty"`
	// Subdistrict represents the [subdistrict] of the object.
	//
	// [subdistrict]: https://en.wikipedia.org/wiki/Subdistrict
	Subdistrict *Feature `yaml:"subdistrict,omitempty"`
	// District represents the [district] of the object.
	//
	// [district]: https://en.wikipedia.org/wiki/District
	District *Feature `yaml:"district,omitempty"`
	// Province represents the [province] of the object.
	//
	// [province]: https://en.wikipedia.org/wiki/Province
	Province *Feature `yaml:"province,omitempty"`
	// State represents the [state] of the object. For the US, uppercase two-letter postal abbreviations (AK, CA, HI, NY, TX, WY, etc.) are used.
	//
	// [state]: https://en.wikipedia.org/wiki/Administrative_division
	State *Feature `yaml:"state,omitempty"`
	// County represents the [county] of the object.
	//
	// [county]: https://en.wikipedia.org/wiki/County
	County *Feature `yaml:"county,omitempty"`
}

func (a *Addr) UnmarshalYAML(unmarshal func(interface{}) error) error {
	temp := &struct {
		HouseNumber        *Feature `yaml:"housenumber,omitempty"`
		HouseName          *Feature `yaml:"housename,omitempty"`
		Flats              *Feature `yaml:"flats,omitempty"`
		ConscriptionNumber *Feature `yaml:"conscriptionnumber,omitempty"`
		Street             *Feature `yaml:"street,omitempty"`
		Place              *Feature `yaml:"place,omitempty"`
		PostCode           *Feature `yaml:"postcode,omitempty"`
		City               *Feature `yaml:"city,omitempty"`
		Country            *Feature `yaml:"country,omitempty"`
		PostBox            *Feature `yaml:"postbox,omitempty"`
		Full               *Feature `yaml:"full,omitempty"`

		Hamlet      *Feature `yaml:"hamlet,omitempty"`
		Suburb      *Feature `yaml:"suburb,omitempty"`
		Subdistrict *Feature `yaml:"subdistrict,omitempty"`
		District    *Feature `yaml:"district,omitempty"`
		Province    *Feature `yaml:"province,omitempty"`
		State       *Feature `yaml:"state,omitempty"`
		County      *Feature `yaml:"county,omitempty"`
	}{}
	if err := unmarshal(temp); err != nil {
		return err
	}

	val := reflect.ValueOf(a).Elem()
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)

		tempField := reflect.ValueOf(temp).Elem().FieldByName(field.Name)
		if tempField.IsValid() && !tempField.IsNil() {
			fieldVal.Set(tempField)
			feature := fieldVal.Interface().(*Feature)
			feature.Tag = strings.ToLower(field.Name)
			feature.Prefix = "addr"
		}
	}

	return nil
}

func (a *Addr) ToOQL() string {
	var queryParts []string

	val := reflect.ValueOf(a).Elem()
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		fieldVal := val.Field(i)
		if fieldVal.IsValid() && !fieldVal.IsNil() {
			queryParts = append(queryParts, fieldVal.Interface().(*Feature).ToOQL())
		}
	}

	return strings.Join(queryParts, "")
}

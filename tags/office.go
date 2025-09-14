package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// OfficeType represents the valid office types.
// https://wiki.openstreetmap.org/wiki/Key:office
type OfficeType string

const (
	// OfficeAccountant represents an office for an accountant.
	OfficeAccountant OfficeType = "accountant"
	// OfficeAdvertisingAgency represents a service-based business dedicated to creating, planning, and handling advertising.
	OfficeAdvertisingAgency OfficeType = "advertising_agency"
	// OfficeAirline represents an office for an airline company.
	OfficeAirline OfficeType = "airline"
	// OfficeArchitect represents an office for an architect or group of architects.
	OfficeArchitect OfficeType = "architect"
	// OfficeAssociation represents an office of a non-profit organisation, society, e.g. student, sport, consumer, automobile, bike association, etc.
	OfficeAssociation OfficeType = "association"
	// OfficeBroadcaster represents a radio or television broadcasting station.
	OfficeBroadcaster OfficeType = "broadcaster"
	// OfficeChamber represents seat of a professional chamber, such as bar associations or chambers of commerce.
	OfficeChamber OfficeType = "chamber"
	// OfficeCharity represents an office of a charitable organization.
	OfficeCharity OfficeType = "charity"
	// OfficeCompany represents an office of a private company.
	OfficeCompany OfficeType = "company"
	// OfficeConstructionCompany represents an office of a building construction company.
	OfficeConstructionCompany OfficeType = "construction_company"
	// OfficeConsulting represents an office for a consulting firm, providing expert professional advice to other companies or organisations.
	OfficeConsulting OfficeType = "consulting"
	// OfficeCourier represents a courier delivery service, which is neither a post office nor a national post.
	OfficeCourier OfficeType = "courier"
	// OfficeCoworking represents an office where people can go to work (typically requires a fee); not limited to a single employer.
	OfficeCoworking OfficeType = "coworking"
	// OfficeDiplomatic represents an embassy, diplomatic mission, consulate or liaison office of a foreign government or parastatal entity in a host country.
	OfficeDiplomatic OfficeType = "diplomatic"
	// OfficeEducationalInstitution represents an office for an educational institution.
	OfficeEducationalInstitution OfficeType = "educational_institution"
	// OfficeEmploymentAgency represents an office for an employment service.
	OfficeEmploymentAgency OfficeType = "employment_agency"
	// OfficeEnergySupplier represents an office for an energy supplier.
	OfficeEnergySupplier OfficeType = "energy_supplier"
	// OfficeEngineer represents an office for an engineer or group of engineers.
	OfficeEngineer OfficeType = "engineer"
	// OfficeEstateAgent represents a place where you can rent or buy a house.
	OfficeEstateAgent OfficeType = "estate_agent"
	// OfficeEventManagement represents an office offering support for event planning and management.
	OfficeEventManagement OfficeType = "event_management"
	// OfficeFinancial represents an office of a company in the financial sector.
	OfficeFinancial OfficeType = "financial"
	// OfficeFinancialAdvisor represents a professional who offers financial planning and sells financial services to clients.
	OfficeFinancialAdvisor OfficeType = "financial_advisor"
	// OfficeForestry represents a forestry office.
	OfficeForestry OfficeType = "forestry"
	// OfficeFoundation represents an office of a foundation.
	OfficeFoundation OfficeType = "foundation"
	// OfficeGeodesist represents an office of scientists studying Earth shape.
	OfficeGeodesist OfficeType = "geodesist"
	// OfficeGongo represents an office of a government-organized non-governmental organization.
	OfficeGongo OfficeType = "gongo"
	// OfficeGovernment represents an office of a (supra)national, regional or local government agency or department.
	OfficeGovernment OfficeType = "government"
	// OfficeGraphicDesign represents an office of a graphic designer.
	OfficeGraphicDesign OfficeType = "graphic_design"
	// OfficeGuide represents an office of tour guides, mountain guides, hiking guides, dive guides, etc.
	OfficeGuide OfficeType = "guide"
	// OfficeHarbourMaster represents the Harbourmaster's office.
	OfficeHarbourMaster OfficeType = "harbour_master"
	// OfficeInsurance represents an office at which you can take out insurance policies.
	OfficeInsurance OfficeType = "insurance"
	// OfficeIt represents an office for an IT company.
	OfficeIt OfficeType = "it"
	// OfficeLawyer represents an office for a lawyer.
	OfficeLawyer OfficeType = "lawyer"
	// OfficeLogistics represents an office for a company offering logistics.
	OfficeLogistics OfficeType = "logistics"
	// OfficeMovingCompany represents an office which offers a relocation service.
	OfficeMovingCompany OfficeType = "moving_company"
	// OfficeNewspaper represents an office of a newspaper.
	OfficeNewspaper OfficeType = "newspaper"
	// OfficeNgo represents an office for a non-profit, non-governmental organisation (NGO).
	OfficeNgo OfficeType = "ngo"
	// OfficeNotary represents an office for a notary public (common law).
	OfficeNotary OfficeType = "notary"
	// OfficePolitician represents a politician's office.
	OfficePolitician OfficeType = "politician"
	// OfficePoliticalParty represents an office of a political party.
	OfficePoliticalParty OfficeType = "political_party"
	// OfficePropertyManagement represents office of a property rental company for residential apartments or commercial office space or any other property.
	OfficePropertyManagement OfficeType = "property_management"
	// OfficePublisher represents an office of a publisher.
	OfficePublisher OfficeType = "publisher"
	// OfficeQuango represents an office of a quasi-autonomous non-governmental organisation.
	OfficeQuango OfficeType = "quango"
	// OfficeReligion represents an office of a community of faith.
	OfficeReligion OfficeType = "religion"
	// OfficeResearch represents an office for research and development.
	OfficeResearch OfficeType = "research"
	// OfficeSecurity represents an office for private security guards.
	OfficeSecurity OfficeType = "security"
	// OfficeSurveyor represents an office of a person doing technical surveys, such as land surveys or risk and damage evaluations of properties and equipment.
	OfficeSurveyor OfficeType = "surveyor"
	// OfficeTaxAdvisor represents an office for a financial expert specially trained in tax law.
	OfficeTaxAdvisor OfficeType = "tax_advisor"
	// OfficeTelecommunication represents an office for a telecommunication company.
	OfficeTelecommunication OfficeType = "telecommunication"
	// OfficeTransport represents an office of a freight forwarder or transport company.
	OfficeTransport OfficeType = "transport"
	// OfficeTravelAgent represents an office of a travel agent.
	OfficeTravelAgent OfficeType = "travel_agent"
	// OfficeTutoring represents an office of tutor or tutoring centre.
	OfficeTutoring OfficeType = "tutoring"
	// OfficeUnion represents an office of a trade union, an association of workers forming a bargaining unit.
	OfficeUnion OfficeType = "union"
	// OfficeUniversity represents an office providing services to university students or staff, or the office of a department / faculty in a university.
	OfficeUniversity OfficeType = "university"
	// OfficeVisa represents an office of an organisation or business which offers visa assistance.
	OfficeVisa OfficeType = "visa"
	// OfficeWaterUtility represents the office for a water utility company or water board.
	OfficeWaterUtility OfficeType = "water_utility"
)

var validOfficeTypes = map[string]OfficeType{
	"accountant":              OfficeAccountant,
	"advertising_agency":      OfficeAdvertisingAgency,
	"airline":                 OfficeAirline,
	"architect":               OfficeArchitect,
	"association":             OfficeAssociation,
	"broadcaster":             OfficeBroadcaster,
	"chamber":                 OfficeChamber,
	"charity":                 OfficeCharity,
	"company":                 OfficeCompany,
	"construction_company":    OfficeConstructionCompany,
	"consulting":              OfficeConsulting,
	"courier":                 OfficeCourier,
	"coworking":               OfficeCoworking,
	"diplomatic":              OfficeDiplomatic,
	"educational_institution": OfficeEducationalInstitution,
	"employment_agency":       OfficeEmploymentAgency,
	"energy_supplier":         OfficeEnergySupplier,
	"engineer":                OfficeEngineer,
	"estate_agent":            OfficeEstateAgent,
	"event_management":        OfficeEventManagement,
	"financial":               OfficeFinancial,
	"financial_advisor":       OfficeFinancialAdvisor,
	"forestry":                OfficeForestry,
	"foundation":              OfficeFoundation,
	"geodesist":               OfficeGeodesist,
	"gongo":                   OfficeGongo,
	"government":              OfficeGovernment,
	"graphic_design":          OfficeGraphicDesign,
	"guide":                   OfficeGuide,
	"harbour_master":          OfficeHarbourMaster,
	"insurance":               OfficeInsurance,
	"it":                      OfficeIt,
	"lawyer":                  OfficeLawyer,
	"logistics":               OfficeLogistics,
	"moving_company":          OfficeMovingCompany,
	"newspaper":               OfficeNewspaper,
	"ngo":                     OfficeNgo,
	"notary":                  OfficeNotary,
	"politician":              OfficePolitician,
	"political_party":         OfficePoliticalParty,
	"property_management":     OfficePropertyManagement,
	"publisher":               OfficePublisher,
	"quango":                  OfficeQuango,
	"religion":                OfficeReligion,
	"research":                OfficeResearch,
	"security":                OfficeSecurity,
	"surveyor":                OfficeSurveyor,
	"tax_advisor":             OfficeTaxAdvisor,
	"telecommunication":       OfficeTelecommunication,
	"transport":               OfficeTransport,
	"travel_agent":            OfficeTravelAgent,
	"tutoring":                OfficeTutoring,
	"union":                   OfficeUnion,
	"university":              OfficeUniversity,
	"visa":                    OfficeVisa,
	"water_utility":           OfficeWaterUtility,
}

// Office is used to tag a place of business where administrative or professional work is carried out.
// https://wiki.openstreetmap.org/wiki/Key:office
type Office Feature

func (o *Office) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validOfficeTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validOfficeTypes))
			for k := range validOfficeTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", o.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = o.GetTag()

	*o = Office(f)
	return nil
}

func (o *Office) GetTag() string {
	return "office"
}

func (o *Office) ToOQL() string {
	return Feature(*o).ToOQL()
}

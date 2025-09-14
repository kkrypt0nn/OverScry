package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// HealthcareType represents the valid healthcare types.
// https://wiki.openstreetmap.org/wiki/Key:healthcare
type HealthcareType string

const (
	// HealthcareAlternative represents a facility where alternative or complementary medicine is practiced: e.g. acupuncture, chiropractic, naturopathy, etc.
	HealthcareAlternative HealthcareType = "alternative"
	// HealthcareAudiologist represents a facility where an audiologist identifies and treats hearing problems.
	HealthcareAudiologist HealthcareType = "audiologist"
	// HealthcareBloodBank represents a center where blood gathered as a result of blood donation is stored and preserved for later use in blood transfusion.
	HealthcareBloodBank HealthcareType = "blood_bank"
	// HealthcareBloodDonation represents facility where you can donate blood, plasma and/or platelets, and possibly have stem cell samples taken.
	HealthcareBloodDonation HealthcareType = "blood_donation"
	// HealthcareClinic represents a medical facility, with more staff than a doctor's office, that does not admit inpatients. Used in addition to tag:amenity=clinic.
	HealthcareClinic HealthcareType = "clinic"
	// HealthcareCounselling represents a facility where health counselling is provided, e.g. dietitian, nutrition, addiction.
	HealthcareCounselling HealthcareType = "counselling"
	// HealthcareDentist represents a dentist practice / surgery. Used in addition to amenity=dentist.
	HealthcareDentist HealthcareType = "dentist"
	// HealthcareDialysis represents facility where people can get dialysis.
	HealthcareDialysis HealthcareType = "dialysis"
	// HealthcareDoctor represents a place to get medical attention or a check up from a physician. Used in addition to tag:amenity=doctors.
	HealthcareDoctor HealthcareType = "doctor"
	// HealthcareHospice represents a facility which provides palliative care to terminally ill people and support to their relatives.
	HealthcareHospice HealthcareType = "hospice"
	// HealthcareHospital represents a hospital providing in-patient medical treatment, used in addition to tag:amenity=hospital.
	HealthcareHospital HealthcareType = "hospital"
	// HealthcareLaboratory represents a medical laboratory, a place which performs analysis and diagnostics on body fluids.
	HealthcareLaboratory HealthcareType = "laboratory"
	// HealthcareMidwife represents a facility where healthcare is provided by a midwife, a health professional who cares for mothers and newborns around childbirth.
	HealthcareMidwife HealthcareType = "midwife"
	// HealthcareNurse represents a facility where a trained health professional (nurse) provides healthcare, who is not a physician.
	HealthcareNurse HealthcareType = "nurse"
	// HealthcareOccupationalTherapist represents a facility where an occupational therapist practices.
	HealthcareOccupationalTherapist HealthcareType = "occupational_therapist"
	// HealthcareOptometrist represents an optometrist's office.
	HealthcareOptometrist HealthcareType = "optometrist"
	// HealthcarePhysiotherapist represents a facility where a physiotherapist practices physical therapy (kinesiology, exercise, mobilization, etc).
	HealthcarePhysiotherapist HealthcareType = "physiotherapist"
	// HealthcarePodiatrist represents the office of a podiatrist for the treatment of disorders of the foot, ankle, and lower extremities.
	HealthcarePodiatrist HealthcareType = "podiatrist"
	// HealthcarePsychotherapist represents an office of a psychotherapist or clinical psychologist.
	HealthcarePsychotherapist HealthcareType = "psychotherapist"
	// HealthcareRehabilitation represents medical rehabilitation facility, usually inpatient or residential.
	HealthcareRehabilitation HealthcareType = "rehabilitation"
	// HealthcareSampleCollection represents site or dedicated healthcare facility where samples of blood/urine/etc are obtained or collected for purpose of analyzing them for healthcare diagnostics.
	HealthcareSampleCollection HealthcareType = "sample_collection"
	// HealthcareSpeechTherapist represents a speech therapist, a health specialist who deals with speech, voice, swallowing or hearing impairment.
	HealthcareSpeechTherapist HealthcareType = "speech_therapist"
	// HealthcareVaccinationCentre represents a healthcare facility specifically dedicated to administering vaccinations to individuals, to provide immunisation against infectious diseases.
	HealthcareVaccinationCentre HealthcareType = "vaccination_centre"
)

var validHealthcareTypes = map[string]HealthcareType{
	"alternative":            HealthcareAlternative,
	"audiologist":            HealthcareAudiologist,
	"blood_bank":             HealthcareBloodBank,
	"blood_donation":         HealthcareBloodDonation,
	"clinic":                 HealthcareClinic,
	"counselling":            HealthcareCounselling,
	"dentist":                HealthcareDentist,
	"dialysis":               HealthcareDialysis,
	"doctor":                 HealthcareDoctor,
	"hospice":                HealthcareHospice,
	"hospital":               HealthcareHospital,
	"laboratory":             HealthcareLaboratory,
	"midwife":                HealthcareMidwife,
	"nurse":                  HealthcareNurse,
	"occupational_therapist": HealthcareOccupationalTherapist,
	"optometrist":            HealthcareOptometrist,
	"physiotherapist":        HealthcarePhysiotherapist,
	"podiatrist":             HealthcarePodiatrist,
	"psychotherapist":        HealthcarePsychotherapist,
	"rehabilitation":         HealthcareRehabilitation,
	"sample_collection":      HealthcareSampleCollection,
	"speech_therapist":       HealthcareSpeechTherapist,
	"vaccination_centre":     HealthcareVaccinationCentre,
}

// Healthcare is used to tag facilities which provide healthcare services.
// https://wiki.openstreetmap.org/wiki/Healthcare
type Healthcare Feature

func (h *Healthcare) UnmarshalYAML(unmarshal func(any) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validHealthcareTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validHealthcareTypes))
			for k := range validHealthcareTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid %s value: %q (allowed: %s)", h.GetTag(), val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = h.GetTag()

	*h = Healthcare(f)
	return nil
}

func (h *Healthcare) GetTag() string {
	return "healthcare"
}

func (h *Healthcare) ToOQL() string {
	return Feature(*h).ToOQL()
}

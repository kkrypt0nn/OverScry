package tags

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkrypt0nn/overscry/helpers"
)

// AmenityType represents the valid amenity types.
// https://wiki.openstreetmap.org/wiki/Key:amenity
type AmenityType string

const (
	// Sustenance

	// AmenityBar represents a purpose-built commercial establishment that sells alcoholic drinks to be consumed on the premises. They are characterised by a noisy and vibrant atmosphere, similar to a party and usually don't sell food.
	AmenityBar AmenityType = "bar"
	// AmenityBiergarten represents an open-air area where alcoholic beverages along with food is prepared and served.
	AmenityBiergarten AmenityType = "biergarten"
	// AmenityCafe represents an informal place that offers casual meals and beverages; typically, the focus is on coffee or tea. Also known as a coffeehouse/shop, bistro or sidewalk cafe.
	AmenityCafe AmenityType = "cafe"
	// AmenityFastFood represents a fast food restaurant.
	AmenityFastFood AmenityType = "fast_food"
	// AmenityFoodCourt represents an area with several different restaurant food counters and a shared eating area. Commonly found in malls, airports, etc.
	AmenityFoodCourt AmenityType = "food_court"
	// AmenityIceCream represents an ice cream shop or ice cream parlour. A place that sells ice cream and frozen yoghurt over the counter.
	AmenityIceCream AmenityType = "ice_cream"
	// AmenityPub represents a place selling beer and other alcoholic drinks; may also provide food or accommodation (UK)
	AmenityPub AmenityType = "pub"
	// AmenityRestaurant represents a restaurant (not fast food, see [AmenityFastFood]).
	AmenityRestaurant AmenityType = "restaurant"

	// Education

	// AmenityCollege represents a campus or buildings of an institute of Further Education (also known as continuing education).
	AmenityCollege AmenityType = "college"
	// AmenityDancingSchool represents a dancing school or dance studio.
	AmenityDancingSchool AmenityType = "dancing_school"
	// AmenityDrivingSchool represents a driving school that offers motor vehicle driving lessons.
	AmenityDrivingSchool AmenityType = "driving_school"
	// AmenityFirstAidSchool represents a place where people can go for first aid courses.
	AmenityFirstAidSchool AmenityType = "first_aid_school"
	// AmenityKindergarten represents a facility for children too young for a regular school (also known as preschool, playschool or nursery school).
	AmenityKindergarten AmenityType = "kindergarten"
	// AmenityLanguageSchool represents an educational institution where one studies a foreign language.
	AmenityLanguageSchool AmenityType = "language_school"
	// AmenityLibrary represents a public library (municipal, university, etc.) to borrow books from.
	AmenityLibrary AmenityType = "library"
	// AmenitySurfSchool represents an establishment that teaches surfing.
	AmenitySurfSchool AmenityType = "surf_school"
	// AmenityToyLibrary represents a place to borrow games and toys, or play with them on site.
	AmenityToyLibrary AmenityType = "toy_library"
	// AmenityResearchInstitute represents an establishment endowed for doing research.
	AmenityResearchInstitute AmenityType = "research_institute"
	// AmenityTraining represents a public place where you can get training.
	AmenityTraining AmenityType = "training"
	// AmenityMusicSchool represents a music school, an educational institution specialized in the study, training, and research of music.
	AmenityMusicSchool AmenityType = "music_school"
	// AmenitySchool represents school and grounds - primary, middle and secondary schools.
	AmenitySchool AmenityType = "school"
	// AmenityTrafficPark represents juvenile traffic schools.
	AmenityTrafficPark AmenityType = "traffic_park"
	// AmenityUniversity represents a university campus: an institute of higher education.
	AmenityUniversity AmenityType = "university"

	// Transportation

	// AmenityBicycleParking represents parking for bicycles.
	AmenityBicycleParking AmenityType = "bicycle_parking"
	// AmenityBicycleRepairStation represents general tools for self-service bicycle repairs, usually on the roadside; no service.
	AmenityBicycleRepairStation AmenityType = "bicycle_repair_station"
	// AmenityBicycleRental represents a place to rent a bicycle.
	AmenityBicycleRental AmenityType = "bicycle_rental"
	// AmenityBicycleWash represents a place to clean a bicycle.
	AmenityBicycleWash AmenityType = "bicycle_wash"
	// AmenityBoatRental represents a place to rent a boat.
	AmenityBoatRental AmenityType = "boat_rental"
	// AmenityBoatSharing represents a place to share a boat.
	AmenityBoatSharing AmenityType = "boat_sharing"
	// AmenityBusStation represents a bus station.
	AmenityBusStation AmenityType = "bus_station"
	// AmenityCarRental represents a place to rent a car.
	AmenityCarRental AmenityType = "car_rental"
	// AmenityCarSharing represents a place to share a car.
	AmenityCarSharing AmenityType = "car_sharing"
	// AmenityCarWash represents a place to wash a car.
	AmenityCarWash AmenityType = "car_wash"
	// AmenityCompressedAir represents a device to inflate tires/tyres (e.g. motorcar, bicycle).
	AmenityCompressedAir AmenityType = "compressed_air"
	// AmenityVehicleInspection represents a government vehicle inspection station.
	AmenityVehicleInspection AmenityType = "vehicle_inspection"
	// AmenityChargingStation represents a charging facility for electric vehicles.
	AmenityChargingStation AmenityType = "charging_station"
	// AmenityDriverTraining represents a place for driving training on a closed course.
	AmenityDriverTraining AmenityType = "driver_training"
	// AmenityFerryTerminal represents a ferry terminal/stop, a place where people/cars/etc. can board and leave a ferry.
	AmenityFerryTerminal AmenityType = "ferry_terminal"
	// AmenityFuel represents a petrol station, gas station, or marine fuel station.
	AmenityFuel AmenityType = "fuel"
	// AmenityGritBin represents a container that holds grit or a mixture of salt and grit.
	AmenityGritBin AmenityType = "grit_bin"
	// AmenityMotorcycleParking represents parking for motorcycles.
	AmenityMotorcycleParking AmenityType = "motorcycle_parking"
	// AmenityParking represents a parking area for vehicles.
	AmenityParking AmenityType = "parking"
	// AmenityParkingEntrance represents an entrance or exit to an underground or multi-storey parking facility.
	AmenityParkingEntrance AmenityType = "parking_entrance"
	// AmenityParkingSpace represents a single parking space within a car park.
	AmenityParkingSpace AmenityType = "parking_space"
	// AmenityTaxi represents a place where taxis wait for passengers.
	AmenityTaxi AmenityType = "taxi"
	// AmenityWeighbridge represents a large weight scale to weigh vehicles and goods.
	AmenityWeighbridge AmenityType = "weighbridge"

	// Financial

	// AmenityATM represents an automated teller machine (ATM) or cashpoint: a device that provides the clients of a financial institution with access to financial transactions.
	AmenityATM AmenityType = "atm"
	// AmenityPaymentTerminal represents a self-service payment kiosk/terminal.
	AmenityPaymentTerminal AmenityType = "payment_terminal"
	// AmenityBank represents a bank or credit union: a financial establishment where customers can deposit and withdraw money, take loans, make investments and transfer funds.
	AmenityBank AmenityType = "bank"
	// AmenityBureauDeChange represents a bureau de change, money changer, currency exchange, Wechsel, or cambio – a place to change foreign bank notes and travellers cheques.
	AmenityBureauDeChange AmenityType = "bureau_de_change"
	// AmenityMoneyTransfer represents a place that offers money transfers, especially cash to cash.
	AmenityMoneyTransfer AmenityType = "money_transfer"
	// AmenityPaymentCentre represents a non-bank place, where people can pay bills of public and private services and taxes.
	AmenityPaymentCentre AmenityType = "payment_centre"

	// Healthcare

	// AmenityBabyHatch represents a place where a baby can be, out of necessity, anonymously left to be safely cared for and perhaps adopted.
	AmenityBabyHatch AmenityType = "baby_hatch"
	// AmenityClinic represents a medium-sized medical facility or health centre.
	AmenityClinic AmenityType = "clinic"
	// AmenityDentist represents a dentist practice / surgery.
	AmenityDentist AmenityType = "dentist"
	// AmenityDoctors represents a doctor's practice / surgery.
	AmenityDoctors AmenityType = "doctors"
	// AmenityHospital represents a hospital providing in-patient medical treatment. Often used in conjunction with emergency=* to note whether the medical centre has emergency facilities (A&E (brit.) or ER (am.)).
	AmenityHospital AmenityType = "hospital"
	// AmenityNursingHome represents a home for disabled or elderly persons who need permanent care (discouraged tag, use amenity=social_facility + social_facility=nursing_home).
	AmenityNursingHome AmenityType = "nursing_home"
	// AmenityPharmacy represents a pharmacy: a shop where a pharmacist sells medications.
	AmenityPharmacy AmenityType = "pharmacy"
	// AmenitySocialFacility represents a facility that provides social services: group & nursing homes, workshops for the disabled, homeless shelters, etc.
	AmenitySocialFacility AmenityType = "social_facility"
	// AmenityVeterinary represents a place where a veterinary surgeon, also known as a veterinarian or vet, practices.
	AmenityVeterinary AmenityType = "veterinary"

	// Entertainment, Arts & Culture

	// AmenityArtsCentre represents a venue where a variety of arts are performed or conducted.
	AmenityArtsCentre AmenityType = "arts_centre"
	// AmenityBrothel represents an establishment specifically dedicated to prostitution.
	AmenityBrothel AmenityType = "brothel"
	// AmenityCasino represents a gambling venue with at least one table game (e.g., roulette, blackjack) that takes bets on sporting and other events at agreed upon odds.
	AmenityCasino AmenityType = "casino"
	// AmenityCinema represents a place where films are shown (US: movie theater).
	AmenityCinema AmenityType = "cinema"
	// AmenityCommunityCentre represents a place mostly used for local events, festivities and group activities; including special interest and special age groups.
	AmenityCommunityCentre AmenityType = "community_centre"
	// AmenityConferenceCentre represents a large building that is used to hold a convention.
	AmenityConferenceCentre AmenityType = "conference_centre"
	// AmenityEventsVenue represents a building specifically used for organizing events.
	AmenityEventsVenue AmenityType = "events_venue"
	// AmenityExhibitionCentre represents an exhibition centre.
	AmenityExhibitionCentre AmenityType = "exhibition_centre"
	// AmenityFountain represents a fountain for cultural / decorational / recreational purposes.
	AmenityFountain AmenityType = "fountain"
	// AmenityGambling represents a place for gambling, not being a shop=bookmaker, shop=lottery, amenity=casino, or leisure=adult_gaming_centre. Games that are covered by this definition include bingo and pachinko.
	AmenityGambling AmenityType = "gambling"
	// AmenityLoveHotel represents a love hotel, a type of short-stay hotel operated primarily for the purpose of allowing guests privacy for sexual activities.
	AmenityLoveHotel AmenityType = "love_hotel"
	// AmenityMusicVenue represents an indoor place to hear contemporary live music.
	AmenityMusicVenue AmenityType = "music_venue"
	// AmenityNightclub represents a place to drink and dance (nightclub). The German word is "Disco" or "Discothek". Please don't confuse this with the German "Nachtclub" which is most likely amenity=stripclub.
	AmenityNightclub AmenityType = "nightclub"
	// AmenityPlanetarium represents a planetarium.
	AmenityPlanetarium AmenityType = "planetarium"
	// AmenityPublicBookcase represents a street furniture containing books. Take one or leave one.
	AmenityPublicBookcase AmenityType = "public_bookcase"
	// AmenitySocialCentre represents a place for free and not-for-profit activities.
	AmenitySocialCentre AmenityType = "social_centre"
	// AmenityStage represents a raised platform for performers.
	AmenityStage AmenityType = "stage"
	// AmenityStripclub represents a place that offers striptease or lapdancing (for sexual services use amenity=brothel).
	AmenityStripclub AmenityType = "stripclub"
	// AmenityStudio represents a TV radio or recording studio.
	AmenityStudio AmenityType = "studio"
	// AmenitySwingerclub represents a club where people meet to have a party and group sex.
	AmenitySwingerclub AmenityType = "swingerclub"
	// AmenityTheatre represents a theatre or opera house where live performances occur, such as plays, musicals and formal concerts. Use [AmenityCinema] for movie theaters.
	AmenityTheatre AmenityType = "theatre"

	// Public Service

	// AmenityCourthouse represents a building home to a court of law, where justice is dispensed.
	AmenityCourthouse AmenityType = "courthouse"
	// AmenityFireStation represents a station of a fire brigade.
	AmenityFireStation AmenityType = "fire_station"
	// AmenityPolice represents a police station where police officers patrol from and that is a first point of contact for civilians.
	AmenityPolice AmenityType = "police"
	// AmenityPostBox represents a box for the reception of mail. Alternative mail-carriers can be tagged via operator=*.
	AmenityPostBox AmenityType = "post_box"
	// AmenityPostDepot represents a post depot or delivery office, where letters and parcels are collected and sorted prior to delivery.
	AmenityPostDepot AmenityType = "post_depot"
	// AmenityPostOffice represents a post office building with postal services.
	AmenityPostOffice AmenityType = "post_office"
	// AmenityPrison represents a prison or jail where people are incarcerated before trial or after conviction.
	AmenityPrison AmenityType = "prison"
	// AmenityRangerStation represents a National Park visitor headquarters: official park visitor facility with police, visitor information, permit services, etc.
	AmenityRangerStation AmenityType = "ranger_station"
	// AmenityTownhall represents a building where the administration of a village, town or city may be located, or just a community meeting place.
	AmenityTownhall AmenityType = "townhall"

	// Facilities

	// AmenityBBQ represents a BBQ or Barbecue, a permanently built grill for cooking food, most typically used outdoors by the public.
	AmenityBBQ AmenityType = "bbq"
	// AmenityBench represents a bench to sit down and relax a bit.
	AmenityBench AmenityType = "bench"
	// AmenityDogToilet represents an area designated for dogs to urinate and excrete.
	AmenityDogToilet AmenityType = "dog_toilet"
	// AmenityDressingRoom represents an area designated for changing clothes.
	AmenityDressingRoom AmenityType = "dressing_room"
	// AmenityDrinkingWater represents a place where humans can obtain potable water for consumption.
	AmenityDrinkingWater AmenityType = "drinking_water"
	// AmenityGiveBox represents a small facility where people drop off and pick up various types of items for free sharing and reuse.
	AmenityGiveBox AmenityType = "give_box"
	// AmenityLounge represents a comfortable waiting area for customers, usually found in airports and other transportation hubs.
	AmenityLounge AmenityType = "lounge"
	// AmenityMailroom represents a mailroom for receiving packages or letters.
	AmenityMailroom AmenityType = "mailroom"
	// AmenityParcelLocker represents a machine for picking up and sending parcels.
	AmenityParcelLocker AmenityType = "parcel_locker"
	// AmenityShelter represents a small shelter against bad weather conditions.
	AmenityShelter AmenityType = "shelter"
	// AmenityShower represents a public shower.
	AmenityShower AmenityType = "shower"
	// AmenityTelephone represents a public telephone.
	AmenityTelephone AmenityType = "telephone"
	// AmenityToilets represents public toilets (might require a fee).
	AmenityToilets AmenityType = "toilets"
	// AmenityWaterPoint represents a place where you can get large amounts of drinking water.
	AmenityWaterPoint AmenityType = "water_point"
	// AmenityWateringPlace represents a place where water is contained and animals can drink.
	AmenityWateringPlace AmenityType = "watering_place"

	// Waste Management

	// AmenitySanitaryDumpStation represents a place for depositing human waste from a toilet holding tank.
	AmenitySanitaryDumpStation AmenityType = "sanitary_dump_station"
	// AmenityRecycling represents recycling facilities (bottle banks, etc.).
	AmenityRecycling AmenityType = "recycling"
	// AmenityWasteBasket represents a single small container for depositing garbage that is easily accessible for pedestrians.
	AmenityWasteBasket AmenityType = "waste_basket"
	// AmenityWasteDisposal represents a medium or large disposal bin, typically for bagged up household or industrial waste.
	AmenityWasteDisposal AmenityType = "waste_disposal"
	// AmenityWasteTransferStation represents a waste transfer station, a location that accepts, consolidates and transfers waste in bulk.
	AmenityWasteTransferStation AmenityType = "waste_transfer_station"

	// Others

	// AmenityAnimalBoarding represents a facility where you, paying a fee, can bring your animal for a limited period of time (e.g. for holidays).
	AmenityAnimalBoarding AmenityType = "animal_boarding"
	// AmenityAnimalBreeding represents a facility where animals are bred, usually to sell them.
	AmenityAnimalBreeding AmenityType = "animal_breeding"
	// AmenityAnimalShelter represents a shelter that recovers animals in trouble.
	AmenityAnimalShelter AmenityType = "animal_shelter"
	// AmenityAnimalTraining represents a facility used for non-competitive animal training.
	AmenityAnimalTraining AmenityType = "animal_training"
	// AmenityBakingOven represents an oven used for baking bread and similar, for example inside a building=bakehouse.
	AmenityBakingOven AmenityType = "baking_oven"
	// AmenityClock represents a public visible clock.
	AmenityClock AmenityType = "clock"
	// AmenityCrematorium represents a place where dead human bodies are burnt.
	AmenityCrematorium AmenityType = "crematorium"
	// AmenityDiveCentre represents a dive center, the base location where sports divers usually start scuba diving or make dive guided trips.
	AmenityDiveCentre AmenityType = "dive_centre"
	// AmenityFuneralHall represents a place for holding a funeral ceremony, other than a place of worship.
	AmenityFuneralHall AmenityType = "funeral_hall"
	// AmenityGraveYard represents a smaller place of burial, often with a church nearby.
	AmenityGraveYard AmenityType = "grave_yard"
	// AmenityHuntingStand represents an open or enclosed platform used by hunters at an elevated height above the terrain.
	AmenityHuntingStand AmenityType = "hunting_stand"
	// AmenityInternetCafe represents a place providing internet services to the public.
	AmenityInternetCafe AmenityType = "internet_cafe"
	// AmenityKitchen represents a public kitchen in a facility for use by everyone or customers.
	AmenityKitchen AmenityType = "kitchen"
	// AmenityKneippWaterCure represents an outdoor foot bath facility, popular in German-speaking countries.
	AmenityKneippWaterCure AmenityType = "kneipp_water_cure"
	// AmenityLounger represents an object for people to lie down.
	AmenityLounger AmenityType = "lounger"
	// AmenityMarketplace represents a marketplace where goods and services are traded daily or weekly.
	AmenityMarketplace AmenityType = "marketplace"
	// AmenityMonastery represents the location of a monastery or a building in which monks and nuns live.
	AmenityMonastery AmenityType = "monastery"
	// AmenityMortuary represents a morgue or funeral home used for storing human corpses.
	AmenityMortuary AmenityType = "mortuary"
	// AmenityPhotoBooth represents a stand to create instant photos.
	AmenityPhotoBooth AmenityType = "photo_booth"
	// AmenityPlaceOfMourning represents a room or building where families and friends can come before the funeral to view the body.
	AmenityPlaceOfMourning AmenityType = "place_of_mourning"
	// AmenityPlaceOfWorship represents a church, mosque, temple, etc.
	AmenityPlaceOfWorship AmenityType = "place_of_worship"
	// AmenityPublicBath represents a location where the public may bathe in common, such as a Japanese onsen, Turkish bath, or hot spring.
	AmenityPublicBath AmenityType = "public_bath"
	// AmenityRefugeeSite represents a human settlement sheltering refugees or internally displaced persons.
	AmenityRefugeeSite AmenityType = "refugee_site"
	// AmenityVendingMachine represents a machine selling goods such as food, tickets, or newspapers.
	AmenityVendingMachine AmenityType = "vending_machine"
)

var validAmenityTypes = map[string]AmenityType{
	// Sustenance
	"bar":        AmenityBar,
	"biergarten": AmenityBiergarten,
	"cafe":       AmenityCafe,
	"fast_food":  AmenityFastFood,
	"food_court": AmenityFoodCourt,
	"ice_cream":  AmenityIceCream,
	"pub":        AmenityPub,
	"restaurant": AmenityRestaurant,

	// Education
	"college":            AmenityCollege,
	"dancing_school":     AmenityDancingSchool,
	"driving_school":     AmenityDrivingSchool,
	"first_aid_school":   AmenityFirstAidSchool,
	"kindergarten":       AmenityKindergarten,
	"language_school":    AmenityLanguageSchool,
	"library":            AmenityLibrary,
	"surf_school":        AmenitySurfSchool,
	"toy_library":        AmenityToyLibrary,
	"research_institute": AmenityResearchInstitute,
	"training":           AmenityTraining,
	"music_school":       AmenityMusicSchool,
	"school":             AmenitySchool,
	"traffic_park":       AmenityTrafficPark,
	"university":         AmenityUniversity,

	// Transportation
	"bicycle_parking":        AmenityBicycleParking,
	"bicycle_repair_station": AmenityBicycleRepairStation,
	"bicycle_rental":         AmenityBicycleRental,
	"bicycle_wash":           AmenityBicycleWash,
	"boat_rental":            AmenityBoatRental,
	"boat_sharing":           AmenityBoatSharing,
	"bus_station":            AmenityBusStation,
	"car_rental":             AmenityCarRental,
	"car_sharing":            AmenityCarSharing,
	"car_wash":               AmenityCarWash,
	"compressed_air":         AmenityCompressedAir,
	"vehicle_inspection":     AmenityVehicleInspection,
	"charging_station":       AmenityChargingStation,
	"driver_training":        AmenityDriverTraining,
	"ferry_terminal":         AmenityFerryTerminal,
	"fuel":                   AmenityFuel,
	"grit_bin":               AmenityGritBin,
	"motorcycle_parking":     AmenityMotorcycleParking,
	"parking":                AmenityParking,
	"parking_entrance":       AmenityParkingEntrance,
	"parking_space":          AmenityParkingSpace,
	"taxi":                   AmenityTaxi,
	"weighbridge":            AmenityWeighbridge,

	// Financial
	"atm":              AmenityATM,
	"payment_terminal": AmenityPaymentTerminal,
	"bank":             AmenityBank,
	"bureau_de_change": AmenityBureauDeChange,
	"money_transfer":   AmenityMoneyTransfer,
	"payment_centre":   AmenityPaymentCentre,

	// Healthcare
	"baby_hatch":      AmenityBabyHatch,
	"clinic":          AmenityClinic,
	"dentist":         AmenityDentist,
	"doctors":         AmenityDoctors,
	"hospital":        AmenityHospital,
	"nursing_home":    AmenityNursingHome,
	"pharmacy":        AmenityPharmacy,
	"social_facility": AmenitySocialFacility,
	"veterinary":      AmenityVeterinary,

	// Entertainment, Arts & Culture
	"arts_centre":       AmenityArtsCentre,
	"brothel":           AmenityBrothel,
	"casino":            AmenityCasino,
	"cinema":            AmenityCinema,
	"community_centre":  AmenityCommunityCentre,
	"conference_centre": AmenityConferenceCentre,
	"events_venue":      AmenityEventsVenue,
	"exhibition_centre": AmenityExhibitionCentre,
	"fountain":          AmenityFountain,
	"gambling":          AmenityGambling,
	"love_hotel":        AmenityLoveHotel,
	"music_venue":       AmenityMusicVenue,
	"nightclub":         AmenityNightclub,
	"planetarium":       AmenityPlanetarium,
	"public_bookcase":   AmenityPublicBookcase,
	"social_centre":     AmenitySocialCentre,
	"stage":             AmenityStage,
	"stripclub":         AmenityStripclub,
	"studio":            AmenityStudio,
	"swingerclub":       AmenitySwingerclub,
	"theatre":           AmenityTheatre,

	// Public Service
	"courthouse":     AmenityCourthouse,
	"fire_station":   AmenityFireStation,
	"police":         AmenityPolice,
	"post_box":       AmenityPostBox,
	"post_depot":     AmenityPostDepot,
	"post_office":    AmenityPostOffice,
	"prison":         AmenityPrison,
	"ranger_station": AmenityRangerStation,
	"townhall":       AmenityTownhall,

	// Facilities
	"bbq":            AmenityBBQ,
	"bench":          AmenityBench,
	"dog_toilet":     AmenityDogToilet,
	"dressing_room":  AmenityDressingRoom,
	"drinking_water": AmenityDrinkingWater,
	"give_box":       AmenityGiveBox,
	"lounge":         AmenityLounge,
	"mailroom":       AmenityMailroom,
	"parcel_locker":  AmenityParcelLocker,
	"shelter":        AmenityShelter,
	"shower":         AmenityShower,
	"telephone":      AmenityTelephone,
	"toilets":        AmenityToilets,
	"water_point":    AmenityWaterPoint,
	"watering_place": AmenityWateringPlace,

	// Waste Management
	"sanitary_dump_station":  AmenitySanitaryDumpStation,
	"recycling":              AmenityRecycling,
	"waste_basket":           AmenityWasteBasket,
	"waste_disposal":         AmenityWasteDisposal,
	"waste_transfer_station": AmenityWasteTransferStation,

	// Others
	"animal_boarding":   AmenityAnimalBoarding,
	"animal_breeding":   AmenityAnimalBreeding,
	"animal_shelter":    AmenityAnimalShelter,
	"animal_training":   AmenityAnimalTraining,
	"baking_oven":       AmenityBakingOven,
	"clock":             AmenityClock,
	"crematorium":       AmenityCrematorium,
	"dive_centre":       AmenityDiveCentre,
	"funeral_hall":      AmenityFuneralHall,
	"grave_yard":        AmenityGraveYard,
	"hunting_stand":     AmenityHuntingStand,
	"internet_cafe":     AmenityInternetCafe,
	"kitchen":           AmenityKitchen,
	"kneipp_water_cure": AmenityKneippWaterCure,
	"lounger":           AmenityLounger,
	"marketplace":       AmenityMarketplace,
	"monastery":         AmenityMonastery,
	"mortuary":          AmenityMortuary,
	"photo_booth":       AmenityPhotoBooth,
	"place_of_mourning": AmenityPlaceOfMourning,
	"place_of_worship":  AmenityPlaceOfWorship,
	"public_bath":       AmenityPublicBath,
	"refugee_site":      AmenityRefugeeSite,
	"vending_machine":   AmenityVendingMachine,
}

// Amenity is used to tag facilities used by visitors and residents. For example: toilets, telephones, banks, pharmacies, cafes, parking and schools.
// https://wiki.openstreetmap.org/wiki/Key:amenity
type Amenity Feature

func (a *Amenity) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var f Feature
	if err := unmarshal(&f); err != nil {
		return err
	}

	if f.Value != nil {
		val := strings.ToLower(*f.Value)
		if _, ok := validAmenityTypes[val]; ok {
			f.Value = helpers.Ptr(val)
		} else {
			validKeys := make([]string, 0, len(validAmenityTypes))
			for k := range validAmenityTypes {
				validKeys = append(validKeys, k)
			}
			sort.Strings(validKeys)
			return fmt.Errorf("invalid amenity value: %q (allowed: %s)", val, strings.Join(validKeys, ", "))
		}
	}
	f.Tag = "amenity"

	*a = Amenity(f)
	return nil
}

func (a *Amenity) GetTag() string {
	return "amenity"
}

func (a *Amenity) ToOQL() string {
	return Feature(*a).ToOQL()
}

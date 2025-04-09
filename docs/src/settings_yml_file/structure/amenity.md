# `Amenity`

A [`Feature`](./feature.md) that is used to tag facilities used by visitors and residents. For example: toilets, telephones, banks, pharmacies, cafes, parking and schools.

```yml
amenity:
  value: bar
  match: equals
```

Only the following `value`s are allowed:

### Sustenance

* `bar` - A purpose-built commercial establishment that sells alcoholic drinks to be consumed on the premises. They are characterised by a noisy and vibrant atmosphere, similar to a party and usually don't sell food.
* `biergarten` - An open-air area where alcoholic beverages along with food are prepared and served.
* `cafe` - An informal place that offers casual meals and beverages; typically, the focus is on coffee or tea. Also known as a coffeehouse/shop, bistro, or sidewalk cafe.
* `fast_food` - A fast food restaurant.
* `food_court` - An area with several different restaurant food counters and a shared eating area. Commonly found in malls, airports, etc.
* `ice_cream` - An ice cream shop or ice cream parlour. A place that sells ice cream and frozen yoghurt over the counter.
* `pub` - A place selling beer and other alcoholic drinks; may also provide food or accommodation (UK).
* `restaurant` - A restaurant (not fast food, see `fast_food`).

### Education

* `college` - A campus or buildings of an institute of Further Education (also known as continuing education).
* `dancing_school` - A dancing school or dance studio.
* `driving_school` - A driving school that offers motor vehicle driving lessons.
* `first_aid_school` - A place where people can go for first aid courses.
* `kindergarten` - A facility for children too young for a regular school (also known as preschool, playschool, or nursery school).
* `language_school` - An educational institution where one studies a foreign language.
* `library` - A public library (municipal, university, etc.) where you can borrow books.
* `surf_school` - An establishment that teaches surfing.
* `toy_library` - A place to borrow games and toys or play with them on-site.
* `research_institute` - An establishment endowed for doing research.
* `training` - A public place where you can receive training.
* `music_school` - A music school, an educational institution specializing in the study, training, and research of music.
* `school` - School and grounds - primary, middle, and secondary schools.
* `traffic_park` - Juvenile traffic schools.
* `university` - A university campus: an institute of higher education.

### Transportation

* `bicycle_parking` - Parking for bicycles.
* `bicycle_repair_station` - General tools for self-service bicycle repairs, usually on the roadside; no service.
* `bicycle_rental` - A place to rent a bicycle.
* `bicycle_wash` - A place to clean a bicycle.
* `boat_rental` - A place to rent a boat.
* `boat_sharing` - A place to share a boat.
* `bus_station` - A bus station.
* `car_rental` - A place to rent a car.
* `car_sharing` - A place to share a car.
* `car_wash` - A place to wash a car.
* `compressed_air` - A device to inflate tires (e.g. motorcar, bicycle).
* `vehicle_inspection` - A government vehicle inspection station.
* `charging_station` - A charging facility for electric vehicles.
* `driver_training` - A place for driving training on a closed course.
* `ferry_terminal` - A ferry terminal/stop, a place where people/cars/etc. can board and leave a ferry.
* `fuel` - A petrol station, gas station, or marine fuel station.
* `grit_bin` - A container that holds grit or a mixture of salt and grit.
* `motorcycle_parking` - Parking for motorcycles.
* `parking` - A parking area for vehicles.
* `parking_entrance` - An entrance or exit to an underground or multi-storey parking facility.
* `parking_space` - A single parking space within a car park.
* `taxi` - A place where taxis wait for passengers.
* `weighbridge` - A large weight scale to weigh vehicles and goods.

### Financial

* `atm` - An automated teller machine (ATM) or cashpoint: a device that provides the clients of a financial institution with access to financial transactions.
* `payment_terminal` - A self-service payment kiosk/terminal.
* `bank` - A bank or credit union: a financial establishment where customers can deposit and withdraw money, take loans, make investments, and transfer funds.
* `bureau_de_change` - A bureau de change, money changer, currency exchange, Wechsel, or cambio – a place to change foreign bank notes and travellers cheques.
* `money_transfer` - A place that offers money transfers, especially cash to cash.
* `payment_centre` - A non-bank place, where people can pay bills of public and private services and taxes.

### Healthcare

* `baby_hatch` - A place where a baby can be, out of necessity, anonymously left to be safely cared for and perhaps adopted.
* `clinic` - A medium-sized medical facility or health centre.
* `dentist` - A dentist practice/surgery.
* `doctors` - A doctor's practice/surgery.
* `hospital` - A hospital providing in-patient medical treatment, often used in conjunction with emergency facilities (A&E or ER).
* `nursing_home` - A home for disabled or elderly persons who need permanent care (discouraged tag, use amenity=social_facility + social_facility=nursing_home).
* `pharmacy` - A pharmacy: a shop where a pharmacist sells medications.
* `social_facility` - A facility that provides social services: group & nursing homes, workshops for the disabled, homeless shelters, etc.
* `veterinary` - A place where a veterinary surgeon (vet) practices.

### Entertainment, Arts & Culture

* `arts_centre` - A venue where a variety of arts are performed or conducted.
* `brothel` - An establishment specifically dedicated to prostitution.
* `casino` - A gambling venue with at least one table game (e.g., roulette, blackjack) that takes bets on sporting and other events at agreed-upon odds.
* `cinema` - A place where films are shown (US: movie theater).
* `community_centre` - A place mostly used for local events, festivities, and group activities; including special interest and special age groups.
* `conference_centre` - A large building used to hold a convention.
* `events_venue` - A building specifically used for organizing events.
* `exhibition_centre` - An exhibition centre.
* `fountain` - A fountain for cultural, decorative, or recreational purposes.
* `gambling` - A place for gambling. Games covered by this definition include bingo and pachinko.
* `love_hotel` - A short-stay hotel primarily for allowing guests privacy for sexual activities.
* `music_venue` - An indoor place to hear contemporary live music.
* `nightclub` - A place to drink and dance, often called "Disco" or "Discothek" in Germany (not to be confused with "Nachtclub" for strip clubs).
* `planetarium` - A planetarium.
* `public_bookcase` - A street furniture containing books. Take one or leave one.
* `social_centre` - A place for free and not-for-profit activities.
* `stage` - A raised platform for performers.
* `stripclub` - A place offering striptease or lap dancing (for sexual services, use `brothel`).
* `studio` - A TV, radio, or recording studio.
* `swingerclub` - A club where people meet to have a party and group sex.
* `theatre` - A theatre or opera house where live performances occur, such as plays, musicals, and formal concerts (use `cinema`) for movie theaters).

### Public Service

* `courthouse` - A building home to a court of law, where justice is dispensed.
* `fire_station` - A station of a fire brigade.
* `police` - A police station where police officers patrol from and that is a first point of contact for civilians.
* `post_box` - A box for the reception of mail.
* `post_depot` - A post depot or delivery office, where letters and parcels are collected and sorted prior to delivery.
* `post_office` - A post office building with postal services.
* `prison` - A prison or jail where people are incarcerated before trial or after conviction.
* `ranger_station` - A National Park visitor headquarters: official park visitor facility with police, visitor information, permit services, etc.
* `townhall` - A building where the administration of a village, town, or city may be located, or just a community meeting place.

### Facilities

* `bbq` - A BBQ or Barbecue, a permanently built grill for cooking food, most typically used outdoors by the public.
* `bench` - A bench to sit down and relax a bit.
* `dog_toilet` - An area designated for dogs to urinate and excrete.
* `dressing_room` - An area designated for changing clothes.
* `drinking_water` - A place where humans can obtain potable water for consumption.
* `give_box` - A small facility where people drop off and pick up various types of items for free sharing and reuse.
* `lounge` - A comfortable waiting area for customers, usually found in airports and other transportation hubs.
* `mailroom` - A mailroom for receiving packages or letters.
* `parcel_locker` - A machine for picking up and sending parcels.
* `shelter` - A small shelter against bad weather conditions.
* `shower` - A public shower.
* `telephone` - A public telephone.
* `toilets` - Public toilets (might require a fee).
* `water_point` - A place where you can get large amounts of drinking water.
* `watering_place` - A place where water is contained and animals can drink.

### Waste Management

* `sanitary_dump_station` - A place for depositing human waste from a toilet holding tank.
* `recycling` - Recycling facilities (bottle banks, etc.).
* `waste_basket` - A single small container for depositing garbage that is easily accessible for pedestrians.
* `waste_disposal` - A medium or large disposal bin, typically for bagged up household or industrial waste.
* `waste_transfer_station` - A waste transfer station, a location that accepts, consolidates, and transfers waste in bulk.

### Others

* `animal_boarding` - A facility where you can bring your animal for a limited period of time, usually for holidays.
* `animal_breeding` - A facility where animals are bred, usually to sell them.
* `animal_shelter` - A shelter that recovers animals in trouble.
* `animal_training` - A facility used for non-competitive animal training.
* `baking_oven` - An oven used for baking bread and similar items, often inside a bakehouse.
* `clock` - A public visible clock.
* `crematorium` - A place where dead human bodies are burnt.
* `dive_centre` - A dive center, the base location where sports divers start scuba diving or make guided trips.
* `funeral_hall` - A place for holding a funeral ceremony, other than a place of worship.
* `grave_yard` - A smaller place of burial, often with a church nearby.
* `hunting_stand` - An open or enclosed platform used by hunters at an elevated height above the terrain.
* `internet_cafe` - A place providing internet services to the public.
* `kitchen` - A public kitchen in a facility for use by everyone or customers.
* `kneipp_water_cure` - An outdoor foot bath facility, popular in German-speaking countries.
* `lounger` - An object for people to lie down.
* `marketplace` - A marketplace where goods and services are traded daily or weekly.
* `monastery` - The location of a monastery or a building where monks and nuns live.
* `mortuary` - A morgue or funeral home used for storing human corpses.
* `photo_booth` - A stand to create instant photos.
* `place_of_mourning` - A room or building where families and friends can view the body before the funeral.
* `place_of_worship` - A church, mosque, temple, etc.
* `public_bath` - A location where the public may bathe together, such as a Japanese onsen or Turkish bath.
* `refugee_site` - A human settlement sheltering refugees or internally displaced persons.
* `vending_machine` - A machine selling goods such as food, tickets, or newspapers.

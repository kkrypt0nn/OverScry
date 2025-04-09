# `Building`

A [`Feature`](./feature.md) that is used to tag individual buildings or groups of connected buildings.

```yml
building:
  value: apartments
  match: equals
  levels:
    value: 2
    match: equals
```

The following fields are allowed

## `architecture`

A [`Feature`](./feature.md) representing the architectural style of a building.

Only the following `value`s are allowed:

* `islamic`
* `mamluk`
* `romanesque`
* `gothic`
* `renaissance`
* `mannerism`
* `ottoman`
* `baroque`
* `rococo`
* `empire`
* `moorish revival`
* `neoclassicism`
* `georgian`
* `victorian`
* `historicism`
* `neo-romanesque`
* `neo-byzantine`
* `neo-gothic`
* `neo-renaissance`
* `neo-baroque`
* `art_nouveau`
* `eclectic`
* `functionalism`
* `cubism`
* `new_objectivity`
* `art_deco`
* `modern`
* `amsterdam_school`
* `international_style`
* `constructivism`
* `stalinist_neoclassicism`
* `brutalist`
* `postmodern`
* `contemporary`

## `fireproof`

A [`Feature`](./feature.md) representing the fire-resistance information.

Only the following `value`s are allowed:

* `yes`
* `no`

## `flats`

A [`Feature`](./feature.md) representing the number of residential units (flats, apartments) in an apartment building (`apartments`), residential building (`residential`), house (`house`), detached house (`detached`) or similar building.

## `levels`

A [`Feature`](./feature.md) representing the number of visible levels (floors) in the building as used in the [Simple 3D buildings](https://wiki.openstreetmap.org/wiki/Simple_3D_Buildings) scheme.

## `soft_storey`

A [`Feature`](./feature.md) representing a building where any one level is significantly more flexible (less stiff) than those above and below it.

Only the following `value`s are allowed:

* `yes`
* `no`
* `reinforced`

## `value`

Only the following `value`s are allowed:

### Accomodation

* `apartments` - A building arranged into individual dwellings, often on separate floors. May also have retail outlets on the ground floor.
* `barracks` - Buildings built to house military personnel or laborers.
* `bungalow` - A single-storey detached small house, often called a Dacha.
* `cabin` - A small, roughly built house, typically with a wood exterior and usually found in rural areas. [Learn more](https://en.wikipedia.org/wiki/Cabin)
* `detached` - A detached house, a free-standing residential building usually housing a single family.
* `annexe` - A small self-contained apartment, cottage, or small residential building on the same property as the main residential unit.
* `dormitory` - A shared building intended for college/university students.
* `farm` - A residential building on a farm (farmhouse).
* `ger` - A permanent or seasonal round yurt, also known as a ger.
* `hotel` - A building designed with separate rooms available for overnight accommodation.
* `house` - A dwelling unit inhabited by a single household.
* `houseboat` - A boat used primarily as a home.
* `residential` - A general tag for a building used primarily for residential purposes.
* `semidetached_house` - A residential house that shares a common wall with another on one side.
* `static_caravan` - A mobile home that is semi-permanently left on a single site.
* `stilt_house` - A building raised on piles over the surface of the soil or a body of water.
* `terrace` - A linear row of residential dwellings, each of which normally has its own entrance.
* `tree_house` - A small hut or room built on tree posts or on a natural tree.
* `trullo` - A stone hut with a conical roof.

### Commercial

* `commercial` - A building for non-specific commercial activities, not necessarily an office building.
* `industrial` - A building for industrial purposes. Use `warehouse` if the purpose is known to be primarily for storage/distribution.
* `kiosk` - A small one-room retail building.
* `office` - An office building.
* `retail` - A building primarily used for selling goods that are sold to the public.
* `supermarket` - A building constructed to house a self-service large-area store.
* `warehouse` - A building primarily intended for the storage of goods or as part of a distribution system.

### Religious

* `religious` - A building that is related to religion, but not specific to any particular structure type.
* `cathedral` - A building that was built as a cathedral.
* `chapel` - A building that was built as a chapel.
* `church` - A building that was built as a church.
* `kingdom_hall` - A building that was built as a [Kingdom Hall](https://en.wikipedia.org/wiki/Kingdom_Hall) for Jehovah's Witnesses.
* `monastery` - A building constructed as a [monastery](https://en.wikipedia.org/wiki/Monastery), often with multiple structures.
* `mosque` - A building that was erected as a mosque.
* `presbytery` - A building where priests live and work.
* `shrine` - A building that was built as a shrine.
* `synagogue` - A building that was built as a synagogue.
* `temple` - A building that was built as a temple.

### Civic/amenity

* `bakehouse` - A building built for baking bread, often used with `amenity.baking_oven`.
* `bridge` - A building used as a bridge, often referred to as a skyway.
* `civic` - A building created to house civic amenities such as community centres, libraries, and town halls.
* `college` - A building designed for a college.
* `fire_station` - A building constructed to house firefighting equipment and personnel.
* `government` - A government building, including municipal and regional offices.
* `gatehouse` - An entry control point building, typically located at the entrance to a city or compound.
* `hospital` - A building erected for use as a hospital.
* `kindergarten` - A building used as a kindergarten.
* `museum` - A building designed as a museum.
* `public` - A building constructed to be accessible to the general public, such as town halls or police stations.
* `school` - A building constructed as a school.
* `toilets` - A toilet block building.
* `train_station` - A building constructed to be a train station, even if now repurposed.
* `transportation` - A building related to public transport, often tagged with transport-related tags.
* `university` - A building designed for a university.

### Agricultural/plant production

* `barn` - An agricultural building used for storage and as a covered workplace.
* `conservatory` - A building or room with glass or tarpaulin roofing, used as an indoor garden or sunroom.
* `cowshed` - A building for housing cows, commonly found on farms.
* `farm_auxiliary` - A building on a farm that is not a dwelling, such as barns or storage structures.
* `greenhouse` - A [greenhouse](https://en.wikipedia.org/wiki/Greenhouse), a structure used for growing plants, typically covered with glass or plastic.
* `slurry_tank` - A circular building used to hold slurry (liquid animal excreta).
* `stable` - A building constructed as a stable for housing horses.
* `sty` - A building for raising domestic pigs, often found on farms. [Learn more](https://en.wikipedia.org/wiki/Sty)
* `livestock` - A building for housing or raising other livestock (apart from cows, horses, or pigs).

### Sports

* `grandstand` - The main stand at racecourses or sports grounds, often roofed and commanding the best view for spectators.
* `pavilion` - A sports [pavilion](https://en.wikipedia.org/wiki/Pavilion), typically with changing rooms, storage areas, and possibly a space for functions & events.
* `riding_hall` - A building constructed specifically as a riding hall.
* `sports_hall` - A building constructed specifically as a sports hall.
* `sports_centre` - A building constructed as a sports centre, often with various sports facilities.
* `stadium` - A building constructed as a stadium, including those that may have been abandoned and repurposed.

### Storage

* `allotment_house` - A small outbuilding for short visits in an allotment garden.
* `boathouse` - A building used for the storage of boats.
* `hangar` - A building used for the storage of airplanes, helicopters, or spacecraft. [Learn more](https://en.wikipedia.org/wiki/Hangar)
* `hut` - A small and crude shelter, possibly used as a shed or residential building of low quality. [Learn more](https://en.wikipedia.org/wiki/Hut)
* `shed` - A simple, single-storey structure often used for storage, hobbies, or as a workshop. [Learn more](https://en.wikipedia.org/wiki/Shed)

### Cars

* `carport` - A covered structure used to offer limited protection to vehicles, usually without four walls. [Learn more](https://en.wikipedia.org/wiki/Carport)
* `garage` - A building suitable for the storage of one or more motor vehicles. [Learn more](https://en.wikipedia.org/wiki/Garage_(residential))
* `garages` - A building consisting of multiple discrete storage spaces for different owners or tenants.
* `parking` - A structure purpose-built for parking cars.

### Power/technical buildings

* `digester` - A bioreactor used for the production of biogas from biomass.
* `service` - A small unmanned building with machinery like pumps or transformers.
* `tech_cab` - A small prefabricated cabin used for the air-conditioned accommodation of technology.
* `transformer_tower` - A tall building comprising a distribution transformer, often connected to a power line.
* `water_tower` - A water tower used for storing water at height for distribution.
* `storage_tank` - A container that holds liquids, such as a storage tank.
* `silo` - A building for storing bulk materials, such as grain or other agricultural products.

### Other

* `beach_hut` - A small, usually wooden cabin or shelter on popular bathing beaches.
* `bunker` - A hardened military building.
* `castle` - A building constructed as a castle.
* `construction` - A building under construction.
* `container` - A permanent building made from a container.
* `guardhouse` - A post or small building used as a guardhouse.
* `military` - A military building.
* `outbuilding` - A less important building on the same land as a larger building.
* `pagoda` - A building constructed as a pagoda.
* `quonset_hut` - A lightweight prefabricated structure in the shape of a semicircle.
* `roof` - A structure consisting of a roof with open sides, such as a rain shelter.
* `ruins` - A building that is abandoned and in poor repair.
* `ship` - A decommissioned ship or submarine that stays in one place.
* `tent` - A permanently placed tent.
* `tower` - A tower building.
* `triumphal_arch` - A free-standing monumental structure in the shape of an archway.
* `windmill` - A traditional windmill used to mill grain.
* `yes` - A building type that can't be more specifically determined.

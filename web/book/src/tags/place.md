# `Place`

A [`Feature`](../settings/feature.md) that is used to give details about settlements.

```yml
place:
  value: region
  match: equals
  levels:
    value: 2
    match: equals
```

The following fields are allowed

## `population`

A [`Feature`](../settings/feature.md) that indicates a rough number of citizens in a given place.

Any number is allowed.

## `is_in`

A [`Feature`](../settings/feature.md) used to index where a place or feature is.

Some valid examples are:

* `California;CA;USA`
* `USA;CA;California;San Francisco`
* `San Francisco`

## `value`

Only the following `value`s are allowed:

* `country` - A nation state or other high-level national political/administrative area.
* `state` - A large sub-national political/administrative area.
* `region` - Used both as a broad tag for geographic or historical areas with no clear boundary and for distinct administration areas (with specific boundaries) in some countries.
* `province` - A subdivision of a country similar to a state.
* `district` - A district – a type of administrative division that, in some countries, is managed by local government.
* `county` - A county - a geographical region of a country.
* `subdistrict` - A subdistrict - a subdivision of a district used for administrative or other purposes.
* `municipality` - A municipality - single urban administrative division having corporate status.
* `city` - The largest urban settlement or settlements within the territory.
* `borough` - A part in a larger city grouped into administrative unit.
* `suburb` - A part of a town or city with a well-known name and often a distinct identity.
* `quarter` - A quarter is a named, geographically localised place within a suburb of a larger city or within a town, which is bigger than a neighbourhood.
* `neighbourhood` - A neighbourhood is a smaller named, geographically localised place within a suburb of a larger city or within a town or village.
* `city_block` - A named city block, usually surrounded by streets.
* `plot` - A named plot is a tract or parcel of land owned or meant to be owned by some owner.
* `city` - The largest urban settlement or settlements within the territory.
* `town` - An important urban centre, between a village and a city in size.
* `village` - A smaller distinct settlement, smaller than a town with few facilities available with people traveling to nearby towns to access these.
* `hamlet` - A smaller rural community, typically with fewer than 100-200 inhabitants, and little infrastructure.
* `isolated_dwelling` - The smallest kind of settlement (1-2 households).
* `farm` - An individually named farm.
* `allotments` - A separate settlement, which is located outside an officially inhabited locality and has its own addressing.
* `continent` - One of the seven continents: Africa, Antarctica, Asia, Europe, North America, Oceania, South America.
* `island` - Any piece of land that is completely surrounded by water and isolated from other significant landmasses (bigger than 1 km²).
* `islet` - A very small island (smaller than 1 km²).
* `square` - A town or village square: a (typically) paved open space, generally of architectural significance, which is surrounded by buildings in a built-up area such as a city, town or village.
* `locality` - A named place that has no population.
* `polder` - A land area that forms an artificial hydrological entity enclosed by embankments and usually is under sea level.
* `sea` - A large body of salt water part of, or connected to, an ocean.
* `ocean` - The world's five main major oceanic divisions.

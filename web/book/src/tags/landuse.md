# `Landuse`

A [`Feature`](../settings/feature.md) that is used to tag the purpose for which an area of land is being used.

```yml
landuse:
  value: commercial
  match: equals
```

Only the following `value`s are allowed:

### Developed land

* `commercial` - Predominantly commercial businesses and their offices. Commercial businesses which sell goods should be categorised as `retail`. Commercial businesses can sell services on site and may include private Doctor's Surgeries, and those non-government services for mental and physical health, such as a Counselor's or Physiotherapist's practice or Veterinary. Commercial businesses can also include office buildings and business parks which have limited interface with the public and sell their services either on site, or externally. Commercial businesses have low amounts of public foot traffic.
* `construction` - A site which is under active development and construction of a building or structure, including any purposeful alteration to the land or vegetation upon it. Abandoned construction projects and sites should not use this tag.
* `education` - An area predominately used for educational purposes/facilities.
* `fairground` - A site where a [fair](https://en.wikipedia.org/wiki/Fair) takes place.
* `industrial` - Predominantly industrial landuses such as workshops, factories, or warehouses.
* `residential` - Land where people reside; predominantly residential detached (single houses, grouped dwellings), or attached (apartments, flats, units) dwellings. For "Mixed-Use" areas where more than half of the land is residential, tag as residential.
* `retail` - Predominantly retail businesses such as shops. Retail businesses sell physical goods such as food (prepared or grocery), clothing, medicine, stationary, appliances, tools, or other similar physical items. Retail businesses have high amounts of public foot traffic. Retail businesses do not exclusively provide or sell their services. For businesses which sell services see `commercial`.
* `institutional` - Land used for institutional purposes, see [Institution (disambiguation)](https://en.wikipedia.org/wiki/Institution_(disambiguation)).

### Rural and agricultural land

* `aquaculture` - A [piece of land](https://en.wikipedia.org/wiki/Aquaculture) for the farming of freshwater and saltwater organisms such as finfish, molluscs, crustaceans and aquatic plants.
* `allotments` - A piece of land given over to local residents for growing vegetables and flowers.
* `forest` - Managed forest or woodland plantation. Some use this to map an area of trees rather than the use of the land.
* `meadow` - A meadow or pasture: land primarily vegetated by grass and non-woody plants, mainly used for hay or grazing.
* `orchard` - intentional planting of trees or shrubs maintained for food production.
* `vineyard` - A piece of land where grapes are grown.

### Waterbody

* `basin` - An area artificially graded to hold water.
* `reservoir` - A [reservoir](https://en.wikipedia.org/wiki/Reservoir) on Wikipedia.

### Other

* `brownfield` - Describes land scheduled for new development where old buildings have been demolished and cleared.
* `cemetery` - Place for burials. Smaller places (e.g. with a church nearby) may use `amenity.grave_yard` instead.
* `grass` - An area of mown and managed grass not otherwise covered by a more specific tag.
* `greenfield` - Describes land scheduled for new development where there have been no buildings before. A greenfield is scheduled to turn into a construction site.
* `landfill` - Place where waste is dumped.
* `military` - For land areas owned/used by the military for whatever purpose.
* `quarry` - Surface mineral extraction.
* `recreation_ground` - An open green space for general recreation, which may include pitches, nets and so on, usually municipal but possibly also private to colleges or companies.
* `religious` - An area used for religious purposes.
* `village_green` - A village green is a distinctive area of grassy public land in a village centre. Not a generic tag for urban greenery. It is a typical English term – defined separately from 'common land' under the Commons Registration Act 1965 and the Commons Act 2006.
* `winter_sports` - An area dedicated to winter sports (e.g. skiing).

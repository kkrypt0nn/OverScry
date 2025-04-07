# `Addr`

An addr is used to provide address information for a building, facility or any other object that has an address. Example:

```yml
addr:
  housenumber:
    value: 1337
    match: equals
```

The following fields are allowed

### `housenumber`

A [`Feature`](./feature.md) representing the house number (may contain letters, dashes or other characters).

### `housename`

A [`Feature`](./feature.md) representing the name of a house.

### `flats`

A [`Feature`](./feature.md) representing the unit numbers (a range or a list) of the flats or apartments located behind a single entrance door.

### `conscriptionnumber`

A [`Feature`](./feature.md) representing a special kind of housenumber relates to a settlement instead of a street.

Conscription numbers were introduced in the Austro-Hungarian Empire and are still in use in some parts of Europe, sometimes together with street-related housenumbers which are also called orientation numbers.

### `street`

A [`Feature`](./feature.md) representing the name of the respective street. If the street name is very long or nonexistent, the ref of the respective street.

### `place`

A [`Feature`](./feature.md) representing a part of an address which refers to the name of some territorial zone (usually an island, square or very small village) instead of a street. Should not be used together with [`street`](#street).

### `postcode`

A [`Feature`](./feature.md) representing the postal code of the building/area.

### `city`

A [`Feature`](./feature.md) representing the name of the city as given in postal addresses of the building/area.

### `country`

A [`Feature`](./feature.md) representing the ISO 3166-1 alpha-2 two letter country code in upper case.

Example: "DE" for Germany, "CH" for Switzerland, "AT" for Austria, "FR" for France, "IT" for Italy.

### `postbox`

A [`Feature`](./feature.md) representing the postal service [Post Office Box](https://en.wikipedia.org/wiki/Post_office_box) as alternative to addressing using street names.

Example: "PO Box 34"

### `full`

A [`Feature`](./feature.md) representing a full-text, often multi-line, address if you find the structured address fields unsuitable for denoting the address of this particular location. Examples: "Fifth house on the left after the village oak, Smalltown, Smallcountry", or addresses using special delivery names or codes (possibly via an unrelated city name and post code), or PO Boxes.

## For countries using hamlet, subdistrict, district, province, state, county

### `hamlet`

A [`Feature`](./feature.md) representing the hamlet of the object. In France, some addresses use hamlets instead of street names, use the generic [`place`](#place) instead.

### `suburb`

A [`Feature`](./feature.md) representing the name of the settlement. If an address exists several times in a city you have to add it. See Australian definition of [suburb](https://en.wikipedia.org/wiki/Suburb).

### `subdistrict`

A [`Feature`](./feature.md) representing the [subdistrict](https://en.wikipedia.org/wiki/Subdistrict) of the object.

### `district`

A [`Feature`](./feature.md) representing the [district](https://en.wikipedia.org/wiki/District) of the object.

### `province`

A [`Feature`](./feature.md) representing the [province](https://en.wikipedia.org/wiki/Province) of the object.

### `state`

A [`Feature`](./feature.md) representing the [state](https://en.wikipedia.org/wiki/Administrative_division) of the object. For the US, uppercase two-letter postal abbreviations (AK, CA, HI, NY, TX, WY, etc.) are used.

### `county`

A [`Feature`](./feature.md) representing the [county](https://en.wikipedia.org/wiki/County) of the object.

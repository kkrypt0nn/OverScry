# `Aeroway`

A [`Feature`](../settings/feature.md) that is used to tag aerodromes, airfields other ground facilities that support the operation of airplanes and helicopters.

```yml
aeroway:
  value: aerodrome
  match: equals
```

Only the following `value`s are allowed:

* `aerodrome` - An aerodrome, airport or airfield.
* `aircraft_crossing` - A point where the flow of traffic is impacted by crossing aircraft.
* `gate` - The bounded space, inside the airport terminal, where passengers wait before boarding their flight.
* `helipad` - A landing area or platform designed for helicopters.
* `heliport` - A special aerodrome built for helicopters.
* `navigationaid` - A facility that supports visual navigation for aircraft.
* `spaceport` - A spaceport or cosmodrome: a site for launching or receiving spacecraft.
* `terminal` - An airport passenger building.
* `windsock` - An object that shows wind direction and speed.

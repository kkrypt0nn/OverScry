# `PublicTransport`

A [`Feature`](../settings/feature.md) that is used to tag features related to public transport. For example: railway stations, bus stops and services.

```yml
public_transport:
  value: stop_position
  match: equals
```

Only the following `value`s are allowed:

* `stop_position` - The position on the street or rails where a public transport vehicle stops.
* `platform` - The place where passengers are waiting for the public transport vehicles.
* `station` - A station is an area designed to access public transport.

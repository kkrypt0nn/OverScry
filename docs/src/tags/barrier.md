# `Barrier`

A [`Feature`](../settings/feature.md) that is used to tag barriers and obstacles that are usually involved by traveling.

```yml
barrier:
  value: kerb
  match: equals
```

Only the following `value`s are allowed:

### Linear barriers

* `kerb` - A stone edging to a pavement or raised path. The right side is considered the bottom, and the left side is the top.

### Access control on highways

* `block` - One or more large immobile block(s) usually barring free access along a way.
* `bollard` - One or more solid (usually concrete or metal) pillar(s) used to control traffic.
* `border_control` - A control point at an international border between two countries.
* `bump_gate` - A drive-through gate used in rural areas to provide a barrier to livestock that does not require the driver to exit the vehicle.
* `bus_trap` - A short section of the roadway where there is a deep dip in the middle to prevent passage by some traffic.
* `cattle_grid` - A hole in the road surface covered in a series of bars that allow wheeled vehicles but not animals to cross.
* `chain` - A chain used to prevent motorised vehicles.
* `cycle_barrier` - A barrier along a path that slows or prevents access for bicycle users.
* `debris` - Debris blocking a road.
* `entrance` - An opening or gap in a barrier.
* `full-height_turnstile` - A full-height turnstile.
* `gate` - A section in a wall or fence which can be opened to allow access.
* `hampshire_gate` - A section of wire fence which can be removed temporarily.
* `height_restrictor` - A height restrictor which prevents access of vehicles higher than a set limit.
* `horse_stile` - A horse stile allows pedestrians and horses to cross a gap through a fence.
* `jersey_barrier` - A barrier made of heavy prefabricated blocks.
* `kissing_gate` - A type of gate where you have to go into an enclosure and open a gate to get through.
* `lift_gate` - A bar or pole pivoted (rotates upwards to open) in such a way as to allow the boom to block vehicular access through a controlled point.
* `log` - A log blocking access.
* `motorcycle_barrier` - A barrier along a path preventing access by motorcycles.
* `rope` - A flexible barrier made of rope.
* `sally_port` - A secure, controlled entryway to a fortification or prison.
* `spikes` - Spikes on the ground preventing unauthorized access.
* `stile` - A structure which provides people a passage through or over a boundary via steps, ladders or narrow gaps.
* `sump_buster` - A barrier to stop cars (two tracked vehicles with less than a certain ground clearance and width between tracks).
* `swing_gate` - A gate consisting of a bar or pole pivoted (rotates sidewards to open) in such a way as to allow the boom to block vehicular access through a controlled point.
* `toll_booth` - A place where a road usage toll or fee is collected.
* `turnstile` - A turnstile, a passage on foot designed to allow one person at a time to pass.
* `yes` - An unspecified barrier. If possible, use a more specific value.

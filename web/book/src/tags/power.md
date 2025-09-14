# `Power`

A [`Feature`](../settings/feature.md) that is used to tag electrical power generation and distributions systems.

```yml
power:
  value: catenary_mast
  match: equals
```

Only the following `value`s are allowed:

* `catenary_mast` - A pole supporting the overhead wires used to supply electricity to vehicles equipped with a pantograph such as trams and trains.
* `compensator` - A static power device used to ensure power quality and electrical network resilience.
* `connection` - A free-standing electrical connection between two or more power lines or cables.
* `converter` - A device to convert power between alternating and direct current electrical power: often, but not only, over high voltage networks.
* `generator` - A device which converts one form of energy to another, for example, an electrical generator.
* `heliostat` - A mirror of a heliostat device.
* `insulator` - An electrical insulator which connects a power line to a support structure and prevents grounding.
* `inverter` - A device to convert power from direct current to alternating current.
* `pole` - A single pole supporting power lines, often a wood, steel, or concrete mast designed to carry minor power lines.
* `portal` - A supporting structure for power lines, composed of vertical legs with cables between them attached to a horizontal crossarm.
* `substation` - A facility which controls the flow of electricity in a power network with transformers, switchgear or compensators.
* `switch` - A device which allows electrical network operators to power up & down lines and transformers in substations or along the power grid.
* `terminal` - A point of connection where an overhead power line ends on a building or wall; for example, when connecting it to an indoor substation.
* `tower` - A tower or pylon carrying high voltage electricity cables. Often constructed from steel latticework but tubular or solid pylons are also used.
* `transformer` - A device for stepping up or down electric voltage. Large power transformers are typically located inside substations.

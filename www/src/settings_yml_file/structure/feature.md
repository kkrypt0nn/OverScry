# `Feature`

An example feature with the supported fields looks like

```yml
value: 1337
match: equals
```

The following fields are allowed

### `value`

A string representing the value of the feature.

### `match`

The **matching method** of the feature, which can be one of the following:

* `equals` (`=`), default
* `not_equals` (`!=`)
* `regex` (`~`)

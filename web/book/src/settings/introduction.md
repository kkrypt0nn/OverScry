# Settings YML File

The settings YML file is the way to tell OverScry what specific features/tags, areas, etc. should use when generating your queries.

A lot of work is put into so that more arguments and features/tags are supported. Please be patient for upcoming changes.

## Structure

An example YML file looks like

```yml
version: 0.0.7-dev
author: "Krypton (@kkrypt0nn)"
description: "A query to get apartments with 2 levels"
node:
  building:
    value: apartments
    levels:
      value: 2
```

The following fields are allowed

### `version`

A string representing the version of OverScry that supports the YML file structure.

### `author`

A string representing the author of the YML file.

### `description`

A string describing the YML file and the resulting generated query.

### `node`

A [`Node`](./node.md) with its fields.

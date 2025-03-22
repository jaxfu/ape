# `Props`

Describes and gives a name to the shape of a field's value.

### **Structure:**

- ...metadata
- `type` [***required***]
  - `"int"`
  - `"uint"`
  - `"float"`
  - `"text"`
  - `"bool"`
  - `"blob"`
  - `"map"`
- `array` [optional | default=`false`]
  - `true` | `false`
- `opts` [optional | default=*type-specific*]
  - ...type-specific fields

### **Example:**

TOML:

```toml
name = "Username"
category = "Users"
description = "alphanumeric unique user identifier"
type = "text"

[opts]
max_length = 16
alnum = true
```

JSON:

```json
{
  "category": "Users",
  "description": "alphanumeric unique user identifier",
  "name": "Username",
  "type": "text",
  "opts": {
    "max_length": 16,
    "alnum": true
  }
}
```

---

### _Types_:

Types are used to describe the shape of the value of a `Prop`. These are
the core types:

- `int`: signed integer
- `uint`: unsigned integer
- `float`: floating point number
- `text`: utf-8 string
- `bool`: boolean
- `blob`: arbitrary chunk of binary data
- `map`: `key`:`value` structure with dynamic `keys` and `values`

---

### _Opts_:

Only the core type needs to be specified, but the `opts` object allows for further specificity to be given. When generating code, the highest possible level of type specificity for the target language will be used.

- `int`:

  - size: [`8` | `16` | `32` | `64`]
  - min: int
  - max: int

- `uint`:

  - size: [`8` | `16` | `32` | `64`]
  - min: int
  - max: int

- `float`:

  - precision: [`single` | `double`]
  - size: [`8` | `16` | `32` | `64`]
  - min: float
  - max: float

- `text`:

  - min_chars: int
  - max_chars: int
  - regex: string
  - alpha (only alphabetic): [`true` | `false`]
  - alnum (only alphanumeric): [`true` | `false`]
  - num (only numeric): [`true` | `false`]

- `blob`:

  - min_size (in bytes): int
  - max_size (in bytes): int

- ***

### _Arrays_:

The `array` field defines whether the prop's value should be an array of the type. It defaults to `false` if not specified.

---

### References

`Prop` references copy the type data of the referenced `Prop` under the name declared in the `Prop` field.

- e.g. `user_id = "@Users.Props.Username"` would use the `Username` type data under the key `user_id`.

`Props` are referenced by category first, like so:
`@CATEGORY_NAME`.`Props`.`Prop_Name`

- e.g. `@Main.Props.Date`, `@Users.Props.Username`

---

### **Child Props:**

Props can declare child props to an indefinite depth like so:

```toml
[props.prop_name.child_prop_name...]
type = "int|float|text|bool|blob"
array = "true|false"
```

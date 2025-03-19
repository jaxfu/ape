# `Props`

Describes and gives a name to the shape of a field's

### **Structure:**

```toml
[props.prop_name]
type = "int|float|text|bool|blob"
array = "true|false"
description = "prop description"
# ... type-specific
```

These are the "type" fields that all props share:

- `type` **[required]**

  - `int`: integer
  - `float`: floating point number
  - `text`: utf-8 string
  - `bool`: boolean
  - `blob`: arbitrary chunk of binary data

- `array` **[optional | default=false]**

  - `true` or `false`
  - declares if prop is an array
  - if not given, will default to false

- `description` **[optional]**

  - short description of prop

---

### **Child Props:**

Props can declare child props to an indefinite depth like so:

```toml
[props.prop_name.child_prop_name...]
type = "int|float|text|bool|blob"
array = "true|false"
```

---

### **Arrays:**

To declare a prop as an array of the specified type, declare the field `array` as `true`. You can optionally add constraints when `array` is `true`:

```toml
[props.prop_name]
type = "float"

# optional
array = "true"
array_min_length = "some_int"
array_max_length = "some_int"
```

---

### References

`Prop` references copy the type data of the referenced `Prop` under the name declared in the `Prop` field.

- e.g. `user_id = "@Users.Props.Username"` would use the `Username` type data under the key `user_id`.

`Props` are referenced by category first, like so:
`@CATEGORY_NAME`.`Props`.`Prop_Name`

- e.g. `@Main.Props.Date`, `@Users.Props.Username`

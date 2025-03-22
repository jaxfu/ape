# Ape: Simplifying API Design

<p align="center">
	<img src="docs/ape.svg">
</p>

---

### v0.1 Notes:

- **_Ape_** currently uses `TOML` as the main definition language. `JSON` is also usable, but the docs are designed around use of `TOML`.
- By v0.5, **_Ape_** will use its own configuration language, which will be a more concise version of `TOML`.

---

### **Ape** is...

- used to describe the behavior of an API by defining its data schemas, routes and interactions between them.

- built following the idea of composabililty, where each structure is composed of those
  beneath it in generality (i.e. `Props` are composed of `Types`, `Objects` are composed of `Props`, etc...).

- built to be portable, meaning it can be used to generate other popular forms of data schema documentation, such as `Protobuf` and `OpenApi`, as well as generate relevant CRUD `SQL` commands and data structures in any language that has an available driver.

- compatible with [`github.com/jaxfu/tisane`](https://github.com/jaxfu/tisane) (wip), which can spin up an entire backend from an **_Ape_** project definition.

---

### _Core Components_

These are the `4 core components` of **_Ape_**, ordered by increasing generality/composability:

1. `Props`
2. `Objects`
3. `Routes`
4. `Actions`

---

### _Metadata_:

All components contain the "metadata" fields in their top level, which are:

- `name` [***required***]
  - alphanumeric identifier for the component
  - utf-8
  - must be case-insensitively unique among the same kind of component
    in the same category
- `category` [optional | default=NULL]
  - top-level organization for components
- `description` [optional | default=NULL]
  - utf-8 string describing the component

---

### 1. `Props`:

`Props` define the shape of the `key`:`value` pairs used as fields to construct an `Object`. All `Props` contain the metadata fields, `type` and `array` fields. The `opts` field is an object consisting of further specificity and constraints specific to each type.

Structure:

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

Example:

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

### 2. `Objects`:

A piece of structured data consisting of a collection of `Props`. It consists of the metadata fields, and an object that contains the `Props` that make up the `Object`. `Props` on an `Object` can either have their type data defined directly, or they can reference existing `Props` to copy their type data.

Structure:

- ...metadata
- `Props` [optional | default=NULL]

Example:

- TOML:

```toml
name = "Todo"
category = "Main"
description = "Todo object created by users"

[props]
user_id = "@Users.Props.Username" # Prop reference

[props.message]
type = "text"
max_chars = 256

[props.due_date]
type = "int"
signed = false

[props.notifications]
type = "text"
array = true
max_chars = 128
```

- JSON:

```json
{
  "name": "Todo",
  "category": "Main",
  "description": "Todo object created by users",
  "props": {
    "user_id": "@Users.Types.Username",
    "message": {
      "type": "text",
      "max_chars": 256
    },
    "due_date": {
      "type": "int",
      "signed": false
    },
    "notifications": {
      "type": "text",
      "array": true,
      "max_chars": 128
    }
  }
}
```

---

### _Prop References_:

- `Props` are referenced by category first, like so:
  `@CATEGORY_NAME`.`Props`.`Prop_Name`
  - e.g. `@Main.Props.Date`, `@Users.Props.Username`

---

### 3. `Routes`:

Structure:

- ...metadata
- `request`:
  - `headers`
  - `props`
- `responses`
  - `response` objects contain two components:
    - `status_code`
    - `props`

Example:

- TOML:

```toml
name = "Create"
category = "Todos"
description = "user submits new Todo"

[request]
url = "/todos"
method = "POST"
props = "@Main.Objects.Todo"

[request.headers]
Authorization = "BEARER auth_token"
Content-Type = "application/json"



[responses]

[responses.success]
status_code = 201
[responses.success.props]
todo_id = "@Todos.Props.todo_id"

[responses.failure]
status_code = 400
[responses.failure.props.error_message]
type = "text"
max_length = 128
```

- JSON:

```json
{
  "name": "Create",
  "category": "Todos",
  "description": "user submits new Todo",
  "request": {
    "url": "/todos",
    "method": "POST",
    "headers": {
      "Authorization": "BEARER auth_token",
      "Content-Type": "application/json"
    },
    "props": "@Main.Objects.Todo"
  },
  "responses": {
    "failure": {
      "status_code": 400,
      "props": {
        "error_message": {
          "max_length": 128,
          "type": "text"
        }
      }
    },
    "success": {
      "status_code": 201,
      "props": {
        "id": "@Todos.Props.todo_id"
      }
    }
  }
}
```

---

### 4. `Actions`:

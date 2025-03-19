# `Routes`

Structures containing metadata, an object defining the expected shape of the `request`, and an object defining the shape of the different possible `responses`.

### **Structure:**

- ...metadata
- `request` [optional | default=NULL]:
  - `url` [optional | default="/"]
  - `method` [optional | default="GET"]
  - `headers` [optional | default=NULL]
  - `props` [optional | default=NULL]
- `responses` [optional | default=NULL]:
  - `key`:`response` pairs
  - `response` objects contain two components:
    - `status_code` [***required***]
    - `props` [optional | default=NULL]

### **Example:**

```toml
name = "Create"
category = "todos"
description = "user submits new todo"

[request]
url = "/todos"
method = "POST"

[request.headers]
Authorization = "BEARER auth_token"
Content-Type = "application/json"

[request.props.prop_name]
type = "text"

[responses.success]
status_code = 201
[responses.success.props.prop_name]
type = "int"

[responses.failure]
status_code = 400
[responses.failure.props]
type = "text"
```

---

### **Request:**

The `request` object contains two fields:

- `headers`: set headers on request object
- `props`: same as `Object`

```toml
# ...metadata

[request]
[request.headers]
Authorization = "BEARER auth_token"
Content-Type = "application/json"

[request.props.prop_name]
type = "text"

# ...responses
```

### **Responses:**

The `responses` object contains two fields:

- `status_code`: HTTP status code associated with response type
- `props`

```toml
# ...metadata

# ...request

[responses]

[responses.success]
status_code = 201
[responses.success.props.prop_name]
type = "text"

[responses.failure]
status_code = 400
[responses.failure.props]
type = "text"
```

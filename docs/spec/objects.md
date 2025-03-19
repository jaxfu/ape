# `Objects`

Structured data consisting of metadata and `Props`.

### **Structure:**

- ...metadata
- `props` [optional | default=NULL]

### **Example:**

TOML:

```toml
name = "Todo"
category = "Main"
description = "todo object created by users"

[props]
username = "@Users.Types.Username"

[props.message]
type = "text"

[props.due_date]
type = "int"

[props.notifications]
type = "text"
array = true
max_length = 128
```

JSON:

```json
{
	"name": "Todo",
	"category": "Main",
	"description": "todo object created by users",
	"props": {
		"username": "@Users.Types.Username",
		"due_date": {
			"type": "int",
			"signed": false
		},
		"message": {
			"type": "text",
			"max_length": 256
		},
		"notifications": {
			"type": "text",
			"array": true,
			"max_length": 128
		}
	}
}
```

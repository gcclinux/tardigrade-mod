# Flexible Fields Feature

## Overview

The flexible fields feature allows you to store records with any number of custom fields, not just the fixed `key` and `data` structure.

## New Structure

```go
type FlexStruct struct {
    Id     int               `json:"id"`
    Key    string            `json:"key"`
    Fields map[string]string `json:"fields"`
}
```

## Usage Examples

### 1. Add Record with Map

```go
tar := tardigrade.Tardigrade{}
tar.AddFlexField("user:2", map[string]string{
    "name":     "ricardo wagemaker",
    "status":   "married",
    "location": "london",
}, "mydb.db")
```

### 2. Add Record with Variadic Arguments (Easier)

```go
tar.AddFlexFieldVariadic("app:3", "mydb.db",
    "cost", "299",
    "billing", "monthly",
    "patch", "17",
    "color", "blue",
    "os", "linux",
    "mode", "auto",
)
```

### 3. Retrieve Full Record

```go
// Get as JSON
result := tar.SelectFlexByID(2, "json", "mydb.db")
fmt.Println(result)

// Output:
// {
//   "id": 2,
//   "key": "user:2",
//   "fields": {
//     "name": "ricardo wagemaker",
//     "status": "married",
//     "location": "london"
//   }
// }
```

### 4. Get Specific Field Value

```go
name := tar.GetFlexField(2, "name", "mydb.db")
fmt.Println(name) // Output: ricardo wagemaker
```

### 5. List All Fields in a Record

```go
fields := tar.ListFlexFields(2, "mydb.db")
fmt.Println(fields) // Output: [name status location]
```

### 6. Search Records

```go
format, results := tar.SelectFlexSearch("london married", "json", "mydb.db")
// Returns all records containing both "london" AND "married"
```

### 7. Modify Record

```go
tar.ModifyFlexField(2, "user:2", map[string]string{
    "name":     "ricardo wagemaker",
    "status":   "married",
    "location": "paris",  // Changed
    "age":      "35",     // New field
}, "mydb.db")
```

## New Functions

```go
// Add record with map
func (*Tardigrade).AddFlexField(key string, fields map[string]string, db string) bool

// Add record with variadic args (key1, value1, key2, value2, ...)
func (*Tardigrade).AddFlexFieldVariadic(key string, db string, keyValuePairs ...string) bool

// Retrieve record by ID
func (*Tardigrade).SelectFlexByID(id int, format string, db string) string

// Search records
func (*Tardigrade).SelectFlexSearch(search, format string, db string) (string, []byte)

// Get specific field value
func (*Tardigrade).GetFlexField(id int, fieldName string, db string) string

// Modify record
func (*Tardigrade).ModifyFlexField(id int, key string, fields map[string]string, db string) (string, bool)

// List all field names
func (*Tardigrade).ListFlexFields(id int, db string) []string
```

## Formats Supported

- `raw` - Single-line JSON
- `json` - Pretty-printed JSON
- `id` - Just the ID
- `key` - Just the key
- `fields` - Just the fields map

## Storage Format

Records are stored as newline-delimited JSON:

```json
{"id":1,"key":"user:1","fields":{"name":"ricardo wagemaker","status":"married","location":"london"}}
{"id":2,"key":"app:2","fields":{"cost":"299","billing":"monthly","patch":"17","color":"blue","os":"linux","mode":"auto"}}
```

## Compatibility

- Flexible records can coexist with standard records in different databases
- Use standard functions (AddField) for simple key-value pairs
- Use flexible functions (AddFlexField) for complex multi-field records
- Both use the same database management functions (CreateDB, DeleteDB, etc.)


# Tardigrade-Mod

> A lightweight, file-based NoSQL database library for Go applications

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-See%20LICENSE-blue.svg)](LICENSE)

**Version:** 0.3.0  
**Updated:** Sun Jan 18 09:38:18 PM GMT 2026

## Overview

Tardigrade-Mod is a simple, zero-dependency NoSQL database solution designed for small to medium-sized Go applications. It provides persistent key-value storage using JSON serialization with a file-based approach, making it perfect for applications that need data persistence without the complexity of a full database server.

### Key Features

- 🚀 **Zero Dependencies** - Pure Go with only standard library
- 📁 **File-Based** - Single-file database for easy deployment
- 🔍 **Search Capable** - Multi-keyword search functionality
- 🔐 **Encryption Ready** - Built-in AES encryption utilities
- 🖥️ **Cross-Platform** - Works on Linux, macOS, and Windows
- 📊 **Multiple Formats** - Output data in raw, JSON, or field-specific formats
- 🗄️ **Multi-Database** - Support for multiple database files per application
- 🔧 **Flexible Schema** - Store records with any number of custom fields

## Quick Start

### Installation

```bash
go get github.com/gcclinux/tardigrade-mod
```

### Basic Example (Simple Key-Value)

```go
package main

import (
    "fmt"
    "github.com/gcclinux/tardigrade-mod"
)

func main() {
    tar := tardigrade.Tardigrade{}
    
    // Create database
    tar.CreateDB("myapp.db")
    
    // Add data
    tar.AddField("user:1", "John Doe", "myapp.db")
    tar.AddField("user:2", "Jane Smith", "myapp.db")
    
    // Retrieve data
    result := tar.SelectByID(1, "json", "myapp.db")
    fmt.Println(result)
    
    // Search
    format, results := tar.SelectSearch("John", "json", "myapp.db")
    fmt.Println(format, string(results))
}
```

### Flexible Fields Example (Multiple Fields)

```go
package main

import (
    "fmt"
    "github.com/gcclinux/tardigrade-mod"
)

func main() {
    tar := tardigrade.Tardigrade{}
    tar.CreateDB("flexible.db")
    
    // Add record with multiple fields
    tar.AddFlexFieldVariadic("user:1", "flexible.db",
        "name", "ricardo wagemaker",
        "status", "married",
        "location", "london")
    
    // Add record with many fields
    tar.AddFlexFieldVariadic("app:2", "flexible.db",
        "cost", "299",
        "billing", "monthly",
        "patch", "17",
        "color", "blue",
        "os", "linux",
        "mode", "auto")
    
    // Retrieve full record
    result := tar.SelectFlexByID(1, "json", "flexible.db")
    fmt.Println(result)
    
    // Get specific field
    name := tar.GetFlexField(1, "name", "flexible.db")
    fmt.Println("Name:", name)
}
```

## Documentation

- [DESIGN.md](DESIGN.md) - Architecture, design patterns, and technical specifications
- [FLEXIBLE.md](FLEXIBLE.md) - Flexible fields feature guide and examples

## API Reference

### Structure and Available Functions

#### Core Structure
```go
type Tardigrade struct{}
```

#### Standard Functions (Fixed Schema: id, key, data)
```go
func (*Tardigrade).AddField(key string, data string, db string) bool
func (*Tardigrade).CountSize(db string) int
func (*Tardigrade).CreateDB(db string) (msg string, status bool)
func (*Tardigrade).CreatedDBCopy(db string) (msg string, status bool)
func (*Tardigrade).DeleteDB(db string) (msg string, status bool)
func (*Tardigrade).EmptyDB(db string) (msg string, status bool)
func (*Tardigrade).FirstField(f string, db string) string
func (*Tardigrade).FirstXFields(count int, format string, db string) (string, []byte)
func (*Tardigrade).GetUpdated() (updated string)
func (*Tardigrade).GetVersion() (release string)
func (*Tardigrade).LastField(f string, db string) string
func (*Tardigrade).LastXFields(count int, format string, db string) (string, []byte)
func (*Tardigrade).ModifyField(id int, k string, v string, db string) (msg string, status bool)
func (*Tardigrade).RemoveField(id int, db string) (string, bool)
func (*Tardigrade).SelectByID(id int, f string, db string) string
func (*Tardigrade).UniqueID(db string) int
func (*Tardigrade).SelectSearch(search, format string, db string) (string, []byte)
```

#### Flexible Field Functions (Variable Schema)
```go
func (*Tardigrade).AddFlexField(key string, fields map[string]string, db string) bool
func (*Tardigrade).AddFlexFieldVariadic(key string, db string, keyValuePairs ...string) bool
func (*Tardigrade).SelectFlexByID(id int, format string, db string) string
func (*Tardigrade).SelectFlexSearch(search, format string, db string) (string, []byte)
func (*Tardigrade).GetFlexField(id int, fieldName string, db string) string
func (*Tardigrade).ModifyFlexField(id int, key string, fields map[string]string, db string) (string, bool)
func (*Tardigrade).ListFlexFields(id int, db string) []string
```

#### Utility Functions
```go
func (*Tardigrade).MyMarshal(t interface{}) ([]byte, error)
func (*Tardigrade).MyIndent(v interface{}, prefix, indent string) ([]byte, error) 
func (*Tardigrade).MyEncode(b []byte) string
func (*Tardigrade).MyDecode(s string) []byte
func (*Tardigrade).MyEncrypt(text, Password string) (string, error)
func (*Tardigrade).MyDecrypt(text, Password string) (string, error)
```

## Detailed Usage Guide

### Flexible Fields (NEW)

For records with multiple custom fields, use the flexible field functions. See [FLEXIBLE.md](docs/FLEXIBLE.md) for complete documentation.

#### AddFlexFieldVariadic

Adds a record with any number of custom fields.

**Signature:** `AddFlexFieldVariadic(key string, db string, keyValuePairs ...string) bool`

```go
Example 1: Multiple fields
	tar := tardigrade.Tardigrade{}
	tar.AddFlexFieldVariadic("user:1", "mydb.db",
		"name", "ricardo wagemaker",
		"status", "married",
		"location", "london")

Example 2: Many fields
	tar.AddFlexFieldVariadic("app:2", "mydb.db",
		"cost", "299",
		"billing", "monthly",
		"patch", "17",
		"color", "blue",
		"os", "linux",
		"mode", "auto")

Result:
	true | false
```

#### SelectFlexByID

Retrieves a flexible record by ID.

**Signature:** `SelectFlexByID(id int, format string, db string) string`

**Formats:** `raw` | `json` | `id` | `key` | `fields`

```go
Example:
	tar := tardigrade.Tardigrade{}
	result := tar.SelectFlexByID(1, "json", "mydb.db")
	fmt.Println(result)

Result:
{
  "id": 1,
  "key": "user:1",
  "fields": {
    "name": "ricardo wagemaker",
    "status": "married",
    "location": "london"
  }
}
```

#### GetFlexField

Retrieves a specific field value from a record.

**Signature:** `GetFlexField(id int, fieldName string, db string) string`

```go
Example:
	tar := tardigrade.Tardigrade{}
	name := tar.GetFlexField(1, "name", "mydb.db")
	fmt.Println(name)

Result:
	ricardo wagemaker
```

#### ListFlexFields

Lists all field names in a record.

**Signature:** `ListFlexFields(id int, db string) []string`

```go
Example:
	tar := tardigrade.Tardigrade{}
	fields := tar.ListFlexFields(1, "mydb.db")
	fmt.Println(fields)

Result:
	[name status location]
```

### Database Management

#### CreateDB

Creates a database file if it doesn't exist.

**Signature:** `CreateDB(db string) (msg string, status bool)`
```
Example 1: (ignore return)
	tar := tardigrade.Tardigrade{}
	tar.CreateDB(db_name)

Example 2 (capture return):
	tar := tardigrade.Tardigrade{}
	msg, status := tar.CreateDB(db_name)
	fmt.Println(msg, status)

Return:
	Created: <full_path>/tardigrade.db true
	Exist: <full_path>/tardigrade.db false

```

#### DeleteDB

⚠️ **WARNING:** Permanently deletes the database file.

**Signature:** `DeleteDB(db string) (msg string, status bool)`
```
Example 1: (ignore return)
	tar := tardigrade.Tardigrade{}
	tar.DeleteDB(db_name)

Example 2 (capture return):
	tar := tardigrade.Tardigrade{}
	msg, status := tar.DeleteDB(db_name)
	fmt.Println(msg, status)

Return:
	Removed: <full_path>/tardigrade.db true
	Unavailable: <full_path>/tardigrade.db false

```
#### CreatedDBCopy

Creates a backup copy of the database in the user's home directory.

**Signature:** `CreatedDBCopy(db string) (msg string, status bool)`

```
Example 1: (ignore return)
	tar := tardigrade.Tardigrade{}
	tar.CreatedDBCopy(db_name)

Example 2 (capture return):
	tar := tardigrade.Tardigrade{}
	msg, status := tar.CreatedDBCopy(db_name)
	fmt.Println(msg, status)

Return:
	Copy: <full_path>/tardigradecopy.db true
	Failed: database tardigrade.db missing! false
	Failed: buffer error failed to create database! false
	Failed: permission error failed to create database! false

```

#### EmptyDB

⚠️ **WARNING:** Destroys all data in the database while preserving the file.

**Signature:** `EmptyDB(db string) (msg string, status bool)` 

```
Example 1: (ignore return)
	tar := tardigrade.Tardigrade{}
	tar.EmptyDB(db_name)

Example 2 (capture return):
	tar := tardigrade.Tardigrade{}
	msg, status := tar.EmptyDB(db_name)
	fmt.Println(msg, status)

Return:
	Empty: database now clean! true
	Failed: no permission to re-create! false
	Missing: could not find database false! false

```

### CRUD Operations

#### AddField

Adds a new record to the database with auto-incrementing ID.

**Signature:** `AddField(key string, data string, db string) bool`

```
Example 1: (ignore return)
	tar := tardigrade.Tardigrade{}
	tar.AddField("New string Entry", "string of data representing a the value", "db_name")

Example 2 (capture return):
	tar := tardigrade.Tardigrade{}
	status := tar.AddField("New string Entry", "string of data representing a the value", "db_name")
	fmt.Println(status)

Return:
	true | false

```

#### CountSize

Returns the total number of records in the database.

**Signature:** `CountSize(db string) int`

````
Example (capture return):
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.CountSize("db_name"))

Result:
	44
````

#### FirstField

Retrieves the first record in the database.

**Signature:** `FirstField(format string, db string) string`

**Formats:** `raw` | `json` | `id` | `key` | `value`

```
Example 1: (true | failed)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.FirstField("raw", "db_name"))

Result: 
	{"id":1,"key":"one","data":"string data test"}
	Failed: database tardigrade.db is empty!
	Failed: database tardigrade.db missing!

Example 2: (true)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.FirstField("json","db_name"))

Result:
{
        "id": 1,
        "key": "New string Entry",
        "data": "string of data representing a the value"
}
```

#### LastField

Retrieves the last record in the database.

**Signature:** `LastField(format string, db string) string`

**Formats:** `raw` | `json` | `id` | `key` | `value`

```
Example 1: (true | failed)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.FirstField("raw", "db_name"))

Result: 
	{"id":44,"key":"New Entry","data":"string of data representing a the value"}
	Failed: database tardigrade.db is empty!
	Failed: database tardigrade.db missing!

Example 2: (true)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.LastField("value", "db_name"))

Result:
	string of data representing a the value

Example 3: (true)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.LastField("key", "db_name"))

Result:
	New Entry

Example: 4 (true)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.LastField("json", "db_name"))

Result:
{
        "id": 44,
        "key": "New Entry",
        "data": "string of data representing a the value"
}
```

#### SelectByID

Retrieves a specific record by its ID.

**Signature:** `SelectByID(id int, format string, db string) string`

**Formats:** `raw` | `json` | `id` | `key` | `value`

```
Example 1: (true)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.SelectByID(10, "raw", "db_name"))

Result:
	{"id":10,"key":"Roman","data":"string of data representing a the value of X"}

Example 2: (false)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.SelectByID(100, "raw", "db_name"))

Result:
	Record 100 is empty!

Example 3: (true)
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.SelectByID(25, "json", "db_name"))

Result:
{
        "id": 25,
        "key": "New string Entry 23",
        "data": "string of data representing a the value"
}
```

#### UniqueID

Returns the last used ID (useful for auto-increment logic).

**Signature:** `UniqueID(db string) int`

```
Example: (always true)
	tar := Tardigrade{}
	fmt.Println(tar.UniqueID("db_name"))

Result:
	52
```


#### FirstXFields

Retrieves the first X records from the database.

**Signature:** `FirstXFields(count int, format string, db string) (string, []byte)`

```
Example:
	tar := tardigrade.Tardigrade{}
	var received = tar.FirstXFields(2, "db_name")

	type MyStruct struct {
		Id   int
		Key  string
		Data string
	}

	bytes := received
	var data []MyStruct
	size := len(data)
	json.Unmarshal(bytes, &data)

	if size == 1 {
		fmt.Printf("id: %v, key: %v, data: %s", data[0].Id, data[0].Key, data[0].Data)
	} else {
		for x := range data {
			fmt.Printf("id: %v, key: %v, data: %s", data[x].Id, data[x].Key, data[x].Data)
			fmt.Println()
		}
	}

Result:
	id: 1, key: New string Entry, data: string of data representing a the value
	id: 2, key: New string Entry 0, data: string of data representing a the value
```

#### LastXFields

Retrieves the last X records from the database.

**Signature:** `LastXFields(count int, format string, db string) (string, []byte)`

```
Example 1: (always true)
	tar := tardigrade.Tardigrade{}
	var received = tar.LastXFields(2, "db_name")

	type MyStruct struct {
		Id   int
		Key  string
		Data string
	}

	bytes := received
	var data []MyStruct
	size := len(data)
	json.Unmarshal(bytes, &data)

	if size == 1 {
		fmt.Printf("id: %v, key: %v, data: %s", data[0].Id, data[0].Key, data[0].Data)
	} else {
		for x := range data {
			fmt.Printf("id: %v, key: %v, data: %s", data[x].Id, data[x].Key, data[x].Data)
			fmt.Println()
		}
	}

Result:
	id: 51, key: New string Entry 49, data: string of data representing a the value
	id: 52, key: New string Entry 50, data: string of data representing a the value
```

#### RemoveField

Deletes a specific record by ID.

**Signature:** `RemoveField(id int, db string) (string, bool)`

```
Example 1: (true | false)
	tar := tardigrade.Tardigrade{}
	msg, status := tar.RemoveField(2, "db_name")
	fmt.Println(msg, status)

Result:
	{"id":2,"key":"New string Entry 0","data":"string of data representing a the value"} true
	Record 2 is empty! false
	Database tardigrade.db is empty! false
```

#### ModifyField

Updates an existing record with new key and data values.

**Signature:** `ModifyField(id int, key string, value string, db string) (msg string, status bool)`

```
Example 1: (true)
	tar := tardigrade.Tardigrade{}
	change, status := tar.ModifyField(2, "Updated key 2", "with new Updated data set with and new inforation", "db_name")
	fmt.Println(change, status)

Result:
	{"id":2,"key":"Updated key 2","data":"with new Updated data set with and new inforation"} true

Example 2: (false)
	tar := tardigrade.Tardigrade{}
	change, status := tar.ModifyField(100, "Updated key 2", "with new Updated data set with and new inforation", "db_name")
	fmt.Println(change, status)

Result:
	Record 100 is empty! false

```
#### SelectSearch

Searches for records matching all provided keywords (AND logic).

**Signature:** `SelectSearch(search string, format string, db string) (string, []byte)`

**Search:** Comma or space-separated keywords (case-insensitive)
```
Example:
	package main

	import (
		"encoding/json"
		"fmt"
		"strconv"
		"strings"

		"github.com/gcclinux/tardigrade-mod"
	)

	func main() {
		tar := tardigrade.Tardigrade{}
		status := tar.AddField("New string Entry word1", "string of data representing a the word2", "db_name")
		fmt.Println(status)

		type MyStruct struct {
			Id   int
			Key  string
			Data string
		}

		var format, received = tar.SelectSearch("word1,word", "json", "db_name")
		bytes := received
		var data []MyStruct
		json.Unmarshal(bytes, &data)

		if (strings.Contains(string(received), "Database") && strings.Contains(string(received), "missing")) || (strings.Contains(string(received), "Database") && strings.Contains(string(received), "empty")) {
			fmt.Println(string(received))
			fmt.Println()
		}

		for x := range data {
			if format == "json" {
				out, _ := json.MarshalIndent(&data[x], "", "  ")
				fmt.Printf("%v", string(out))
				fmt.Println()
			} else if format == "value" {
				fmt.Println(string(data[x].Data))
				fmt.Println()
			} else if format == "raw" {
				fmt.Printf("id: %d, key: %v, data: %s\n", data[x].Id, data[x].Key, data[x].Data)
			} else if format == "key" {
				fmt.Printf("%v\n", data[x].Key)
			} else if format == "id" {
				fmt.Println(strconv.Itoa(data[x].Id))
				fmt.Println()
			} else {
				fmt.Printf("Invalid format provided!")
			}
		}
	}
Result:
	{
	"Id": 1,
	"Key": "New string Entry word1",
	"Data": "string of data representing a the word2"
	}

```

### Utility Functions

#### Serialization & Encoding

- `MyMarshal(t interface{}) ([]byte, error)` - JSON marshal with HTML escape disabled
- `MyIndent(v interface{}, prefix, indent string) ([]byte, error)` - Pretty-print JSON
- `MyEncode(b []byte) string` - Base64 encoding
- `MyDecode(s string) []byte` - Base64 decoding

#### Encryption (AES-CFB)

- `MyEncrypt(text, password string) (string, error)` - Encrypt text with password
- `MyDecrypt(text, password string) (string, error)` - Decrypt text with password

**Note:** Password must be 16, 24, or 32 bytes for AES compatibility.

#### Version Information
```
Example:
	package main

	import (
		"fmt"

		"github.com/gcclinux/tardigrade-mod"
	)

	func main() {
		tar := tardigrade.Tardigrade{}

		fmt.Println(tar.GetUpdated())
		fmt.Println(tar.GetVersion())
	}

Result:
	Sat 4 Mar 18:56:11 GMT 2023
	0.2.0
```


## Performance Considerations

- **Optimal for:** < 100,000 records, < 100 MB file size
- **Search:** O(n) linear scan (no indexing)
- **Updates/Deletes:** Full file rewrite operation
- **Concurrency:** Single writer recommended

For detailed performance characteristics, see [DESIGN.md](DESIGN.md).

## Use Cases

✅ **Ideal For:**
- Configuration storage
- Session management
- Application caching
- Logging and audit trails
- Prototyping and development
- CLI tools and utilities
- Embedded systems

❌ **Not Recommended For:**
- High-concurrency applications
- Large datasets (>1M records)
- Real-time analytics
- Applications requiring ACID transactions

## Release Notes

```
** release 0.0.1 - Initial version
** release 0.0.2 - Updated README.md and corrected some issues.
** release 0.0.3 - Modified to use structure method
** release 0.0.4 - Converted tardigrade app to tardigrade-mod
** release 0.1.0 - Several functions added from tardigrade app here
** release 0.1.2 - Bug fix returning string in lower case (fixed)
** release 0.1.3 - Bug fix function set to lower case was unaccessible
** release 0.1.4 - Bug fix storing string with encoder.SetEscapeHTML(false)
** release 0.2.0 - Added 2 new functions to Tardigrade main struct
** release 0.2.1 - Minor bug fix inntroduced in previous version
** release 0.2.3 - Working Progress enabling data encryption
** release 0.2.5 - Modified functions to include database name so an app can have more than 1 db
** release 0.3.0 - Added flexible field support for variable schema records
```

## Roadmap
### Planned Features

```go
// Encrypted field operations (coming soon)
func (*Tardigrade).AddCryptField(key string, data string, db string) bool
func (*Tardigrade).SelectByIDdecrypt(id int, f string, db string) string
```

## Contributing

Contributions are welcome! Please ensure:
- Code follows Go conventions and best practices
- All functions include proper error handling
- Documentation is updated for new features
- Examples are provided

## License

See [LICENSE](LICENSE) file for details.

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/gcclinux/tardigrade-mod).

---

**Made with ❤️ for the Go community**

## This is the mod version of the is small and simple no-SQL database app for small GO apps.
*updated:  Sun 10 Sep 18:54:19 BST 2023*<br>
*release:  0.2.5*

<br>

## Getting Started
>go get [github.com/gcclinux/tardigrade-mod](http://github.com/gcclinux/tardigrade-mod)

<BR>

Current structure and available functions()

```
type Tardigrade struct{}

func (*Tardigrade).AddField(key string, data string, db string) bool
func (*Tardigrade).CountSize(db string) int
func (*Tardigrade).CreateDB(db string) (msg string, status bool)
func (*Tardigrade).CreatedDBCopy(db string) (msg string, status bool)
func (*Tardigrade).DeleteDB(db string) (msg string, status bool)
func (*Tardigrade).EmptyDB(db string) (msg string, status bool)
func (*Tardigrade).FirstField(f string, db string) string
func (*Tardigrade).FirstXFields(count int, db string) []byte
func (*Tardigrade).GetUpdated() (updated string)
func (*Tardigrade).GetVersion() (release string)
func (*Tardigrade).LastField(f string, db string) string
func (*Tardigrade).LastXFields(count int, db string) []byte
func (*Tardigrade).ModifyField(id int, k string, v string, db string) (msg string, status bool)
func (*Tardigrade).RemoveField(id int, db string) (string, bool)
func (*Tardigrade).SelectByID(id int, f string, db string) string
func (*Tardigrade).UniqueID(db string) int
func (*Tardigrade).SelectSearch(search, format string, db string) (string, []byte)
func (*Tardigrade).MyMarshal(t interface{}) ([]byte, error)
func (*Tardigrade).MyIndent(v interface{}, prefix, indent string) ([]byte, error) 
func (*Tardigrade).MyEncode(b []byte) string
func (*Tardigrade).MyDecode(s string) []byte
func (*Tardigrade).MyEncrypt(text, Password string) (string, error)
func (*Tardigrade).MyDecrypt(text, Password string) (string, error)
```

# HOW-TO-USE

<BR>

**CreateDB - This function will create a database file if it does not exist and return true | false**
>function: CreateDB(db string)
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

**DeleteDB - WARNING - this function delete the database file return true | false**
>function: DeleteDB(db string)
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
**CreatedDBCopy creates a copy of the Database and store in UserHomeDir()**
>function: CreatedDBCopy(db string)

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

**EmptyDB function - WARNING - this will destroy the database and all data stored in it!**

>function: EmptyDB(db string) 

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

**AddField() function take in ((key)string, (Value) string, db string) and add to database.**

>function: AddField()

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

**CountSize() function will return number of rows in the gojsondb.db**

>function: CountSize(db string)

````
Example (capture return):
	tar := tardigrade.Tardigrade{}
	fmt.Println(tar.CountSize("db_name"))

Result:
	44
````

**FirstField func returns the first entry of gojsondb.db in all formats \[ raw | json | id | key | value ] specify format required**

>function: FirstField(db string)

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

**LastField() func returns the last entry in multi-format \[ raw | json | id | key | value ]**

>function: LastField(db string)

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

**SelectByID func returns an entry string for a specific id in all formats \[ raw | json | id | key | value ]**
>function: SelectByID(db string)

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

**UniqueID function returns an int for the last used UniqueID**
>function: UniqueID(db string)

```
Example: (always true)
	tar := Tardigrade{}
	fmt.Println(tar.UniqueID("db_name"))

Result:
	52
```


**FirstXFields returns last X number of entries from db in byte[] format**
>function: FirstXFields(db string)

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

**LastXFields returns last X number of entries from db in values byte[] format**
>function: LastXFields(db string)

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

**RemoveField function takes an unique field id as an input and remove the matching field entry**
>function: RemoveField(db string)

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

**ModifyField function takes ID, Key, Value and update row = ID with new information provided**
> ModifyField(2, "Updated key", "Updated data set with new inforation", "db_name")

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
**SelectSearch function takes in a single or multiple words(comma, or space separated) and format type, Returns true values in all formats**
>SelectSearch("patern1,pattern2","json", "db_name")
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

Additional couple of informational functions
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


RELEASE NOTE:

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
```

OUTSTANDING:
```
** Write and share additional functions
func (*Tardigrade).AddCryptField(key string, data string, db string) bool
func (*Tardigrade).SelectByIDdecrypt(id int, f string, db string) string
```
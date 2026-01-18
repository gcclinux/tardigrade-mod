package tardigrade

// Updated - Sun Jan 18 09:38:18 PM GMT 2026
// Version - 0.3.0

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FlexStruct supports variable number of fields
type FlexStruct struct {
	Id     int               `json:"id"`
	Key    string            `json:"key"`
	Fields map[string]string `json:"fields"`
}

// AddFlexField adds a record with variable fields
// Usage: tar.AddFlexField("user:2", map[string]string{"name": "ricardo", "status": "married", "city": "london"}, "mydb.db")
func (tar *Tardigrade) AddFlexField(key string, fields map[string]string, db string) bool {
	if !tar.fileExists(db) {
		tar.CreateDB(db)
		if !tar.fileExists(db) {
			return false
		}
	}

	id := tar.UniqueID(db) + 1
	record := FlexStruct{
		Id:     id,
		Key:    key,
		Fields: fields,
	}

	response, err := tar.MyMarshal(record)
	CheckError("AddFlexField", err)

	file, err := os.OpenFile(db, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	CheckError("AddFlexField", err)
	defer file.Close()
	file.Write(response)

	return true
}

// AddFlexFieldVariadic adds a record with variadic string arguments
// Usage: tar.AddFlexFieldVariadic("user:2", "mydb.db", "name", "ricardo", "status", "married", "city", "london")
func (tar *Tardigrade) AddFlexFieldVariadic(key string, db string, keyValuePairs ...string) bool {
	if len(keyValuePairs)%2 != 0 {
		return false // Must have even number of arguments (key-value pairs)
	}

	fields := make(map[string]string)
	for i := 0; i < len(keyValuePairs); i += 2 {
		fields[keyValuePairs[i]] = keyValuePairs[i+1]
	}

	return tar.AddFlexField(key, fields, db)
}

// SelectFlexByID retrieves a flexible record by ID
func (tar *Tardigrade) SelectFlexByID(id int, format string, db string) string {
	regx := fmt.Sprintf("\"id\":%v,", id)
	src := db

	if !tar.fileExists(src) {
		return fmt.Sprintf("Database %s missing!", src)
	}

	fInfo, _ := os.Stat(src)
	if fInfo.Size() <= 1 {
		return fmt.Sprintf("Database %s is empty!", src)
	}

	file, err := os.Open(src)
	CheckError("SelectFlexByID", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, regx) {
			var s FlexStruct
			err = json.Unmarshal([]byte(line), &s)
			CheckError("SelectFlexByID", err)

			switch format {
			case "json":
				out, _ := tar.MyIndent(&s, "", "  ")
				return string(out)
			case "raw":
				return line
			case "key":
				return s.Key
			case "id":
				return strconv.Itoa(s.Id)
			case "fields":
				out, _ := tar.MyMarshal(s.Fields)
				return string(out)
			default:
				return "Invalid format! Use: raw, json, id, key, fields"
			}
		}
	}

	return fmt.Sprintf("Record %v is empty!", id)
}

// SelectFlexSearch searches flexible records
func (tar *Tardigrade) SelectFlexSearch(search, format string, db string) (string, []byte) {
	search = strings.ToLower(search)
	search = strings.ReplaceAll(search, " ", ",")
	keywords := strings.Split(search, ",")

	if !tar.fileExists(db) {
		return format, []byte(fmt.Sprintf("Database %s missing!", db))
	}

	fInfo, _ := os.Stat(db)
	if fInfo.Size() <= 1 {
		return format, []byte(fmt.Sprintf("Database %s is empty!", db))
	}

	var results []FlexStruct
	file, err := os.Open(db)
	CheckError("SelectFlexSearch", err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		matchAll := true

		for _, keyword := range keywords {
			if !strings.Contains(line, keyword) {
				matchAll = false
				break
			}
		}

		if matchAll {
			var record FlexStruct
			err = json.Unmarshal([]byte(scanner.Text()), &record)
			CheckError("SelectFlexSearch", err)
			results = append(results, record)
		}
	}

	output, err := tar.MyMarshal(results)
	CheckError("SelectFlexSearch", err)
	return format, output
}

// GetFlexField retrieves a specific field value from a record
func (tar *Tardigrade) GetFlexField(id int, fieldName string, db string) string {
	result := tar.SelectFlexByID(id, "raw", db)
	if strings.Contains(result, "empty") || strings.Contains(result, "missing") {
		return result
	}

	var record FlexStruct
	err := json.Unmarshal([]byte(result), &record)
	CheckError("GetFlexField", err)

	if value, exists := record.Fields[fieldName]; exists {
		return value
	}
	return fmt.Sprintf("Field '%s' not found in record %d", fieldName, id)
}

// ModifyFlexField updates a flexible record
func (tar *Tardigrade) ModifyFlexField(id int, key string, fields map[string]string, db string) (string, bool) {
	before := tar.SelectFlexByID(id, "raw", db)
	if strings.Contains(before, "Record") && strings.Contains(before, "empty") {
		return before, false
	}

	record := FlexStruct{
		Id:     id,
		Key:    key,
		Fields: fields,
	}

	after, _ := tar.MyMarshal(&record)
	afterStr := strings.TrimSpace(string(after))

	input, err := os.ReadFile(db)
	CheckError("ModifyFlexField", err)

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, before) {
			lines[i] = afterStr
		}
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(db, []byte(output), 0644)
	CheckError("ModifyFlexField", err)

	return tar.SelectFlexByID(id, "raw", db), true
}

// ListFlexFields returns all field names from a record
func (tar *Tardigrade) ListFlexFields(id int, db string) []string {
	result := tar.SelectFlexByID(id, "raw", db)
	if strings.Contains(result, "empty") || strings.Contains(result, "missing") {
		return []string{}
	}

	var record FlexStruct
	err := json.Unmarshal([]byte(result), &record)
	CheckError("ListFlexFields", err)

	fields := make([]string, 0, len(record.Fields))
	for key := range record.Fields {
		fields = append(fields, key)
	}
	return fields
}

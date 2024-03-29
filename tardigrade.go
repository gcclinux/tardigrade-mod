package tardigrade

// Updated Sat  4 Mar 18:56:11 GMT 2023

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// MyStruct contains the structure of the data stored into the tardigrade.db!
type MyStruct struct {
	Id   int    `json:"id"`
	Key  string `json:"key"`
	Data string `json:"data"`
}

func (tar *Tardigrade) GetOS() rune {
	PATH_SEPARATOR := '/'
	if runtime.GOOS == "windows" {
		PATH_SEPARATOR = '\\'
	} else if runtime.GOOS == "linux" {
		PATH_SEPARATOR = '/'
	} else if runtime.GOOS == "darwin" {
		PATH_SEPARATOR = '/'
	} else {
		log.Println("unknown")
	}
	return PATH_SEPARATOR
}

// AddField take in (key, sprint) (data, string) and add to tardigrade.db
func (tar *Tardigrade) AddField(key, data string, db string) bool {

	if !tar.fileExists(db) {
		tar.CreateDB(db)
		if !tar.fileExists(db) {
			return false
		}
	}

	id := tar.UniqueID(db) + 1
	var getStruct = MyStruct{}
	getStruct.Id = id
	getStruct.Key = key
	getStruct.Data = data

	response, err := tar.MyMarshal(getStruct)
	CheckError("Marshal", err)

	file, err := os.OpenFile(db, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	CheckError("O_APPEND", err)
	//file.WriteString("\n")
	file.Write(response)

	return true
}

// RemoveField function takes an unique field id as an input and remove the matching field entry
func (tar *Tardigrade) RemoveField(id int, db string) (string, bool) {

	status := true
	msg := ""

	src := db
	if !tar.fileExists(src) {
		return (fmt.Sprintf("Database %s missing!", src)), false
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return (fmt.Sprintf("Database %s is empty!", src)), false
		} else {
			line := tar.SelectByID(id, "raw", db)

			if strings.Contains(line, "Record") && strings.Contains(line, "empty") {
				status = false
				msg = line
			} else {
				msg = line
				fpath := src
				f, err := os.Open(fpath)
				CheckError("RemoveField(1)", err)

				var bs []byte
				buf := bytes.NewBuffer(bs)

				scanner := bufio.NewScanner(f)
				for scanner.Scan() {
					if scanner.Text() != line {
						_, err := buf.Write(scanner.Bytes())
						CheckError("RemoveField(2)", err)
						_, err = buf.WriteString("\n")
						CheckError("RemoveField(3)", err)
					}
				}
				if err := scanner.Err(); err != nil {
					CheckError("RemoveField(4)", err)
				}

				err = os.WriteFile(fpath, buf.Bytes(), 0666)
				CheckError("RemoveField(5)", err)
				f.Close()
			}
		}
	}
	return msg, status
}

// SelectByID function returns an entry string for a specific id in all formats [ raw | json | id | key | value ]
func (tar *Tardigrade) SelectByID(id int, f string, db string) string {

	regx := fmt.Sprintf("\"id\":%v,", id)

	result := ""
	src := db
	if !tar.fileExists(src) {
		return (fmt.Sprintf("Database %s missing!", src))
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return (fmt.Sprintf("Database %s is empty!", src))
		} else {
			line := ""
			file, err := os.Open(src)
			CheckError("SelectByID(1)", err)
			defer file.Close()
			var r io.Reader = file
			sc := bufio.NewScanner(r)
			for sc.Scan() {
				if strings.Contains(sc.Text(), regx) {
					line = sc.Text()
				}
			}
			if len(line) == 0 {
				return (fmt.Sprintf("Record %v is empty!", id))
			} else {
				var s MyStruct
				in := []byte(line)
				err = json.Unmarshal(in, &s)
				CheckError("SelectByID(2)", err)

				if f == "json" {
					out, _ := tar.MyIndent(&s, "", "  ")
					result = string(out)
				} else if f == "value" {
					result = string(s.Data)
				} else if f == "raw" {
					result = line
				} else if f == "key" {
					result = string(s.Key)
				} else if f == "id" {
					result = strconv.Itoa(s.Id)
				} else {
					result = "Invalid format provided!"
				}
			}
		}
	}
	return result
}

// ModifyField function takes ID, Key, Value and update row = ID with new information provided
func (tar *Tardigrade) ModifyField(id int, k, v string, db string) (msg string, status bool) {

	status = true
	src := db
	if !tar.fileExists(src) {
		return (fmt.Sprintf("Database %s missing!", src)), status
	} else {

		before := tar.SelectByID(id, "raw", db)
		if strings.Contains(before, "Record") && strings.Contains(before, "empty!") {
			status = false
			return before, status
		}
		var s MyStruct
		s.Id = id
		s.Key = k
		s.Data = v
		out, _ := tar.MyMarshal(&s)
		after := string(out)

		input, err := os.ReadFile(src)
		CheckError("ModifyField(1)", err)
		lines := strings.Split(string(input), "\n")

		for i, line := range lines {
			if strings.Contains(line, before) {
				lines[i] = after
			}
		}
		output := strings.Join(lines, "\n")
		err = os.WriteFile(src, []byte(output), 0644)
		CheckError("ModifyField(2)", err)

		msg = tar.SelectByID(id, "raw", db)
		return msg, status
	}
}

// CountSize will return number of rows in the tardigrade.db
func (tar *Tardigrade) CountSize(db string) int {

	src := db
	f, err := os.Open(src)
	CheckError("CountSize(1)", err)

	defer f.Close()
	var r io.Reader = f
	var count int
	const lineBreak = '\n'
	buf := make([]byte, bufio.MaxScanTokenSize)
	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}
		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}
	fInfo, _ := os.Stat(src)
	fsize := fInfo.Size()
	if fsize > 2 && count == 0 {
		count = 1
	}
	return count
}

// UniqueID function returns an int for the last used UniqueID to AutoIncrement in the AddField()
func (tar *Tardigrade) UniqueID(db string) int {
	lastID := 0
	src := db
	if !tar.fileExists(src) {
		return lastID
	} else {
		lastID, _ = strconv.Atoi(tar.LastField("id", db))
	}

	return lastID
}

// FirstXFields returns first X number of entries from database in byte[] format
// Example: (0.1.2)
// specify number of fields X and format [ raw | json | id | key | value ] to return FirstXFields(2)
func (tar *Tardigrade) FirstXFields(count int, format string, db string) (string, []byte) {

	var allRecord []byte

	src := db
	if !tar.fileExists(src) {
		return format, []byte(fmt.Sprintf("Database %s missing!", src))
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return format, []byte(fmt.Sprintf("Database %s is empty!", src))
		} else {
			var allRecords []MyStruct
			xFields := new(MyStruct)
			var tmpStruct MyStruct
			lastLine := 0
			start := 1
			end := count
			line := ""

			file, err := os.Open(src)
			CheckError("FirstXFields(1)", err)

			defer file.Close()
			var r io.Reader = file
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				lastLine++
				if lastLine >= start && lastLine <= end {
					line = sc.Text()
					in := []byte(line)

					err = json.Unmarshal(in, &tmpStruct)
					CheckError("FirstXFields(2)", err)

					xFields.Id = tmpStruct.Id
					xFields.Key = string(tmpStruct.Key)
					xFields.Data = string(tmpStruct.Data)
					allRecords = append(allRecords, *xFields)
				}
			}
			allRecord, err = tar.MyMarshal(allRecords)
			CheckError("FirstXFields(3)", err)
		}
	}
	return format, allRecord
}

// LastXFields returns last X numbers of entries from db in byte[] format
//
// Example:
// specify number of fields to return LastXFields(2)
func (tar *Tardigrade) LastXFields(count int, format string, db string) (string, []byte) {

	var allRecord []byte

	src := db
	if !tar.fileExists(src) {
		return format, []byte(fmt.Sprintf("Database %s missing!", src))
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return format, []byte(fmt.Sprintf("Database %s is empty!", src))
		} else {
			var allRecords []MyStruct
			var lastLine, start, end = 0, 0, 0
			line := ""

			if count == 1 {
				count = 0
			}

			if tar.CountSize(db) < count {
				count = tar.CountSize(db)
			} else if count >= 2 {
				count = count - 1
			}

			xFields := new(MyStruct)
			var tmpStruct MyStruct

			start = tar.CountSize(db) - count
			end = tar.CountSize(db)

			file, err := os.Open(src)
			CheckError("LastXFields(1)", err)

			defer file.Close()
			var r io.Reader = file
			sc := bufio.NewScanner(r)
			for sc.Scan() {
				lastLine++
				if lastLine >= start && lastLine <= end {
					line = sc.Text()
					in := []byte(line)

					err = json.Unmarshal(in, &tmpStruct)
					CheckError("LastXFields(2)", err)

					xFields.Id = tmpStruct.Id
					xFields.Key = string(tmpStruct.Key)
					xFields.Data = string(tmpStruct.Data)
					allRecords = append(allRecords, *xFields)
				}
			}
			allRecord, err = tar.MyMarshal(allRecords)
			CheckError("LastXFields(3)", err)
		}
	}
	return format, allRecord
}

// FirstField returns the first entry in the database in all formats [ raw | json | id | key | value ],
// must specify format required Example: FirstField("json")
func (tar *Tardigrade) FirstField(f string, db string) string {

	result := ""

	src := db
	if !tar.fileExists(src) {
		return fmt.Sprintf("Database %s missing!", src)
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return fmt.Sprintf("Database %s is empty!", src)
		} else {
			lastLine := 0
			line := ""
			file, err := os.Open(src)

			CheckError("FirstField(1)", err)
			defer file.Close()
			var r io.Reader = file
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				lastLine++
				if lastLine == 1 {
					line = sc.Text()
				}
			}
			var s MyStruct
			in := []byte(line)
			err = json.Unmarshal(in, &s)
			CheckError("FirstField(2)", err)

			if f == "json" {
				out, _ := tar.MyIndent(&s, "", "  ")
				result = string(out)
			} else if f == "value" {
				result = string(s.Data)
			} else if f == "raw" {
				result = line
			} else if f == "key" {
				result = string(s.Key)
			} else if f == "id" {
				result = strconv.Itoa(s.Id)
			} else {
				result = "Invalid format provided!"
			}
		}
	}
	return result
}

// LastField returns the last entry of the database in all formats [ raw | json | id | key | value ] specify format required
func (tar *Tardigrade) LastField(f string, db string) string {

	result := ""

	src := db
	if !tar.fileExists(src) {
		return fmt.Sprintf("Database %s missing!", src)
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return fmt.Sprintf("Database %s is empty!", src)
		} else {
			lastLine := 0
			line := ""
			file, err := os.Open(src)

			CheckError("LastField(1)", err)
			defer file.Close()
			var r io.Reader = file
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				lastLine++
				if lastLine == tar.CountSize(db) {
					line = sc.Text()
				}
			}
			var s MyStruct
			in := []byte(line)
			err = json.Unmarshal(in, &s)
			CheckError("LastField(2)", err)

			if f == "json" {
				out, _ := tar.MyIndent(&s, "", "  ")
				result = string(out)
			} else if f == "value" {
				result = string(s.Data)
			} else if f == "raw" {
				result = line
			} else if f == "key" {
				result = string(s.Key)
			} else if f == "id" {
				result = strconv.Itoa(s.Id)
			} else {
				result = "Invalid format provided!"
			}
		}
	}
	return result
}

// SelectSearch function takes in a single or multiple words(comma,separated) and format type, Returns the format [ raw | json | id | key | value ] and []bytes array with result
// search will need to match ALL words for it to be true and return result.
func (tar *Tardigrade) SelectSearch(search, format string, db string) (string, []byte) {
	search = strings.ToLower(search)
	search = strings.ReplaceAll(search, " ", ",")
	split := strings.Split(search, ",")
	size := len(split)

	var allRecord []byte

	src := db
	if !tar.fileExists(src) {
		return format, []byte(fmt.Sprintf("Database %s missing!", src))
	} else {
		fInfo, _ := os.Stat(src)
		fsize := fInfo.Size()
		if fsize <= 1 {
			return format, []byte(fmt.Sprintf("Database %s is empty!", src))
		} else {
			var allRecords []MyStruct
			xFields := new(MyStruct)
			var tmpStruct MyStruct
			line := ""

			file, err := os.Open(src)
			CheckError("SelectSearch(1)", err)

			defer file.Close()
			var r io.Reader = file
			sc := bufio.NewScanner(r)
			containsAll := true

			for sc.Scan() {
				line = strings.ToLower(sc.Text())
				for i := 0; i < size; i++ {
					for x := 0; x < size; x++ {
						if !strings.Contains(line, split[x]) {
							containsAll = false
						}
					}
				}
				if containsAll {
					in := []byte(sc.Text())
					err = json.Unmarshal(in, &tmpStruct)
					CheckError("SelectSearch(2)", err)

					xFields.Id = tmpStruct.Id
					xFields.Key = string(tmpStruct.Key)
					xFields.Data = string(tmpStruct.Data)
					allRecords = append(allRecords, *xFields)
				}
				containsAll = true

			}
			allRecord, err = tar.MyMarshal(allRecords)
			CheckError("SelectSearch(3)", err)
		}
	}
	return format, allRecord

}

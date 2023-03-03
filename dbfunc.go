package tardigrade

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CreatedDBCopy creates a copy of the Database and store in UserHomeDir().
func (tar *Tardigrade) CreatedDBCopy() (msg string, status bool) {
	status = true
	dirname, err := os.UserHomeDir()
	CheckError("CreatedDBCopy(0)", err)
	target := "tardigradecopy.db"

	src := DBFile
	if !tar.fileExists(src) {
		msg = fmt.Sprintf("Failed: database %s missing!", src)
		return msg, false
	}
	fin, err := os.Open(src)
	CheckError("CreatedDBCopy(1)", err)
	defer fin.Close()

	PATH_SEPARATOR := tar.GetOS()
	dst := fmt.Sprintf("%s%s%s", dirname, string(PATH_SEPARATOR), target)
	buf := make([]byte, 1024)
	tmp, err := os.Create(dst)
	CheckError("CreatedDBCopy(2)", err)
	defer tmp.Close()

buffering:
	for {
		n, err := fin.Read(buf)
		if err != nil && err != io.EOF {
			CheckError("CreatedDBCopy(3)", err)
			msg = "Failed: buffer error failed to create database!"
			return msg, false
		}

		if n == 0 {
			fin.Close()
			tmp.Close()
			break buffering
		}

		if _, err := tmp.Write(buf[:n]); err != nil {
			CheckError("CreatedDBCopy(4)", err)
			msg = "Failed: permission error failed to create database!"
			return msg, false
		}
	}
	msg = fmt.Sprintf("Copy: %s", dst)
	return msg, true
}

// CreateDB - This function will create a database file if it does not exist and return true | false
func (tar *Tardigrade) CreateDB() (msg string, status bool) {
	status = true
	fname := DBFile
	pwd, _ := filepath.Abs(fname)
	if !tar.fileExists(fname) {
		_, err := os.Create(fname)
		CheckError("CreateDB(2)", err)
		if !tar.fileExists(fname) {
			CheckError("CreateDB(3)", err)
			status = false
			return fmt.Sprintf("Failed: %v", pwd), status
		}
	} else {
		status = false
		return fmt.Sprintf("Exist: %v", pwd), status
	}
	return fmt.Sprintf("Created: %v", pwd), status
}

// DeleteDB - WARNING - this function delete the database file return true | false
func (tar *Tardigrade) DeleteDB() (msg string, status bool) {
	fname := DBFile
	status = true
	pwd, _ := filepath.Abs(fname)
	if tar.fileExists(fname) {
		delete := os.Remove(fname)
		CheckError("DeleteDB(1)", delete)
		if tar.fileExists(fname) {
			status = false
			return fmt.Sprintf("Failed: %v", pwd), status
		}
	} else {
		status = false
		return fmt.Sprintf("Unavailable: %v", pwd), status
	}
	return fmt.Sprintf("Removed: %v", pwd), status
}

// fileExists function will check if the database exists and return true / false
func (tar *Tardigrade) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// EmptyDB function - WARNING - this will destroy the database and all data stored in it!
func (tar *Tardigrade) EmptyDB() (msg string, status bool) {
	_, status = tar.DeleteDB()
	if status {
		_, status = tar.CreateDB()
		if !status {
			status = false
			msg = "Failed: no permission to re-create!"
			return msg, status
		}
	} else {
		status = false
		msg = "Missing: could not find database!"
		return msg, status
	}
	msg = "Empty: database now clean!"
	return msg, status
}

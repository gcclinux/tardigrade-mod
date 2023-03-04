package tardigrade

// Built Sat 4 Mar 12:32:07 GMT 2023

import (
	"bytes"
	"encoding/json"
)

// MyMarshal function is adapted to SetEscapeHTML to false before encoding
func (tar *Tardigrade) MyMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// MyIndent function is adapted to SetEscapeHTML to false before encoding and indenting
func (tar *Tardigrade) MyIndent(v interface{}, prefix, indent string) ([]byte, error) {
	b, err := tar.MyMarshal(v)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	err = json.Indent(&buffer, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

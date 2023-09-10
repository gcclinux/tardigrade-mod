package tardigrade

// Built Sat 4 Mar 12:32:07 GMT 2023

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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

// MyEncode returns the base64 encoding of source
func (tar *Tardigrade) MyEncode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// MyDecode returns the bytes represented by the base64 string s
func (tar *Tardigrade) MyDecode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

var bytez = []byte{33, 45, 67, 28, 75, 15, 26, 77, 97, 25, 28, 91, 55, 31, 44, 69}

// Encrypt method is to encrypt or hide any classified text
func (tar *Tardigrade) MyEncrypt(text, Password string) (string, error) {
	block, err := aes.NewCipher([]byte(Password))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytez)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return tar.MyEncode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func (tar *Tardigrade) MyDecrypt(text, Password string) (string, error) {
	block, err := aes.NewCipher([]byte(Password))
	if err != nil {
		return "", err
	}
	cipherText := tar.MyDecode(text)
	cfb := cipher.NewCFBDecrypter(block, bytez)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

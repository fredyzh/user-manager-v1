package v1

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
)

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// aesCBCEncrypt
func aesCBCEncrypt(rawData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//
	blockSize := block.BlockSize()
	rawData = pkcs7Padding(rawData, blockSize)
	//
	cipherText := make([]byte, blockSize+len(rawData))
	//block 16
	iv := cipherText[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	//block
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[blockSize:], rawData)

	return cipherText, nil
}

func aesCBCDecrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	if len(encryptData) < blockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := encryptData[:blockSize]
	encryptData = encryptData[blockSize:]

	// CBC mode always works in whole blocks.
	if len(encryptData)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	//
	encryptData = pkcs7UnPadding(encryptData)
	return encryptData, nil
}

func sha1Hash2(key []byte) []byte {
	h := sha1.New()
	h.Write(key)
	hashData := h.Sum(nil)
	keyBuffer := bytes.NewBuffer(hashData)

	h.Reset()
	h.Write(hashData)
	keyBuffer.Write(h.Sum(nil))

	return keyBuffer.Bytes()[:32]
}

// Encrypt aes
func Encrypt(data, key []byte) ([]byte, error) {
	return aesCBCEncrypt(data, sha1Hash2(key))
}

// Decrypt aes
func Decrypt(data, key []byte) ([]byte, error) {
	return aesCBCDecrypt(data, sha1Hash2(key))
}

// Encrypt aes
func EncryptToString(plan string, key []byte) (string, error) {
	data, err := aesCBCEncrypt([]byte(plan), sha1Hash2(key))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// Decrypt aes
func DecryptFromString(encrypt string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}

	plan, err := aesCBCDecrypt(data, sha1Hash2(key))
	if err != nil {
		return "", err
	}
	return string(plan), nil
}

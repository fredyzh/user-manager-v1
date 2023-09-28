package v1

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var secret = "jwt-test-secrets"

func TestAESCBC(t *testing.T) {
	data := []byte("abcdefg")
	key := []byte("123456")

	should := require.New(t)

	cipherData, err := Encrypt(data, key)
	should.NoError(err)
	t.Logf("cipher data: %s", cipherData)

	rawData, err := Decrypt(cipherData, key)
	should.NoError(err)
	t.Logf("raw data: %s", rawData)

	should.Equal(data, []byte(rawData))
}

func TestEncryptString(t *testing.T) {
	data := "abcdefg"
	key := []byte("123456")

	should := require.New(t)

	cipherData, err := EncryptToString(data, key)
	should.NoError(err)
	t.Logf("cipher data: %s", cipherData)

	rawData, err := DecryptFromString(cipherData, key)
	should.NoError(err)
	t.Logf("raw data: %s", rawData)

	should.Equal(data, rawData)
}

func TestParseToken(t *testing.T) {
	should := require.New(t)

	// var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyLW1hbmFnZXJfMDAwMCIsImV4cCI6MTY5NTgzNjEwMiwiaWF0IjoxNjk1ODM1MjAyLCJpc3MiOiJVc2VyTWFuYWdlciIsIm5hbWUiOiJhcHBfYWRtaW5AdGVzdC5jb20iLCJzdWIiOiI2NTEyZTliZGM3YjVlNDliNTEzNjU3MzMiLCJ0eXAiOiJKV1QifQ.WMUeo392SZcIwLhN2FcXmhbeifRH8ZUgET-m6qYJAgM"

	tkn, ep, err := ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ1c2VyLW1hbmFnZXJfMDAwMCIsImV4cCI6MTY5NTg0NzQ3MiwiaWF0IjoxNjk1ODQwNDgyLCJpc3MiOiJVc2VyTWFuYWdlciIsIm5hbWUiOiJhcHBfYWRtaW5AdGVzdC5jb20iLCJzdWIiOiI2NTEyZTliZGM3YjVlNDliNTEzNjU3MzMiLCJ0eXAiOiJKV1QifQ.J3UvVyumnL3MdGKewYslIIr8uQ3mVnhhYqEL_WtK1rg", secret)
	should.NoError(err)
	log.Println(ep)
	log.Println(tkn)
}

package v1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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

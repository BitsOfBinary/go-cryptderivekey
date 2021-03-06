package cryptderivekey

import (
	"bytes"
	"testing"
)

func TestMD5CryptDeriveKey(t *testing.T) {
	md5Test, _ := CryptDeriveKey([]byte("abcd"), "md5")

	result := []byte{0x16, 0x76, 0xaa, 0x88, 0x60, 0xd8, 0x0c, 0xe0, 0xc2, 0x00, 0xad, 0x16, 0xb4, 0xd9, 0x45, 0xfb, 0xb5, 0x54, 0xb7, 0xc4, 0x87, 0x06, 0x3a, 0xf0, 0x69, 0xaf, 0x18, 0xc8, 0xc2, 0xac, 0x10, 0x15}

	if !bytes.Equal(md5Test, result) {
		t.Errorf("CryptDeriveKey for MD5 did not return the correct result.")
	}
}

func TestSHA1CryptDeriveKey(t *testing.T) {
	sha1Test, _ := CryptDeriveKey([]byte("abcd"), "sha1")

	result := []byte{0xf6, 0x12, 0xd1, 0xdc, 0x34, 0x6c, 0x8c, 0xf8, 0x9d, 0xdf, 0xe5, 0x40, 0x61, 0xeb, 0x90, 0xa9, 0x4d, 0x72, 0x34, 0x0e, 0xce, 0x5a, 0xe6, 0xb8, 0x68, 0x42, 0x03, 0x92, 0xe4, 0xeb, 0x78, 0xa9, 0xc9, 0xbc, 0x13, 0x3e, 0x6a, 0x73, 0xd4, 0x54}

	if !bytes.Equal(sha1Test, result) {
		t.Errorf("CryptDeriveKey for SHA1 did not return the correct result.")
	}
}

func TestInvalidHashType(t *testing.T) {
	_, err := CryptDeriveKey([]byte("abcd"), "test")

	if err == nil {
		t.Errorf("Invalid hash type accepted.")
	}
}

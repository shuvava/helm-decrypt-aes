package helm

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	if len(src) == 0 {
		return src
	}
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

// DecryptAES decrypts the given encrypted (by Helm) string using the provided key
func DecryptAES(encryptedString string, keyString string) ([]byte, error) {
	key := []byte(keyString)
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return nil, err
	}

	// Helm using only 32 bytes key
	for len(key) < 32 {
		key = append(key, 0)
	}
	if len(key) > 32 {
		key = key[:32]
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return PKCS5UnPadding(ciphertext), nil
}

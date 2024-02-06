package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesDecrypt(str string, newKey string, newIV string) ([]byte, error) {

	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(newKey))

	if err != nil {
		return nil, err
	}

	var blockSize = block.BlockSize()
	newIV = newIV[:blockSize]

	blockMode := cipher.NewCBCDecrypter(block, []byte(newIV))

	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)

	origData = PKCS7UnPadding(origData)

	return origData, nil

}

func AseEncrypt(str []byte, newKey []byte, newIV []byte) (string, error) {

	block, err := aes.NewCipher(newKey)
	if err != nil {
		return "", err
	}
	var blockSize = block.BlockSize()
	newIV = newIV[:blockSize]

	src, err := PKCS7Padding(str, blockSize)
	if err != nil {
		return "", err
	}

	var dst = make([]byte, len(src))

	var mode = cipher.NewCBCEncrypter(block, newIV)
	mode.CryptBlocks(dst, src)

	return base64.StdEncoding.EncodeToString(dst), nil
}

func PKCS7Padding(src []byte, blockSize int) ([]byte, error) {
	var pSize = blockSize - len(src)%blockSize
	var pText = bytes.Repeat([]byte{byte(pSize)}, pSize)
	return append(src, pText...), nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)

	unPadding := int(origData[length-1])

	return origData[:(length - unPadding)]
}

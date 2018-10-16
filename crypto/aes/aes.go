package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const (
	ivDefValue = "Gogoods, go good"
)

// AES加密, 使用CBC模式，注意key必须为16/24/32位长度，iv初始化向量为非必需参数
func Encrypt(plainText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	plainText = PKCS5Padding(plainText, blockSize)
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(ivDefValue)
	}
	blockMode := cipher.NewCBCEncrypter(block, ivValue)
	ciphertext := make([]byte, len(plainText))
	blockMode.CryptBlocks(ciphertext, plainText)

	return ciphertext, nil
}

// AES解密, 使用CBC模式，填充方式为PKCS5Padding，注意key必须为16/24/32位长度，iv初始化向量为非必需参数
func Decrypt(cipherText []byte, key []byte, iv ...[]byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(cipherText) < blockSize {
		return nil, errors.New("cipherText too short")
	}
	ivValue := ([]byte)(nil)
	if len(iv) > 0 {
		ivValue = iv[0]
	} else {
		ivValue = []byte(ivDefValue)
	}
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("cipherText is not a multiple of the block size")
	}
	blockModel := cipher.NewCBCDecrypter(block, ivValue)
	plainText := make([]byte, len(cipherText))
	blockModel.CryptBlocks(plainText, cipherText)
	plainText = PKCS5UnPadding(plainText)

	return plainText, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

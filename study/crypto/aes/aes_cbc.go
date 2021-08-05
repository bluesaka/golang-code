package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

// AesEncryptCBC AES加密CBC模式
func AesEncryptCBC(s string) (encryptStr string, err error) {
	origData := []byte(s)
	// 加密的密钥 16, 24, 32位
	iv := []byte("ABCDEFGHIJKLMNOP")

	// 生成加密用的block
	block, err := aes.NewCipher(iv)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()

	// PKCS5明文补码
	origData = PKCS5Padding(origData, blockSize)

	// CBC加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encryptData := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptData, origData)

	// hex encode
	encryptStr = hex.EncodeToString(encryptData)
	return
}

// AesDecryptCBC AES解密CBC模式
func AesDecryptCBC(s string) (decryptStr string, err error) {
	// hex decode
	encrypted, err := hex.DecodeString(s)
	if err != nil {
		return
	}

	// 解密的密钥
	iv := []byte("ABCDEFGHIJKLMNOP")

	// 生成加密用的block
	block, err := aes.NewCipher(iv)
	if err != nil {
		return
	}

	// CBC解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	blockMode.CryptBlocks(decrypted, encrypted)

	// 明文减码算法
	decrypted = PKCS5UnPadding(decrypted)
	decryptStr = string(decrypted)
	return
}

// PKCS5Padding 明文补码算法
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

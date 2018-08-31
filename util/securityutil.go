package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

//DesEncode des加密
func DesEncode(key string, iv string, content string) (string, error) {
	key += "00000000"
	date := []byte(key)
	desmodel, err := des.NewCipher(date[0:8])
	if err != nil {
		return "", err
	}
	des := []byte(content)
	des = pKCS5Padding(des, desmodel.BlockSize())
	bm := cipher.NewCBCEncrypter(desmodel, []byte(iv)[0:8])
	dst := make([]byte, len(des))
	bm.CryptBlocks(dst, des)
	result := base64.StdEncoding.EncodeToString(dst)
	return result, nil
}
func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)

}

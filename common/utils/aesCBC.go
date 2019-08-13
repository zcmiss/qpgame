package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//加密
func PswEncrypt(src, sKey, ivParameter string) string {
	key := []byte(sKey)
	iv := []byte(ivParameter)
	result, err := Aes128Encrypt([]byte(src), key, iv)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(result)
}

//解密
func PswDecrypt(src, sKey, ivParameter string) string {
	key := []byte(sKey)
	iv := []byte(ivParameter)
	var result []byte
	var err error
	result, err = base64.URLEncoding.DecodeString(src)
	if err != nil {
		panic(err)
	}
	origData, err := Aes128Decrypt(result, key, iv)
	if err != nil {
		panic(err)
	}
	return string(origData)
}
func Aes128Encrypt(origData, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCEncrypter(block, IV[:blockSize])
	origData = padding(origData, blockSize)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//加密字符不足16位的时候需要进行位数补齐
func padding(origData []byte, blockSize int) []byte {
	plainTextLength := len(origData)
	if plainTextLength%blockSize != 0 {
		plainTextLength = plainTextLength + (blockSize - plainTextLength%blockSize)
	}
	plainText := make([]byte, plainTextLength)
	for i, v := range origData {
		plainText[i] = v
	}
	return plainText
}

func unPadding(origData []byte) []byte {
	newOrigData := make([]byte, 0)
	for _, i := range origData {
		if i == 0 {
			break
		}
		newOrigData = append(newOrigData, i)
	}
	return newOrigData
}

func Aes128Decrypt(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unPadding(origData)
	return origData, nil
}

//func main(){
//	key := "4929e54f0f1ae929"
//	iv := "4929e54f0f1ae929"
//	encodingString := PswEncrypt("abccccd",key,iv)
//	decodingString := PswDecrypt(encodingString,key,iv);
//	fmt.Printf("AES-128-CBC\n加密：%s\n解密：%s\n",encodingString,decodingString)
//}

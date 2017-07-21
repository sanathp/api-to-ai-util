package vcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

const (
	secretAesKey = "!a#l=i?-et-#u$k|"
)

var vcryptConfig VcrptConfig

type VcrptConfig struct {
	SecretAesKey string
}

func InitVcrypt(configData VcrptConfig) {
	//secretAesKey should come from config
	vcryptConfig = configData
	//TODO: in future rotate these keys for security
	testString := "api.ai is future"

	encrypted, err := HexaAesEncrypt(testString)

	if err != nil {
		panic(err)
	}
	decrptyedStr, err := HexaAesDecrypt(encrypted)

	if err != nil {
		panic(err)
	}

	if decrptyedStr != testString {
		panic("Validation failed, vcrypt not working")
	}
}

func HexaAesEncrypt(inputstring string) (string, error) {
	text := []byte(inputstring)
	key := []byte(vcryptConfig.SecretAesKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return hex.EncodeToString(ciphertext), nil
}

func HexaAesDecrypt(hexastring string) (string, error) {
	text, err0 := hex.DecodeString(hexastring)
	key := []byte(vcryptConfig.SecretAesKey)
	if err0 != nil {
		return "", err0
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}

	return string(data), nil
}

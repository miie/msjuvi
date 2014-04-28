package crypt

import (
    "io"
    "errors"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "../logger"
    "../random"
    "../hash"
)

// Encrypts text using AES-256 with key. Key can be of arbitrary length (it's md5 hashed).
// Returns iv + encrypted string (iv is the first 16 bytes).
func AESEncrypt(key, text string) (string, error) {
    // Start by md5 hashing the key to always get a 32 chars string (meaning AES-256)
    // without the need to pad or cut the key.
    key, err := hash.MD5Hash(key)
    if err != nil {
    	logger.LogWarning("error when getting md5 hash. error: ", err)
    	return "", err
    }
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        logger.LogWarning("error when aes.NewCipher. error: ", err)
        return "", err
    }
    b := encodeBase64([]byte(text))
    ciphertext := make([]byte, len(b))
    randomStr, err := random.RandString(aes.BlockSize)
    if err != nil {
    	logger.LogWarning("error when getting random string. error: ", err)
    	return "", err
    }
    iv := []byte(randomStr)
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        logger.LogWarning("error when io.ReadFull. error: ", err)
        return "", err
    }
    cfb := cipher.NewCFBEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext, []byte(b))
    return (string(iv) + string(ciphertext)), nil
}

// Decrypts a AES-256 encrypted string. Internally key is md5 hashed (which means it needs 
// to be md5 hashed at encryption). The function assumes that the first 16 bytes of the encrypted
// text is the iv.
func AESDecrypt(key, text string) (string, error) {
	iv := []byte(text[:aes.BlockSize])
	key, err := hash.MD5Hash(key)
	if err != nil {
		logger.LogWarning("error when getting md5 hash for key. error: ", err)
		return "", err
	}
	txt := []byte(text[aes.BlockSize:])
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        logger.LogWarning("error when aes.NewCipher. error: ", err)
        return "", err
    }
    if len(txt) < aes.BlockSize {
        logger.LogWarning("error, cipher txt < aes.BlockSize.")
        return "", errors.New("error, cipher txt < aes.BlockSize.")
    }
    cfb := cipher.NewCFBDecrypter(block, iv)
    cfb.XORKeyStream(txt, txt)
    s, err := decodeBase64(string(txt))
    if err != nil {
    	logger.LogWarning("error when base64 decoding string. error: ", err)
    	return "", err
    }
    return string(s), nil
}

func encodeBase64(b []byte) string {
    return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) (data []byte, err error) {
    data, err = base64.StdEncoding.DecodeString(s)
    if err != nil {
        logger.LogWarning("error when decoding string. string & error: ", s, err)
    }
    return
}
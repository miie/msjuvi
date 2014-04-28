package crypt

import (
	"testing"
	"fmt"
	//"bytes"
	//"crypto/aes"
)

func TestAESEncrypt(t *testing.T) {
	s, err := AESEncrypt("testpassdsafasfasfeavarvaervarvaregvaervaeravarearaar", "test string to encrypt åäö")
	if err != nil {
		t.Errorf("got error: ", err)
	}
	fmt.Printf("TestAESEncrypt: encrypted string: %x\n", s)
}

func TestAESDecrypt(t *testing.T) {
	key := "a very very very very secsdaffafdafåäöäöåälä"
	strtoencryptdecrypt := "test string to encrypt åäö"
	
	fmt.Println("TestAESDecrypt: string to encrypt: ", strtoencryptdecrypt)
	s, err := AESEncrypt(key, strtoencryptdecrypt)
	fmt.Printf("TestAESDecrypt: encrypted string: %x\n", s)

	// Decryption should succeed
	ds, err := AESDecrypt(key, s)
	if err != nil {
		t.Errorf("1. got error when decrypting. err: ", err)
		return
	}

	if ds != strtoencryptdecrypt {
		t.Errorf("decrypt string does not match. stoencrypt & sdecrypt:", strtoencryptdecrypt, ds)
	}
	fmt.Println("TestAESDecrypt: decrypted string: ", ds)

	// Decryption should fail
	ds, err = AESDecrypt("wrong key", s)
	if err == nil {
		t.Errorf("expecting error!", err)
		return
	}

	if ds == strtoencryptdecrypt {
		t.Errorf("decrypt string does match (which is shouldn't). stoencrypt & sdecrypt:", strtoencryptdecrypt, ds)
	}
}

func BenchmarkEncrypt(b *testing.B) {
    key := "a very very very very secret key"
    plaintext := "some really really really long plaintext"
    for i := 0; i < b.N; i++ {
        AESEncrypt(key, plaintext)
    }
}

func BenchmarkDecrypt(b *testing.B) {
    key := "a very very very very secret key"
    plaintext := "some really really really long plaintext"
    for i := 0; i < b.N; i++ {
        b, _ := AESEncrypt(key, plaintext)
        AESDecrypt(key, b)
    }
}
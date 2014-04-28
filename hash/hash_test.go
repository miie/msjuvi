package hash

import (
	"testing"
)

func TestMD5(t *testing.T) {
	s, err := MD5Hash("test string to hash åäö")
	if err != nil {
		t.Errorf("TestMD5: got error: ", err)
	}
	if s != "1315c4500a0640b4657dc72e14374210" {
		t.Errorf("got unexpected hash. expected & got: ", "1315c4500a0640b4657dc72e14374210", s)
	}
}

func TestRipemd160(t *testing.T) {
	s, err := Ripemd160Hash("test string to hash åäö")
	if err != nil {
		t.Errorf("TestMD5: got error: ", err)
	}
	if s != "d4067ee6c3fd49b6220dbfa3ae86c540839b6ba5" {
		t.Errorf("got unexpected hash. expected & got: ", "d4067ee6c3fd49b6220dbfa3ae86c540839b6ba5", s)
	}
}

func TestSha256(t *testing.T) {
	s, err := Sha256Hash("test string to hash åäö")
	if err != nil {
		t.Errorf("TestMD5: got error: ", err)
	}
	if s != "d0f5267421c31dd77f9c590a439f07f93dd562b77a98eff17ca2479007d09c3e" {
		t.Errorf("got unexpected hash. expected & got: ", "d0f5267421c31dd77f9c590a439f07f93dd562b77a98eff17ca2479007d09c3e", s)
	}
}
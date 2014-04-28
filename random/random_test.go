package random

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	s, err := RandString(100)
	if err != nil {
		t.Errorf("error: ", err)
	}
	fmt.Println("TestRandString: got random string: ", s)
}

func TestRandInt(t *testing.T) {
	i, err := RandInt(100)
	if err != nil {
		t.Errorf("error: ", err)
	}
	fmt.Println("TestRandInt: got random int: ", i)
}

func TestPseudoRandInt(t *testing.T) {
	i, err := PseudoRandInt(100)
	if err != nil {
		t.Errorf("error: ", err)
	}
	fmt.Println("TestPseudoRandInt: got random int: ", i)
}

func TestRandIntMinMax(t *testing.T) {
	for i := 0; i < 100000; i++ {
		i, err := RandIntMinMax(10, 15)
		if err != nil {
			t.Errorf("error: ", err)
			return
		} else if i < 10 {
			t.Errorf("i < 10")
			return
		} else if i > 14 {
			t.Errorf("i > 14")
			return
		}
	}
}
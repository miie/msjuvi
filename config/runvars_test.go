package config

import (
	"fmt"
	"testing"
)

func Test0(t *testing.T) {

	SetVarsString("key0", "key0str")
	s, ok := GetVarsString("key0")
	if !ok {
		t.Errorf("1. Test0: did not get ok")
	}
	if s != "key0str" {
		t.Errorf("2. Test0: did not get key0str. got: %s", s)
	}
	fmt.Println("1. key0: ", s)

	SetVarsInteger("key1", 99)
	i, ok := GetVarsInteger("key1")
	if !ok {
		t.Errorf("3. Test0: did not get ok")
	}
	if i != 99 {
		t.Errorf("4. Test0: did not get 99. got: ", i)
	}
	fmt.Println("2. key1: ", i)

	SetVarsBoolean("key2", true)
	b, ok := GetVarsBoolean("key2")
	if !ok {
		t.Errorf("5. Test0: did not get ok")
	}
	if b != true {
		t.Errorf("6. Test0: did not get true. got: ", b)
	}
	fmt.Println("3. key2: ", b)

	vars := GetVars()
	i, _ = GetVarsInteger("key1")
	fmt.Println("4. key1: ", i)
	vars.SetVarsInt("key1", 101)
	i, _ = GetVarsInteger("key1")
	fmt.Println("5. key1: ", i)

	InitCfgVars("./tst.cfg")

	s, ok = GetVarsString("default:mailresultlist")
	if !ok {
		t.Errorf("could not get default:mailresultlist")
	} else {
		fmt.Println("default:mailresultlist: ", s)
	}

	i, ok = GetVarsInteger("sectionwithtype:optionwithtype1")
	if !ok {
		t.Errorf("could not get sectionwithtype:optionwithtype1")
	} else {
		fmt.Println("sectionwithtype:optionwithtype1: ", i)
	}

	s, ok = GetVarsString("noofpostgresconns:maxnoofconns")
	if !ok {
		t.Errorf("could not get noofpostgresconns:maxnoofconns")
	} else {
		fmt.Println("noofpostgresconns:maxnoofconns: ", s)
	}

	b, ok = GetVarsBoolean("sectionwithtype:optionwithtype2")
	if !ok {
		t.Errorf("could not get optionwithtype2:bool")
	} else {
		fmt.Println("optionwithtype2:bool: ", b)
	}

	s, ok = GetVarsString("sectionwithtype:optionwithtype0")
	if ok {
		t.Errorf("found optionwithtype0:string which we shouldn't")
	}

	s, ok = GetVarsString("optionwithtype0")
	if ok {
		t.Errorf("found optionwithtype0 which we shouldn't")
	}
}

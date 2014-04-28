package config

import (
	//"fmt"
	"sync"
	"strings"
	"errors"
	"../logger"
	"github.com/miie/goconfig"
)

type RunVars struct {
	TstInt	int
	TstStr	string
	TstBool	bool
}

type Vars struct {
	strvalues 	map[string]string
	intvalues 	map[string]int
	boolvalues 	map[string]bool
	sync.Mutex
}

func New() *Vars {
	return &Vars{make(map[string]string), make(map[string]int), make(map[string]bool), sync.Mutex{}}
}

var vars = New()

func GetVars() *Vars {
	return vars
}

func SetVarsString(key, s string) {
	vars.SetVarsStr(key, s)
}

func GetVarsString(key string) (s string, ok bool) {
	s, ok = vars.GetVarsStr(key)
	return
}

func (v *Vars) SetVarsStr(key, s string) {
	v.Lock()
	defer v.Unlock()
	v.strvalues[key] = s
}

func (v *Vars) GetVarsStr(key string) (s string, ok bool) {
	s, ok = v.strvalues[key]
	return
}

func SetVarsInteger(key string, i int) {
	vars.SetVarsInt(key, i)
}

func GetVarsInteger(key string) (i int, ok bool) {
	i, ok = vars.GetVarsInt(key)
	return
}

func (v *Vars) SetVarsInt(key string, i int) {
	v.Lock()
	defer v.Unlock()
	v.intvalues[key] = i
}

func (v *Vars) GetVarsInt(key string) (i int, ok bool) {
	i, ok = v.intvalues[key]
	return
}

func SetVarsBoolean(key string, b bool) {
	vars.SetVarsBool(key, b)
}

func GetVarsBoolean(key string) (b bool, ok bool) {
	b, ok = vars.GetVarsBool(key)
	return
}

func (v *Vars) SetVarsBool(key string, b bool) {
	v.Lock()
	defer v.Unlock()
	v.boolvalues[key] = b
}

func (v *Vars) GetVarsBool(key string) (b bool, ok bool) {
	b, ok = v.boolvalues[key]
	return
}

func setvar(option *string, section *string, f *goconfig.ConfigFile) {
	typed := strings.Split(*option, "|")
	lentyped := len(typed)
	if lentyped > 2 {
		logger.LogWarning("found more than one | in option, please user <optionname|type[int,str,bool]=variable>. option: ", *option)
	} else if lentyped == 2 {
		switch typed[1] {
			case "str":
				s, err := f.GetString(*section, *option)
				if err != nil {
					logger.LogWarning("error when getting string for section & option: ", *section, *option)
				}
				SetVarsString(*section + ":" + typed[0], s)
			case "int":
				i, err := f.GetInt64(*section, *option)
				if err != nil {
					logger.LogWarning("error when getting \"int string\" for section & option: ", *section, *option)
				}
				SetVarsInteger(*section + ":" + typed[0], int(i))
			case "bool":
				b, err := f.GetBool(*section, *option)
				if err != nil {
					logger.LogWarning("error when getting \"bool string\" for section & option: ", *section, *option)
				}
				SetVarsBoolean(*section + ":" + typed[0], b)
			default:
				logger.LogWarning("error, got typed option with unknown type. section & option: ", *section, *option)
		}
	} else {
		s, err := f.GetString(*section, *option)
		if err != nil {
			logger.LogWarning("error when getting string for section & option: ", *section, *option)
		}
		SetVarsString(*section + ":" + typed[0], s)
	}
}

func InitCfgVars(configfilepath string) error {
	f, err := GetConfFile(configfilepath)
	if err != nil {
		logger.LogWarning("error when getting conf file. cannot initialize cfg vars.")
		return errors.New("error when getting conf file. err: " + err.Error())
	}
	sections := f.GetSections()
	for sectionsindex := range sections {
		options, err := f.GetOptions(sections[sectionsindex], false)
		if err != nil {
			logger.LogWarning("error when getting options for section. section: ", sections[sectionsindex])
		}
		for optionsindex := range options {
			setvar(&options[optionsindex], &sections[sectionsindex], f)
		}
	}
	return nil
}

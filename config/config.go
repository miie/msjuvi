package config

import (
	"github.com/miie/msjuvi/logger"
	"github.com/miie/goconfig"
)

func GetConfFile(configfilepath string) (c *goconfig.ConfigFile, err error) {
	c, err = goconfig.ReadConfigFile(configfilepath)
	if err != nil {
		logger.LogWarning("could not open config file: " + configfilepath + ", got error: " + err.Error())
	}
	return
}

func GetString(f *goconfig.ConfigFile, key string, section string) (string, error) {
	s, err := f.GetString(section, key)
	if err != nil {
		logger.LogWarning("config: could not get key, got error: " + err.Error())
	}
	return s, err
}

func GetInt64(f *goconfig.ConfigFile, key string, section string) (int64, error) {
	i, err := f.GetInt64(section, key)
	if err != nil {
		logger.LogWarning("config: could not get key, got error: " + err.Error())
	}
	return i, err
}

func GetBool(f *goconfig.ConfigFile, key string, section string) (bool, error) {
	b, err := f.GetBool(section, key)
	if err != nil {
		logger.LogWarning("config: could not get key, got error: " + err.Error())
	}
	return b, err
}

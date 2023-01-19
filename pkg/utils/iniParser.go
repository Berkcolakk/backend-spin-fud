package iniParser

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func GetValueFromINI(section string, key string) (string, error) {
	cfg, err := ini.Load("../config/config.ini")
	if err != nil {
		return "", fmt.Errorf("Fail to read file: %v", err)
	}

	value := cfg.Section(section).Key(key).String()
	if value == "" {
		return "", fmt.Errorf("Key not found in section")
	}
	return value, nil
}

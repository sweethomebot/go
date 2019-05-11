package core

import (
	"os"
	"io/ioutil"
)

func ArrayContains(arr []string, key string) bool {
	for _, a := range arr {
		if a == key {
			return true
		}
	}
	return false
}

func HasKey(dic map[string]string, key string) bool {
	_, ok := dic[key]
	return ok
}

func HasKeys(dic map[string]string, keys []string) bool {
	for i := 0; i < len(keys); i++ {
		_, ok := dic[keys[i]]
		if !ok {
			return false
		}
	}
	return true
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FileGetContents(ch Channel, path string) (string, bool) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		LogE(ch, "Could not read "+path+" file!")
		return "", false
	}
	return string(data), true
}
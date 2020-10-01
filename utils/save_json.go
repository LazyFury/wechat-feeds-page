package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// SaveJSON 保存json
func SaveJSON(v interface{}) (path string, err error) {

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return
	}

	dir, err := GetDir("./static", "")
	if err != nil {
		return
	}
	path = filepath.Join(dir, "temp.json")
	j, err := os.Create(path)
	if err != nil {
		return
	}
	_, err = j.Write(jsonBytes)
	return
}

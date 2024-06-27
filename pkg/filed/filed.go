package filed

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func JsonFile(jsonData any, savePathName string) error {
	if savePathName == "" {
		return errors.New("file save path must")
	}
	if filepath.Ext(savePathName) == "" {
		savePathName = savePathName + ".json"
	}
	jsonStr, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	return os.WriteFile(savePathName, jsonStr, 0644)
}

package filed

import (
	"encoding/json"
	"errors"
	"os"
)

func JsonFile(jsonData any, savePathName string) error {
	if savePathName == "" {
		return errors.New("file save path must")
	}
	jsonStr, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	return os.WriteFile(savePathName, jsonStr, 0644)
}

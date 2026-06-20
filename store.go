package main

import (
	"encoding/json"
	"os"
)

const dataFile = "data.json"

func loadData() (*DataStore, error) {
	data := &DataStore{}
	raw, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return nil, err
	}
	if len(raw) == 0 {
		return data, nil
	}
	err = json.Unmarshal(raw, data)
	return data, err
}

func saveData(data *DataStore) error {
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, raw, 0644)
}

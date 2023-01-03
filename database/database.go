package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Database struct {
	data []map[string]interface{}
}

func New() *Database {
	return &Database{
		data: make([]map[string]interface{}, 0),
	}
}

func (db *Database) Set(key string, value interface{}) {
	entry := make(map[string]interface{})
	entry[key] = value
	db.data = append(db.data, entry)
}

func (db *Database) Get(key string) interface{} {
	for _, entry := range db.data {
		if value, ok := entry[key]; ok {
			return value
		}
	}
	return nil
}

func (db *Database) Delete(key string) {
	for i, entry := range db.data {
		if _, ok := entry[key]; ok {
			db.data = append(db.data[:i], db.data[i+1:]...)
			break
		}
	}
}

func (db *Database) Save(filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	var existingData []map[string]interface{}
	jsonData, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, &existingData); err != nil {
		return err
	}
	existingData = append(existingData, db.data...)

	jsonData, err = json.Marshal(existingData)
	if err != nil {
		return err
	}
	if _, err = f.WriteAt(jsonData, 0); err != nil {
		return err
	}
	return nil
}

func (db *Database) Load(filename string) error {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonData, &db.data)
}

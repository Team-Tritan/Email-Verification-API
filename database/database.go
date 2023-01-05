package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Database struct {
	data map[string]interface{}
}

func New(filename string) *Database {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	data := map[string]interface{}{}
	json.Unmarshal(jsonData, &data)

	return &Database{
		data: data,
	}

}

func (db *Database) Set(key string, value interface{}) {
	db.data[key] = value
}

func (db *Database) Get(key string) interface{} {
	return db.data[key]

}

func (db *Database) Delete(key string) {
	delete(db.data, key)
}

func (db *Database) Save(filename string) error {
	file, _ := json.MarshalIndent(db.data, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)

	return nil
}

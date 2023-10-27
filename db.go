package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type Vault struct {
  Key   string  `json:"key"`
  Data  string  `json:"data"`
}

type Database struct {
  List    []Vault   `json:"list"` 
}

var db Database

func SaveDatabase() {
  data, err := json.Marshal(&db)
  if err != nil {
    log.Fatalf("Error saving database: %s", err.Error())
    os.Exit(1)
  }

  err = os.Mkdir("db", os.ModePerm)
  if err != nil && !os.IsExist(err) {
    log.Fatalf("Error creating database directory: %s", err)
    os.Exit(1)
  }

  if os.WriteFile(path.Join("db", "db.json"), data, 0666) != nil {
    log.Fatalf("Error saving database: %s", err.Error())
    os.Exit(1)
  }
}

func LoadDatabase() {
  data, err := os.ReadFile(path.Join("db", "db.json"))
  if os.IsNotExist(err) {
    db = Database{}
    return
  } 

  if err != nil {
    log.Fatalf("Error loading database: %s", err.Error())
  }

  err = json.Unmarshal(data, &db)
  if err != nil {
    log.Fatalf("Error loading database: %s", err.Error())
    os.Exit(1)
  }
}

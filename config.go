package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
  Port            string  `json:"port"`
  Password        string  `json:"password"`
  MaxVaultSize    int     `json:"max_vault_size"`
  MaxVaultCount   int     `json:"max_vault"`
}

var cfg Config

func ReadConfig(){
  content, err := os.ReadFile("config.json")
  if err != nil {
    log.Fatal("Cannot read config.json")
    os.Exit(1)
  }

  err = json.Unmarshal(content, &cfg)
  if err != nil {
    log.Fatal("Cannot read config.json")
    os.Exit(1)
  }
}

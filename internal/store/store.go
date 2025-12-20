package store

import (
	"encoding/json"
	"log"
	"os"
)

type PortfolioData struct {
	Ping struct {
		Message string `json:"message"`
	} `json:"ping"`
	Skills      map[string]string `json:"skills"`
	Projects    map[string]string `json:"projects"`
	Experiences map[string]string `json:"experiences"`
	Contact     map[string]string `json:"contact"`
	Blog        struct {
		URL string `json:"url"`
	} `json:"blog"`
	About      map[string]string `json:"about"`
	EasterEggs struct {
		Egg1 map[string]string `json:"egg1"`
		Egg2 map[string]string `json:"egg2"`
	} `json:"easterEggs"`
	Help map[string]string `json:"help"`
}

var Data *PortfolioData

func Init() error {
	file, err := os.ReadFile("../../database/main.json")
	if err != nil {
		return err
	}

	Data = &PortfolioData{}
	if err := json.Unmarshal(file, Data); err != nil {
		return err
	}

	log.Println("Portfolio data loaded successfully")
	return nil
}

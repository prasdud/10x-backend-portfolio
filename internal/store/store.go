package store

import (
	"encoding/json"
	"log"
	"os"
)

type Skill struct {
	Name     string `json:"name"`
	Level    string `json:"level"`
	Category string `json:"category"`
}

type Project struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	Status       string   `json:"status"`
}

type Experience struct {
	ID          string `json:"id"`
	Company     string `json:"company"`
	Position    string `json:"position"`
	Duration    string `json:"duration"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type PortfolioData struct {
	Ping struct {
		Message string `json:"message"`
	} `json:"ping"`
	Skills      []Skill           `json:"skills"`
	Projects    []Project         `json:"projects"`
	Experiences []Experience      `json:"experiences"`
	Contact     map[string]string `json:"contact"`
	Blog        struct {
		URL string `json:"url"`
	} `json:"blog"`
	About      map[string]string `json:"about"`
	EasterEggs struct {
		Egg1 map[string]string `json:"egg1"`
		Egg2 map[string]string `json:"egg2"`
	} `json:"easterEggs"`
	Help           map[string]string `json:"help"`
	Certifications []string          `json:"certifications"`
	Achievements   []string          `json:"achievements"`
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

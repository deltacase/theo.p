package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Pour la question 5 : remplacez /api/games par /api/broken, puis par /api/slow
const gamesURL = "http://localhost:8081/api/games"
const tournamentsURL = "http://localhost:8081/api/tournaments"

type Game struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Year   int      `json:"year"`
	Studio string   `json:"studio"`
	Genres []string `json:"genres"`
}

type Tournament struct {
	ID     int     `json:"id"`
	GameID int     `json:"gameId"`
	Name   string  `json:"name"`
	City   string  `json:"city"`
	Date   string  `json:"date"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}

var client = &http.Client{Timeout: 3 * time.Second}

func FetchGames() ([]Game, error) {
	resp, err := client.Get(gamesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("l'API a répondu %d", resp.StatusCode)
	}

	var games []Game
	err = json.NewDecoder(resp.Body).Decode(&games)
	return games, err
}

func FetchTournaments() ([]Tournament, error) {
	resp, err := client.Get(tournamentsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("l'API a répondu %d", resp.StatusCode)
	}

	var tournaments []Tournament
	err = json.NewDecoder(resp.Body).Decode(&tournaments)
	return tournaments, err
}

func afficherJeux() {
	games, err := FetchGames()
	if err != nil {
		log.Fatal("Impossible de récupérer les jeux : ", err)
	}

	tournaments, err := FetchTournaments()
	if err != nil {
		log.Fatal("Impossible de récupérer les tournois : ", err)
	}

	byGame := make(map[int][]Tournament)
	for _, t := range tournaments {
		byGame[t.GameID] = append(byGame[t.GameID], t)
	}

	for _, g := range games {
		fmt.Printf("%-25s %d   %d tournoi(s)\n", g.Name, g.Year, len(byGame[g.ID]))
	}
}

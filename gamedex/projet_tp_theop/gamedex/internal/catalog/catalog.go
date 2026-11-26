package catalog

import (
	"slices"
	"strings"

	"gamedex/internal/api"
)

func Filter(games []api.Game, query string) []api.Game {
	result := []api.Game{}
	for _, g := range games {
		if strings.Contains(strings.ToLower(g.Name), strings.ToLower(query)) {
			result = append(result, g)
		}
	}
	return result
}

func FindGame(games []api.Game, id int) (api.Game, bool) {
	for _, g := range games {
		if g.ID == id {
			return g, true
		}
	}
	return api.Game{}, false
}

func TournamentsOf(tournaments []api.Tournament, gameID int) []api.Tournament {
	result := []api.Tournament{}
	for _, t := range tournaments {
		if t.GameID == gameID {
			result = append(result, t)
		}
	}
	return result
}

func Genres(games []api.Game) []string {
	result := []string{}
	for _, g := range games {
		for _, genre := range g.Genres {
			if !slices.Contains(result, genre) {
				result = append(result, genre)
			}
		}
	}
	slices.Sort(result)
	return result
}

func SortByLikes(games []api.Game, likes func(int) int) {
	slices.SortStableFunc(games, func(a, b api.Game) int {
		return likes(b.ID) - likes(a.ID)
	})
}

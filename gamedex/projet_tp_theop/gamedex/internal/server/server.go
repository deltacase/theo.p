package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gamedex/internal/api"
	"gamedex/internal/catalog"
	"gamedex/internal/store"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

type Server struct {
	Games       []api.Game
	Tournaments []api.Tournament
}

type JeuListe struct {
	Game  api.Game
	Likes int
}

type ListePage struct {
	Recherche string
	Genre     string
	Genres    []string
	Games     []JeuListe
}

type JeuPage struct {
	Game        api.Game
	Tournaments []api.Tournament
	Likes       int
}

type ErreurPage struct {
	Code    int
	Message string
}

func (s *Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.accueil)
	mux.HandleFunc("GET /jeux", s.liste)
	mux.HandleFunc("GET /jeux/{id}", s.detail)
	mux.HandleFunc("POST /jeux/{id}/like", s.like)

	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	return mux
}

func show(w http.ResponseWriter, page string, data any) {
	err := tmpl.ExecuteTemplate(w, page, data)
	if err != nil {
		log.Println(err)
	}
}

func showError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	show(w, "erreur.html", ErreurPage{Code: code, Message: message})
}

func (s *Server) accueil(w http.ResponseWriter, r *http.Request) {
	show(w, "accueil.html", nil)
}

func (s *Server) liste(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	genre := r.URL.Query().Get("genre")

	games, err := api.SearchGames(q, genre)
	if err != nil {
		log.Println(err)
		showError(w, http.StatusBadGateway, "Le service des jeux ne répond pas.")
		return
	}

	catalog.SortByLikes(games, store.Likes)

	liste := []JeuListe{}
	for _, g := range games {
		liste = append(liste, JeuListe{Game: g, Likes: store.Likes(g.ID)})
	}

	show(w, "liste.html", ListePage{
		Recherche: q,
		Genre:     genre,
		Genres:    catalog.Genres(s.Games),
		Games:     liste,
	})
}

func (s *Server) detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		showError(w, http.StatusBadRequest, "Identifiant invalide.")
		return
	}

	game, ok := catalog.FindGame(s.Games, id)
	if !ok {
		showError(w, http.StatusNotFound, "Ce jeu n'existe pas.")
		return
	}

	show(w, "jeu.html", JeuPage{
		Game:        game,
		Tournaments: catalog.TournamentsOf(s.Tournaments, id),
		Likes:       store.Likes(id),
	})
}

func (s *Server) like(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		showError(w, http.StatusBadRequest, "Identifiant invalide.")
		return
	}

	if _, ok := catalog.FindGame(s.Games, id); !ok {
		showError(w, http.StatusNotFound, "Ce jeu n'existe pas.")
		return
	}

	store.AddLike(id)
	http.Redirect(w, r, "/jeux/"+strconv.Itoa(id), http.StatusSeeOther)
}

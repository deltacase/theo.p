package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gamedex/internal/api"
	"gamedex/internal/catalog"
)

var tmpl = template.Must(template.ParseGlob("web/templates/*.html"))

type Server struct {
	Games       []api.Game
	Tournaments []api.Tournament
}

type ListePage struct {
	Recherche string
	Games     []api.Game
}

type JeuPage struct {
	Game        api.Game
	Tournaments []api.Tournament
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
	games := catalog.Filter(s.Games, q)
	show(w, "liste.html", ListePage{Recherche: q, Games: games})
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

	tournaments := catalog.TournamentsOf(s.Tournaments, id)
	show(w, "jeu.html", JeuPage{Game: game, Tournaments: tournaments})
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"gamedex/internal/api"
	"gamedex/internal/server"
)

func main() {
	port := flag.String("port", "8080", "port du serveur")
	flag.Parse()

	games, err := api.FetchGames()
	if err != nil {
		log.Fatal("Impossible de charger les jeux : ", err)
	}
	tournaments, err := api.FetchTournaments()
	if err != nil {
		log.Fatal("Impossible de charger les tournois : ", err)
	}

	srv := &server.Server{Games: games, Tournaments: tournaments}

	fmt.Println("GameDex sur http://localhost:" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, srv.Routes()))
}

# Mini TP 1

**3.** `/api/games/999` → 404 (jeu introuvable). `/api/games/abc` → 400 (l'id doit être un nombre).

**4.** Le jeu 4 (Street Fighter 6) a 4 tournois : Marseille, Paris, Tokyo, Bruxelles.

**5.** Jeux de course : Rocket League, Trackmania, Mario Kart 8 Deluxe. Jeux de 2018 : Celeste, Dead Cells.

**6.** `/api/stats` : sans clé → 401, clé `eleve` → 403, clé `prof` → 200.

**Bonus.** `/api/broken` → 500, tout de suite. `/api/slow` → 200, après 5 secondes.
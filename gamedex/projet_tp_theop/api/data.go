package main

// Données de démonstration : les jeux sont réels, les tournois sont inventés.

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

var games = []Game{
	{1, "Celeste", 2018, "Maddy Makes Games", []string{"plateforme", "indé"}},
	{2, "Hades", 2020, "Supergiant Games", []string{"rogue-like", "action"}},
	{3, "Rocket League", 2015, "Psyonix", []string{"sport", "course"}},
	{4, "Street Fighter 6", 2023, "Capcom", []string{"combat"}},
	{5, "Stardew Valley", 2016, "ConcernedApe", []string{"simulation", "indé"}},
	{6, "League of Legends", 2009, "Riot Games", []string{"moba", "stratégie"}},
	{7, "Valorant", 2020, "Riot Games", []string{"fps", "tactique"}},
	{8, "Dead Cells", 2018, "Motion Twin", []string{"rogue-like", "action", "indé"}},
	{9, "Trackmania", 2020, "Nadeo", []string{"course"}},
	{10, "Mario Kart 8 Deluxe", 2017, "Nintendo", []string{"course", "party"}},
	{11, "Hollow Knight", 2017, "Team Cherry", []string{"metroidvania", "action", "indé"}},
	{12, "Minecraft", 2011, "Mojang", []string{"sandbox", "survie"}},
}

var tournaments = []Tournament{
	{1, 3, "Lyon Rocket Cup", "Lyon", "2027-02-13", 45.7640, 4.8357},
	{2, 3, "Aerial Masters", "Madrid", "2027-04-03", 40.4168, -3.7038},
	{3, 3, "Boost Open", "Copenhague", "2027-06-19", 55.6761, 12.5683},
	{4, 4, "Fight Night Marseille", "Marseille", "2027-01-30", 43.2965, 5.3698},
	{5, 4, "Combo Breaker EU", "Paris", "2027-03-20", 48.8566, 2.3522},
	{6, 4, "Rising Dragon", "Tokyo", "2027-08-07", 35.6762, 139.6503},
	{7, 4, "Hadoken Series", "Bruxelles", "2027-05-15", 50.8503, 4.3517},
	{8, 6, "Rift Clash", "Berlin", "2027-03-06", 52.5200, 13.4050},
	{9, 6, "Summoner Cup", "Lyon", "2027-05-22", 45.7640, 4.8357},
	{10, 6, "Baron Masters", "Séoul", "2027-10-16", 37.5665, 126.9780},
	{11, 7, "Spike Rush Lille", "Lille", "2027-02-27", 50.6292, 3.0573},
	{12, 7, "Clutch Masters", "Londres", "2027-04-24", 51.5074, -0.1278},
	{13, 7, "Defuse Open", "Barcelone", "2027-07-10", 41.3874, 2.1686},
	{14, 9, "Shortcut Challenge", "Bordeaux", "2027-03-13", 44.8378, -0.5792},
	{15, 9, "Nadeo Speed Cup", "Montpellier", "2027-06-05", 43.6108, 3.8767},
	{16, 10, "Banana Cup", "Nantes", "2027-02-06", 47.2184, -1.5536},
	{17, 10, "Rainbow Road Party", "Genève", "2027-04-10", 46.2044, 6.1432},
	{18, 1, "Speedrun Summit", "Toulouse", "2027-09-18", 43.6047, 1.4442},
	{19, 8, "Prisoners Run", "Bordeaux", "2027-11-06", 44.8378, -0.5792},
	{20, 12, "Build Battle", "Amsterdam", "2027-07-24", 52.3676, 4.9041},
}

package picks

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	dsn string
	db  *sql.DB
}

func NewStore(dataSourceName string) (s *Store, err error) {
	s = &Store{
		dsn: dataSourceName,
	}
	if s.db, err = sql.Open("sqlite3", dataSourceName); err != nil {
		return
	}
	err = s.Setup()
	return
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) SavePick(userId int, p *Pick) (err error) {
	query := `INSERT OR REPLACE INTO picks
		(user_id, game_id, pick_value)
		VALUES
		(?,       ?,       ?         )
	`
	_, err = s.db.Exec(query, userId, p.GameId, p.Value)
	return
}

func (s *Store) SavePickSet(u *User, p *PickSet) (err error) {
	return
}

func (s *Store) SaveLine(id string, spread, overUnder float64) (err error) {
	query := `UPDATE games
		SET
			game_spread        = ?,
			game_over_under    = ?,
			game_lines_updated = NOW()
		WHERE
			game_id            = ?
	`
	_, err = s.db.Exec(query, spread, overUnder, id)
	return
}

func (s *Store) SaveGame(g *Game) (err error) {
	id := g.Id
	if g.Id == "" {
		id = GameId(g.HomeId, g.AwayId, g.Start)
	}
	// Stage 1: Try to update existing game
	update := `UPDATE games SET game_score_home = ?, game_score_away = ?, game_quarter = ? WHERE game_id = ?`
	result, err := s.db.Exec(update, g.HomeScore, g.AwayScore, string(g.Quarter), id)
	if err != nil {
		return
	}
	if a, err := result.RowsAffected(); err != nil || a == 1 {
		return err
	}

	// Stage 2: Insert new game data
	insert := `INSERT OR REPLACE INTO games 
		(game_id, nfl_event_id, stadium_id, team_id_home, team_id_away, game_score_home, game_score_away, game_start, game_quarter, game_week, game_year)
		VALUES
		(?,       ?,            ?,          ?,            ?,            ?,               ?,               ?,          ?,            ?,         ?        )
	`
	_, err = s.db.Exec(insert, id, g.EventId, g.HomeId, g.HomeId, g.AwayId, g.HomeScore, g.AwayScore, g.Start, string(g.Quarter), g.Week, g.Year)
	return
}

func (s *Store) Setup() (err error) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS stadiums (
			stadium_id    TEXT PRIMARY KEY,
			stadium_name  TEXT NOT NULL,
			stadium_city  TEXT NOT NULL,
			stadium_state TEXT NOT NULL,
			stadium_turf  TEXT NOT NULL,
			stadium_roof  TEXT NOT NULL
		)`,
		`INSERT OR IGNORE INTO stadiums VALUES
			("KC",  "Arrowhead Stadium", "Kansas City", "Missouri", "Grass", "Open"),
			("DAL", "AT&T Stadium", "Arlington", "Texas", "Matrix RealGrass", "Retractable"),
			("CAR", "Bank of America Stadium", "Charlotte", "North Carolina", "Grass", "Open"),
			("SEA", "CenturyLink Field", "Seattle", "Washington", "FieldTurf", "Open"),
			("STL", "Edward Jones Dome", "St. Louis", "Missouri", "AstroTurf 3D", "Domed"),
			("JAC", "EverBank Field", "Jacksonville", "Florida", "Bermuda Grass", "Open"),
			("WAS", "FedExField", "Landover", "Maryland", "Bermuda Grass", "Open"),
			("CLE", "FirstEnergy Stadium", "Cleveland", "Ohio", "Kentucky Bluegrass", "Open"),
			("DET", "Ford Field", "Detroit", "Michigan", "FieldTurf", "Domed"),
			("ATL", "Georgia Dome", "Atlanta", "Georgia", "FieldTurf", "Domed"),
			("NE",  "Gillette Stadium", "Foxborough", "Massachusetts", "FieldTurf", "Open"),
			("PIT", "Heinz Field", "Pittsburgh", "Pennsylvania", "Grass", "Open"),
			("GB",  "Lambeau Field", "Green Bay", "Wisconsin", "Desso GrassMaster", "Open"),
			("SF",  "Levi's Stadium", "Santa Clara", "California", "Bermuda / Ryegrass ", "Open"),
			("PHI", "Lincoln Financial Field", "Philadelphia", "Pennsylvania", "Desso GrassMaster", "Open"),
			("TEN", "LP Field", "Nashville", "Tennessee", "Bermuda Grass", "Open"),
			("IND", "Lucas Oil Stadium", "Indianapolis", "Indiana", "FieldTurf", "Retractable"),
			("BAL", "M&T Bank Stadium", "Baltimore", "Maryland", "Sportexe Turf", "Open"),
			("NO",  "Mercedes-Benz Superdome", "New Orleans", "Louisiana", "Synthetic Turf", "Domed"),
			("NYG",  "MetLife Stadium", "East Rutherford", "New Jersey", "Synthetic Turf", "Open"),
			("NYJ",  "MetLife Stadium", "East Rutherford", "New Jersey", "Synthetic Turf", "Open"),
			("HOU", "NRG Stadium", "Houston", "Texas", "Bermuda Grass", "Retractable"),
			("OAK", "O.co Coliseum", "Oakland", "California", "Grass", "Open"),
			("CIN", "Paul Brown Stadium", "Cincinnati", "Ohio", "Synthetic Turf", "Open"),
			("SD",  "Qualcomm Stadium", "San Diego", "California", "Grass", "Open"),
			("BUF", "Ralph Wilson Stadium", "Orchard Park", "New York", "A-Turf Titan", "Open"),
			("TB",  "Raymond James Stadium", "Tampa", "Florida", "Bermuda Grass", "Open"),
			("CHI", "Soldier Field", "Chicago", "Illinois", "Grass", "Open"),
			("DEN", "Sports Authority Field at Mile High", "Denver", "Colorado", "Desso GrassMaster", "Open"),
			("MIA", "Sun Life Stadium", "Miami Gardens", "Florida", "Athletic Grass", "Open"),
			("MIN", "TCF Bank Stadium", "Minneapolis", "Minnesota", "FieldTurf", "Open"),
			("ARI", "University of Phoenix Stadium", "Glendale", "Arizona", "Bermuda Grass", "Retractable")
		`,
		`CREATE TABLE IF NOT EXISTS teams (
			team_id       TEXT PRIMARY KEY,
			team_city     TEXT NOT NULL,
			team_name     TEXT NOT NULL,
			team_league   TEXT NOT NULL,
			team_division TEXT NOT NULL
		)`,
		`INSERT OR IGNORE INTO teams VALUES
			("ARI", "Arizona",       "Cardinals",  "NFC", "W"),
			("ATL", "Atlanta",       "Falcons",    "NFC", "S"),
			("BAL", "Baltimore",     "Ravens",     "AFC", "N"),
			("BUF", "Buffalo",       "Bills",      "AFC", "E"),
			("CAR", "Carolina",      "Panthers",   "NFC", "S"),
			("CHI", "Chicago",       "Bears",      "NFC", "N"),
			("CIN", "Cincinnati",    "Bengals",    "AFC", "N"),
			("CLE", "Cleveland",     "Browns",     "AFC", "N"),
			("DAL", "Dallas",        "Cowboys",    "NFC", "E"),
			("DEN", "Denver",        "Broncos",    "AFC", "W"),
			("DET", "Detroit",       "Lions",      "NFC", "N"),
			("GB",  "Green Bay",     "Packers",    "NFC", "N"),
			("HOU", "Houston",       "Texans",     "AFC", "S"),
			("IND", "Indianapolis",  "Colts",      "AFC", "S"),
			("JAC", "Jacksonville",  "Jaguars",    "AFC", "S"),
			("KC",  "Kansas City",   "Chiefs",     "AFC", "W"),
			("MIA", "Miami",         "Dolphins",   "AFC", "E"),
			("MIN", "Minnesota",     "Vikings",    "NFC", "N"),
			("NO",  "New Orleans",   "Saints",     "NFC", "S"),
			("NE",  "New England",   "Patriots",   "AFC", "E"),
			("NYG", "New York",      "Giants",     "NFC", "E"),
			("NYJ", "New York",      "Jets",       "AFC", "E"),
			("OAK", "Oakland",       "Raiders",    "AFC", "W"),
			("PHI", "Philadelphia",  "Eagles",     "NFC", "E"),
			("PIT", "Pittsburgh",    "Steelers",   "AFC", "N"),
			("SD",  "San Diego",     "Chargers",   "AFC", "W"),
			("SEA", "Seattle",       "Seahawks",   "NFC", "W"),
			("SF",  "San Francisco", "49ers",      "NFC", "W"),
			("STL", "St. Louis",     "Rams",       "NFC", "W"),
			("TB" , "Tampa Bay",     "Buccaneers", "NFC", "S"),
			("TEN", "Tennessee",     "Titans",     "AFC", "S"),
			("WAS", "Washington",    "Redskins",   "NFC", "E")
		`,
		`CREATE TABLE IF NOT EXISTS games (
			game_id            TEXT PRIMARY KEY,
			nfl_event_id       INTEGER NOT NULL DEFAULT 0,
			stadium_id         TEXT NOT NULL REFERENCES stadiums (stadium_id),
			team_id_away       TEXT NOT NULL REFERENCES teams (team_id),
			team_id_home       TEXT NOT NULL REFERENCES teams (team_id),
			game_score_away    INTEGER NOT NULL DEFAULT 0,
			game_score_home    INTEGER NOT NULL DEFAULT 0,
			game_start         TIMESTAMP,
			game_quarter       TEXT NOT NULL DEFAULT "",
			game_week          INTEGER NOT NULL DEFAULT 0,
			game_year          INTEGER NOT NULL DEFAULT 0,
			game_spread        FLOAT(5,2) NOT NULL DEFAULT 0.0,
			game_over_under    FLOAT(5,2) NOT NULL DEFAULT 0.0,
			game_lines_updated DATETIME
		)`,
		`INSERT OR IGNORE INTO games (game_id, game_spread, game_over_under) VALUES
			("GBvSEA@20140904",   -5.5, 46.0),
			("NOvATL@20140907",      3, 51.5),
			("MINvSTL@20140907",    -4, 44.0),
			("CLEvPIT@20140907",  -6.5, 41.5),
			("JACvPHI@20140907", -10.5, 52.5),
			("OAKvNYJ@20140907",    -5, 40.0),
			("CINvBAL@20140907",  -1.5, 43.0),
			("BUFvCHI@20140907",    -7, 47.5),
			("WASvHOU@20140907",  -2.5, 45.5),
			("TENvKC@20140907",     -4, 43.5),
			("NEvMIA@20140907",      5, 47.5),
			("CARvTB@20140907",     -2, 39.0),
			("SFvDAL@20140907",      5, 51.0),
			("INDvDEN@20140907",  -7.5, 55.5),
			("NYGvDET@20140908",  -5.5, 47.0),
			("SDvARI@20140908",     -3, 45.0),
			("PITvBAL@20140911",  -2.5, 44.0),
			("MIAvBUF@20140914",     1, 43.0),
			("JACvWAS@20140914",  -5.5, 43.0),
			("DALvTEN@20140914",  -3.5, 49.0),
			("ARIvNYG@20140914",     2, 43.0),
			("NEvMIN@20140914",    3.5, 49.0),
			("NOvCLE@20140914",    6.5, 48.0),
			("ATLvCIN@20140914",    -5, 48.0),
			("DETvCAR@20140914",  -2.5, 44.0),
			("STLvTB@20140914",   -5.5, 37.0),
			("SEAvSD@20140914",      6, 45.0),
			("HOUvOAK@20140914",     3, 39.0),
			("NYJvGB@20140914",   -8.5, 46.0),
			("KCvDEN@20140914",  -13.5, 52.0),
			("CHIvSF@20140914",     -7, 48.0),
			("PHIvIND@20140915",    -3, 54.0),
			("TBvATL@20140918",   -6.5, 45.0),
			("SDvBUF@20140921",   -2.5, 44.0),
			("DALvSTL@20140921",     1, 45.0),
			("WASvPHI@20140921",  -6.5, 50.0),
			("HOUvNYG@20140921",   2.5, 41.0),
			("MINvNO@20140921",  -10.5, 50.0),
			("TENvCIN@20140921",    -7, 43.0),
			("BALvCLE@20140921",   1.5, 42.0),
			("GBvDET@20140921",     -2, 53.0),
			("INDvJAC@20140921",     7, 45.0),
			("OAKvNE@20140921",  -14.5, 47.0),
			("SFvARI@20140921",      3, 42.0),
			("DENvSEA@20140921",  -4.5, 49.0),
			("KCvMIA@20140921",   -4.5, 42.0),
			("PITvCAR@20140921",  -3.5, 42.0),
			("CHIvNYJ@20140922",  -2.5, 45.0),
			("NYGvWAS@20140925",    -4, 46.5),
			("GBvCHI@20140928",    1.5, 49.5),
			("BUFvHOU@20140928",    -3, 41.0),
			("TENvIND@20140928",  -7.5, 46.0),
			("CARvBAL@20140928",  -3.5, 41.0),
			("DETvNYJ@20140928",   1.5, 45.0),
			("TBvPIT@20140928",   -7.5, 44.5),
			("MIAvOAK@20140928",     4, 40.5),
			("JACvSD@20140928",  -13.5, 45.0),
			("ATLvMIN@20140928",     3, 47.0),
			("PHIvSF@20140928",     -5, 51.0),
			("NOvDAL@20140928",      3, 53.5),
			("NEvKC@20140929",     3.5, 46.0),
			("MINvGB@20141002",     -9, 48.0),
			("CHIvCAR@20141005",  -2.5, 46.0),
			("CLEvTEN@20141005",     1, 44.0),
			("STLvPHI@20141005",    -7, 48.0),
			("ATLvNYG@20141005",    -4, 50.5),
			("TBvNO@20141005",     -10, 48.5),
			("HOUvDAL@20141005",    -6, 46.5),
			("BUFvDET@20141005",    -7, 43.5),
			("BALvIND@20141005",  -3.5, 49.0),
			("PITvJAC@20141005",   6.5, 46.5),
			("ARIvDEN@20141005",    -7, 49.0),
			("KCvSF@20141005",    -6.5, 44.5),
			("NYJvSD@20141005",     -7, 44.0),
			("CINvNE@20141005",    1.5, 46.0),
			("SEAvWAS@20141006",   7.5, 46.0),
			("INDvHOU@20141009",     3, 47.0),
			("DENvNYJ@20141012",     9, 48.0),
			("PITvCLE@20141012",    -2, 47.0),
			("JACvTEN@20141012",    -6,  0.0),
			("GBvMIA@20141012",    3.5, 49.5),
			("DETvMIN@20141012",     2,  0.0),
			("CARvCIN@20141012",    -7, 43.5),
			("NEvBUF@20141012",      3,  0.0),
			("BALvTB@20141012",    3.5, 43.5),
			("SDvOAK@20141012",    7.5, 43.5),
			("CHIvATL@20141012",    -3, 54.0),
			("DALvSEA@20141012",    -8, 47.5),
			("WASvARI@20141012",  -3.5,  0.0),
			("NYGvPHI@20141012",  -2.5, 50.5),
			("SFvSTL@20141013",    3.5, 44.0),
			("NYJvNE@20141016",    -10, 45.0),
			("ATLvBAL@20141019",    -7, 49.5),
			("TENvWAS@20141019",  -5.5, 46.5),
			("SEAvSTL@20141019",     7, 43.5),
			("CLEvJAC@20141019",     6, 45.0),
			("CINvIND@20141019",  -3.5, 50.0),
			("MINvBUF@20141019",  -5.5, 42.5),
			("MIAvCHI@20141019",    -4, 49.0),
			("NOvDET@20141019",     -3, 48.5),
			("CARvGB@20141019",     -7, 49.0),
			("KCvSD@20141019",      -4, 45.5),
			("ARIvOAK@20141019",     4, 44.0),
			("NYGvDAL@20141019",  -6.5, 48.5),
			("SFvDEN@20141019",   -6.5, 51.5),
			("HOUvPIT@20141020",  -3.5, 45.0),
			("HOUvPIT@20141020",  -3.5, 45.0),
			("SDvDEN@20141023",     -8, 52.5),
			("DETvATL@20141026",     4, 47.0),
			("STLvKC@20141026",     -7, 44.0),
			("HOUvTEN@20141026",   2.5, 43.0),
			("MINvTB@20141026",     -3, 42.0),
			("SEAvCAR@20141026",   5.5, 45.5),
			("BALvCIN@20141026",     1, 46.0),
			("MIAvJAC@20141026",     6, 43.5),
			("CHIvNE@20141026",     -6, 51.0),
			("BUFvNYJ@20141026",    -3, 41.0),
			("PHIvARI@20141026",  -2.5, 48.5),
			("OAKvCLE@20141026",    -7, 43.5),
			("INDvPIT@20141026",   3.5, 49.0),
			("GBvNO@20141026",       0, 56.5),
			("WASvDAL@20141027",   -10, 50.5),
			("WASvDAL@20141027",   -10, 50.5),
			("NOvCAR@20141030",      3, 49.5),
			("TBvCLE@20141102",     -7, 44.0),
			("ARIvDAL@20141102",    -3, 45.0),
			("PHIvHOU@20141102",   2.5, 49.0),
			("NYJvKC@20141102",    -10, 42.0),
			("JACvCIN@20141102", -11.5, 43.5),
			("SDvMIA@20141102",     -1, 45.0),
			("WASvMIN@20141102",  -1.5, 45.5),
			("STLvSF@20141102",    -10, 44.0),
			("DENvNE@20141102",    3.5, 56.0),
			("OAKvSEA@20141102",   -15, 43.5),
			("BALvPIT@20141102",     1, 48.5),
			("INDvNYG@20141103",   3.5, 52.0),
			("CLEvCIN@20141106",  -6.5, 45.5),
			("KCvBUF@20141109",    2.5, 42.0),
			("MIAvDET@20141109",    -3, 44.0),
			("DALvJAC@20141109",     7, 45.0),
			("SFvNO@20141109",      -5, 49.5),
			("TENvBAL@20141109",   -10, 44.0),
			("PITvNYJ@20141109",   6.5, 45.5),
			("ATLvTB@20141109",      2, 46.0),
			("DENvOAK@20141109",  12.5, 50.0),
			("STLvARI@20141109",  -7.5, 43.5),
			("NYGvSEA@20141109",  -9.5, 45.5)
		`,
		`CREATE TABLE IF NOT EXISTS users (
			user_id   SERIAL PRIMARY KEY,
			user_name TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS picks (
			pick_id    SERIAL PRIMARY KEY,
			user_id    SERIAL NOT NULL REFERENCES users (user_id),
			game_id    TEXT NOT NULL REFERENCES odds (game_id),
			pick_value TEXT NOT NULL
		)`,
	}
	for _, query := range queries {
		if _, err = s.db.Exec(query); err != nil {
			return
		}
	}
	return
}

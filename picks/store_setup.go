package picks

import (
	"fmt"
	"strings"
)

func (s *Store) Setup() (err error) {
	queries := []struct {
		Table string
		Data  string
	}{
		{
			`CREATE TABLE IF NOT EXISTS config (
				year   INTEGER NOT NULL,
				week   INTEGER NOT NULL,
				season VARCHAR(8) NOT NULL
			)`,
			`INSERT INTO config VALUES (2014, 11, 'REG')`,
		},
		{
			`CREATE TABLE IF NOT EXISTS stadiums (
				stadium_id    VARCHAR(3) NOT NULL UNIQUE,
				stadium_name  TEXT NOT NULL,
				stadium_city  TEXT NOT NULL,
				stadium_state TEXT NOT NULL,
				stadium_turf  TEXT NOT NULL,
				stadium_roof  TEXT NOT NULL
			)`,
			`INSERT INTO stadiums VALUES
				('KC',  'Arrowhead Stadium', 'Kansas City', 'Missouri', 'Grass', 'Open'),
				('DAL', 'AT&T Stadium', 'Arlington', 'Texas', 'Matrix RealGrass', 'Retractable'),
				('CAR', 'Bank of America Stadium', 'Charlotte', 'North Carolina', 'Grass', 'Open'),
				('SEA', 'CenturyLink Field', 'Seattle', 'Washington', 'FieldTurf', 'Open'),
				('STL', 'Edward Jones Dome', 'St. Louis', 'Missouri', 'AstroTurf 3D', 'Domed'),
				('JAC', 'EverBank Field', 'Jacksonville', 'Florida', 'Bermuda Grass', 'Open'),
				('WAS', 'FedExField', 'Landover', 'Maryland', 'Bermuda Grass', 'Open'),
				('CLE', 'FirstEnergy Stadium', 'Cleveland', 'Ohio', 'Kentucky Bluegrass', 'Open'),
				('DET', 'Ford Field', 'Detroit', 'Michigan', 'FieldTurf', 'Domed'),
				('ATL', 'Georgia Dome', 'Atlanta', 'Georgia', 'FieldTurf', 'Domed'),
				('NE',  'Gillette Stadium', 'Foxborough', 'Massachusetts', 'FieldTurf', 'Open'),
				('PIT', 'Heinz Field', 'Pittsburgh', 'Pennsylvania', 'Grass', 'Open'),
				('GB',  'Lambeau Field', 'Green Bay', 'Wisconsin', 'Desso GrassMaster', 'Open'),
				('SF',  'Levi''s Stadium', 'Santa Clara', 'California', 'Bermuda / Ryegrass ', 'Open'),
				('PHI', 'Lincoln Financial Field', 'Philadelphia', 'Pennsylvania', 'Desso GrassMaster', 'Open'),
				('TEN', 'LP Field', 'Nashville', 'Tennessee', 'Bermuda Grass', 'Open'),
				('IND', 'Lucas Oil Stadium', 'Indianapolis', 'Indiana', 'FieldTurf', 'Retractable'),
				('BAL', 'M&T Bank Stadium', 'Baltimore', 'Maryland', 'Sportexe Turf', 'Open'),
				('NO',  'Mercedes-Benz Superdome', 'New Orleans', 'Louisiana', 'Synthetic Turf', 'Domed'),
				('NYG', 'MetLife Stadium', 'East Rutherford', 'New Jersey', 'Synthetic Turf', 'Open'),
				('NYJ', 'MetLife Stadium', 'East Rutherford', 'New Jersey', 'Synthetic Turf', 'Open'),
				('HOU', 'NRG Stadium', 'Houston', 'Texas', 'Bermuda Grass', 'Retractable'),
				('OAK', 'O.co Coliseum', 'Oakland', 'California', 'Grass', 'Open'),
				('CIN', 'Paul Brown Stadium', 'Cincinnati', 'Ohio', 'Synthetic Turf', 'Open'),
				('SD',  'Qualcomm Stadium', 'San Diego', 'California', 'Grass', 'Open'),
				('BUF', 'Ralph Wilson Stadium', 'Orchard Park', 'New York', 'A-Turf Titan', 'Open'),
				('TB',  'Raymond James Stadium', 'Tampa', 'Florida', 'Bermuda Grass', 'Open'),
				('CHI', 'Soldier Field', 'Chicago', 'Illinois', 'Grass', 'Open'),
				('DEN', 'Sports Authority Field at Mile High', 'Denver', 'Colorado', 'Desso GrassMaster', 'Open'),
				('MIA', 'Sun Life Stadium', 'Miami Gardens', 'Florida', 'Athletic Grass', 'Open'),
				('MIN', 'TCF Bank Stadium', 'Minneapolis', 'Minnesota', 'FieldTurf', 'Open'),
				('ARI', 'University of Phoenix Stadium', 'Glendale', 'Arizona', 'Bermuda Grass', 'Retractable')
			`,
		},
		{
			`CREATE TABLE IF NOT EXISTS teams (
				team_id       VARCHAR(3) NOT NULL UNIQUE,
				team_city     VARCHAR(16) NOT NULL,
				team_name     VARCHAR(16) NOT NULL,
				team_league   CHAR(3) NOT NULL,
				team_division CHAR(1) NOT NULL
			)`,
			`INSERT INTO teams VALUES
				('ARI', 'Arizona',       'Cardinals',  'NFC', 'W'),
				('ATL', 'Atlanta',       'Falcons',    'NFC', 'S'),
				('BAL', 'Baltimore',     'Ravens',     'AFC', 'N'),
				('BUF', 'Buffalo',       'Bills',      'AFC', 'E'),
				('CAR', 'Carolina',      'Panthers',   'NFC', 'S'),
				('CHI', 'Chicago',       'Bears',      'NFC', 'N'),
				('CIN', 'Cincinnati',    'Bengals',    'AFC', 'N'),
				('CLE', 'Cleveland',     'Browns',     'AFC', 'N'),
				('DAL', 'Dallas',        'Cowboys',    'NFC', 'E'),
				('DEN', 'Denver',        'Broncos',    'AFC', 'W'),
				('DET', 'Detroit',       'Lions',      'NFC', 'N'),
				('GB',  'Green Bay',     'Packers',    'NFC', 'N'),
				('HOU', 'Houston',       'Texans',     'AFC', 'S'),
				('IND', 'Indianapolis',  'Colts',      'AFC', 'S'),
				('JAC', 'Jacksonville',  'Jaguars',    'AFC', 'S'),
				('KC',  'Kansas City',   'Chiefs',     'AFC', 'W'),
				('MIA', 'Miami',         'Dolphins',   'AFC', 'E'),
				('MIN', 'Minnesota',     'Vikings',    'NFC', 'N'),
				('NO',  'New Orleans',   'Saints',     'NFC', 'S'),
				('NE',  'New England',   'Patriots',   'AFC', 'E'),
				('NYG', 'New York',      'Giants',     'NFC', 'E'),
				('NYJ', 'New York',      'Jets',       'AFC', 'E'),
				('OAK', 'Oakland',       'Raiders',    'AFC', 'W'),
				('PHI', 'Philadelphia',  'Eagles',     'NFC', 'E'),
				('PIT', 'Pittsburgh',    'Steelers',   'AFC', 'N'),
				('SD',  'San Diego',     'Chargers',   'AFC', 'W'),
				('SEA', 'Seattle',       'Seahawks',   'NFC', 'W'),
				('SF',  'San Francisco', '49ers',      'NFC', 'W'),
				('STL', 'St. Louis',     'Rams',       'NFC', 'W'),
				('TB' , 'Tampa Bay',     'Buccaneers', 'NFC', 'S'),
				('TEN', 'Tennessee',     'Titans',     'AFC', 'S'),
				('WAS', 'Washington',    'Redskins',   'NFC', 'E')
			`,
		},
		{
			`CREATE TABLE IF NOT EXISTS games (
				game_id            TEXT PRIMARY KEY,
				nfl_event_id       INTEGER NOT NULL DEFAULT 0,
				stadium_id         TEXT NOT NULL REFERENCES stadiums (stadium_id),
				team_id_away       VARCHAR(3) NOT NULL REFERENCES teams (team_id),
				team_id_home       VARCHAR(3) NOT NULL REFERENCES teams (team_id),
				game_score_away    INTEGER NOT NULL DEFAULT 0,
				game_score_home    INTEGER NOT NULL DEFAULT 0,
				game_score_updated TIMESTAMP WITH TIME ZONE,
				game_start         TIMESTAMP WITH TIME ZONE,
				game_quarter       VARCHAR(2) NOT NULL DEFAULT 'P',
				game_timeleft      INTERVAL,
				game_posession     VARCHAR(3),
				game_week          INTEGER NOT NULL DEFAULT 0,
				game_year          INTEGER NOT NULL DEFAULT 0,
				game_season        VARCHAR(8) NOT NULL DEFAULT 'REG',
				game_spread        NUMERIC(5,2) NOT NULL DEFAULT 0.00,
				game_over_under    NUMERIC(5,2) NOT NULL DEFAULT 0.00,
				game_line_updated  TIMESTAMP WITH TIME ZONE
			)`,
			``,
		},
		{
			`CREATE TABLE IF NOT EXISTS users (
				user_id        SERIAL PRIMARY KEY,
				user_name      TEXT NOT NULL,
				user_pin       CHAR(4) NOT NULL DEFAULT '0000',
				user_admin     BOOLEAN NOT NULL DEFAULT FALSE,
				user_last_seen TIMESTAMP WITH TIME ZONE
			)`,
			``,
		},
		{
			`CREATE TABLE IF NOT EXISTS picks (
				pick_id    SERIAL PRIMARY KEY,
				user_id    SERIAL NOT NULL REFERENCES users (user_id),
				game_id    TEXT NOT NULL REFERENCES games (game_id),
				pick_value TEXT NOT NULL,
				pick_added TIMESTAMP WITH TIME ZONE,
				UNIQUE(user_id, game_id)
			)`,
			``,
		},
	}

	if _, err = s.db.Exec("SET TIME ZONE 'America/New_York'"); err != nil {
		return
	}

	for _, query := range queries {
		create := query.Table[:strings.Index(query.Table, " (")]
		fmt.Printf("%s\n", create)
		if _, err = s.db.Exec(query.Table); err != nil {
			return
		}
		if query.Data == "" {
			fmt.Println("    No data")
			continue
		}
		// Check for existing data. If none found, insert.
		table := create[strings.LastIndex(create, " "):]
		var count int
		if err = s.db.QueryRow("SELECT COUNT(*) FROM " + table).Scan(&count); err != nil {
			return
		}
		if count > 0 {
			fmt.Println("    Already populated")
			continue
		}
		fmt.Println("    Inserting data")
		if _, err = s.db.Exec(query.Data); err != nil {
			return
		}
	}
	return
}

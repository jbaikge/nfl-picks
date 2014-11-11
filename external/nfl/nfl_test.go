package nfl

import (
	"encoding/xml"
	"testing"
	"time"
)

var timeLayout = "2006-01-02 15:04"

var samplePreGame = `<?xml version="1.0" encoding="UTF-8"?>
<ss>
	<gms gd="0" w="10" y="2014" t="R">
		<g eid="2014110600" gsis="56304" d="Thu" t="8:25" q="P" k="" h="CIN" hnn="bengals"    hs="" v="CLE" vnn="browns"      vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110901" gsis="56306" d="Sun" t="1:00" q="P" k="" h="BUF" hnn="bills"      hs="" v="KC"  vnn="chiefs"      vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110902" gsis="56307" d="Sun" t="1:00" q="P" k="" h="DET" hnn="lions"      hs="" v="MIA" vnn="dolphins"    vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110903" gsis="56308" d="Sun" t="1:00" q="P" k="" h="JAC" hnn="jaguars"    hs="" v="DAL" vnn="cowboys"     vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110904" gsis="56309" d="Sun" t="1:00" q="P" k="" h="NO"  hnn="saints"     hs="" v="SF"  vnn="49ers"       vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110900" gsis="56305" d="Sun" t="1:00" q="P" k="" h="BAL" hnn="ravens"     hs="" v="TEN" vnn="titans"      vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110905" gsis="56310" d="Sun" t="1:00" q="P" k="" h="NYJ" hnn="jets"       hs="" v="PIT" vnn="steelers"    vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110906" gsis="56311" d="Sun" t="1:00" q="P" k="" h="TB"  hnn="buccaneers" hs="" v="ATL" vnn="falcons"     vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110907" gsis="56312" d="Sun" t="4:05" q="P" k="" h="OAK" hnn="raiders"    hs="" v="DEN" vnn="broncos"     vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110908" gsis="56313" d="Sun" t="4:25" q="P" k="" h="ARI" hnn="cardinals"  hs="" v="STL" vnn="rams"        vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110909" gsis="56314" d="Sun" t="4:25" q="P" k="" h="SEA" hnn="seahawks"   hs="" v="NYG" vnn="giants"      vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014110910" gsis="56315" d="Sun" t="8:30" q="P" k="" h="GB"  hnn="packers"    hs="" v="CHI" vnn="bears"       vs="" p="" rz="" ga="" gt="REG"/>
		<g eid="2014111000" gsis="56316" d="Mon" t="8:30" q="P" k="" h="PHI" hnn="eagles"     hs="" v="CAR" vnn="panthers"    vs="" p="" rz="" ga="" gt="REG"/>
	</gms>
</ss>`

var samplePreGameExpect = []struct {
	Time string
	Home string
	Away string
}{
	{"2014-11-06 20:25", "CIN", "CLE"},
	{"2014-11-09 13:00", "BUF", "KC"},
	{"2014-11-09 13:00", "DET", "MIA"},
	{"2014-11-09 13:00", "JAC", "DAL"},
	{"2014-11-09 13:00", "NO", "SF"},
	{"2014-11-09 13:00", "BAL", "TEN"},
	{"2014-11-09 13:00", "NYJ", "PIT"},
	{"2014-11-09 13:00", "TB", "ATL"},
	{"2014-11-09 16:05", "OAK", "DEN"},
	{"2014-11-09 16:25", "ARI", "STL"},
	{"2014-11-09 16:25", "SEA", "NYG"},
	{"2014-11-09 20:30", "GB", "CHI"},
	{"2014-11-10 20:30", "PHI", "CAR"},
}

var sampleActiveGame = `<?xml version="1.0" encoding="UTF-8"?>
<ss>
	<gms w="10" y="2014" t="R" gd="1" bph="5">
		<g eid="2014110600" gsis="56304" d="Thu" t="8:25" q="2" k="01:03" h="CIN" hnn="bengals"    hs="3" v="CLE" vnn="browns"   vs="17" p="CIN" rz="0" ga="" gt="REG"/>
		<g eid="2014110900" gsis="56305" d="Sun" t="1:00" q="P"           h="BAL" hnn="ravens"     hs="0" v="TEN" vnn="titans"   vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110901" gsis="56306" d="Sun" t="1:00" q="P"           h="BUF" hnn="bills"      hs="0" v="KC"  vnn="chiefs"   vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110902" gsis="56307" d="Sun" t="1:00" q="P"           h="DET" hnn="lions"      hs="0" v="MIA" vnn="dolphins" vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110903" gsis="56308" d="Sun" t="1:00" q="P"           h="JAC" hnn="jaguars"    hs="0" v="DAL" vnn="cowboys"  vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110904" gsis="56309" d="Sun" t="1:00" q="P"           h="NO"  hnn="saints"     hs="0" v="SF"  vnn="49ers"    vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110905" gsis="56310" d="Sun" t="1:00" q="P"           h="NYJ" hnn="jets"       hs="0" v="PIT" vnn="steelers" vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110906" gsis="56311" d="Sun" t="1:00" q="P"           h="TB"  hnn="buccaneers" hs="0" v="ATL" vnn="falcons"  vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110907" gsis="56312" d="Sun" t="4:05" q="P"           h="OAK" hnn="raiders"    hs="0" v="DEN" vnn="broncos"  vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110908" gsis="56313" d="Sun" t="4:25" q="P"           h="ARI" hnn="cardinals"  hs="0" v="STL" vnn="rams"     vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110909" gsis="56314" d="Sun" t="4:25" q="P"           h="SEA" hnn="seahawks"   hs="0" v="NYG" vnn="giants"   vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014110910" gsis="56315" d="Sun" t="8:30" q="P"           h="GB"  hnn="packers"    hs="0" v="CHI" vnn="bears"    vs="0"          rz="0" ga="" gt="REG"/>
		<g eid="2014111000" gsis="56316" d="Mon" t="8:30" q="P"           h="PHI" hnn="eagles"     hs="0" v="CAR" vnn="panthers" vs="0"          rz="0" ga="" gt="REG"/>
	</gms>
</ss>`

var samplePostGame = `<?xml version="1.0" encoding="UTF-8"?>
<ss>
	<gms gd="0" w="8" y="2014" t="R">
		<g eid="2014102300" gsis="56276" d="Thu" t="8:25" q="F"  k="" h="DEN" hnn="broncos"    hs="35" v="SD"  vnn="chargers" vs="21" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102600" gsis="56277" d="Sun" t="9:30" q="F"  k="" h="ATL" hnn="falcons"    hs="21" v="DET" vnn="lions"    vs="22" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102607" gsis="56284" d="Sun" t="1:00" q="FO" k="" h="TB"  hnn="buccaneers" hs="13" v="MIN" vnn="vikings"  vs="19" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102608" gsis="56285" d="Sun" t="1:00" q="F"  k="" h="TEN" hnn="titans"     hs="16" v="HOU" vnn="texans"   vs="30" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102601" gsis="56278" d="Sun" t="1:00" q="F"  k="" h="CAR" hnn="panthers"   hs="9"  v="SEA" vnn="seahawks" vs="13" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102602" gsis="56279" d="Sun" t="1:00" q="F"  k="" h="CIN" hnn="bengals"    hs="27" v="BAL" vnn="ravens"   vs="24" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102603" gsis="56280" d="Sun" t="1:00" q="F"  k="" h="JAC" hnn="jaguars"    hs="13" v="MIA" vnn="dolphins" vs="27" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102604" gsis="56281" d="Sun" t="1:00" q="F"  k="" h="KC"  hnn="chiefs"     hs="34" v="STL" vnn="rams"     vs="7"  p="" rz="" ga="" gt="REG"/>
		<g eid="2014102605" gsis="56282" d="Sun" t="1:00" q="F"  k="" h="NE"  hnn="patriots"   hs="51" v="CHI" vnn="bears"    vs="23" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102606" gsis="56283" d="Sun" t="1:00" q="F"  k="" h="NYJ" hnn="jets"       hs="23" v="BUF" vnn="bills"    vs="43" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102609" gsis="56286" d="Sun" t="4:05" q="F"  k="" h="ARI" hnn="cardinals"  hs="24" v="PHI" vnn="eagles"   vs="20" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102610" gsis="56287" d="Sun" t="4:25" q="F"  k="" h="CLE" hnn="browns"     hs="23" v="OAK" vnn="raiders"  vs="13" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102611" gsis="56288" d="Sun" t="4:25" q="F"  k="" h="PIT" hnn="steelers"   hs="51" v="IND" vnn="colts"    vs="34" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102612" gsis="56289" d="Sun" t="8:30" q="F"  k="" h="NO"  hnn="saints"     hs="44" v="GB"  vnn="packers"  vs="23" p="" rz="" ga="" gt="REG"/>
		<g eid="2014102700" gsis="56290" d="Mon" t="8:30" q="FO" k="" h="DAL" hnn="cowboys"    hs="17" v="WAS" vnn="redskins" vs="20" p="" rz="" ga="" gt="REG"/>
	</gms>
</ss>`

func TestSamplePreGame(t *testing.T) {
	ss := new(ScoreStrip)
	if err := xml.Unmarshal([]byte(samplePreGame), ss); err != nil {
		t.Fatalf("xml.Unmarshal: %s", err)
	}
	if v := ss.GameSet.Year; v != 2014 {
		t.Fatalf("Year invalid: %d", v)
	}
	if v := len(ss.GameSet.Games); v != 13 {
		t.Fatalf("len(games): %d; expected 13", v)
	}
	for i, g := range ss.GameSet.Games {
		exp := samplePreGameExpect[i]
		if v := g.Start().Format(timeLayout); exp.Time != v {
			t.Errorf("[%d] Time Exp: %s; Got: %s", i, exp.Time, v)
		}
		if v := g.Home; exp.Home != v {
			t.Errorf("[%d] Home Exp: %s; Got: %s", i, exp.Home, v)
		}
		if v := g.Away; exp.Away != v {
			t.Errorf("[%d] Away Exp: %s; Got: %s", i, exp.Away, v)
		}
	}
}

func TestSampleActiveGame(t *testing.T) {
	ss := new(ScoreStrip)
	if err := xml.Unmarshal([]byte(sampleActiveGame), ss); err != nil {
		t.Fatalf("xml.Unmarshal: %s", err)
	}
	g := ss.GameSet.Games[0]
	if v := g.Quarter; v != "2" {
		t.Fatalf("Quarter Exp %s Got %s", "2", v)
	}
	if v := g.TimeLeft(); v != time.Minute+3*time.Second {
		t.Fatalf("TimeLeft Got %s", v)
	}
	if v := g.Posession; v != "CIN" {
		t.Fatalf("Posession Got %s", v)
	}
}

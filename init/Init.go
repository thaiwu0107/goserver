package init

import (
	Fmt "fmt"
	Ioutil "io/ioutil"

	"git.com.ggttoo44/src/config"
	Jsoniter "github.com/json-iterator/go"
)

// Rank7 結構
type Rank7 struct {
	Type7     string `json:"Type7"`
	Rank      int    `json:"Rank"`
	Type5Ch   string `json:"Type5Ch"`
	Type5En   string `json:"Type5En"`
	CardPoint []int  `json:"CardPoint"`
}

// Rank5 結構
type Rank5 struct {
	Rank      int    `json:"Rank"`
	Type5Ch   string `json:"Type5Ch"`
	Type5En   string `json:"Type5En"`
	CardPoint []int  `json:"CardPoint"`
}

// Card 結構
type Card struct {
	NameCh     string `json:"NameCh"`
	NameCh2    string `json:"NameCh2"`
	NameEn     string `json:"NameEn"`
	NameCode   string `json:"NameCode"`
	Suit       string `json:"Suit"`
	SuitCh     string `json:"SuitCh"`
	SuitEn     string `json:"SuitEn"`
	SuitCode   string `json:"SuitCode"`
	Point      string `json:"Point"`
	PointEn    string `json:"PointEn"`
	PointCode  string `json:"PointCode"`
	PointCode2 string `json:"PointCode2"`
}

// RankTable7CF 已經讀取完的轉成MAP格式
var RankTable7CF map[string]*Rank7

// RankTable7CNF 已經讀取完的轉成MAP格式
var RankTable7CNF map[string]*Rank7

// RankTable5CF 已經讀取完的轉成MAP格式
var RankTable5CF map[string]*Rank5

// RankTable5CNF 已經讀取完的轉成MAP格式
var RankTable5CNF map[string]*Rank5

// CardNameTable 已經讀取完的轉成MAP格式
var CardNameTable map[string]*Card

// Start 要初始化的方法都放進來
func Start() {
	config.FristInit()
	InitJSONToMap()
}

// InitJSONToMap 負責初始化JSON轉成MAP格式
func InitJSONToMap() {
	jsonFileCardNameTable, err1 := Ioutil.ReadFile("./json/CardNameTable.json")
	if err1 != nil {
		Fmt.Println(err1)
	}
	jsonFileRankTable7CF, err2 := Ioutil.ReadFile("./json/RankTable7CF.json")
	if err2 != nil {
		Fmt.Println(err2)
	}
	jsonFileRankTable7CNF, err3 := Ioutil.ReadFile("./json/RankTable7CNF.json")
	if err3 != nil {
		Fmt.Println(err3)
	}
	jsonFileRankTable5CF, err4 := Ioutil.ReadFile("./json/RankTable5CF.json")
	if err4 != nil {
		Fmt.Println(err4)
	}
	jsonFileRankTable5CNF, err5 := Ioutil.ReadFile("./json/RankTable5CNF.json")
	if err5 != nil {
		Fmt.Println(err5)
	}
	Fmt.Println("Successfully Ioutil.ReadFile")
	var json = Jsoniter.ConfigCompatibleWithStandardLibrary
	error1 := json.Unmarshal(jsonFileCardNameTable, &CardNameTable)
	if error1 != nil {
		Fmt.Println("Unmarshal CardNameTable failed, ", error1)
		return
	}
	error2 := json.Unmarshal(jsonFileRankTable7CF, &RankTable7CF)
	if error1 != nil {
		Fmt.Println("Unmarshal RankTable7CF failed, ", error2)
		return
	}
	error3 := json.Unmarshal(jsonFileRankTable7CNF, &RankTable7CNF)
	if error3 != nil {
		Fmt.Println("Unmarshal RankTable7CNF failed, ", error3)
		return
	}
	error4 := json.Unmarshal(jsonFileRankTable5CF, &RankTable5CF)
	if error4 != nil {
		Fmt.Println("Unmarshal RankTable5CF failed, ", error4)
		return
	}
	error5 := json.Unmarshal(jsonFileRankTable5CNF, &RankTable5CNF)
	if error5 != nil {
		Fmt.Println("Unmarshal RankTable5CNF failed, ", error5)
		return
	}
	Fmt.Println("Successfully All Json To Map")
}

package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"weather4Go/seniverse"

	_ "github.com/go-sql-driver/mysql"
	"github.com/json-iterator/go"
)

const (
	baseURL   string = "https://api.seniverse.com/v3/weather/"
	nowAPIURL string = "now.json?"
	key       string = "di7ey9jsktwjlwse"
	location  string = "beijing"
	lang      string = "zh-Hans"
	unit      string = "c"
)

func dbConn() (db *sql.DB) {

	const dbDriver = "mysql"
	const dbUser = "root"
	const dbPass = "12345"
	const dbURL = "39.106.122.56"
	const dbPort = "3306"
	const dbName = "weather"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbURL, dbPort, dbName)

	db, err := sql.Open(dbDriver, dataSource)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {

	nowRequestURL := fmt.Sprintf("%s%skey=%s&location=%s&language=%s&unit=%s",
		baseURL,
		nowAPIURL,
		key,
		location,
		lang,
		unit)

	response, err := http.Get(nowRequestURL)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(response.Body)

	var results = seniverse.Results{}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	//fmt.Println(string(body))

	json.UnmarshalFromString(string(body), &results)

	resultLocation := results.Results[0].Location
	resultNow := results.Results[0].Now

	locationId := resultLocation.ID
	locationName := resultLocation.Name
	locationCountry := resultLocation.Country
	locationPath := resultLocation.Path
	locationTimezone := resultLocation.Timezone
	locationTimezoneOffset := resultLocation.TimezoneOffset
	nowTest := resultNow.Text
	nowCode := resultNow.Code
	nowTempature := resultNow.Temperature
	lastUpdate := results.Results[0].LastUpdate

	db := dbConn()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO now_result(location_id, location_name, location_country, location_path, location_timezone, location_timezone_offset, now_text, now_code, now_tempature, last_update) VALUES (?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(locationId, locationName, locationCountry, locationPath, locationTimezone, locationTimezoneOffset, nowTest, nowCode, nowTempature, lastUpdate)
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}

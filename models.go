package main

import (
	"database/sql"
	"fmt"
	"time"
	"log"

	_ "github.com/lib/pq"
)

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS cities
(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	latitude FLOAT(8) NOT NULL,
	longitude FLOAT(8) NOT NULL
)
`
const createtempquery = `
CREATE TABLE IF NOT EXISTS temprature
(
	city_id SERIAL PRIMARY KEY,
	max FLOAT(8) NOT NULL,
	min FLOAT(8) NOT NULL,
	time_stamp TIMESTAMPTZ NOT NULL
)`

var db *sql.DB


type city struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type temprature struct {
	City_ID int `json:"city_id"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Timestamp time.Time `json:"timestamp"`
}

func InitDB() *sql.DB{

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

		var err error

		db, err = sql.Open("postgres",psqlInfo)

		if err != nil {
			log.Panic(err)
		}

		if err = db.Ping(); err != nil {
			log.Panic(err)
		}
		if _, err := db.Exec(tableCreationQuery); err != nil {
			log.Fatal(err)
		}

		if _, err := db.Exec(createtempquery); err != nil {
			log.Fatal(err)
		}
		return db

}

func CheckName(c city) int {

	db := InitDB()

	query := fmt.Sprintf(`
	SELECT name FROM cities WHERE name = '%s'
	`,c.Name)

	// fmt.Println(c.Name)

	var name string

	err := db.QueryRow(query).Scan(&name)

	if err == sql.ErrNoRows {
		return 1
	}
	
	if err != nil {
		
		fmt.Println("Error occured while executing query")

		return 0
	}


	return 0


}

func InsertCity(c city) (int, error) {


	db := InitDB()

	statement := fmt.Sprintf("INSERT INTO cities(name, latitude, longitude) VALUES('%s', %f, %f) RETURNING id", c.Name, c.Latitude, c.Longitude)

	var id int
	
	err := db.QueryRow(statement).Scan(&id)

	fmt.Println(id)
	
    if err != nil {
		fmt.Println("error creating city")
        return 0, err
	}
	
	
    return id, nil

}

func UpdateMyCity(c city) error{
	db := InitDB()

	statement := fmt.Sprintf(`
	UPDATE cities SET
	name='%s',latitude=%f,longitude=%f
	WHERE id=%d
	`,c.Name,c.Latitude,c.Longitude,c.ID)

	_, err := db.Exec(statement)

	return err
}


func DeleteMyCity(c city) error {
	db := InitDB()

	statement := fmt.Sprintf("DELETE FROM cities WHERE id=%d", c.ID)
    _, err := db.Exec(statement)
    return err
	
}

func GetMyCity(c city) (city, error){
	db := InitDB()

	statement := fmt.Sprintf(`
	SELECT name,latitude,longitude
	FROM cities
	WHERE id = %d
	`,c.ID)

	rows, err := db.Query(statement)

	if err != nil {
		fmt.Println("error occured")
		return c,err
	}

	defer rows.Close()

	var c1 city

	c1.ID = c.ID

	for rows.Next(){
		err := rows.Scan(&c1.Name,&c1.Latitude,c1.Longitude)
		if err != nil {
			return c1, nil
		}
	}

	return c1, nil
}

func AddTemp(t temprature) (temprature,error){

	db := InitDB()
	Time_Stamp := time.Now()
	statement := fmt.Sprintf(`
	INSERT INTO temprature(city_id,max,min,time_stamp)
	VALUES (%d,%f,%f,%T)
	`,t.City_ID,t.Max,t.Min,time.Now())

	fmt.Println(statement)
	_, err := db.Exec(statement)

	if err != nil {
		return t, err
	}

	t.Timestamp = Time_Stamp
	
	return t, nil

}



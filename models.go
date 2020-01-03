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
CREATE TABLE IF NOT EXISTS temperature
(
	city_id INT NOT NULL,
	max FLOAT(8) NOT NULL,
	min FLOAT(8) NOT NULL,
	time_stamp timestamptz NOT NULL DEFAULT now()
)`

const createwebhook = `
CREATE TABLE IF NOT EXISTS webhooks
(
	id SERIAL PRIMARY KEY,
	city_id INT NOT NULL,
	callback_url VARCHAR NOT NULL
)`

var db *sql.DB


type city struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type temperature struct {
	City_ID int `json:"city_id"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
	Timestamp time.Time `json:"timestamp"`
}

type forecast struct {
	City_ID int `json:"city_id"`
	Max int `json:"max"`
	Min int `json:"min"`
	Sample int `json:"sample"`
}


type webhook struct {
	ID int `json:"id"`
	City_ID int `json:"city_id"`
	CallbackUrl string `json:"callback_url"`
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
		if _, err := db.Exec(createwebhook); err != nil {
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

func AddTemp(t temperature) (temperature,error){

	db := InitDB()

	var Time_Stamp time.Time

	statement := fmt.Sprintf(`
	INSERT INTO temperature(city_id,max,min)
	VALUES (%d,%f,%f) RETURNING time_stamp
	`,t.City_ID,t.Max,t.Min)


	err := db.QueryRow(statement).Scan(&Time_Stamp)

	if err != nil {
		return t, err
	}

	t.Timestamp = Time_Stamp
	
	return t, nil

}

func ForecastData(t temperature) (forecast,error){
	db := InitDB()

	statement := fmt.Sprintf(`
	select max,min from temperature where city_id = %d
	`,t.City_ID)

	rows, err := db.Query(statement)

	var f forecast

	if err != nil {
		fmt.Println("error occured")
		return f,err
	}

	defer rows.Close()

	var maxlist []int
	var minlist []int

	for rows.Next(){
		var max int
		var min int
		err := rows.Scan(&max,&min)

		if err != nil {
			fmt.Println(err.Error())
		}

		maxlist = append(maxlist,max)
		minlist = append(minlist,min)
	}

	

	f.Sample = len(maxlist)

	for _,i := range maxlist {
		f.Max += i
	}

	for _,i := range minlist {
		f.Min += i
	}

	f.Max = f.Max/f.Sample
	f.Min = f.Min/f.Sample


	return f,nil

}

func AddWebhook(w webhook)(webhook, error){
	db := InitDB()

	statement := fmt.Sprintf(`
	INSERT INTO webhooks(city_id,callback_url)
	VALUES (%d,'%s')
	RETURNING id
	`,w.City_ID,w.CallbackUrl)

	
	err := db.QueryRow(statement).Scan(&w.ID)

	if err != nil {

		return w, err
	}

	// fmt.Println(w)

	return w, nil


}

func GetWebhook(w webhook) (webhook, error){
	db := InitDB()

	statement := fmt.Sprintf(`
	SELECT city_id,callback_url
	FROM webhooks
	WHERE id = %d
	`,w.ID)

	rows, err := db.Query(statement)

	if err != nil {
		fmt.Println("error occured")
		return w,err
	}

	defer rows.Close()

	var w1 webhook

	w1.ID = w.ID

	for rows.Next(){
		err := rows.Scan(&w1.City_ID,&w1.CallbackUrl)
		if err != nil {
			return w1, nil
		}
	}

	return w1, nil
}

func DeleteWebhook(w webhook) error {
	db := InitDB()

	statement := fmt.Sprintf("DELETE FROM webhooks WHERE id=%d", w.ID)
    _, err := db.Exec(statement)
    return err
}

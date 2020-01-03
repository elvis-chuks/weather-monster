# weather-monster

## Setting Up the Project

Make sure you have golang and postgressql installed

Run 
``` bash
go get github.com/gorilla/mux
go get github.com/lib/pq
```

## Postgressql connection details

```go
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "weathermonster"
	)
```

## Create database with 
``` sql
CREATE DATABASE weathermonster
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;
```
Create tables

``` sql
CREATE TABLE IF NOT EXISTS cities
(
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	latitude VARCHAR(50) NOT NULL,
	longitude VARCHAR(50) NOT NULL
)

CREATE TABLE IF NOT EXISTS temperature
(
	city_id INT NOT NULL,
	max FLOAT(8) NOT NULL,
	min FLOAT(8) NOT NULL,
	time_stamp timestamptz NOT NULL DEFAULT now()
)

CREATE TABLE IF NOT EXISTS webhooks
(
	id SERIAL PRIMARY KEY,
	city_id INT NOT NULL,
	callback_url VARCHAR NOT NULL
)

```

## Running Tests


Run 
```bash
	go test -v
```

## Running the App

To run this app simple run 

```bash  
go build
weather.exe
```



# Comments

When testing endpoints that need the mux.vars method you might experience getting null or 0 values for you id or city_id
Read more here http://mrgossett.com/post/mux-vars-problem/

To bypass this issue
I wrote minor tests in a python file
So if you have a python environment run the tests in that file.




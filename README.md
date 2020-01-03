# weather-monster

## Getting Dependencies

Make sure you have golang and postgressql installed

Run 
``` bash
go get github.com/gorilla/mux
go get github.com/lib/pq
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





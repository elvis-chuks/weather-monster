package main

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "weathermonster"
	)



func main(){
	fmt.Println("Hello world")
	InitDB()
	initialize()
}
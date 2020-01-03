package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "strconv"
	"github.com/gorilla/mux"
)

func initialize(){
	fmt.Println("Initializing application")
	router := mux.NewRouter()

	router.HandleFunc("/cities",CreateCity).Methods("POST")
	router.HandleFunc("/cities/{id:[0-9]+}", UpdateCity).Methods("PATCH")
	router.HandleFunc("/cities/{id:[0-9]+}", DeleteCity).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",router))

}

func CreateCity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var c city

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&c); err != nil {

		respondWithError(w,http.StatusBadRequest,"Invalid Request Body")

		return

	}

	defer r.Body.Close()

	find := CheckName(c)

	if find == 1 {

		fmt.Println("City Name wasnt found go ahead",find)

		id, err := InsertCity(c)

		if err != nil {

			http.Error(w,http.StatusText(500),500)
			
			return
		}
		c.ID = id
		response, _ := json.Marshal(c)
		w.WriteHeader(200)
		w.Write(response)
	}
	if find == 0 {
		fmt.Println("City Name was found dont post to the database",find)

		respondWithJSON(w, 200, map[string]string{"error": "City already exists"})
	}

	

}

func UpdateCity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	vars := mux.Vars(r)
	
	id, err := strconv.Atoi(vars["id"])

	fmt.Println("this is id",id)

	if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid City ID")
        return
    }

	var c city

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&c); err != nil {

		respondWithError(w,http.StatusBadRequest,"Invalid Request Body")

		return

	}

	defer r.Body.Close()

	c.ID = id

	err = UpdateMyCity(c)

	if err != nil{
		http.Error(w,http.StatusText(500),500)
	}

	response, _ := json.Marshal(c)
	w.WriteHeader(200)
	w.Write(response)

}

func DeleteCity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	
	id, err := strconv.Atoi(vars["id"])

	fmt.Println("You are deleting this id:",id)

	if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid City ID")
        return
    }

	var c city

	c.ID = id

	resp, err := GetMyCity(c)

	err = DeleteMyCity(c)

	if err != nil{
		http.Error(w,http.StatusText(500),500)
	}

	

	fmt.Println(resp)

	if err != nil {
		fmt.Println("An error occured")
	}



	response, _ := json.Marshal(resp)
	w.WriteHeader(200)
	w.Write(response)

}


func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
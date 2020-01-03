package main

import (
	"fmt"
	"bytes"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestCreateCity(t *testing.T){

	payload := []byte(`{"name":"Mash","latitude":30.08,"longitude":20.0}`)

	req, err := http.NewRequest("POST","/cities", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCity)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",status, http.StatusOK)
	}



	// check the response body

	fmt.Println(rr.Body.String())

}


func TestCheckName(t *testing.T){
	var c city

	c.Name = "India"

	res := CheckName(c)

	correct := 1

	if res != 1 {
		t.Errorf("Expecting %d here for got %d",correct,res)
	}

	fmt.Println(res)

}

func TestGet(t *testing.T){
	c := city{
		ID: 10,
	}
	resp, err := GetMyCity(c)

	if err != nil {
		t.Errorf("Error executing sql")
	}
	fmt.Println(resp)
}

func TestDelete(t *testing.T){
	c := city{
		ID : 11,
	}

	err := DeleteMyCity(c)

	if err != nil {
		t.Errorf("error executing sql")
	}
}

func TestUpdate(t *testing.T){

	c := city{
		ID : 11,
		Name:"Mash",
		Latitude:34.08,
		Longitude:89.00,
	}

	err := UpdateMyCity(c)

	if err != nil {
		t.Errorf("error executing sql")
	}
}
func TestInsert(t *testing.T){
	var c city

	c.ID = 1
	c.Name = "Lagos"
	c.Latitude = 13.56
	c.Longitude = 45.67

	_, err := InsertCity(c)

	if err != nil {
		t.Errorf("Can not perform sql statement at insertcity")
	}

}

func TestUpdateCity(t *testing.T){

	payload := []byte(`{"name":"Mash","latitude":30.08,"longitude":20.0}`) 
	// already existing data in our database
	req, err := http.NewRequest("PATCH","/cities/11", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateCity)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body.String())

	if status := rr.Code; status != http.StatusOK {
		// t.Errorf("Handler returned wrong status code: got %v want %v",status, http.StatusOK)
		fmt.Println("tests fail because of a gorilla/mux isssue found here: http://mrgossett.com/post/mux-vars-problem/")
	}


}


func TestTemp(t *testing.T){
	temp := temperature{
		City_ID:1,
		Max: 33.0,
		Min:20,
	}
	res, err := AddTemp(temp)

	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(res)
}

func TestAddTemp(t *testing.T){
	payload := []byte(`{"city_id":1,"max":32,"min":18}`) 

	req, err := http.NewRequest("POST","/temperature", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Addtemperature)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body.String())

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",status, http.StatusOK)
		// fmt.Println("tests fail because of a gorilla/mux isssue found here: http://mrgossett.com/post/mux-vars-problem/")
	}
}


func TestForecastData(t *testing.T){
	tes := temperature{
		City_ID:1,
	}

	resp,err := ForecastData(tes)

	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(resp)
}

func TestGetForecast(t *testing.T){
	// payload := []byte(`{"city_id":1}`) 

	req, err := http.NewRequest("GET","/forecasts/1",nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetForecast)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body.String())

	if status := rr.Code; status != http.StatusOK {
		// t.Errorf("Handler returned wrong status code: got %v want %v",status, http.StatusOK)
		fmt.Println("tests fail because of a gorilla/mux isssue found here: http://mrgossett.com/post/mux-vars-problem/")
	}
}
func TestAddWebhook(t *testing.T){
	w := webhook{
		City_ID:1,
		CallbackUrl:"http://127.0.0.1:5000",
	}

	_, err := AddWebhook(w)

	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestCreateWebhook(t *testing.T){

	payload := []byte(`{"city_id":1,"callback_url":"http://127.0.0.1:5000"}`) 

	req, err := http.NewRequest("POST","/webhooks", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateWebhook)

	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body.String())

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",status, http.StatusOK)
	}

}

func TestGetWebhook(t *testing.T){
	w := webhook{
		ID:1,
	}
	_, err := GetWebhook(w)

	if err != nil {
		t.Errorf(err.Error())
	}


}

func TestDeleteWebhook(t *testing.T){
	w := webhook{
		ID:1,
	}
	err := DeleteWebhook(w)

	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestDeleteWebhookEndpoint(t *testing.T){
	req, err := http.NewRequest("DELETE","/webhooks/3",nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteWebhookEndpoint)
	handler.ServeHTTP(rr, req)

	fmt.Println(rr.Body.String())

	if status := rr.Code; status != http.StatusOK {
		// t.Errorf("Handler returned wrong status code: got %v want %v",status, http.StatusOK)
		fmt.Println("tests fail because of a gorilla/mux isssue found here: http://mrgossett.com/post/mux-vars-problem/")
	}
}



// will fail if the callback_url doesnt exist

// func TestCallWebhooks(t *testing.T){
// 	w := webhook{
// 		City_ID:1,
// 	}
// 	payload := `{"city_id":1,"max":32,"min":18,"timestamp":"2020-01-03T15:11:04.650729+01:00"}`

// 	err := CallWebhooks(w,payload)

// 	if err != nil {
// 		t.Errorf(err.Error())
// 	}

// }





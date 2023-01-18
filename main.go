// main.go
// for reference look this url, gives you the fine refrence about how to use http request https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html

/*
* go lang rest api.
* no static data view.
* gives yout fine undrstanding about http and its sub-resources
* For database -- MYSQL
*
 */
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
// type Article struct {
// 	Id      string `json:"Id"`
// 	Title   string `json:"Title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

type Payload struct {
	Id        string `json:"id"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Status    string `json:"status"`
}

// type user struct {
// 	Id        string
// 	Latitude  string
// 	Longitude string
// 	Status    string
// }

// var userDetials user

// var Articles []Article
var DB *sql.DB

// var allRecords []Payload

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAlldata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAlldata")
	// query := fmt.Sprintf("SELECT id, latitude, longitude, status FROM user_address where status = 1")
	status := "1"
	fmt.Printf("SELECT id, latitude, longitude, status FROM users_address WHERE status = '1'")
	results, err := DB.Query("SELECT id, latitude, longitude, status FROM users_address WHERE status = ?", status)

	// fmt.Println(rows)
	if err != nil {
		panic(err.Error())
	}

	// var payload Payload

	// var emailData map[string]interface{}

	// returnPayload := make(map[string]interface{})

	// var finalPayload []string
	// finalPayload := [...]string{}
	//new code

	// for results.Next() {
	// 	// fmt.Println("here")
	// 	// fmt.Println(result)
	// 	err = results.Scan(&payload.Id, &payload.Latitude, &payload.Longitude, &payload.Status)
	// 	if err != nil {
	// 		fmt.Printf("Error mapping DB_ROW data to local varialble %v", err)
	// 	}

	// 	returnPayload["id"] = payload.Id
	// 	returnPayload["latitude"] = payload.Latitude
	// 	returnPayload["longitude"] = payload.Longitude
	// 	returnPayload["status"] = payload.Status
	// // userDetials = append(userDetials.Id, payload.Id)
	// fmt.Println("return payload", returnPayload)
	// finalPayload = append(finalPayload, returnPayload)

	// fmt.Println("final payload", finalPayload)
	// jsonData, _ := json.Marshal(returnPayload)
	// finalPayload = append(finalPayload, returnPayload)

	// fmt.Println("jsonData:", string(jsonData))
	// fmt.Println("returnPayload", returnPayload)
	// finalPayload = append(finalPayload, payload.Id, payload.Latitude, payload.Longitude, payload.Status)
	// finalPayload = append(finalPayload, returnPayload)

	// using interface mapping for all result

	// finalPayload = append(finalPayload, returnPayload)

	// emailData = map[string]interface{}{
	// 	"id":        payload.Id,
	// 	"latitude":  payload.Latitude,
	// 	"longitude": payload.Longitude,
	// 	"status":    payload.Status,
	// }

	w.Header().Set("Content-Type", "application/json")

	var userData []Payload

	for results.Next() {
		var payload Payload
		// fmt.Println("here")
		// fmt.Println(result)
		err = results.Scan(&payload.Id, &payload.Latitude, &payload.Longitude, &payload.Status)
		if err != nil {
			fmt.Printf("Error mapping DB_ROW data to local varialble %v", err)
		}

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// fmt.Println(returnPayload)
		// jsonFinalData, _ := json.Marshal(returnPayload)
		// w.Write(jsonFinalData)

		userData = append(userData, payload)

	}

	json.NewEncoder(w).Encode(userData)
	// fmt.Println("emailData", emailData)
	// fmt.Println(finalPayload...)
	// fmt.Println("==return", payload)
	// rawData := []byte(returnPayload)
	// fmt.Println("finalPayload", string(finalPayload))
	// fmt.Println("data of struct", payload)
	// fmt.Println("--final payload--", finalPayload)
	// jsonFinalData, _ := json.Marshal(finalPayload)

	// json.NewEncoder(w).Encode(string(jsonFinalData))

	// jsonData, _ := json.Marshal(finalPayload)
	// payloadUnmarshalerr := json.Unmarshal([]byte(jsonData), &finalPayload)
	// fmt.Println("--json data-- ", jsonData)
	// fmt.Println("----------------------------")
	// fmt.Println("unmarshal data", string(jsonData), PayloadUnmarshal)
	// w.Header().Set("Content-Type", "application/json")
	// if payloadUnmarshalerr != nil {
	// 	fmt.Println(payloadUnmarshalerr)
	// }
	// jsonData = []byte(fmt.Sprintf(string(jsonData)))
	// w.Write(finalPayload)
	// fmt.Println("==returnPayload", finalPayload)
	// for k, value := range returnPayload {
	// 	fmt.Println("key:", k)
	// 	finalPayload = append(finalPayload, value)
	// }

	// fmt.Println("--finalpayload--", finalPayload)
	// fmt.Println("---array ---- ", a)
	// jsonData, _ := json.Marshal(payload)
	// fmt.Println("jsondata:", string(jsonData))
	// fmt.Println("==rturnpayload===", returnPayload)

	// fmt.Println("==returnpayload==", returnPayload)
	// jsonData, _ := json.Marshal(returnPayload)
	// bytes := []byte(jsonData)
	// fmt.Println("===json data===", bytes)
	// fmt.Println(jsonData)
	// fmt.Printf("%s\n", payload.Latitude)
	// json.NewEncoder(w).Encode(payload.Latitude)
	// defer DB.Close()
	// json.NewEncoder(w).Encode(Articles)

	// w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	fmt.Println(payload)
	// 	jsonFinalData, _ := json.Marshal(payload)
	// 	w.Write(jsonFinalData)
}

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: returnSingleArticle")
// 	// vars := mux.Vars(r)
// 	// // key := vars["id"]

// 	// for _, article := range Articles {
// 	// 	if article.Id == key {
// 	// 		json.NewEncoder(w).Encode(article)
// 	// 	}
// 	// }
// }

func createNewData(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	// var data []Payload
	// payload := []Payload{}

	w.Header().Set("Content-Type", "application/json")
	message := make(map[string]string)
	var payload Payload
	// var status bool
	reqBody, _ := io.ReadAll(r.Body)
	// fmt.Println(json.Marshal(reqBody))
	// fmt.Print(reflect.TypeOf(reqBody))
	// json.Unmarshal(reqBody, &payload)
	// finalpayload, err := json.Marshal(reqBody)

	err := json.Unmarshal([]byte(string(reqBody)), &payload)

	if err != nil {
		fmt.Println(err.Error())
	}
	count := 0
	fmt.Print("select count(latitude) where latitude = ? and longitude = ?", payload.Latitude, payload.Longitude)
	rows, err := DB.Query("select count(latitude) from users_address where latitude = ? and longitude = ?", payload.Latitude, payload.Longitude)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		count++
	}
	// fmt.Println("total count ", count)
	if count == 0 {
		updateDetail, err := DB.Prepare("INSERT INTO users_address(latitude, longitude, status) VALUES(?,?,?)")
		// as we unmarshal the given response body in the struct mentioned above we can can directly access the value through its
		// fmt.Print(reflect.TypeOf(payload))
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		}

		updateDetail.Exec(payload.Latitude, payload.Longitude, payload.Status)

		log.Println("INSERT: latitude: " + payload.Latitude + " | longitude : " + payload.Longitude + " | status : " + payload.Status)

		message["status"] = "true"
		message["response"] = "Data inserted successfully !"

		json.NewEncoder(w).Encode(message)
	} else {
		message["status"] = "false"
		message["response"] = "Data already exists in the database."

		json.NewEncoder(w).Encode(message)
	}
	//
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// json.Unmarshal(finalpayload, &payload)

	// data = append(data, payload)

	// fmt.Print(data)
	// fmt.Print(string(reqBody))
	// var article Article
	// json.Unmarshal(reqBody, &article)
	// // update our global Articles array to include
	// // our new Article
	// Articles = append(Articles, article)

	// json.NewEncoder(w).Encode(article)
}

func deleteUserDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUserDetails")

	// w.Header().Set("Content-Type", "application/json")

	// var payload Payload
	message := make(map[string]string)
	// reqBody, _ := io.ReadAll(r.Body)
	// err := json.Unmarshal([]byte(string(reqBody)), &payload)
	// id := r.URL.Query().Get("id")
	//get last segement of the url string using path.base and get path using r.url.path
	id := path.Base(r.URL.Path)

	// emp := r.URL.Query().Get("id")
	// id := payload.Id
	delUserData, err := DB.Prepare("DELETE FROM users_address WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delUserData.Exec(strconv.Atoi(id))

	message["status"] = "true"
	message["response"] = "Data deleted successfully."
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
	// vars := mux.Vars(r)
	// id := vars["id"]

	// for index, article := range Articles {
	// 	if article.Id == id {
	// 		Articles = append(Articles[:index], Articles[index+1:]...)
	// 	}
	// }

}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/v1/getalldata", returnAlldata).Methods("GET")
	myRouter.HandleFunc("/api/v1/insertdetails", createNewData).Methods("POST")
	myRouter.HandleFunc("/api/v1/deletedata/{id}", deleteUserDetails).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

// get connecton db generic function here ...

func getconnection() {
	db, err := sql.Open("mysql", "root:A@g!ggn$di7@tcp(localhost:3306)/userdb")
	if err != nil {
		panic(err.Error())
	}
	DB = db
	fmt.Println("Success!")
}

func main() {
	getconnection()
	// Articles = []Article{
	// 	Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	// 	Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	// }
	handleRequests()
}

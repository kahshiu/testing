package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/another", anotherHandler)
	log.Println(fmt.Sprintf("Server running on http://localhost%s", ":4000"))
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalf("could not run the server %v", err)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//db, err := sql.Open("postgres", "postgres://postgres:test123@localhost:5432/caretaker")
	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=caretaker user=postgres password=test123 sslmode=disable")
	if err != nil {
		log.Fatal("failed database connection", err)
	}
	defer db.Close()

	type User struct {
		TestingID   int
		TestingName string
		//testing_int  int
	}
	var myUser User
	//userSql := "select testing_id, testing_name, testing_int from testing limit 1"
	userSQL := "select testing_id, testing_name from testing limit 1"

	err = db.QueryRow(userSQL).Scan(&myUser.TestingID, &myUser.TestingName)
	//err = db.QueryRow(userSql).Scan(&myUser.testing_id, &myUser.testing_name, &myUser.testing_int)
	if err != nil {
		log.Fatal("failed database query", err)
	}

	fmt.Println(myUser.TestingID)
	fmt.Println(myUser.TestingName)
	var byteArray []byte
	byteArray, err = json.Marshal(myUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	w.Write(byteArray)
	//w.Write([]byte("hello from home another one handler"))

	// ** sample working **//
	// type FruitBasket struct {
	// 	Name    string
	// 	Fruit   []string
	// 	Id      int64  `json:"ref"`
	// 	private string // An unexported field is not encoded.
	// 	Created time.Time
	// }
	// basket := FruitBasket{
	// 	Name:    "Standard",
	// 	Fruit:   []string{"Apple", "Banana", "Orange"},
	// 	Id:      999,
	// 	private: "Second-rate",
	// 	Created: time.Now(),
	// }
	// jsonData, err := json.Marshal(basket)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(string(jsonData))
	// w.Write(jsonData)
	// ** sample working **//

	//query := r.URL.Query()
	// route, present := query["route"
	// if !present || len(route) == 0 {
	// 	fmt.Println("route is missing")
	// }

	// var m map[string]interface{}{
	// 	"f": contactHandler
	// }

	// for k,v := range m {
	// 	switch k {
	// 	case "f":
	// 	}
	// }

}

func anotherHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from anotehr handler"))

}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from contact handler"))
}

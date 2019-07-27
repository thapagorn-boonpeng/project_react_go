package main

import (
    "database/sql"
    "fmt"
    "strconv"
	"log"
    "net/http"
    "encoding/json"
	"io/ioutil"
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
    Id      string `json:"Id"`
    Firstname   string `json:"Firstname"`
    Lastname    string `json:"Lastname"`
    Email string `json:"Email"`
	Gender string `json:"Gender"`
	Age string `json:"Age"`
}

var Users []User

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/allusers", searchAllUser)
	myRouter.HandleFunc("/searchuserbyid/{id}", searchAllUserbyID)
	myRouter.HandleFunc("/searchuserbyfirstname/{firstname}",searchAllUserbyFirstName)
	myRouter.HandleFunc("/searchuserbylastname/{lastname}",searchAllUserbyLastName)
	myRouter.HandleFunc("/searchuserbyemail/{email}",searchAllUserbyEmail)
	myRouter.HandleFunc("/searchuserbygender/{gender}",searchAllUserbyGender)
	myRouter.HandleFunc("/searchuserbyage/{age}",searchAllUserbyAge)
	myRouter.HandleFunc("/searchuserbyagebetween/{fage}/{tage}",searchAllUserbyAgeBetween)
	
	myRouter.HandleFunc("/addnewuser", addNewUser).Methods("POST")
	myRouter.HandleFunc("/edituser", editUser).Methods("POST")
	myRouter.HandleFunc("/deleteuser", deleteUser).Methods("POST")
	
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Full Stack Test Project HomePage!")
    fmt.Println("Full Stack Test Project")
}

func searchAllUser(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUser")
	readDataAllUser("")
    json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyID(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyID")
	
	vars := mux.Vars(r)
    key := vars["id"]

	readDataAllUser("WHERE id = " + key)
	json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyFirstName(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyFirstName")
	
	vars := mux.Vars(r)
    key := vars["firstname"]

	readDataAllUser("WHERE first_name = '" + key + "'")
	json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyLastName(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyLastName")
	
	vars := mux.Vars(r)
    key := vars["lastname"]

	readDataAllUser("WHERE last_name = '" + key + "'")
	json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyEmail(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyEmail")
	
	vars := mux.Vars(r)
    key := vars["email"]

	readDataAllUser("WHERE email = '" + key + "'")
	json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyGender(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyGender")
	
	vars := mux.Vars(r)
    key := vars["gender"]

	readDataAllUser("WHERE gender = '" + key + "'")
	json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyAge(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyAge")
	
	vars := mux.Vars(r)
    key := vars["age"]

	readDataAllUser("WHERE age = " + key )
	json.NewEncoder(w).Encode(Users)
}

func searchAllUserbyAgeBetween(w http.ResponseWriter, r *http.Request){
    fmt.Println("searchAllUserbyAgeBetween")
	
	vars := mux.Vars(r)
    fkey := vars["fage"]
	tkey := vars["tage"]

	readDataAllUser("WHERE age BETWEEN " + fkey + " AND " + tkey )
	json.NewEncoder(w).Encode(Users)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addNewUser")   
    reqBody, _ := ioutil.ReadAll(r.Body)
    var adduser User 
    json.Unmarshal(reqBody, &adduser)
	
	database, _ := sql.Open("sqlite3", "./user.db")
	var newid int
	rows, _ := database.Query("SELECT MAX(id)+1 AS id FROM users")
	for rows.Next() {
        rows.Scan(&newid)
	}
	fmt.Println("User ID :" + strconv.Itoa(newid))
	statement, _ := database.Prepare("INSERT INTO users (id, first_name, last_name, email, gender, age) VALUES (?, ?, ?, ?, ?, ?)")
	statement.Exec(strconv.Itoa(newid), adduser.Firstname, adduser.Lastname, adduser.Email, adduser.Gender, adduser.Age)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("editUser")   
    reqBody, _ := ioutil.ReadAll(r.Body)
    var edituser User 
    json.Unmarshal(reqBody, &edituser)
	
	database, _ := sql.Open("sqlite3", "./user.db")

	fmt.Println("User ID :" + edituser.Id)
	statement, _ := database.Prepare("UPDATE users SET first_name = ?,last_name = ?,email = ?,gender = ?,age = ? WHERE id = ?")
	statement.Exec(edituser.Firstname, edituser.Lastname, edituser.Email, edituser.Gender, edituser.Age, edituser.Id)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("editUser")   
    reqBody, _ := ioutil.ReadAll(r.Body)
    var deluser User 
    json.Unmarshal(reqBody, &deluser)
	
	database, _ := sql.Open("sqlite3", "./user.db")

	fmt.Println("User ID :" + deluser.Id)
	statement, _ := database.Prepare("DELETE FROM users WHERE id = ?")
	statement.Exec(deluser.Id)
}


func readDataAllUser(wherekey string) {
    Users = nil
	database, _ := sql.Open("sqlite3", "./user.db")
    rows, _ := database.Query("SELECT  id, first_name, last_name, email, gender, age FROM users " + wherekey)
	var id int
    var first_name string
    var last_name string
    var email string
    var gender string
    var age int
    for rows.Next() {
        rows.Scan(&id, &first_name, &last_name, &email, &gender, &age)
		Users = append(Users, User{Id: strconv.Itoa(id), Firstname: first_name, Lastname: last_name, Email: email, Gender: gender, Age: strconv.Itoa(age)})
    }
}

func main() {
	fmt.Println("Full Stack Test Project")
    handleRequests()
	
}
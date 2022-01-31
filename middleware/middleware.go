package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api
	"os"       // used to read the environment variable

	"github.com/manohar-rajawat/universityhousing/models" // models package where User schema is defined

	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq" // postgres golang driver
)

type response struct {
	Status string
	Data   []models.Univeristy
}

// create connection with postgres db
func createConnection() *sql.DB {
	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// GetAllUser will return all the users
func GetUniversity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	params := mux.Vars(r)
	users, err := getUniversity(params["name"])

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}
	res := response{
		Status: "OK",
		Data:   users,
	}
	// send all the users as response
	json.NewEncoder(w).Encode(res)
}

// get one user from the DB by its userid
func getUniversity(name string) ([]models.Univeristy, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()
	var universites []models.Univeristy

	// create the select sql query
	sqlStatement := `SELECT university.id,name,logo.url FROM university inner join logo on LOWER(university.name) like LOWER($1) and logo.universityid = university.id limit 10`

	// execute the sql statement
	rows, err := db.Query(sqlStatement, "%"+name+"%")
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var university models.Univeristy
		// unmarshal the row object to user
		err = rows.Scan(&university.ID, &university.Name, &university.URL)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		universites = append(universites, university)

	}

	// return empty user on error
	return universites, err
}

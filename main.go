package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http" // library that provides us with code for creating HTTP server and request response logic
	"os"
	"strconv"

	"github.com/Robbie08/GoTaskManager/objects"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus" // library that helps with loging and monitoring
)

var database *sql.DB

func main() {
	setup()
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/taskmanager/api/v1/addtask", addTask)
	// TODO: Add functionallity to DELETE, UPDATE, GET task
	http.ListenAndServe(":8080", nil)
}

// Default handle for when server is spun up. You can think of this
// as the home page
func defaultPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		/* Server our clients website */
		log.Info("Someone hit the homepage")

	default:
		fmt.Println("This service only supports GET and POST requests")
	}
}

// This handle will add our POST request task to the database
func addTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Info("Someone hit the addTask")
		decoder := json.NewDecoder(r.Body)
		var t objects.Task
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Info("Assignee: " + t.Assignee)
		log.Info("Task: " + t.Title)
		// TODO: Handle POST Request

		/* Add  */
		task := new(objects.Task)
		task.Init(t.Title, t.DateCreated, t.DateDue, t.Assignee)
		task.AddTask(database)

		// TODO: Respond to client, letting them know we added the task correctly (maybe return something from AddTask?)

	default:
		fmt.Println("This service only supports POST requests")
	}
}

// This function is in charge of gracefully shutting down
// our HTTP Server to prevent any external access to the pi
func shutdown(w http.ResponseWriter, r *http.Request) {
	defer database.Close()
	log.Info("Shutting server down...")
	os.Exit(0) // our system exited without any errors
}

func setup() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	log.Info("=============== Setting up Database and Environment ===============")
	psqlInfo := fmt.Sprintf("host=%s port=%d  user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Info("Opening connection with PostgreSQL")
	// OPEN CONNECTION WITH DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Info("Opened connection with PostgreSQL")
	log.Info("Testing connection with PostgreSQL")

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	database = db
	log.Info("Established a successful connection!")
	log.Info("=============== Finished Setting up Database and Environment ===============")
}

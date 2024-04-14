package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8081"

var counts int

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication service")

	//connect to db
	conn := connectToDB()
	if conn == nil {
		log.Fatal("Can not connect to postgres")
	}

	//set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	log.Println("Server started on port ", webPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Error occurred ", err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "host=localhost port=5434 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
	}
	fmt.Println("Config: ", dsn)
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not yet ready... Error: ", err)
			counts++
		} else {
			log.Println("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds...Count is ", counts)
		time.Sleep(2 * time.Second)
		continue
	}
}

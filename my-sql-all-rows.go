package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	id       int
	name     string
	location string
}

func main() {

	// connect to database.
	db, err := sql.Open("mysql", "root@/development")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM development.employees where id=?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// We'll store all the people we return from the db in a slice.
	people := []person{}

	for rows.Next() {

		// create a struct to store the values from the row.
		p := person{}

		// scan copies the values from the current row into the destination.
		// we create pointers to values we expect to find in the current row.
		err := rows.Scan(&p.id, &p.name, &p.location)
		if err != nil {
			log.Fatal(err)
			continue
		}

		// use the built in append function to append the new value into the people slice.
		people = append(people, p)
	}

	// if there are any errors when scanning the rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(people[0].name)

}

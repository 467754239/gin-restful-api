package models

import (
	"log"

	db "github.com/467754239/db-api/database"
)

type Person struct {
	Id        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

//func (p *Person) GetPersons() (persons []Person, err error) {
func (p *Person) GetPersons() ([]Person, error) {
	/*
		persons = make([]Person, 0)
		rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
		defer rows.Close()

		if err != nil {
			return
		}

		for rows.Next() {
			var person Person
			rows.Scan(&person.Id, &person.FirstName, &person.LastName)
			persons = append(persons, person)
		}
		if err = rows.Err(); err != nil {
			return
		}
		return
	*/
	persons := []Person{}
	err := db.SqlDB.Select(&persons, "SELECT * FROM person")
	if err != nil {
		return persons, err
	}
	return persons, nil

}

func (p *Person) GetPerson() (Person, error) {
	var person Person
	err := db.SqlDB.Get(&person, "SELECT * FROM person WHERE id = ?", p.Id)
	if err != nil {
		return person, err
	}
	return person, nil
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM person WHERE id = ?", p.Id)
	if err != nil {
		log.Fatal(err)
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE person SET first_name=?, last_name=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

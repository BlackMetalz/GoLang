package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "kienlt"
	dbPass := "Asd@123123"
	dbName := "go_blog"

	dbc, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return dbc
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	selDB, err := db.Query("SELECT * FROM employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	result := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)

		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		result = append(result, emp)
	}

	tmpl.ExecuteTemplate(w, "Index", result)
	defer db.Close()

}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	empId := r.URL.Query().Get("id")
	selDB, err := db.Query("Select * from employee where id=?", empId)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()

}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	empId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * from employee where id=?", empId)
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insertForm, err := db.Prepare("INSERT INTO employee(name,city) values (?,?)")
		if err != nil {
			panic(err.Error())
		}

		insertForm.Exec(name, city)
		log.Println("Insert: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)

}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("id")

		insertForm, err := db.Prepare("Update employee set name=?, city=? where id=? ")
		if err != nil {
			panic(err.Error())
		}

		insertForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city + "With ID: " + id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	empId := r.URL.Query().Get("id")
	deleteForm, err := db.Prepare("delete from employee where id=?")

	if err != nil {
		panic(err.Error())
	}

	deleteForm.Exec(empId)
	log.Println("Deleted Employee ID: " + empId)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8000",nil)
}

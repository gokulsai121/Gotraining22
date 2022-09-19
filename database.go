package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func dbconnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@123@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	return db, nil
}

func insert(db *sql.DB) {
	var n, e string
	var c, i int
	fmt.Printf("Enter Name, Contact, Id, Email\n")
	fmt.Scan(&n, &c, &i, &e)
	insert, err := db.Query("INSERT INTO test VALUES (?,?,?,?)", n, c, i, e)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted into user tables")
}

func delete(db *sql.DB) {
	var i int
	fmt.Println("Enter ID to delete the record")
	fmt.Scan(&i)
	delete, err := db.Query("DELETE FROM test WHERE id=?", i)
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()
	fmt.Println("Successfully deleted user details")
}

func update(db *sql.DB, i int) {
	var a int
	fmt.Println("What detials to updated (1. Name, 2. Contact, 3. Email")
	fmt.Scan(&a)
	if a == 1 {
		var n string
		fmt.Println("Enter name")
		fmt.Scan(&n)
		sql, err := db.Query("UPDATE test SET name= ? WHERE id = ?", n, i)
		if err != nil {
			panic(err.Error())
		}
		defer sql.Close()
		fmt.Println("Successfully updated name")
	}
	if a == 2 {
		var c int
		fmt.Println("Enter Contact")
		fmt.Scan(&c)
		sql, err := db.Query("UPDATE test SET contact= ? WHERE id = ?", c, i)
		if err != nil {
			panic(err.Error())
		}
		defer sql.Close()
		fmt.Println("Successfully updated contact")
	} else {
		var e int
		fmt.Println("Enter Email")
		fmt.Scan(&e)
		sql, err := db.Query("UPDATE test SET email= ? WHERE id = ?", e, i)
		if err != nil {
			panic(err.Error())
		}
		defer sql.Close()
		fmt.Println("Successfully updated email")
	}
}
func main() {
	var i, e int
	db, err := dbconnection()
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("Create Table test(name varchar(255), contact int(10), id int, email varchar(255));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("1.Insert\n2.Delete\n3.Update")
	fmt.Scan(&i)
	switch i {
	case 1:
		insert(db)
	case 2:
		delete(db)
	case 3:
		fmt.Println("Enter Id to update")
		fmt.Scan(&e)
		update(db, e)
	}
}

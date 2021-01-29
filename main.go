package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root@/gosample")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM users")
    defer rows.Close()
    if err != nil {
        panic(err.Error())
    }

    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            panic(err.Error())
        }
        fmt.Println(id, name)
    }
}
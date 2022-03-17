package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type prize struct {
	item       int
	name       string
	table      string
	image      string
	redeem_url string
	created_at time.Time
}

func main() {
	db, err := sql.Open("mysql", "A0006_web@accu-db:zvaGw7eu@tcp(accu-db.mysql.database.azure.com:3306)/pchome_test_finish?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(1000)

	sql := "SELECT `item`, `name`, `image`, `table`, `redeem_url` FROM `prize` WHERE `status` = ?;"
	rs, err := db.Query(sql, "Y")
	if err != nil {
		log.Fatal(err)
	}

	var prizes []prize
	for rs.Next() {
		var r prize
		err = rs.Scan(&r.item, &r.name, &r.image, &r.table, &r.redeem_url)
		if err != nil {
			log.Fatal(err)
		}
		prizes = append(prizes, r)
	}
	rs.Close()

	for i := range prizes {
		fmt.Println(prizes[i].name)
	}

	defer db.Close()
}

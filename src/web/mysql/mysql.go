package mysql

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Prize struct {
	item       int
	name       string
	table      string
	image      string
	redeem_url string
	created_at time.Time
}


func Db(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "A0006_web@accu-db:zvaGw7eu@tcp(accu-db.mysql.database.azure.com:3306)/pchome_test_finish?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.Ping()

	sql := "SELECT `item`, `name`, `image`, `table`, `redeem_url` FROM `prize` WHERE `status` = ?;"
	rs, err := db.Query(sql, "Y")
	if err != nil {
		log.Fatal(err)
	}

	var prizes []Prize
	for rs.Next() {
		var r Prize
		err = rs.Scan(&r.item, &r.name, &r.image, &r.table, &r.redeem_url)
		if err != nil {
			log.Fatal(err)
		}
		prizes = append(prizes, r)
	}
	rs.Close()

	for i := range prizes {
		// fmt.Println(prizes[i].name)

		js := []byte(prizes[i].name)
		w.Write(js)
	}

}

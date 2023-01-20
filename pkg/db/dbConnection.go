package dbConnection

import (
	"database/sql"
	"fmt"

	jsonFileReader "spin-fud/pkg/utils"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "5128950"
// 	dbname   = "SpinFud"
// )

func DBConnect() {
	db, err := sql.Open("postgres", jsonFileReader.GetJSONFile("pkg/config/config.json", "ConnectionString"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	// Veritabanına sorgu yollar
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var UserName string
		fmt.Println(rows.Columns())
		if err := rows.Scan(&UserName); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(UserName)
	}
}

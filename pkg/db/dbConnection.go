package dbConnection

import (
	"database/sql"
	"fmt"

	iniParser "rest-api-test/pkg/utils"

	_ "github.com/lib/pq"
)

func DBConnect() {
	// veritabanına bağlanmak için gerekli olan ayarlar,
	value, err := iniParser.GetValueFromINI("Database", "ConnectionString")
	fmt.Println(err)
	fmt.Println(value)
	if err != nil {
		return
	}
	connStr := value
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Veritabanına sorgu yollar
	rows, err := db.Query("SELECT * FROM User")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	fmt.Println(rows)
	// Sonuçları yazdır
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(name)
	}
}

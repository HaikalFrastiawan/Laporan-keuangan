package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/laporan_keuangan")
	if err != nil { panic(err) }
	rows, err := db.Query("DESCRIBE transactions")
	if err != nil { panic(err) }
	for rows.Next() {
		var Field, Type, Null, Key, Default, Extra sql.NullString
		rows.Scan(&Field, &Type, &Null, &Key, &Default, &Extra)
		fmt.Printf("%s %s\n", Field.String, Type.String)
	}
}

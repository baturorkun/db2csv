package pkg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"strings"
)

var db *gorm.DB

type Cols []string
type Data []Cols


func SetupDB() {
	var err error

	switch strings.ToLower(DatabaseSetting.Type) {

		case "sqlite3":
			db, err = gorm.Open("sqlite3",
					DatabaseSetting.SqliteFile)
			db.SingularTable(true)

		case "postgres":

			db, err = gorm.Open(DatabaseSetting.Type,
				fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
					DatabaseSetting.Host,
					DatabaseSetting.Port,
					DatabaseSetting.User,
					DatabaseSetting.Password,
					DatabaseSetting.DbName,
					DatabaseSetting.SslMode,
				))

		case "mysql":

			db, err = gorm.Open(DatabaseSetting.Type,
				fmt.Sprintf("%s:%s@%s:%d/%s?charset=utf8&parseTime=True&loc=Local",
					DatabaseSetting.User,
					DatabaseSetting.Password,
					DatabaseSetting.Host,
					DatabaseSetting.Port,
					DatabaseSetting.DbName,
				))

	}


	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}


}


func Populate() (cols Cols, data Data) {

	rows, _ := db.Raw(AppSetting.Sql).Rows()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		//fmt.Printf(">>>>> %#v\n", result)
		data = append(data,result)
	}

	return
}



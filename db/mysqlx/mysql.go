package mysqlx

import (
	"fmt"

	"gorm.io/gorm"
)

func Truncate(db *gorm.DB, table string) {
	db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
}

func DropDatabase(db *gorm.DB, database string) {
	db.Exec(fmt.Sprintf("DROP DATABASE %s", database))
}

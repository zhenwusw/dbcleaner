package engine

import (
	"database/sql"
	"fmt"
)

// MySQL dbEngine
type MySQL struct {
	db *sql.DB
}

// NewMySQLEngine returns Mysql engine that knows how to truncate a table
func NewMySQLEngine(db *sql.DB) *MySQL {
	return &MySQL{
		db: db,
	}
}

func (mysql *MySQL) Truncate(table string) error {
	tx, err := mysql.db.Begin()
	if err != nil {
		return err
	}

	cmds := []string{
		"SET FOREIGN_KEY_CHECKS = 0",
		fmt.Sprintf("TRUNCATE %s", table),
		"SET FOREIGN_KEY_CHECKS = 1",
	}

	for _, cmd := range cmds {
		if _, err := tx.Exec(cmd); err != nil {
			return tx.Rollback()
		}
	}

	return tx.Commit()
}

func (mysql *MySQL) Close() error {
	return mysql.db.Close()
}

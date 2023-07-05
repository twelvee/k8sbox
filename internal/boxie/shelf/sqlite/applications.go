package sqlite

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

func createApplication(app structs.Application, boxName string) error {
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	sqlStmt := "insert into applications ('name', 'chart', 'box_name') values (?, ?, ?)"

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	_, err = stmt.Exec(app.Name, app.Chart, boxName)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

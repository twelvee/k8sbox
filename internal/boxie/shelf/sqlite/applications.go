package sqlite

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

func createApplication(app structs.Application, force bool, boxID int64) error {
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

	sqlStmt := "insert into applications ('name', 'chart', 'box_id') values (?, ?, ?)"
	if force {
		sqlStmt = "insert into applications ('name', 'chart', 'box_id') values (?, ?, ?) on conflict do update set name=?, chart=?, box_id=?"
	}

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

	if !force {
		_, err = stmt.Exec(app.Name, app.Chart, boxID)
	} else {
		_, err = stmt.Exec(app.Name, app.Chart, boxID, app.Name, app.Chart, boxID)
	}
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

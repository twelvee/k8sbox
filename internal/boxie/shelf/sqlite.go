// Package shelf contains database methods that handle REST API and CLI creation commands
package shelf

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	_ "modernc.org/sqlite"
)

func createSqliteTables() {
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	sqlStmt := `
	create table if not exists boxes (
	    id integer not null primary key,
	    name varchar(255) not null unique,
	    namespace varchar(255) null,
	    box_type varchar(16) default "helm" not null,
	    helm_chart text not null,
	    helm_values text not null,
		created_at integer(4) not null default (strftime('%s','now'))
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func putBoxIntoSQLite(box structs.Box, force bool) error {
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	sqlStmt := "insert into boxes ('name', 'namespace', 'box_type', 'helm_chart', 'helm_values') values (?, ?, ?, ?, ?)"
	if force {
		sqlStmt = "insert into boxes ('name', 'namespace', 'box_type', 'helm_chart', 'helm_values') values (?, ?, ?, ?, ?) on conflict do update set namespace=?, box_type=?, helm_chart=?, helm_values=?"
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
		_, err = stmt.Exec(box.Name, box.Namespace, box.Type, box.Chart, box.Values)
	} else {
		_, err = stmt.Exec(box.Name, box.Namespace, box.Type, box.Chart, box.Values, box.Namespace, box.Type, box.Chart, box.Values)
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

func deleteBoxFromSQLite(name string) error {
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	sqlStmt := "delete from boxes where name=?"

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
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getBoxesFromSQLite() ([]structs.Box, error) {
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	rows, err := db.Query("select name, namespace, box_type, created_at as created from boxes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var boxes []structs.Box
	for rows.Next() {
		var name string
		var namespace string
		var box_type string
		var created string

		err = rows.Scan(&name, &namespace, &box_type, &created)
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, structs.Box{Name: name, Namespace: namespace, Type: box_type, Created: created})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return boxes, nil
}

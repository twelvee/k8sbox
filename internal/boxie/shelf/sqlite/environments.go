package sqlite

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

func getEnvironment(name string) (structs.Box, error) {
	var box structs.Box
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return box, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select name, namespace, box_type, helm_chart, helm_values, id from boxes where name = ?")
	if err != nil {
		return box, err
	}
	defer stmt.Close()

	var boxID string

	err = stmt.QueryRow(name).Scan(&box.Name, &box.Namespace, &box.Type, &box.Chart, &box.Values, &boxID)
	if err != nil {
		return box, err
	}

	stmt, err = db.Prepare("select name, chart from applications where box_id = ?")
	if err != nil {
		return box, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(boxID)
	if err != nil {
		return box, err
	}
	var applications []structs.Application
	for rows.Next() {
		var name string
		var chart string

		err = rows.Scan(&name, &chart)
		if err != nil {
			return box, err
		}
		applications = append(applications, structs.Application{Name: name, Chart: chart})
	}
	err = rows.Err()
	if err != nil {
		return box, err
	}

	box.Applications = applications

	return box, nil
}

func putEnvironment(box structs.Box, force bool) error {
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

	for _, a := range box.Applications {
		err = createApplication(a, box.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteEnvironment(name string) error {
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

func getEnvironments() ([]structs.Box, error) {
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return nil, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	rows, err := db.Query("select id, name, namespace, box_type, created_at as created from boxes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var boxes []structs.Box
	for rows.Next() {
		var boxID string
		var name string
		var namespace string
		var box_type string
		var created string

		err = rows.Scan(&boxID, &name, &namespace, &box_type, &created)
		if err != nil {
			return nil, err
		}

		stmt, err := db.Prepare("select name, chart from applications where box_id = ?")
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		appRows, err := stmt.Query(boxID)
		if err != nil {
			return nil, err
		}

		var applications []structs.Application
		for appRows.Next() {
			var name string
			var chart string

			err = rows.Scan(&name, &chart)
			if err != nil {
				return nil, err
			}
			applications = append(applications, structs.Application{Name: name, Chart: chart})
		}
		err = rows.Err()
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, structs.Box{Name: name, Namespace: namespace, Type: box_type, Created: created, Applications: applications})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return boxes, nil
}

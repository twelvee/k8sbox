package sqlite

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

func getEnvironment(request structs.GetEnvironmentRequest) (structs.Environment, error) {
	var environment structs.Environment
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return environment, err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select name, namespace, user_id, cluster_name, created_at, status from environments where name = ?")
	if err != nil {
		return environment, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(request.Name).Scan(&environment.Name, &environment.Namespace, &environment.UserID, &environment.ClusterName, &environment.CreatedAt, &environment.Status)
	if err != nil {
		return environment, err
	}

	// get environment applications
	stmtApps, err := db.Prepare("select box_name, chart, created_at, application_name from environment_applications where environment_name = ?")
	if err != nil {
		return environment, err
	}
	defer stmtApps.Close()

	rows, err := stmtApps.Query(environment.Name)
	if err != nil {
		return environment, err
	}
	var applications []structs.EnvironmentApplication
	for rows.Next() {
		var boxName string
		var chart string
		var createdAt string
		var appName string

		err = rows.Scan(&boxName, &chart, &createdAt, &appName)
		if err != nil {
			return environment, err
		}
		applications = append(applications, structs.EnvironmentApplication{BoxName: boxName, Chart: chart, EnvironmentName: environment.Name, CreatedAt: createdAt, Name: appName})
	}
	err = rows.Err()
	if err != nil {
		return environment, err
	}
	environment.EnvironmentApplications = applications

	// get environment variables
	stmtVars, err := db.Prepare("select name, value, created_at from environment_variables where environment_name = ?")
	if err != nil {
		return environment, err
	}
	defer stmtVars.Close()

	rowsVars, err := stmtVars.Query(environment.Name)
	if err != nil {
		return environment, err
	}

	m := make(map[string]string)
	for rowsVars.Next() {
		var name string
		var value string
		// todo: make struct to fit created at field
		var createdAt string

		err = rowsVars.Scan(&name, &value, &createdAt)
		if err != nil {
			return environment, err
		}
		m[name] = value
	}
	err = rowsVars.Err()
	if err != nil {
		return environment, err
	}

	environment.VariablesMap = m

	return environment, nil
}

func putEnvironment(environment structs.Environment, user structs.User) error {
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := "insert into environments ('name', 'namespace', 'cluster_name', 'user_id') values (?, ?, ?, ?)"

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(environment.Name, environment.Namespace, environment.ClusterName, user.ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		return err
	}
	for _, box := range environment.Boxes {
		for k, a := range box.HelmRender {
			err = createEnvironmentApplication(structs.EnvironmentApplication{
				BoxName:         box.Name,
				Chart:           a,
				EnvironmentName: environment.Name,
				Name:            k,
			})
			if err != nil {
				return err
			}
		}
	}

	for k, v := range environment.VariablesMap {
		sqlStmt = "insert into environment_variables ('name', 'value', 'environment_name') values (?, ?, ?)"

		tx, err = db.Begin()
		if err != nil {
			return err
		}
		stmt, err = tx.Prepare(sqlStmt)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(k, v, environment.Name)
		if err != nil {
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}

	defer stmt.Close()

	return nil
}

func updateEnvironmentStatus(request structs.UpdateEnvironmentStatusRequest) error {
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

	sqlStmt := "update environments set status=? where name=?;"

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
	_, err = stmt.Exec(request.Status, request.Name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func deleteEnvironment(request structs.DeleteEnvironmentRequest) error {
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

	sqlStmt := "delete from environments where name=?;"

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
	_, err = stmt.Exec(request.Name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	sqlStmt = "delete from environment_variables where environment_name=?;"

	tx, err = db.Begin()
	if err != nil {
		return err
	}
	stmt, err = tx.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)
	_, err = stmt.Exec(request.Name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	sqlStmt = "delete from environment_applications where environment_name=?;"

	tx, err = db.Begin()
	if err != nil {
		return err
	}
	stmt, err = tx.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)
	_, err = stmt.Exec(request.Name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getEnvironments() ([]structs.Environment, error) {
	var environments []structs.Environment
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return environments, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	rows, err := db.Query("select name, namespace, user_id, cluster_name, created_at, status from environments")
	if err != nil {
		return environments, err
	}
	defer rows.Close()
	for rows.Next() {
		var environment structs.Environment

		err = rows.Scan(&environment.Name, &environment.Namespace, &environment.UserID, &environment.ClusterName, &environment.CreatedAt, &environment.Status)
		if err != nil {
			return environments, err
		}

		stmt, err := db.Prepare("select chart, box_name, created_at, application_name from environment_applications where environment_name = ?")
		if err != nil {
			return environments, err
		}
		defer stmt.Close()

		appRows, err := stmt.Query(environment.Name)
		if err != nil {
			return environments, err
		}

		var applications []structs.EnvironmentApplication
		for appRows.Next() {
			var envApp structs.EnvironmentApplication
			err = appRows.Scan(&envApp.Chart, &envApp.BoxName, &envApp.CreatedAt, &envApp.Name)
			if err != nil {
				return environments, err
			}
			envApp.EnvironmentName = environment.Name
			applications = append(applications, envApp)
		}
		err = appRows.Err()
		if err != nil {
			return environments, err
		}
		environment.EnvironmentApplications = applications

		// get environment variables
		varsStmt, err := db.Prepare("select name, value, created_at from environment_variables where environment_name = ?")
		if err != nil {
			return environments, err
		}
		defer varsStmt.Close()

		varsRows, err := varsStmt.Query(environment.Name)
		if err != nil {
			return environments, err
		}

		m := make(map[string]string)
		for varsRows.Next() {
			var name string
			var value string
			// todo: make struct to fit created at field
			var createdAt string

			err = varsRows.Scan(&name, &value, &createdAt)
			if err != nil {
				return environments, err
			}
			m[name] = value
		}
		err = varsRows.Err()
		if err != nil {
			return environments, err
		}
		environment.VariablesMap = m

		environments = append(environments, environment)
	}
	err = rows.Err()
	if err != nil {
		return environments, err
	}

	return environments, nil
}

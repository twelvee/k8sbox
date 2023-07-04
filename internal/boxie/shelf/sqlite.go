// Package shelf contains database methods that handle REST API and CLI creation commands
package shelf

import (
	"database/sql"
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"github.com/twelvee/boxie/pkg/boxie/utils"
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
	create table if not exists users(
	  	id integer not null primary key,
	  	name varchar(255) not null,
	    email varchar(255) not null unique,
	    password text null,
	    created_at integer(4) not null default (strftime('%s','now')),
	    updated_at integer(4) not null default (strftime('%s','now')),
	    invite_code varchar(255) null,
	    token text null
	);

	create table if not exists boxes (
	    id integer not null primary key,
	    name varchar(255) not null unique,
	    namespace varchar(255) null,
	    box_type varchar(16) default "helm" not null,
	    helm_chart text null,
	    helm_values text null,
		created_at integer(4) not null default (strftime('%s','now'))
	);

	create table if not exists applications(
	  	id integer not null primary key,
	  	name varchar(255) not null,
	    chart text not null,
	    box_id integer not null,
	    created_at integer(4) not null default (strftime('%s','now')),
	    
	    CONSTRAINT fk_box
			FOREIGN KEY (box_id)
			REFERENCES boxes (id)
			ON DELETE CASCADE
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func getBoxFromSQLite(name string) (structs.Box, error) {
	var box structs.Box
	db, err := sql.Open("sqlite", currentConnectionDSN)
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
	var r sql.Result
	if !force {
		r, err = stmt.Exec(box.Name, box.Namespace, box.Type, box.Chart, box.Values)
	} else {
		r, err = stmt.Exec(box.Name, box.Namespace, box.Type, box.Chart, box.Values, box.Namespace, box.Type, box.Chart, box.Values)
	}
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	boxID, err := r.LastInsertId()
	if err != nil {
		return err
	}

	for _, a := range box.Applications {
		err = insertApplicationToSQLite(a, force, boxID)
		if err != nil {
			return err
		}
	}

	return nil
}

func insertApplicationToSQLite(app structs.Application, force bool, boxID int64) error {
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

func createUserSQLite(request structs.CreateUserRequest) (string, error) {
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	sqlStmt := "insert into users ('name', 'email', 'invite_code') values (?, ?, ?)"

	tx, err := db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return "", err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	inviteCode := utils.GetShortID(8)

	_, err = stmt.Exec(request.Name, request.Email, inviteCode)

	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return inviteCode, nil
}

func deleteUserSQLite(request structs.DeleteUserRequest) error {
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

	sqlStmt := "delete from users where id=?"

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

	_, err = stmt.Exec(request.ID)

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getUserSQLite(token string) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return user, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select id, name, email, created_at, updated_at from users where token = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(token).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	user.Token = token
	user.Password = "hidden"

	return user, nil
}

func checkInviteCodeSQLite(code string) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return user, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select id, name, email, created_at, updated_at, invite_code from users where invite_code = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(code).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.InviteCode)
	if err != nil {
		return user, err
	}

	return user, nil
}

func setUserPasswordSQLite(code string, password string) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return user, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select id, name, email, created_at, updated_at, invite_code from users where invite_code = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(code).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.InviteCode)
	if err != nil {
		return user, err
	}

	p, err := hashPassword(password)
	if err != nil {
		return user, err
	}

	user.Password = p
	user.Token = generateSecureToken(32)

	sqlStmt := "update users SET password=?, token=?, invite_code=null where id=?"

	tx, err := db.Begin()
	if err != nil {
		return user, err
	}
	stmt, err = tx.Prepare(sqlStmt)
	if err != nil {
		return user, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	_, err = stmt.Exec(user.Password, user.Token, user.ID)

	if err != nil {
		return user, err
	}

	err = tx.Commit()
	if err != nil {
		return user, err
	}

	return user, nil
}

func createTokenSQLite(request structs.LoginRequest) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", currentConnectionDSN)
	if err != nil {
		return user, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select id, name, email, created_at, updated_at, password from users where email = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	var password string
	err = stmt.QueryRow(request.Email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &password)
	if err != nil {
		return user, err
	}

	if !checkPasswordHash(request.Password, password) {
		return user, fmt.Errorf("Incorrect password")
	}

	user.Token = generateSecureToken(32)

	// Update user token in DB
	sqlStmt := "update users SET token=? where id=?"

	tx, err := db.Begin()
	if err != nil {
		return user, err
	}
	stmt, err = tx.Prepare(sqlStmt)
	if err != nil {
		return user, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	_, err = stmt.Exec(user.Token, user.ID)

	if err != nil {
		return user, err
	}

	err = tx.Commit()
	if err != nil {
		return user, err
	}

	return user, nil
}

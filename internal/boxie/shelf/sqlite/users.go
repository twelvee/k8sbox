package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"github.com/twelvee/boxie/pkg/boxie/utils"
)

func createUser(request structs.CreateUserRequest) (string, error) {
	db, err := sql.Open("sqlite", connectionDSN)
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

func deleteUser(request structs.DeleteUserRequest) error {
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

func getUser(token string) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", connectionDSN)
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

func getUsers() ([]structs.User, error) {
	var users []structs.User
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return users, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	rows, err := db.Query("select id, name, email, created_at, updated_at from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userID int32
		var name string
		var email string
		var created string
		var updated string

		err = rows.Scan(&userID, &name, &email, &created, &updated)
		if err != nil {
			return nil, err
		}

		users = append(users, structs.User{Name: name, Email: email, ID: userID, CreatedAt: created, UpdatedAt: updated})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func checkInviteCode(code string) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", connectionDSN)
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

func setUserPassword(code string, password string) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", connectionDSN)
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

	p, err := utils.HashPassword(password)
	if err != nil {
		return user, err
	}

	user.Password = p
	user.Token = utils.GenerateSecureToken(32)

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

func getSetupRequired() (bool, error) {
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return false, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	rows, err := db.Query("select id, name from users")
	defer rows.Close()
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return false, nil
	}
	return true, nil
}

func createToken(request structs.LoginRequest) (structs.User, error) {
	var user structs.User
	db, err := sql.Open("sqlite", connectionDSN)
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

	if !utils.CheckPasswordHash(request.Password, password) {
		return user, fmt.Errorf("Incorrect password")
	}

	user.Token = utils.GenerateSecureToken(32)

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

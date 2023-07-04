// Package shelf contains database methods that handle REST API and CLI creation commands
package shelf

import (
	"fmt"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"strings"
)

var currentConnectionType string
var currentConnectionDSN string

type ConnectionType string

const (
	CONNECTION_SQLITE ConnectionType = "sqlite"
)

var default_sqlite_dsn string = "./shelf_sqlite"

// Shelf is a boxes database
type Shelf struct {
	ShelfType       string
	ShelfDSN        string
	PutBox          func(box structs.Box, force bool) error
	DeleteBox       func(string) error
	GetBoxes        func() ([]structs.Box, error)
	GetBox          func(string) (structs.Box, error)
	CreateUser      func(request structs.CreateUserRequest) (string, error)
	GetUser         func(string) (structs.User, error)
	DeleteUser      func(request structs.DeleteUserRequest) error
	CreateToken     func(request structs.LoginRequest) (structs.User, error)
	AcceptInvite    func(code string) (structs.User, error)
	SetUserPassword func(code string, password string) (structs.User, error)
}

// NewShelf will return a new instance of shelf
func NewShelf(connectionType string, dsn string) Shelf {
	currentConnectionType = connectionType
	currentConnectionDSN = dsn
	if len(strings.TrimSpace(currentConnectionType)) == 0 {
		currentConnectionType = string(CONNECTION_SQLITE)
	}
	if len(strings.TrimSpace(currentConnectionDSN)) == 0 {
		if currentConnectionType == string(CONNECTION_SQLITE) {
			currentConnectionDSN = default_sqlite_dsn
		} else {
			panic(fmt.Errorf("Failed to init a shelf. DSN string is empty."))
		}
	}

	if currentConnectionType == string(CONNECTION_SQLITE) {
		// Will throw panic on error
		createSqliteTables()
	}

	return Shelf{
		ShelfType:       connectionType,
		ShelfDSN:        dsn,
		PutBox:          putBox,
		DeleteBox:       deleteBox,
		GetBoxes:        getBoxes,
		GetBox:          getBox,
		CreateUser:      createUser,
		GetUser:         getUser,
		DeleteUser:      deleteUser,
		CreateToken:     createToken,
		AcceptInvite:    acceptInvite,
		SetUserPassword: setUserPassword,
	}
}

func getBox(name string) (structs.Box, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return getBoxFromSQLite(name)
	}
	return structs.Box{}, fmt.Errorf("Failed to open shelf.")
}

func putBox(box structs.Box, force bool) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return putBoxIntoSQLite(box, force)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func deleteBox(boxName string) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return deleteBoxFromSQLite(boxName)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func getBoxes() ([]structs.Box, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return getBoxesFromSQLite()
	}
	return nil, fmt.Errorf("Failed to open shelf.")
}

func createUser(request structs.CreateUserRequest) (string, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return createUserSQLite(request)
	}
	return "", fmt.Errorf("Failed to open shelf.")
}

func deleteUser(request structs.DeleteUserRequest) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return deleteUserSQLite(request)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func getUser(token string) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return getUserSQLite(token)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func createToken(request structs.LoginRequest) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return createTokenSQLite(request)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func acceptInvite(code string) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return checkInviteCodeSQLite(code)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func setUserPassword(code string, password string) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return setUserPasswordSQLite(code, password)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

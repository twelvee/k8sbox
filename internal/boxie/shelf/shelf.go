// Package shelf contains database methods that handle REST API and CLI creation commands
package shelf

import (
	"fmt"
	"github.com/twelvee/boxie/internal/boxie/shelf/sqlite"
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
	ShelfType          string
	ShelfDSN           string
	ShelfSQLiteAdapter *sqlite.SQLite

	// Setup
	GetSetupRequired func() (bool, error)

	// Boxes
	PutBox    func(box structs.Box, force bool) error
	DeleteBox func(string) error
	GetBoxes  func() ([]structs.Box, error)
	GetBox    func(string) (structs.Box, error)

	// Users
	CreateUser      func(request structs.CreateUserRequest) (string, error)
	GetUser         func(string) (structs.User, error)
	DeleteUser      func(request structs.DeleteUserRequest) error
	CreateToken     func(request structs.LoginRequest) (structs.User, error)
	AcceptInvite    func(code string) (structs.User, error)
	SetUserPassword func(code string, password string) (structs.User, error)
	GetUsers        func() ([]structs.User, error)

	// Clusters
	PutCluster           func(cluster structs.Cluster, force bool) error
	GetClusters          func() ([]structs.Cluster, error)
	GetCluster           func(request structs.GetClusterRequest) (structs.Cluster, error)
	DeleteCluster        func(request structs.DeleteClusterRequest) error
	SetClusterConnection func(cluster structs.Cluster) (bool, error)
	UpdateCluster        func(cluster structs.Cluster) error
}

var adapter *sqlite.SQLite

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
		adapter = sqlite.NewSQLite(currentConnectionDSN)
	}

	return Shelf{
		ShelfType:          connectionType,
		ShelfDSN:           dsn,
		ShelfSQLiteAdapter: adapter,

		GetSetupRequired: getSetupRequired,

		PutBox:    putBox,
		DeleteBox: deleteBox,
		GetBoxes:  getBoxes,
		GetBox:    getBox,

		CreateUser:      createUser,
		GetUser:         getUser,
		DeleteUser:      deleteUser,
		CreateToken:     createToken,
		AcceptInvite:    acceptInvite,
		SetUserPassword: setUserPassword,
		GetUsers:        getUsers,

		GetClusters:          getClusters,
		GetCluster:           getCluster,
		DeleteCluster:        deleteCluster,
		PutCluster:           putCluster,
		SetClusterConnection: setClusterConnection,
		UpdateCluster:        updateCluster,
	}
}

func getSetupRequired() (bool, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetSetupRequired()
	}
	return false, fmt.Errorf("Failed to open shelf.")
}

func getBox(name string) (structs.Box, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetBox(name)
	}
	return structs.Box{}, fmt.Errorf("Failed to open shelf.")
}

func putBox(box structs.Box, force bool) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.PutBox(box, force)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func deleteBox(boxName string) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.DeleteBox(boxName)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func getBoxes() ([]structs.Box, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetBoxes()
	}
	return nil, fmt.Errorf("Failed to open shelf.")
}

func createUser(request structs.CreateUserRequest) (string, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.CreateUser(request)
	}
	return "", fmt.Errorf("Failed to open shelf.")
}

func deleteUser(request structs.DeleteUserRequest) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.DeleteUser(request)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func getUser(token string) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetUser(token)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func createToken(request structs.LoginRequest) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.CreateToken(request)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func acceptInvite(code string) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.CheckInviteCode(code)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func setUserPassword(code string, password string) (structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.SetUserPassword(code, password)
	}
	return structs.User{}, fmt.Errorf("Failed to open shelf.")
}

func getUsers() ([]structs.User, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetUsers()
	}
	return nil, fmt.Errorf("Failed to open shelf.")
}

func putCluster(cluster structs.Cluster, force bool) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.PutCluster(cluster, force)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func updateCluster(cluster structs.Cluster) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.UpdateCluster(cluster)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func getClusters() ([]structs.Cluster, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetClusters()
	}
	return nil, fmt.Errorf("Failed to open shelf.")
}

func getCluster(request structs.GetClusterRequest) (structs.Cluster, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.GetCluster(request)
	}
	return structs.Cluster{}, fmt.Errorf("Failed to open shelf.")
}

func deleteCluster(request structs.DeleteClusterRequest) error {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.DeleteCluster(request)
	}
	return fmt.Errorf("Failed to open shelf.")
}

func setClusterConnection(cluster structs.Cluster) (bool, error) {
	if currentConnectionType == string(CONNECTION_SQLITE) {
		return adapter.SetClusterConnection(cluster)
	}
	return false, fmt.Errorf("Failed to open shelf.")
}

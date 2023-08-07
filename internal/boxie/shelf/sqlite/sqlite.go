// Package shelf contains database methods that handle REST API and CLI creation commands
package sqlite

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	_ "modernc.org/sqlite"
)

var connectionDSN string

func createSqliteTables(conn string) {
	connectionDSN = conn
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		panic(err)
	}
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

	create table if not exists clusters(
	  	name varchar(255) not null unique,
	    kubeconfig text not null,
	    is_active boolean not null default false,
	    created_at integer(4) not null default (strftime('%s','now'))
	);

	create table if not exists boxes (
	    name varchar(255) not null unique,
	    namespace varchar(255) null,
	    box_type varchar(16) default "helm" not null,
	    helm_chart text null,
	    helm_values text null,
		created_at integer(4) not null default (strftime('%s','now'))
	);

	create table if not exists applications(
	  	name varchar(255) not null,
	    chart text not null,
	    box_name varchar(255) not null,
	    created_at integer(4) not null default (strftime('%s','now')),
	    
	    CONSTRAINT fk_box_name
			FOREIGN KEY (box_name)
			REFERENCES boxes (name)
			ON DELETE CASCADE
	);

	create table if not exists environments(
	  	name varchar(255) not null unique,
	    namespace varchar(255) default 'default',
	    user_id integer not null,
	    cluster_name varchar(255) not null,
	    status integer not null default 0,
	    created_at integer(4) not null default (strftime('%s','now')),
	    
	    CONSTRAINT fk_user
			FOREIGN KEY (user_id)
			REFERENCES users (id)
			ON DELETE CASCADE,
	    CONSTRAINT fk_cluster_name
			FOREIGN KEY (cluster_name)
			REFERENCES clusters (name)
			ON DELETE CASCADE
	);

	create table if not exists environment_applications(
	    environment_name varchar(255) not null,
	    application_name varchar(255) not null,
	    box_name varchar(255) not null,
	    chart varchar(255) not null,
	    created_at integer(4) not null default (strftime('%s','now')),
	    CONSTRAINT fk_environment_name
			FOREIGN KEY (environment_name)
			REFERENCES environments (name)
			ON DELETE CASCADE,
	    CONSTRAINT fk_box_name
			FOREIGN KEY (box_name)
			REFERENCES boxes (name)
			ON DELETE NO ACTION
	);

	create table if not exists environment_variables(
	    environment_name varchar(255) not null,
	    name varchar(255) not null,
	    value text not null,
	    created_at integer(4) not null default (strftime('%s','now')),
	    CONSTRAINT fk_environment_name
			FOREIGN KEY (environment_name)
			REFERENCES environments (name)
			ON DELETE CASCADE
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	db.Close()
}

type SQLite struct {
	connectionDNS string

	CreateSQLiteTables func(conn string)
	GetSetupRequired   func() (bool, error)

	// Users
	CreateUser      func(request structs.CreateUserRequest) (string, error)
	DeleteUser      func(request structs.DeleteUserRequest) error
	GetUser         func(token string) (structs.User, error)
	GetUsers        func() ([]structs.User, error)
	CheckInviteCode func(code string) (structs.User, error)
	SetUserPassword func(code string, password string) (structs.User, error)
	CreateToken     func(request structs.LoginRequest) (structs.User, error)

	// Boxes
	GetBox    func(name string) (structs.Box, error)
	PutBox    func(box structs.Box, force bool) error
	UpdateBox func(box structs.Box) error
	DeleteBox func(name string) error
	GetBoxes  func() ([]structs.Box, error)

	// Applications
	CreateApplications func(app structs.Application, boxName string) error

	// Clusters
	GetCluster           func(request structs.GetClusterRequest) (structs.Cluster, error)
	PutCluster           func(cluster structs.Cluster, force bool) error
	DeleteCluster        func(request structs.DeleteClusterRequest) error
	GetClusters          func() ([]structs.Cluster, error)
	SetClusterConnection func(cluster structs.Cluster) (bool, error)
	UpdateCluster        func(cluster structs.Cluster) error

	// Environments
	GetEnvironment          func(request structs.GetEnvironmentRequest) (structs.Environment, error)
	GetEnvironments         func() ([]structs.Environment, error)
	PutEnvironment          func(environment structs.Environment, user structs.User) error
	DeleteEnvironment       func(request structs.DeleteEnvironmentRequest) error
	UpdateEnvironmentStatus func(request structs.UpdateEnvironmentStatusRequest) error
}

func NewSQLite(conn string) *SQLite {
	createSqliteTables(conn)
	return &SQLite{
		connectionDNS:           connectionDSN,
		CreateSQLiteTables:      createSqliteTables,
		GetSetupRequired:        getSetupRequired,
		CreateApplications:      createApplication,
		CreateUser:              createUser,
		DeleteUser:              deleteUser,
		GetUsers:                getUsers,
		GetUser:                 getUser,
		CheckInviteCode:         checkInviteCode,
		SetUserPassword:         setUserPassword,
		CreateToken:             createToken,
		GetBox:                  getBox,
		GetBoxes:                getBoxes,
		DeleteBox:               deleteBox,
		PutBox:                  putBox,
		UpdateBox:               updateBox,
		GetCluster:              getCluster,
		GetClusters:             getClusters,
		DeleteCluster:           deleteCluster,
		PutCluster:              putCluster,
		SetClusterConnection:    setClusterConnection,
		UpdateCluster:           updateCluster,
		GetEnvironment:          getEnvironment,
		GetEnvironments:         getEnvironments,
		PutEnvironment:          putEnvironment,
		DeleteEnvironment:       deleteEnvironment,
		UpdateEnvironmentStatus: updateEnvironmentStatus,
	}
}

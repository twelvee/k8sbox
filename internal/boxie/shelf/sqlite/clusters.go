package sqlite

import (
	"database/sql"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

func getCluster(request structs.GetClusterRequest) (structs.Cluster, error) {
	var cluster structs.Cluster
	db, err := sql.Open("sqlite", connectionDSN)
	if err != nil {
		return cluster, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	stmt, err := db.Prepare("select id, name, kubeconfig, is_active, created_at from clusters where name = ?")
	if err != nil {
		return cluster, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(request.Name).Scan(&cluster.ID, &cluster.Name, &cluster.Kubeconfig, &cluster.IsActive, &cluster.CreatedAt)
	if err != nil {
		return cluster, err
	}

	return cluster, nil
}

func putCluster(cluster structs.Cluster, force bool) error {
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

	sqlStmt := "insert into clusters ('name', 'kubeconfig') values (?, ?)"
	if force {
		sqlStmt = "insert into clusters ('name', 'kubeconfig') values (?, ?) on conflict do update set kubeconfig=?"
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
		_, err = stmt.Exec(cluster.Name, cluster.Kubeconfig)
	} else {
		_, err = stmt.Exec(cluster.Name, cluster.Kubeconfig, cluster.Kubeconfig)
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

func deleteCluster(request structs.DeleteClusterRequest) error {
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

	sqlStmt := "delete from clusters where name=?"

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

	return nil
}

func getClusters() ([]structs.Cluster, error) {
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

	rows, err := db.Query("select name, kubeconfig, is_active, created_at from clusters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var clusters []structs.Cluster
	for rows.Next() {
		var name string
		var kubeconfig string
		var isActive bool
		var created string

		err = rows.Scan(&name, &kubeconfig, &isActive, &created)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, structs.Cluster{Name: name, Kubeconfig: kubeconfig, IsActive: isActive, CreatedAt: created})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return clusters, nil
}
